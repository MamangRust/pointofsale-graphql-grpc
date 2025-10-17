package graph

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/mapper/response/graphql"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/upload_image"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthGraphql        AuthHandleGraphql
	RoleGraphql        RoleHandleGraphql
	UserGraphql        UserHandleGraphql
	MerchantGraphql    MerchantHandleGraphql
	CategoryGraphql    CategoryHandleGraphql
	CashierGraphql     CashierHandleGraphql
	ProductGraphql     ProductHandleGraphql
	OrderGraphql       OrderHandleGraphql
	OrderItemGraphql   OrderItemHandleGraphql
	TransactionGraphql TransactionHandleGraphql
}

type GRPCClients struct {
	AuthClient        pb.AuthServiceClient
	RoleClient        pb.RoleServiceClient
	UserClient        pb.UserServiceClient
	MerchantClient    pb.MerchantServiceClient
	CategoryClient    pb.CategoryServiceClient
	CashierClient     pb.CashierServiceClient
	ProductClient     pb.ProductServiceClient
	OrderClient       pb.OrderServiceClient
	OrderItemClient   pb.OrderItemServiceClient
	TransactionClient pb.TransactionServiceClient
}

func NewResolver(
	clients *GRPCClients,
	mapper *graphql.GraphqlMapper,
	imageUpload upload_image.ImageUploads,
) *Resolver {
	return &Resolver{
		AuthGraphql: AuthHandleGraphql{
			AuthClient: clients.AuthClient,
			Mapping:    mapper.AuthGraphqlMapper,
		},
		RoleGraphql: RoleHandleGraphql{
			RoleClient: clients.RoleClient,
			Mapping:    mapper.RoleGraphqlMapper,
		},
		UserGraphql: UserHandleGraphql{
			UserClient: clients.UserClient,
			Mapping:    mapper.UserGraphqlMapper,
		},
		MerchantGraphql: MerchantHandleGraphql{
			MerchantClient: clients.MerchantClient,
			Mapping:        mapper.MerchantGraphqlMapper,
		},
		CategoryGraphql: CategoryHandleGraphql{
			CategoryClient: clients.CategoryClient,
			Mapping:        mapper.CategoryGraphqlMapper,
		},
		CashierGraphql: CashierHandleGraphql{
			CashierClient: clients.CashierClient,
			Mapping:       mapper.CashierGraphqlMapper,
		},
		ProductGraphql: ProductHandleGraphql{
			ProductClient: clients.ProductClient,
			Mapping:       mapper.ProductGraphqlMapper,
			ImageUpload:   imageUpload,
		},
		OrderGraphql: OrderHandleGraphql{
			OrderClient: clients.OrderClient,
			Mapping:     mapper.OrderGraphqlMapper,
		},
		OrderItemGraphql: OrderItemHandleGraphql{
			OrderItemClient: clients.OrderItemClient,
			Mapping:         mapper.OrderItemGraphqlMapper,
		},
		TransactionGraphql: TransactionHandleGraphql{
			TransactionClient: clients.TransactionClient,
			Mapping:           mapper.TransactionGraphqlMapper,
		},
	}
}

type AuthHandleGraphql struct {
	AuthClient pb.AuthServiceClient
	Mapping    graphql.AuthGraphqlMapper
}

type RoleHandleGraphql struct {
	RoleClient pb.RoleServiceClient
	Mapping    graphql.RoleGraphqlMapper
}

type UserHandleGraphql struct {
	UserClient pb.UserServiceClient
	Mapping    graphql.UserGraphqlMapper
}

type MerchantHandleGraphql struct {
	MerchantClient pb.MerchantServiceClient
	Mapping        graphql.MerchantGraphqlMapper
}

type CashierHandleGraphql struct {
	CashierClient pb.CashierServiceClient
	Mapping       graphql.CashierGraphqlMapper
}

type CategoryHandleGraphql struct {
	CategoryClient pb.CategoryServiceClient
	Mapping        graphql.CategoryGraphqlMapper
}

type ProductHandleGraphql struct {
	ProductClient pb.ProductServiceClient
	Mapping       graphql.ProductGraphqlMapper
	ImageUpload   upload_image.ImageUploads
}

type OrderHandleGraphql struct {
	OrderClient pb.OrderServiceClient
	Mapping     graphql.OrderGraphqlMapper
}

type OrderItemHandleGraphql struct {
	OrderItemClient pb.OrderItemServiceClient
	Mapping         graphql.OrderItemGraphqlMapper
}

type TransactionHandleGraphql struct {
	TransactionClient pb.TransactionServiceClient
	Mapping           graphql.TransactionGraphqlMapper
}
