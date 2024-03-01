package tasks

import (
	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	WardCollect = "ward:collect"
	Status      = "status:check"
)

func NewWardCollectTask() (*asynq.Task, error) {
	return asynq.NewTask(WardCollect, nil), nil
}

func NewStatusTask() (*asynq.Task, error) {
	return asynq.NewTask(Status, nil), nil
}
