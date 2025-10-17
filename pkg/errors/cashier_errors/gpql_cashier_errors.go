package cashier_errors

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"
)

var (
	ErrGraphqlFailedInvalidId         = response.NewGraphqlError("error", "Invalid ID", int(http.StatusBadRequest))
	ErrGraphqlFailedInvalidMerchantId = response.NewGraphqlError("error", "Invalid merchant ID", int(http.StatusBadRequest))
	ErrGraphqlFailedInvalidYear       = response.NewGraphqlError("error", "Invalid year", int(http.StatusBadRequest))
	ErrGraphqlFailedInvalidMonth      = response.NewGraphqlError("error", "Invalid month", int(http.StatusBadRequest))

	ErrGraphqlValidateCreateCashier = response.NewGraphqlError("error", "validation failed: invalid create cashier request", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateCashier = response.NewGraphqlError("error", "validation failed: invalid update cashier request", int(http.StatusBadRequest))
)
