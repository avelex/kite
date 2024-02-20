package player

import (
	"context"
	"fmt"
	"math"

	"github.com/avelex/kite/internal/entity"
)

type PlayerRepository interface {
	AccountIDByNickname(ctx context.Context, nickname string) (int64, error)
	AllPlayerWards(ctx context.Context, accountID int64) ([]entity.PlayerWard, error)
}

// TODO: добавить KV ник-аккаунт
type service struct {
	repo PlayerRepository
}

func NewService(pr PlayerRepository) *service {
	return &service{
		repo: pr,
	}
}

func (s *service) PlayerWards(ctx context.Context, nickname string) (entity.PlayerWardsView, error) {
	accID, err := s.repo.AccountIDByNickname(ctx, nickname)
	if err != nil {
		return entity.PlayerWardsView{}, err
	}

	wards, err := s.repo.AllPlayerWards(ctx, accID)
	if err != nil {
		return entity.PlayerWardsView{}, err
	}

	radiant, dire := calculateWardsFreqBySides(wards)

	return entity.PlayerWardsView{
		Nickname:  nickname,
		AccountID: accID,
		Radiant:   radiant,
		Dire:      dire,
	}, nil
}

func calculateWardsFreqBySides(wards []entity.PlayerWard) (map[int][]entity.Ward, map[int][]entity.Ward) {
	frequencies := make(map[string][]entity.Ward)

	for _, w := range wards {
		roundedX := roundToNearestMultiple(float64(w.X), 5)
		roundedY := roundToNearestMultiple(float64(w.Y), 5)
		key := fmt.Sprintf("[%d %d]", roundedX, roundedY)
		frequencies[key] = append(frequencies[key], entity.Ward{
			Time:      w.Time,
			Type:      string(w.Type),
			X:         roundedX,
			Y:         roundedY,
			IsRadiant: w.IsRadiant,
		})
	}

	radiant := make(map[int][]entity.Ward)
	dire := make(map[int][]entity.Ward)

	for _, v := range frequencies {
		ward := v[0]
		hit := float64(len(v)) / float64(len(wards)) * 100
		roundedHit := roundToNearestMultiple(hit, 10)

		if ward.IsRadiant {
			radiant[roundedHit] = append(radiant[roundedHit], v...)
		} else {
			dire[roundedHit] = append(dire[roundedHit], v...)
		}
	}

	return radiant, dire
}

func roundToNearestMultiple(num, base float64) int {
	return int(math.Round(num/base) * base)
}
