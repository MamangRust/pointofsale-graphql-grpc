package product_errors

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"
)

var (
	ErrGraphqlInvalidID = response.NewGraphqlError("error", "invalid ID", int(http.StatusBadRequest))

	ErrGraphqlValidateCreateProduct = response.NewGraphqlError("error", "validation failed: invalid create product request", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateProduct = response.NewGraphqlError("error", "validation failed: invalid update product request", int(http.StatusBadRequest))
)
