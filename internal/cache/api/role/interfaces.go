package role_cache

import (
	"context"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/model"
)

type RoleQueryCache interface {
	SetCachedRoles(ctx context.Context, req *model.FindAllRoleInput, res *model.APIResponsePaginationRole)
	SetCachedRoleById(ctx context.Context, res *model.APIResponseRole)
	SetCachedRoleByUserId(ctx context.Context, userId int, res *model.APIResponsesRole)
	SetCachedRoleActive(ctx context.Context, req *model.FindAllRoleInput, res *model.APIResponsePaginationRoleDeleteAt)
	SetCachedRoleTrashed(ctx context.Context, req *model.FindAllRoleInput, res *model.APIResponsePaginationRoleDeleteAt)

	GetCachedRoles(ctx context.Context, req *model.FindAllRoleInput) (*model.APIResponsePaginationRole, bool)
	GetCachedRoleByUserId(ctx context.Context, userId int) (*model.APIResponsesRole, bool)
	GetCachedRoleById(ctx context.Context, id int) (*model.APIResponseRole, bool)
	GetCachedRoleActive(ctx context.Context, req *model.FindAllRoleInput) (*model.APIResponsePaginationRoleDeleteAt, bool)
	GetCachedRoleTrashed(ctx context.Context, req *model.FindAllRoleInput) (*model.APIResponsePaginationRoleDeleteAt, bool)
}

type RoleCommandCache interface {
	DeleteCachedRole(ctx context.Context, id int)
}
