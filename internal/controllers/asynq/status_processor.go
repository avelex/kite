package asynq

import (
	"context"

	"github.com/avelex/kite/internal/controllers/asynq/tasks"
	"github.com/avelex/kite/logger"
	"github.com/hibiken/asynq"
)

type statusProcessor struct {
}

func NewStatusProcessor() *statusProcessor {
	return &statusProcessor{}
}

func (p *statusProcessor) Register(m *asynq.ServeMux) {
	m.Handle(tasks.Status, p)
}

func (p *statusProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	logger := logger.LoggerFromContext(ctx)
	logger.Info("Start processing status  task")

	return nil
}
