package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/logger"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func InvalidAccessToken() error {
	return fmt.Errorf("invalid access token")
}

func HandleApiError(w http.ResponseWriter, err error, traceID string) {
	if err == nil {
		return
	}

	var apiErr *AppError
	if errors.As(err, &apiErr) {
		response := ErrorResponse{
			Status:      "error",
			Message:     apiErr.Message,
			Code:        apiErr.Code,
			TraceID:     traceID,
			Retryable:   apiErr.Retryable,
			Validations: apiErr.Validations,
		}
		w.WriteHeader(apiErr.Code)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := ErrorResponse{
		Status:  "error",
		Message: "An internal server error occurred",
		Code:    http.StatusInternalServerError,
		TraceID: traceID,
	}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(response)
}

type HttpFunction func(w http.ResponseWriter, r *http.Request) error

type HttpHandler interface {
	Handle(method string, handler HttpFunction) http.HandlerFunc
	HandleApiErrorWithTracing(w http.ResponseWriter, r *http.Request, err error, span trace.Span, method string)
}

type httpHandler struct {
	observability observability.TraceLoggerObservability
	logger        logger.LoggerInterface
}

func NewHttpHandler(observability observability.TraceLoggerObservability, logger logger.LoggerInterface) HttpHandler {
	return &httpHandler{
		observability: observability,
		logger:        logger,
	}
}

func (h *httpHandler) Handle(method string, handler HttpFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span, end, status, logSuccess := h.observability.StartTracingAndLogging(
			r.Context(),
			method,
			attribute.String("path", r.URL.Path),
			attribute.String("method", r.Method),
		)

		r = r.WithContext(ctx)

		defer func() {
			end(status)
		}()

		err := handler(w, r)
		if err != nil {
			status = "error"
			h.HandleApiErrorWithTracing(w, r, err, span, method)
			return
		}

		logSuccess("Request completed successfully")
	}
}

func (h *httpHandler) HandleApiErrorWithTracing(w http.ResponseWriter, r *http.Request, err error, span trace.Span, method string) {
	traceID := span.SpanContext().TraceID().String()

	h.logger.Error(
		fmt.Sprintf("API error in %s", method),
		zap.Error(err),
		zap.String("trace.id", traceID),
		zap.String("path", r.URL.Path),
		zap.String("method", r.Method),
	)

	span.SetAttributes(
		attribute.String("trace.id", traceID),
		attribute.String("error", err.Error()),
	)
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())

	HandleApiError(w, err, traceID)
}

type ErrorResponse struct {
	Status      string            `json:"status"`
	Message     string            `json:"message"`
	Code        int               `json:"code"`
	TraceID     string            `json:"trace_id,omitempty"`
	Retryable   bool              `json:"retryable,omitempty"`
	Validations []ValidationError `json:"validations,omitempty"`
}

func NewBadRequestError(message string) *AppError {
	return ErrBadRequest.WithMessage(message)
}

func NewValidationError(validations []ValidationError) *AppError {
	return ErrValidationFailed.WithValidations(validations)
}

func NewNotFoundError(resource string) *AppError {
	return ErrNotFound.WithMessage(fmt.Sprintf("%s not found", resource))
}

func NewConflictError(message string) *AppError {
	return ErrConflict.WithMessage(message)
}

func NewInternalError(err error) *AppError {
	return ErrInternal.WithInternal(err)
}

func NewServiceUnavailableError(service string) *AppError {
	return ErrServiceUnavailable.WithMessage(fmt.Sprintf("%s is temporarily unavailable", service))
}
