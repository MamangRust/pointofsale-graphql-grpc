package transaction_errors

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"
)

var (
	ErrGraphqlInvalidID         = response.NewGraphqlError("error", "invalid ID", int(http.StatusBadRequest))
	ErrGraphqlInvalidMonth      = response.NewGraphqlError("error", "invalid month", int(http.StatusBadRequest))
	ErrGraphqlInvalidYear       = response.NewGraphqlError("error", "invalid year", int(http.StatusBadRequest))
	ErrGraphqlInvalidMerchantId = response.NewGraphqlError("error", "invalid merchant ID", int(http.StatusBadRequest))

	ErrGraphqlValidateCreateTransaction = response.NewGraphqlError("error", "validation failed: invalid create transaction request", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateTransaction = response.NewGraphqlError("error", "validation failed: invalid update transaction request", int(http.StatusBadRequest))
)
