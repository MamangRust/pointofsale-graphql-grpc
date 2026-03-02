package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/cache"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/handler/graph"
	graphql "github.com/MamangRust/pointofsale-graphql-grpc/internal/mapper"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/middlewares"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/permission"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/auth"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/dotenv"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/logger"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/observability"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/otel"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/resilience"
	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/upload_image"
	"github.com/grafana/pyroscope-go"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/ast"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

const (
	defaultClientAddr             = "localhost:50051"
	defaultWindowSizeClient       = 16 * 1024 * 1024
	defaultKeepaliveTimeClient    = 20 * time.Second
	defaultKeepaliveTimeoutClient = 5 * time.Second
)

var (
	addr = flag.String("addr", defaultClientAddr, "the gRPC server address to grpcConnect to")
)

type Client struct {
	Logger       logger.LoggerInterface
	TokenManager *auth.Manager
	Resolver     *graph.Resolver

	Port         string
	CacheStore   *cache.CacheStore
	Redis        *redis.Client
	Telemetry    *otel.Telemetry
	GRPCgrpcConn *grpc.ClientConn
	cancelTasks  context.CancelFunc
	tasksDone    []<-chan struct{}
}

type ClientConfig struct {
	ServiceName    string
	ServiceVersion string
	Environment    string
	OtelEndpoint   string
	GRPCAddr       string
	ServerPort     string
	AllowedOrigins []string
}

type CacheManager struct {
	cache  *cache.CacheStore
	logger logger.LoggerInterface
}

func NewCacheManager(cache *cache.CacheStore, logger logger.LoggerInterface) *CacheManager {
	return &CacheManager{
		cache:  cache,
		logger: logger,
	}
}

func (cm *CacheManager) StartMonitoring(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		ticker := time.NewTicker(monitoringInterval)
		defer ticker.Stop()

		cm.logger.Info("Cache monitoring task started",
			zap.Duration("interval", monitoringInterval),
		)

		for {
			select {
			case <-ctx.Done():
				cm.logger.Info("Cache monitoring task stopped")
				return
			case <-ticker.C:
				cm.monitor(ctx)
			}
		}
	}()

	return done
}

func (cm *CacheManager) monitor(ctx context.Context) {
	refCount := cm.cache.GetRefCount()

	stats, err := cm.cache.GetStats(ctx)
	if err != nil {
		cm.logger.Error("Failed to get cache stats", zap.Error(err))
		return
	}

	logLevel := zap.InfoLevel
	if refCount > cacheRefCountThreshold {
		logLevel = zap.WarnLevel
	}

	if ce := cm.logger.Check(logLevel, "Cache statistics"); ce != nil {
		ce.Write(
			zap.Int64("ref_count", refCount),
			zap.Int64("total_keys", stats.TotalKeys),
			zap.Float64("hit_rate", stats.HitRate),
			zap.String("memory_used", stats.MemoryUsedHuman),
			zap.Bool("high_ref_count", refCount > cacheRefCountThreshold),
		)
	}
}

func (cm *CacheManager) StartCleanup(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		ticker := time.NewTicker(cleanupInterval)
		defer ticker.Stop()

		cm.logger.Info("Cache cleanup task started",
			zap.Duration("interval", cleanupInterval),
		)

		for {
			select {
			case <-ctx.Done():
				cm.logger.Info("Cache cleanup task stopped")
				return
			case <-ticker.C:
				cm.cleanup(ctx)
			}
		}
	}()

	return done
}

func (cm *CacheManager) cleanup(ctx context.Context) {
	cm.logger.Info("Starting periodic cache cleanup")

	statsBefore, err := cm.cache.GetStats(ctx)
	if err != nil {
		cm.logger.Error("Failed to get cache stats before cleanup", zap.Error(err))
		statsBefore = nil
	}

	scanned, err := cm.cache.ClearExpired(ctx)
	if err != nil {
		cm.logger.Error("Cache cleanup failed", zap.Error(err))
		return
	}

	statsAfter, err := cm.cache.GetStats(ctx)
	if err != nil {
		cm.logger.Error("Failed to get cache stats after cleanup", zap.Error(err))
		statsAfter = nil
	}

	logFields := []zap.Field{
		zap.Int64("scanned_keys", scanned),
		zap.Int64("ref_count", cm.cache.GetRefCount()),
	}

	if statsBefore != nil && statsAfter != nil {
		keysRemoved := statsBefore.TotalKeys - statsAfter.TotalKeys
		logFields = append(logFields,
			zap.Int64("keys_before", statsBefore.TotalKeys),
			zap.Int64("keys_after", statsAfter.TotalKeys),
			zap.Int64("keys_removed", keysRemoved),
			zap.String("memory_before", statsBefore.MemoryUsedHuman),
			zap.String("memory_after", statsAfter.MemoryUsedHuman),
		)
	}

	cm.logger.Info("Cache cleanup completed", logFields...)
}

func NewClient(cfg *ClientConfig) (*Client, error) {
	if err := initPyroscope(); err != nil {
		log.Fatal("Failed to initialize pyroscope:", err)
	}

	if cfg == nil {
		cfg = &ClientConfig{
			ServiceName:    "client",
			ServiceVersion: "1.0.0",
			Environment:    getEnv("ENVIRONMENT", "production"),
			OtelEndpoint:   getEnv("OTEL_ENDPOINT", "otel-collector:4317"),
			GRPCAddr:       *addr,
			ServerPort:     "5000",
			AllowedOrigins: []string{"http://localhost:1420"},
		}
	}

	telemetry, err := initTelemetryClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize telemetry: %w", err)
	}

	cacheMetrics, err := observability.NewCacheMetrics("cache")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cache metrics: %w", err)
	}

	appLogger, err := logger.NewLogger(cfg.ServiceName, telemetry.GetLogger())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	if err := dotenv.Viper(); err != nil {
		appLogger.Fatal("Failed to load .env file", zap.Error(err))
	}

	grpcConn, err := grpcConnectToGRPC(cfg.GRPCAddr, appLogger)
	if err != nil {
		return nil, fmt.Errorf("failed to grpcConnect to gRPC server: %w", err)
	}

	appLogger.Info("gRPC grpcConnection established", zap.String("addr", cfg.GRPCAddr))

	tokenManager, err := auth.NewManager(viper.GetString("SECRET_KEY"), appLogger)
	if err != nil {
		appLogger.Fatal("Failed to create token manager", zap.Error(err))
	}

	ctx := context.Background()

	redisClient, err := initRedisClient(ctx, appLogger)
	if err != nil {
		grpcConn.Close()
		return nil, fmt.Errorf("failed to initialize Redis: %w", err)
	}

	cacheStore := cache.NewCacheStore(redisClient, appLogger, cacheMetrics)

	tasksCtx, cancelTasks := context.WithCancel(ctx)
	cacheManager := NewCacheManager(cacheStore, appLogger)

	tasksDone := []<-chan struct{}{
		cacheManager.StartMonitoring(tasksCtx),
		cacheManager.StartCleanup(tasksCtx),
	}

	grpcClients := &graph.GRPCClients{
		AuthClient:        pb.NewAuthServiceClient(grpcConn),
		RoleClient:        pb.NewRoleServiceClient(grpcConn),
		UserClient:        pb.NewUserServiceClient(grpcConn),
		MerchantClient:    pb.NewMerchantServiceClient(grpcConn),
		CategoryClient:    pb.NewCategoryServiceClient(grpcConn),
		CashierClient:     pb.NewCashierServiceClient(grpcConn),
		ProductClient:     pb.NewProductServiceClient(grpcConn),
		OrderClient:       pb.NewOrderServiceClient(grpcConn),
		OrderItemClient:   pb.NewOrderItemServiceClient(grpcConn),
		TransactionClient: pb.NewTransactionServiceClient(grpcConn),
	}

	perm := permission.NewPermission(grpcClients.RoleClient)
	mapper := graphql.NewGraphqlMapper()
	upload_image := upload_image.NewImageUpload(appLogger)

	resolver := graph.NewResolver(&graph.Deps{
		Clients:     grpcClients,
		Mapper:      mapper,
		Logger:      appLogger,
		ImageUpload: upload_image,
		Permission:  perm,
		Cache:       cacheStore,
	})

	port := viper.GetString("PORT")
	if port == "" {
		port = cfg.ServerPort
	}

	return &Client{
		Logger:       appLogger,
		TokenManager: tokenManager,
		Resolver:     resolver,
		Port:         port,
		CacheStore:   cacheStore,
		Redis:        redisClient,
		Telemetry:    telemetry,
		GRPCgrpcConn: grpcConn,
		cancelTasks:  cancelTasks,
		tasksDone:    tasksDone,
	}, nil
}

func (c *Client) Run() error {
	c.Logger.Debug("Starting GraphQL server", zap.String("port", c.Port))

	defer c.Cleanup()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: c.Resolver,
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	resilienceMiddleware := c.initResilience()
	authMiddleware := middlewares.AuthMiddleware(c.TokenManager, c.Logger)
	handlerChain := resilienceMiddleware.Middleware()(authMiddleware(srv))

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	mux.Handle("/query", handlerChain)

	server := &http.Server{
		Addr:    ":" + c.Port,
		Handler: mux,
	}

	c.Logger.Info("GraphQL Playground running", zap.String("url", "http://localhost:"+c.Port))

	idlegrpcConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		c.Logger.Info("Received shutdown signal, shutting down...")

		if c.cancelTasks != nil {
			c.cancelTasks()
		}

		for _, done := range c.tasksDone {
			<-done
		}

		c.Logger.Info("All background tasks stopped")

		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			c.Logger.Error("HTTP server shutdown error", zap.Error(err))
		}

		close(idlegrpcConnsClosed)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}

	<-idlegrpcConnsClosed
	c.Logger.Info("Server stopped")
	return nil
}

func (c *Client) initResilience() *middlewares.ResilienceHttpMiddleware {
	return middlewares.NewResilienceHttpMiddleware(
		resilience.NewLoadMonitor(),
		resilience.NewCircuitBreaker(100, 10, c.Logger),
		resilience.NewRequestLimiter(800, c.Logger),
	)
}

func (c *Client) Cleanup() {
	c.Logger.Info("Cleaning up resources...")

	if c.GRPCgrpcConn != nil {
		if err := c.GRPCgrpcConn.Close(); err != nil {
			c.Logger.Error("Failed to close gRPC grpcConnection", zap.Error(err))
		} else {
			c.Logger.Info("gRPC grpcConnection closed")
		}
	}

	if c.Redis != nil {
		if err := c.Redis.Close(); err != nil {
			c.Logger.Error("Failed to close Redis grpcConnection", zap.Error(err))
		} else {
			c.Logger.Info("Redis grpcConnection closed")
		}
	}

	if c.Telemetry != nil {
		if err := c.Telemetry.Shutdown(context.Background()); err != nil {
			c.Logger.Error("Failed to shutdown telemetry", zap.Error(err))
		} else {
			c.Logger.Info("Telemetry shutdown successfully")
		}
	}

	c.Logger.Info("Cleanup completed")
}

func grpcConnectToGRPC(addr string, logger logger.LoggerInterface) (*grpc.ClientConn, error) {
	logger.Info("grpcConnecting to gRPC server", zap.String("address", addr))

	grpcConn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithInitialConnWindowSize(defaultWindowSizeClient),
		grpc.WithInitialWindowSize(defaultWindowSizeClient),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                defaultKeepaliveTimeClient,
			Timeout:             defaultKeepaliveTimeoutClient,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	logger.Info("gRPC grpcConnection established", zap.String("address", addr))
	return grpcConn, nil
}

func initPyroscope() error {
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "client",
		ServerAddress:   os.Getenv("PYROSCOPE_SERVER"),

		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},

		Tags: map[string]string{
			"service": "grpc-client-echo",
			"env":     os.Getenv("ENV"),
			"version": os.Getenv("VERSION"),
		},
	})
	return err
}

func initTelemetryClient(cfg *ClientConfig) (*otel.Telemetry, error) {
	telemetry := otel.NewTelemetry(otel.Config{
		ServiceName:            cfg.ServiceName,
		ServiceVersion:         cfg.ServiceVersion,
		Environment:            cfg.Environment,
		Endpoint:               cfg.OtelEndpoint,
		Insecure:               true,
		EnableRuntimeMetrics:   true,
		RuntimeMetricsInterval: 15 * time.Second,
	})

	if err := telemetry.Init(context.Background()); err != nil {
		return nil, err
	}

	return telemetry, nil
}

func initRedisClient(ctx context.Context, appLogger logger.LoggerInterface) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", viper.GetString("REDIS_HOST_CLIENT"), viper.GetString("REDIS_PORT_CLIENT"))

	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     viper.GetString("REDIS_PASSWORD_CLIENT"),
		DB:           viper.GetInt("REDIS_DB_CLIENT"),
		DialTimeout:  redisDialTimeout,
		ReadTimeout:  redisReadTimeout,
		WriteTimeout: redisWriteTimeout,
		PoolSize:     redisPoolSize,
		MinIdleConns: redisMinIdleConns,
	})

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping Redis at %s: %w", addr, err)
	}

	appLogger.Info("Redis grpcConnection established",
		zap.String("addr", addr),
		zap.Int("db", viper.GetInt("REDIS_DB_CLIENT")),
	)

	return client, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
