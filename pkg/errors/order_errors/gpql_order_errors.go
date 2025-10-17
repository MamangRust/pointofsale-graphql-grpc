package order_errors

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"
)

var (
	ErrGraphqlInvalidYear             = response.NewGraphqlError("error", "Invalid year", int(http.StatusBadRequest))
	ErrGraphqlInvalidMonth            = response.NewGraphqlError("error", "Invalid month", int(http.StatusBadRequest))
	ErrGraphqlFailedInvalidMerchantId = response.NewGraphqlError("error", "Invalid merchant ID", int(http.StatusBadRequest))
	ErrGraphqlFailedInvalidId         = response.NewGraphqlError("error", "Invalid ID", int(http.StatusBadRequest))

	ErrGraphqlValidateCreateOrder = response.NewGraphqlError("error", "validation failed: invalid create order request", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateOrder = response.NewGraphqlError("error", "validation failed: invalid update order request", int(http.StatusBadRequest))
)
