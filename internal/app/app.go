package app

import (
	"context"
	"time"

	"github.com/avelex/kite/config"
	opendota "github.com/avelex/kite/internal/adapters/webapi/open_dota"
	async_proc "github.com/avelex/kite/internal/controllers/asynq"
	"github.com/avelex/kite/internal/controllers/asynq/tasks"
	http_v1 "github.com/avelex/kite/internal/controllers/http/v1"
	"github.com/avelex/kite/internal/service/player"
	wardcollector "github.com/avelex/kite/internal/service/ward_collector"
	"github.com/avelex/kite/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hibiken/asynq"
)

func Run(ctx context.Context, cfg config.Config) error {
	logger := logger.LoggerFromContext(ctx)

	openDotaAPI := opendota.New(cfg.OpenDotaAPI.Key)

	// Registering Services
	wardCollectorService := wardcollector.NewService(nil, openDotaAPI)
	playerService := player.NewService(nil)

	// Registering Controllers
	v1 := http_v1.New(playerService)
	asynqProcessor := async_proc.NewProcessor(wardCollectorService)

	router := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	api := router.Group("/api")
	api.Use(recover.New())
	api.Use(cors.New())

	v1.RegisterRoutes(api)

	asynqSrv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: cfg.Redis.Addr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency:     10,
			ShutdownTimeout: 10 * time.Second,
		},
	)

	mux := asynq.NewServeMux()
	asynqProcessor.Register(mux)

	logger.Infof("Start HTTP server listening on port :%s", cfg.HTTP.Port)
	go func() {
		if err := router.Listen(cfg.HTTP.Host + ":" + cfg.HTTP.Port); err != nil {
			logger.Error(err)
		}
	}()

	logger.Info("Start Asynq server")
	go func() {
		if err := asynqSrv.Run(mux); err != nil {
			logger.Error(err)
		}
	}()

	time.Sleep(time.Second)
	go startBagroundTasks(cfg)

	<-ctx.Done()

	if err := router.ShutdownWithTimeout(5 * time.Second); err != nil {
		return err
	}

	asynqSrv.Shutdown()

	return nil
}

func startBagroundTasks(cfg config.Config) error {
	logger := logger.LoggerFromContext(context.Background())

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: cfg.Redis.Addr})
	defer client.Close()

	task, err := tasks.NewWardCollectTask()
	if err != nil {
		return err
	}

	info, err := client.Enqueue(task, asynq.ProcessIn(1*time.Hour))
	if err != nil {
		logger.Errorf("failed to enqueue task: %w", err)
		return err
	}

	logger.Infof("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	return nil
}
