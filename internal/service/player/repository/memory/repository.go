package memory

import (
	"context"
	"fmt"

	"github.com/avelex/kite/internal/entity"
)

type playerRepository struct {
	nicknames map[string]int64
}

func New() *playerRepository {
	return &playerRepository{
		nicknames: map[string]int64{
			"tOfu": 16497807,
		},
	}
}

func (r *playerRepository) AccountIDByNickname(ctx context.Context, nickname string) (int64, error) {
	id, ok := r.nicknames[nickname]
	if !ok {
		return 0, fmt.Errorf("not found")
	}

	return id, nil
}

func (r *playerRepository) AllPlayerWards(ctx context.Context, accountID int64) ([]entity.PlayerWard, error) {
	return nil, nil
}
