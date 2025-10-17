package response

import (
	"fmt"
)

func NewErrorResponse(message string, code int) *ErrorResponse {
	return &ErrorResponse{
		Status:  "error",
		Message: message,
		Code:    code,
	}
}

func ToGraphqlErrorFromErrorResponse(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("graphql error: %v", err)
}

func NewGraphqlError(statusText string, message string, code int) error {
	errResp := &ErrorResponse{
		Status:  statusText,
		Message: message,
		Code:    code,
	}

	return fmt.Errorf("graphql error: [%d] %s - %s", errResp.Code, errResp.Status, errResp.Message)
}
