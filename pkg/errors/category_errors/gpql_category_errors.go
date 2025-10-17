package category_errors

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"
)

var (
	ErrGraphqlFailedInvalidId         = response.NewGraphqlError("error", "Invalid ID", int(http.StatusBadRequest))
	ErrGraphqlFailedInvalidMerchantId = response.NewGraphqlError("error", "Invalid Merchant ID", int(http.StatusBadRequest))
	ErrGraphqlFailedInvalidYear       = response.NewGraphqlError("error", "Invalid year", int(http.StatusBadRequest))
	ErrGraphqlFailedInvalidMonth      = response.NewGraphqlError("error", "Invalid month", int(http.StatusBadRequest))

	ErrGraphqlValidateCreateCategory = response.NewGraphqlError("error", "validation failed: invalid create category request", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateCategory = response.NewGraphqlError("error", "validation failed: invalid update category request", int(http.StatusBadRequest))
)
