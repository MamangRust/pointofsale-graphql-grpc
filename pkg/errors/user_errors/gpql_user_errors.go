package user_errors

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"
)

var (
	ErrGraphqlUserNotFound  = response.NewGraphqlError("error", "User not found", int(http.StatusNotFound))
	ErrGraphqlUserInvalidId = response.NewGraphqlError("error", "Invalid User ID", int(http.StatusNotFound))

	ErrGraphqlValidateCreateUser = response.NewGraphqlError("error", "validation failed: invalid create User request", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateUser = response.NewGraphqlError("error", "validation failed: invalid update User request", int(http.StatusBadRequest))
)
