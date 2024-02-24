package player

import (
	"context"
	"fmt"
	"math"

	"github.com/avelex/kite/internal/entity"
)

type PlayerRepository interface {
	AccountIDByNickname(ctx context.Context, nickname string) (int64, error)
	PlayerWards(ctx context.Context, accountID, patch int64, sides, wardsType []string, times [2]uint16) ([]entity.PlayerWard, error)
}

type PatchService interface {
	LatestPatchVersion(ctx context.Context) int64
	PatchVersionFromString(ctx context.Context, patch string) int64
}

type service struct {
	repo     PlayerRepository
	patchSvc PatchService
}

func NewService(pr PlayerRepository, patchSvc PatchService) *service {
	return &service{
		repo:     pr,
		patchSvc: patchSvc,
	}
}

func (s *service) PlayerWards(ctx context.Context, nickname, patch string, sides entity.Sides, wardsType entity.Wards, times []uint16) (entity.PlayerWardsView, error) {
	accID, err := s.repo.AccountIDByNickname(ctx, nickname)
	if err != nil {
		return entity.PlayerWardsView{}, err
	}

	var patchInt int64
	if patch == "latest" {
		patchInt = s.patchSvc.LatestPatchVersion(ctx)
	} else {
		patchInt = s.patchSvc.PatchVersionFromString(ctx, patch)
	}

	wards, err := s.repo.PlayerWards(ctx, accID, patchInt, sides.Slice(), wardsType.Slice(), [2]uint16(times))
	if err != nil {
		return entity.PlayerWardsView{}, err
	}

	view := entity.PlayerWardsView{
		Profile: entity.Profile{
			AccountID: accID,
			Nickname:  nickname,
		},
	}

	if sides.HasDire() {
		obs := make([]entity.WardView, 0)
		if wardsType.HasObserver() {
			obs = append(obs, calculateWardsFreqBySides(wards, entity.DireSide, entity.Observer)...)
		}

		sen := make([]entity.WardView, 0)
		if wardsType.HasSentry() {
			sen = append(sen, calculateWardsFreqBySides(wards, entity.DireSide, entity.Sentry)...)
		}

		view.Dire = &entity.Side{
			Obs: obs,
			Sen: sen,
		}
	}

	if sides.HasRadiant() {
		obs := make([]entity.WardView, 0)
		if wardsType.HasObserver() {
			obs = append(obs, calculateWardsFreqBySides(wards, entity.RadiantSide, entity.Observer)...)
		}

		sen := make([]entity.WardView, 0)
		if wardsType.HasSentry() {
			sen = append(sen, calculateWardsFreqBySides(wards, entity.RadiantSide, entity.Sentry)...)
		}

		view.Radiant = &entity.Side{
			Obs: obs,
			Sen: sen,
		}
	}

	return view, nil
}

func calculateWardsFreqBySides(wards []entity.PlayerWard, side string, wardType entity.WardType) []entity.WardView {
	frequencies := make(map[string][]entity.WardView)
	count := 0

	for _, w := range wards {
		if w.Side != side || w.Type != wardType {
			continue
		}

		roundedX := roundToNearestMultiple(float64(w.X), 5)
		roundedY := roundToNearestMultiple(float64(w.Y), 5)

		key := fmt.Sprintf("[%d %d]", roundedX, roundedY)

		frequencies[key] = append(frequencies[key], entity.WardView{
			X: uint8(roundedX),
			Y: uint8(roundedY),
		})

		count++
	}

	out := make([]entity.WardView, 0)

	for _, v := range frequencies {
		ward := v[0]
		rate := float64(len(v)) / float64(count) * 10000
		out = append(out, entity.WardView{
			X:    ward.X,
			Y:    ward.Y,
			Rate: uint16(rate),
		})
	}

	return out
}

func roundToNearestMultiple(num, base float64) int {
	return int(math.Round(num/base) * base)
}
