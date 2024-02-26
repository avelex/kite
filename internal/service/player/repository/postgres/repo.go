package postgres

import (
	"context"
	"time"

	opendota "github.com/avelex/kite/internal/adapters/webapi/open_dota"
	"github.com/avelex/kite/internal/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB

	openDotaAPI       opendota.API
	nicknameAccountID map[string]int64
}

func New(db *gorm.DB, open opendota.API) *repository {
	return &repository{
		db:                db,
		openDotaAPI:       open,
		nicknameAccountID: make(map[string]int64),
	}
}

func (r *repository) PrepareData(ctx context.Context) error {
	ctxPlayer, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	players, err := r.openDotaAPI.ProPlayers(ctxPlayer)
	if err != nil {
		return err
	}

	for _, pp := range players {
		r.nicknameAccountID[pp.Name] = pp.AccountID
	}

	return nil
}

func (r *repository) AccountIDByNickname(_ context.Context, nickname string) (int64, error) {
	return r.nicknameAccountID[nickname], nil
}

// ID           uint
// AccountID    int64 `gorm:"index"`
// MatchID      int64
// PatchVersion int64
// Time         int64
// X, Y         uint8
// Side         string
// Type         WardType

func (r *repository) PlayerWards(ctx context.Context, accountID, patch int64, sides, wardsType []string, times [2]int16) ([]entity.PlayerWard, error) {
	var out []entity.PlayerWard

	stmt := r.db.WithContext(ctx).Where("account_id = ?", accountID)

	if patch >= 0 {
		stmt = stmt.Where("patch_version = ?", patch)
	}

	if len(sides) == 1 {
		stmt = stmt.Where("side = ?", sides[0])
	}

	if len(wardsType) == 1 {
		stmt = stmt.Where("type = ?", wardsType[0])
	}

	if err := stmt.Where("time BETWEEN ? AND ?", times[0], times[1]).Find(&out).Error; err != nil {
		return nil, err
	}

	return out, nil
}
