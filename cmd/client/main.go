package main

import (
	"fmt"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/app"
	"go.uber.org/zap"
)

func main() {
	client, err := app.NewClient()
	if err != nil {
		client.Logger.Error("Failed to initialize Client",
			zap.String("stage", "initialization"),
			zap.Error(err),
		)
		panic(fmt.Sprintf("❌ Client initialization failed: %v", err))
	}

	if err := client.Run(); err != nil {
		client.Logger.Fatal("Client stopped unexpectedly",
			zap.String("stage", "runtime"),
			zap.Error(err),
		)
	}

	client.Logger.Info("🛑 Client gracefully stopped")
}
