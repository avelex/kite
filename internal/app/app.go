package app

import (
	"context"
	"fmt"
	"time"

	"github.com/avelex/kite/config"
	opendota "github.com/avelex/kite/internal/adapters/webapi/open_dota"
	async_proc "github.com/avelex/kite/internal/controllers/asynq"
	"github.com/avelex/kite/internal/controllers/asynq/tasks"
	http_v1 "github.com/avelex/kite/internal/controllers/http/v1"
	"github.com/avelex/kite/internal/entity"
	"github.com/avelex/kite/internal/service/patch"
	patchMemRepo "github.com/avelex/kite/internal/service/patch/repository/memory"
	"github.com/avelex/kite/internal/service/player"
	playerRepo "github.com/avelex/kite/internal/service/player/repository/postgres"
	wardcollector "github.com/avelex/kite/internal/service/ward_collector"
	wardsRepo "github.com/avelex/kite/internal/service/ward_collector/repository/postgres"
	"github.com/avelex/kite/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

func Run(ctx context.Context, cfg config.Config) error {
	logger := logger.LoggerFromContext(ctx)

	// Init connections to db
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DatabaseName,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return err
	}

	if err := migration(db); err != nil {
		return err
	}

	openDotaAPI := opendota.New(cfg.OpenDotaAPI.Key)

	// Registerging repos
	wardsRepo := wardsRepo.New(db)
	playerRepo := playerRepo.New(db, openDotaAPI)
	patchRepo := patchMemRepo.New(openDotaAPI)

	if err := patchRepo.PrepareData(ctx); err != nil {
		return err
	}

	if err := playerRepo.PrepareData(ctx); err != nil {
		return err
	}

	// Registering Services
	wardCollectorService := wardcollector.NewService(wardsRepo, openDotaAPI)
	patchService := patch.NewService(patchRepo)
	playerService := player.NewService(playerRepo, patchService)

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
			ShutdownTimeout: 5 * time.Second,
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

func startBagroundTasks(cfg config.Config) {
	logger := logger.LoggerFromContext(context.Background())

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: cfg.Redis.Addr})
	defer client.Close()

	task, err := tasks.NewPlayerWardCollectTask(16497807)
	if err != nil {
		return
	}

	info, err := client.Enqueue(task, asynq.Timeout(24*time.Hour))
	if err != nil {
		logger.Errorf("failed to enqueue task: %w", err)
		return
	}

	logger.Infof("Enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

func migration(db *gorm.DB) error {
	if err := db.AutoMigrate(&entity.PlayerWard{}); err != nil {
		return err
	}

	return nil
}
