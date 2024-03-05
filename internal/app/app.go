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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.DatabaseName,
		cfg.Postgres.Port,
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
	wardProcessor := async_proc.NewWardProcessor(wardCollectorService)
	statusProcessor := async_proc.NewStatusProcessor()

	router := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	router.Static("/", cfg.FrontendDir)

	api := router.Group("/api")
	api.Use(recover.New())
	api.Use(cors.New())

	v1.RegisterRoutes(api)

	asynqSrv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: cfg.Redis.Addr},
		asynq.Config{
			ShutdownTimeout: 5 * time.Second,
		},
	)

	mux := asynq.NewServeMux()
	wardProcessor.Register(mux)
	statusProcessor.Register(mux)

	asnyqScheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: cfg.Redis.Addr},
		&asynq.SchedulerOpts{},
	)

	if err := scheduleTask(asnyqScheduler); err != nil {
		return err
	}

	logger.Infof("Start HTTP server listening on port :%s", cfg.HTTP.Port)
	go func() {
		if err := router.Listen(cfg.HTTP.Host + ":" + cfg.HTTP.Port); err != nil {
			logger.Error(err)
		}
	}()

	logger.Info("Start Asynq server")
	go func() {
		if err := asynqSrv.Start(mux); err != nil {
			logger.Error(err)
		}
	}()

	logger.Info("Start Asynq scheduler")
	go func() {
		if err := asnyqScheduler.Start(); err != nil {
			logger.Error(err)
		}
	}()

	<-ctx.Done()

	if err := router.ShutdownWithTimeout(5 * time.Second); err != nil {
		return err
	}

	asynqSrv.Shutdown()
	asnyqScheduler.Shutdown()

	return nil
}

func scheduleTask(s *asynq.Scheduler) error {
	task, err := tasks.NewWardCollectTask()
	if err != nil {
		return err
	}

	if _, err := s.Register("0 0 */3 * *", task, asynq.Timeout(12*time.Hour)); err != nil {
		return err
	}

	return nil
}

func migration(db *gorm.DB) error {
	if err := db.AutoMigrate(&entity.PlayerWard{}); err != nil {
		return err
	}

	return nil
}
