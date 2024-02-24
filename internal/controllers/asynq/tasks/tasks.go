package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	WardCollect       = "ward:collect"
	PlayerWardCollect = "player_ward:collect"
)

func NewWardCollectTask() (*asynq.Task, error) {
	return asynq.NewTask(WardCollect, nil), nil
}

type PlayerWardPayload struct {
	AccountID int64
}

func NewPlayerWardCollectTask(accountID int64) (*asynq.Task, error) {
	payload, err := json.Marshal(PlayerWardPayload{AccountID: accountID})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(PlayerWardCollect, payload), nil
}
