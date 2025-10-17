package orderitem_errors

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/response"
)

var (
	ErrGraphqlInvalidID = response.NewGraphqlError("error", "invalid ID", int(http.StatusBadRequest))
)
