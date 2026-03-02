package middlewares

import (
	"net/http"
	"strconv"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/service"
	mycontext "github.com/MamangRust/pointofsale-graphql-grpc/pkg/context"
)

func RoleMiddleware(allowedRoles []string, roleService service.RoleService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userIDValue := r.Context().Value(mycontext.UserIDContextKey)
			if userIDValue == nil {
				http.Error(w, "User ID not found in context", http.StatusForbidden)
				return
			}

			userIDStr, ok := userIDValue.(string)
			if !ok || userIDStr == "" {
				http.Error(w, "Invalid user ID in context", http.StatusForbidden)
				return
			}

			userIDInt, err := strconv.Atoi(userIDStr)
			if err != nil {
				http.Error(w, "Invalid user ID format", http.StatusBadRequest)
				return
			}

			roleResponse, errResp := roleService.FindByUserId(r.Context(), userIDInt)

			if errResp != nil {
				http.Error(w, "Failed to fetch user role", http.StatusInternalServerError)
				return
			}

			authorized := false
			for _, userRole := range roleResponse {
				for _, allowed := range allowedRoles {
					if userRole.RoleName == allowed {
						authorized = true
						break
					}
				}
				if authorized {
					break
				}
			}

			if !authorized {
				http.Error(w, "You do not have permission to access this resource", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
