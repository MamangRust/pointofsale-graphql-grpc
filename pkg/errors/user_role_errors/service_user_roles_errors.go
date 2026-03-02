package userrole_errors

import (
	"net/http"

	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/errors"
)

var (
	ErrFailedAssignRoleToUser = errors.NewErrorResponse("Failed to assign role to user", http.StatusInternalServerError)
	ErrFailedRemoveRole       = errors.NewErrorResponse("Failed to remove role from user", http.StatusInternalServerError)
)
