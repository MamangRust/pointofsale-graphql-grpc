package product_errors

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/errors"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcInvalidID = errors.NewGrpcError("invalid ID", int(codes.InvalidArgument))

	ErrGrpcValidateCreateProduct = errors.NewGrpcError("validation failed: invalid create product request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateProduct = errors.NewGrpcError("validation failed: invalid update product request", int(codes.InvalidArgument))
)
