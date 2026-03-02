package orderitem_errors

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/errors"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcInvalidID = errors.NewGrpcError("invalid ID", int(codes.InvalidArgument))
)
