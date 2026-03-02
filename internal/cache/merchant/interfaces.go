package merchant_cache

import (
	"context"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/domain/requests"
	db "github.com/MamangRust/pointofsale-graphql-grpc/pkg/database/schema"
)

type MerchantQueryCache interface {
	GetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsRow, *int, bool)
	SetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants, data []*db.GetMerchantsRow, total *int)

	GetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsActiveRow, *int, bool)
	SetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants, data []*db.GetMerchantsActiveRow, total *int)

	GetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsTrashedRow, *int, bool)
	SetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants, data []*db.GetMerchantsTrashedRow, total *int)

	GetCachedMerchant(ctx context.Context, id int) (*db.GetMerchantByIDRow, bool)
	SetCachedMerchant(ctx context.Context, data *db.GetMerchantByIDRow)

	GetCachedMerchantsByUserId(ctx context.Context, id int) ([]*db.GetMerchantByIDRow, bool)
	SetCachedMerchantsByUserId(ctx context.Context, userId int, data []*db.GetMerchantByIDRow)
}

type MerchantCommandCache interface {
	DeleteCachedMerchant(ctx context.Context, id int)
}
