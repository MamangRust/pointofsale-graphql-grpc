package role_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/cache"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/model"
)

const (
	roleAllCacheKey      = "role:all:page:%d:pageSize:%d:search:%s"
	roleByIdCacheKey     = "role:id:%d"
	roleActiveCacheKey   = "role:active:page:%d:pageSize:%d:search:%s"
	roleTrashedCacheKey  = "role:trashed:page:%d:pageSize:%d:search:%s"
	roleByUserIdCacheKey = "role:user_id:%d"

	ttlDefault = 5 * time.Minute
)

type roleQueryCache struct {
	store *cache.CacheStore
}

func NewRoleQueryCache(store *cache.CacheStore) *roleQueryCache {
	return &roleQueryCache{store: store}
}

func (m *roleQueryCache) SetCachedRoles(ctx context.Context, req *model.FindAllRoleInput, res *model.APIResponsePaginationRole) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(roleAllCacheKey, req.Page, req.PageSize, req.Search)
	// Langsung simpan objek ApiResponse
	cache.SetToCache(ctx, m.store, key, res, ttlDefault)
}

func (m *roleQueryCache) SetCachedRoleById(ctx context.Context, res *model.APIResponseRole) {
	if res == nil || res.Data == nil {
		return
	}

	// Asumsi: res.Data memiliki field ID untuk membuat kunci cache
	key := fmt.Sprintf(roleByIdCacheKey, res.Data.ID)
	cache.SetToCache(ctx, m.store, key, res, ttlDefault)
}

func (m *roleQueryCache) SetCachedRoleByUserId(ctx context.Context, userId int, res *model.APIResponsesRole) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(roleByUserIdCacheKey, userId)
	cache.SetToCache(ctx, m.store, key, res, ttlDefault)
}

func (m *roleQueryCache) SetCachedRoleActive(ctx context.Context, req *model.FindAllRoleInput, res *model.APIResponsePaginationRoleDeleteAt) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(roleActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, m.store, key, res, ttlDefault)
}

func (m *roleQueryCache) SetCachedRoleTrashed(ctx context.Context, req *model.FindAllRoleInput, res *model.APIResponsePaginationRoleDeleteAt) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(roleTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, m.store, key, res, ttlDefault)
}

func (m *roleQueryCache) GetCachedRoles(ctx context.Context, req *model.FindAllRoleInput) (*model.APIResponsePaginationRole, bool) {
	key := fmt.Sprintf(roleAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*model.APIResponsePaginationRole](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *roleQueryCache) GetCachedRoleById(ctx context.Context, id int) (*model.APIResponseRole, bool) {
	key := fmt.Sprintf(roleByIdCacheKey, id)

	result, found := cache.GetFromCache[*model.APIResponseRole](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *roleQueryCache) GetCachedRoleByUserId(ctx context.Context, userId int) (*model.APIResponsesRole, bool) {
	key := fmt.Sprintf(roleByUserIdCacheKey, userId)

	result, found := cache.GetFromCache[*model.APIResponsesRole](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *roleQueryCache) GetCachedRoleActive(ctx context.Context, req *model.FindAllRoleInput) (*model.APIResponsePaginationRoleDeleteAt, bool) {
	key := fmt.Sprintf(roleActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*model.APIResponsePaginationRoleDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *roleQueryCache) GetCachedRoleTrashed(ctx context.Context, req *model.FindAllRoleInput) (*model.APIResponsePaginationRoleDeleteAt, bool) {
	key := fmt.Sprintf(roleTrashedCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*model.APIResponsePaginationRoleDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}
