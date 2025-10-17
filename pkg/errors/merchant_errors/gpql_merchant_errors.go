package merchant_errors

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"
)

var (
	ErrGraphqlInvalidID = response.NewGraphqlError("error", "invalid ID", int(http.StatusBadRequest))

	ErrGraphqlValidateCreateMerchant = response.NewGraphqlError("error", "validation failed: invalid create merchant request", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateMerchant = response.NewGraphqlError("error", "validation failed: invalid update merchant request", int(http.StatusBadRequest))
)
