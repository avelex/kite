package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/avelex/kite/config"
	"github.com/avelex/kite/internal/app"
	"github.com/avelex/kite/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := config.InitConfig()
	logger := logger.LoggerFromContext(ctx)

	logger.Info("App run")
	if err := app.Run(ctx, cfg); err != nil {
		logger.Error(err)
	}
	logger.Info("App done!")
}
