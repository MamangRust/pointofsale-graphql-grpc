package refreshtoken_errors

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var ErrGrpcRefreshToken = response.NewGrpcError("error", "refresh token failed", int(codes.Unauthenticated))
