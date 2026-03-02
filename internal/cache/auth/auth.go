package auth_cache

import "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache"

type authMencache struct {
	IdentityCache
	LoginCache
}

type AuthMencache interface {
	IdentityCache
	LoginCache
}

func NewMencache(cacheStore *cache.CacheStore) AuthMencache {
	return &authMencache{
		IdentityCache: NewidentityCache(cacheStore),
		LoginCache:    NewLoginCache(cacheStore),
	}
}
