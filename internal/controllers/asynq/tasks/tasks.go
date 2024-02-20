package tasks

import (
	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	WardCollect = "ward:collect"
)

func NewWardCollectTask() (*asynq.Task, error) {
	return asynq.NewTask(WardCollect, nil), nil
}
