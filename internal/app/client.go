package app

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/handler/graph"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/mapper/response/graphql"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/middlewares"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/auth"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/dotenv"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/logger"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/upload_image"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/ast"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

const defaultPort = "8080"

type Client struct {
	Logger       logger.LoggerInterface
	TokenManager *auth.Manager
	Resolver     *graph.Resolver
	Ctx          context.Context
	Port         string
}

func NewClient() (*Client, error) {
	lg, err := logger.NewLogger()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		lg.Fatal("Failed to connect to server", zap.Error(err))
	}

	if err := dotenv.Viper(); err != nil {
		lg.Fatal("Failed to load .env file", zap.Error(err))
	}

	tokenManager, err := auth.NewManager(viper.GetString("SECRET_KEY"), lg)
	if err != nil {
		lg.Fatal("Failed to create token manager", zap.Error(err))
	}

	ctx := context.Background()

	mapperGraphql := graphql.NewGraphqlMapper()

	upload_image := upload_image.NewImageUpload(lg)

	grpcClient := &graph.GRPCClients{
		AuthClient:        pb.NewAuthServiceClient(conn),
		RoleClient:        pb.NewRoleServiceClient(conn),
		UserClient:        pb.NewUserServiceClient(conn),
		MerchantClient:    pb.NewMerchantServiceClient(conn),
		CategoryClient:    pb.NewCategoryServiceClient(conn),
		CashierClient:     pb.NewCashierServiceClient(conn),
		ProductClient:     pb.NewProductServiceClient(conn),
		OrderClient:       pb.NewOrderServiceClient(conn),
		OrderItemClient:   pb.NewOrderItemServiceClient(conn),
		TransactionClient: pb.NewTransactionServiceClient(conn),
	}

	resolver := graph.NewResolver(
		grpcClient,
		mapperGraphql,
		upload_image,
	)

	port := viper.GetString("PORT")
	if port == "" {
		port = defaultPort
	}

	return &Client{
		Logger:       lg,
		TokenManager: tokenManager,
		Ctx:          ctx,
		Port:         port,
		Resolver:     resolver,
	}, nil
}

func (s *Client) Run() error {
	s.Logger.Debug("Starting GraphQL Client", zap.Any("port", s.Port))

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: s.Resolver,
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", middlewares.AuthMiddleware(s.TokenManager, s.Logger)(srv))

	s.Logger.Debug("GraphQL Playground running at", zap.String("url", "http://localhost:"+s.Port))
	return http.ListenAndServe(":"+s.Port, nil)
}
