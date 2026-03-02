package user_cache

import "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache"

type UserMencache interface {
	UserQueryCache
	UserCommandCache
}

type userMencache struct {
	UserQueryCache
	UserCommandCache
}

func NewUserMencache(cacheStore *cache.CacheStore) UserMencache {
	return &userMencache{
		UserQueryCache:   NewUserQueryCache(cacheStore),
		UserCommandCache: NewUserCommandCache(cacheStore),
	}
}
