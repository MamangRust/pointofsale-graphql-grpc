package refreshtoken_errors

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"
)

var ErrGraphqlRefreshToken = response.NewGraphqlError("error", "refresh token failed", int(http.StatusUnauthorized))
