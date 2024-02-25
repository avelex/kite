package asynq

import (
	"context"
	"fmt"

	"github.com/avelex/kite/internal/controllers/asynq/tasks"
	"github.com/avelex/kite/logger"
	"github.com/hibiken/asynq"
)

type WardCollectorUsecases interface {
	Collect(ctx context.Context) error
	CollectPlayer(ctx context.Context, accountID int64) error
}

type processor struct {
	wc WardCollectorUsecases
}

func NewProcessor(wc WardCollectorUsecases) *processor {
	return &processor{
		wc: wc,
	}
}

func (p *processor) Register(m *asynq.ServeMux) {
	m.Handle(tasks.WardCollect, p)
}

func (p *processor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	logger := logger.LoggerFromContext(ctx)

	logger.Info("Start processing collect wards task")
	if err := p.wc.Collect(ctx); err != nil {
		return fmt.Errorf("failed to collect wards: %v %w", err, asynq.SkipRetry)
	}

	logger.Info("Processing ward collect done!")

	return nil
}
