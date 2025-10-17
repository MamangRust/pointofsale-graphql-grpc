package response

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToGrpcErrorFromErrorResponse(err *ErrorResponse) error {
	if err == nil {
		return nil
	}
	return status.Errorf(codes.Code(err.Code),
		"%s", errors.GrpcErrorToJson(&pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
			Code:    int32(err.Code),
		}),
	)
}

func NewGrpcError(statusText string, message string, code int) error {
	return status.Errorf(codes.Code(code),
		"%s", errors.GrpcErrorToJson(&pb.ErrorResponse{
			Status:  statusText,
			Message: message,
			Code:    int32(code),
		}),
	)
}
