package auth_errors

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"
)

var ErrGraphqlLogin = response.NewGraphqlError(
	"error",
	"login failed: invalid argument provided",
	int(http.StatusBadRequest),
)

var ErrGraphqlGetMe = response.NewGraphqlError(
	"error",
	"get user info failed: unauthenticated",
	int(http.StatusUnauthorized),
)

var ErrGraphqlRegisterToken = response.NewGraphqlError(
	"error",
	"register failed: invalid argument",
	int(http.StatusBadRequest),
)
