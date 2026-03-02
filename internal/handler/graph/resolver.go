package graph

import (
	errorstd "errors"
	"fmt"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/cache"
	auth_cache "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache/api/auth"
	cashier_cache "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache/api/cashier"
	category_cache "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache/api/category"
	merchant_cache "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache/api/merchant"
	order_cache "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache/api/order"
	orderitem_cache "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache/api/order_item"
	product_cache "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache/api/product"
	role_cache "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache/api/role"
	transaction_cache "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache/api/transaction"
	user_cache "github.com/MamangRust/pointofsale-graphql-grpc/internal/cache/api/user"
	graphql "github.com/MamangRust/pointofsale-graphql-grpc/internal/mapper"
	"github.com/go-playground/validator/v10"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/permission"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/errors"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/logger"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/observability"
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
	ResolverHandle     *resolverHandler
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

type Deps struct {
	Clients     *GRPCClients
	Mapper      *graphql.GraphqlMapper
	Logger      logger.LoggerInterface
	Permission  permission.Permission
	Cache       *cache.CacheStore
	ImageUpload upload_image.ImageUploads
}

func NewResolver(deps *Deps) *Resolver {
	observability, _ := observability.NewObservability(
		"graphql-client",
		deps.Logger,
	)

	resolverHandle := NewResolverHandler(observability, deps.Logger)

	authCache := auth_cache.NewMencache(deps.Cache)
	userCache := user_cache.NewUserMencache(deps.Cache)
	roleCache := role_cache.NewRoleMencache(deps.Cache)
	categoryCache := category_cache.NewCategoryMencache(deps.Cache)
	merchantCache := merchant_cache.NewMerchantMencache(deps.Cache)
	cashierCache := cashier_cache.NewCashierMencache(deps.Cache)
	orderItemCache := orderitem_cache.NewOrderItemCache(deps.Cache)
	orderCache := order_cache.NewOrderMencache(deps.Cache)
	productCache := product_cache.NewProductMencache(deps.Cache)
	transactionCache := transaction_cache.NewTransactionMencache(deps.Cache)

	return &Resolver{
		AuthGraphql: AuthHandleGraphql{
			AuthClient: deps.Clients.AuthClient,
			Mapping:    deps.Mapper.AuthGraphqlMapper,
			Logger:     deps.Logger,
			Cache:      authCache,
		},
		RoleGraphql: RoleHandleGraphql{
			RoleClient: deps.Clients.RoleClient,
			Mapping:    deps.Mapper.RoleGraphqlMapper,
			Logger:     deps.Logger,
			Cache:      roleCache,
		},
		UserGraphql: UserHandleGraphql{
			UserClient: deps.Clients.UserClient,
			Mapping:    deps.Mapper.UserGraphqlMapper,
			Logger:     deps.Logger,
			Cache:      userCache,
		},
		MerchantGraphql: MerchantHandleGraphql{
			MerchantClient: deps.Clients.MerchantClient,
			Mapping:        deps.Mapper.MerchantGraphqlMapper,
			Logger:         deps.Logger,
			Cache:          merchantCache,
		},
		CategoryGraphql: CategoryHandleGraphql{
			CategoryClient: deps.Clients.CategoryClient,
			Mapping:        deps.Mapper.CategoryGraphqlMapper,
			Logger:         deps.Logger,
			Cache:          categoryCache,
		},
		CashierGraphql: CashierHandleGraphql{
			CashierClient: deps.Clients.CashierClient,
			Mapping:       deps.Mapper.CashierGraphqlMapper,
			Logger:        deps.Logger,
			Cache:         cashierCache,
		},
		ProductGraphql: ProductHandleGraphql{
			ProductClient: deps.Clients.ProductClient,
			Mapping:       deps.Mapper.ProductGraphqlMapper,
			ImageUpload:   deps.ImageUpload,
			Logger:        deps.Logger,
			Cache:         productCache,
		},
		OrderGraphql: OrderHandleGraphql{
			OrderClient: deps.Clients.OrderClient,
			Mapping:     deps.Mapper.OrderGraphqlMapper,
			Logger:      deps.Logger,
			Cache:       orderCache,
		},
		OrderItemGraphql: OrderItemHandleGraphql{
			OrderItemClient: deps.Clients.OrderItemClient,
			Mapping:         deps.Mapper.OrderItemGraphqlMapper,
			Logger:          deps.Logger,
			Cache:           orderItemCache,
		},
		TransactionGraphql: TransactionHandleGraphql{
			TransactionClient: deps.Clients.TransactionClient,
			Mapping:           deps.Mapper.TransactionGraphqlMapper,
			Logger:            deps.Logger,
			Cache:             transactionCache,
		},
		ResolverHandle: resolverHandle,
	}
}

type AuthHandleGraphql struct {
	AuthClient pb.AuthServiceClient
	Mapping    graphql.AuthGraphqlMapper
	Logger     logger.LoggerInterface
	Cache      auth_cache.AuthMencache
}

type RoleHandleGraphql struct {
	RoleClient pb.RoleServiceClient
	Mapping    graphql.RoleGraphqlMapper
	Logger     logger.LoggerInterface
	Cache      role_cache.RoleMencache
}

type UserHandleGraphql struct {
	UserClient pb.UserServiceClient
	Mapping    graphql.UserGraphqlMapper
	Logger     logger.LoggerInterface
	Cache      user_cache.UserMencache
}

type MerchantHandleGraphql struct {
	MerchantClient pb.MerchantServiceClient
	Mapping        graphql.MerchantGraphqlMapper
	Logger         logger.LoggerInterface
	Cache          merchant_cache.MerchantMenCache
}

type CashierHandleGraphql struct {
	CashierClient pb.CashierServiceClient
	Mapping       graphql.CashierGraphqlMapper
	Logger        logger.LoggerInterface
	Cache         cashier_cache.CashierMencache
}

type CategoryHandleGraphql struct {
	CategoryClient pb.CategoryServiceClient
	Mapping        graphql.CategoryGraphqlMapper
	Logger         logger.LoggerInterface
	Cache          category_cache.CategoryMencache
}

type ProductHandleGraphql struct {
	ProductClient pb.ProductServiceClient
	Mapping       graphql.ProductGraphqlMapper
	ImageUpload   upload_image.ImageUploads
	Logger        logger.LoggerInterface
	Cache         product_cache.ProductMencache
}

type OrderHandleGraphql struct {
	OrderClient pb.OrderServiceClient
	Mapping     graphql.OrderGraphqlMapper
	Logger      logger.LoggerInterface
	Cache       order_cache.OrderMencache
}

type OrderItemHandleGraphql struct {
	OrderItemClient pb.OrderItemServiceClient
	Mapping         graphql.OrderItemGraphqlMapper
	Logger          logger.LoggerInterface
	Cache           orderitem_cache.OrderItemCache
}

type TransactionHandleGraphql struct {
	TransactionClient pb.TransactionServiceClient
	Mapping           graphql.TransactionGraphqlMapper
	Logger            logger.LoggerInterface
	Cache             transaction_cache.TransactionMencache
}

func (h *Resolver) handleGraphQLError(err error, operation string) *errors.AppError {
	if err == nil {
		return nil
	}

	var appErr *errors.AppError
	if errorstd.As(err, &appErr) {
		return appErr
	}

	return errors.NewInternalError(err).WithMessage("Failed to " + operation)
}

func (h *Resolver) parseValidationErrors(err error) []errors.ValidationError {
	var validationErrs []errors.ValidationError

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			validationErrs = append(validationErrs, errors.ValidationError{
				Field:   fe.Field(),
				Message: h.getValidationMessage(fe),
			})
		}
		return validationErrs
	}

	return []errors.ValidationError{
		{
			Field:   "general",
			Message: err.Error(),
		},
	}
}

func (h *Resolver) getValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Must be at least %s", fe.Param())
	case "max":
		return fmt.Sprintf("Must be at most %s", fe.Param())
	case "gte":
		return fmt.Sprintf("Must be greater than or equal to %s", fe.Param())
	case "lte":
		return fmt.Sprintf("Must be less than or equal to %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", fe.Param())
	default:
		return fmt.Sprintf("Validation failed on '%s' tag", fe.Tag())
	}
}
