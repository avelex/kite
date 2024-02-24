package asynq

import (
	"context"
	"encoding/json"
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
	m.Handle(tasks.PlayerWardCollect, p)
}

func (p *processor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	logger := logger.LoggerFromContext(ctx)

	var in tasks.PlayerWardPayload
	if err := json.Unmarshal(t.Payload(), &in); err != nil {
		return err
	}

	logger.Info("Start processing collect wards task")
	if err := p.wc.CollectPlayer(ctx, in.AccountID); err != nil {
		return fmt.Errorf("failed to collect wards: %v %w", err, asynq.SkipRetry)
	}

	logger.Info("Processing ward collect done!")

	return nil
}
