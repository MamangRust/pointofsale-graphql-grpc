package main

import (
	"fmt"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/app"
	"go.uber.org/zap"
)

func main() {
	server, err := app.NewServer()
	if err != nil {
		server.Logger.Error("Failed to initialize server",
			zap.String("stage", "initialization"),
			zap.Error(err),
		)
		panic(fmt.Sprintf("❌ Server initialization failed: %v", err))
	}

	server.Run()
}
