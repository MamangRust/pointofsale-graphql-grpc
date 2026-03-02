package permission

import (
	"context"
	"errors"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
)

type Permission interface {
	HasRole(ctx context.Context, userID int, allowedRoles ...string) (bool, error)
}

type permission struct {
	ctx         context.Context
	roleService pb.RoleServiceClient
}

func NewPermission(roleService pb.RoleServiceClient) *permission {
	return &permission{
		roleService: roleService,
	}
}

func (s *permission) HasRole(ctx context.Context, userID int, allowedRoles ...string) (bool, error) {
	roleResponse, errResp := s.roleService.FindByUserId(ctx, &pb.FindByIdUserRoleRequest{
		UserId: int32(userID),
	})
	if errResp != nil {
		return false, errors.New("failed to fetch user role")
	}

	for _, role := range roleResponse.Data {
		for _, allowed := range allowedRoles {
			if role.Name == allowed {
				return true, nil
			}
		}
	}
	return false, nil
}
