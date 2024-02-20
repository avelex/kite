package wardcollector

import (
	"context"

	opendota "github.com/avelex/kite/internal/adapters/webapi/open_dota"
	"github.com/avelex/kite/internal/entity"
)

// Работает с OpenDota API и БД
// 1) Запись агрегированных данных о вардах в бд
// 2) Сбор данных из OpenDota
type Service interface {
	Collect(ctx context.Context) error
}

type Repository interface {
	SaveWards(ctx context.Context, wards []entity.PlayerWard) error
	PlayerLastSavedMatchID(ctx context.Context, accountID int64) int64
}

type service struct {
	repo        Repository
	openDotaAPI *opendota.API
}

func NewService(r Repository, api *opendota.API) *service {
	return &service{
		repo:        r,
		openDotaAPI: api,
	}
}

// 1) Получить список про игроков
// 2) По каждому игроку найти его матчи, которых еше нет в репо
// 3) По каждому новому матчу собрать информацию о вардах
// 4) С агрегировать данные
// 5) Записать в бд
func (s *service) Collect(ctx context.Context) error {
	proPlayers, err := s.openDotaAPI.ProPlayers(ctx)
	if err != nil {
		return err
	}

	for _, accountID := range proPlayers {
		if err := s.collectPlayer(ctx, accountID); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) CollectPlayer(ctx context.Context, accountID int64) error {
	return s.collectPlayer(ctx, accountID)
}

func (s *service) collectPlayer(ctx context.Context, accountID int64) error {
	lastMatchID := s.repo.PlayerLastSavedMatchID(ctx, accountID)

	matches, err := s.openDotaAPI.PlayerAllMatches(ctx, accountID)
	if err != nil {
		return err
	}

	newMatches := make([]int64, 0, len(matches))

	for _, v := range matches {
		if v.MatchID > lastMatchID {
			newMatches = append(newMatches, v.MatchID)
		}
	}

	wards := make([]entity.PlayerWard, 0, len(newMatches))

	for _, matchID := range newMatches {
		match, err := s.openDotaAPI.Match(ctx, matchID)
		if err != nil {
			continue
		}

		playerStats := match.Player(accountID)

		for _, w := range playerStats.Obs {
			wards = append(wards, entity.PlayerWard{
				AccountID:    accountID,
				MatchID:      matchID,
				PatchVersion: int64(match.Patch),
				Time:         w.Time,
				X:            w.X,
				Y:            w.Y,
				IsRadiant:    w.IsRadiant,
				Type:         entity.Observer,
			})
		}

		for _, w := range playerStats.Sen {
			wards = append(wards, entity.PlayerWard{
				AccountID:    accountID,
				MatchID:      matchID,
				PatchVersion: int64(match.Patch),
				Time:         w.Time,
				X:            w.X,
				Y:            w.Y,
				IsRadiant:    w.IsRadiant,
				Type:         entity.Sentry,
			})
		}
	}

	if err := s.repo.SaveWards(ctx, wards); err != nil {
		return err
	}

	return nil
}
