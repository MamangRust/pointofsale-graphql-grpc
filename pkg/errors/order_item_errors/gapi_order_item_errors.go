package orderitem_errors

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcInvalidID = response.NewGrpcError("error", "invalid ID", int(codes.InvalidArgument))
)
