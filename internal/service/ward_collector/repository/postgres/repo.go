package postgres

import (
	"context"

	"github.com/avelex/kite/internal/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) SaveWards(ctx context.Context, wards []entity.PlayerWard) error {
	if len(wards) == 0 {
		return nil
	}

	reuslt := r.db.WithContext(ctx).Create(&wards)
	if reuslt.Error != nil {
		return reuslt.Error
	}

	return nil
}

func (r *repository) PlayerLastSavedMatchID(ctx context.Context, accountID int64) int64 {
	p := entity.PlayerWard{
		AccountID: accountID,
	}

	r.db.WithContext(ctx).Order("match_id").Last(&p)

	return p.MatchID
}
