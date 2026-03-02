package orderitem_cache

import "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache"

type OrderItemCache interface {
	OrderItemQueryCache
}

type orderItemCache struct {
	OrderItemQueryCache
}

func NewOrderItemCache(store *cache.CacheStore) OrderItemCache {
	return &orderItemCache{
		OrderItemQueryCache: NewOrderItemQueryCache(store),
	}
}
