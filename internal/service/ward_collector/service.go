package wardcollector

import (
	"context"
	"time"

	opendota "github.com/avelex/kite/internal/adapters/webapi/open_dota"
	"github.com/avelex/kite/internal/entity"
	"github.com/avelex/kite/logger"
)

type Repository interface {
	SaveWards(ctx context.Context, wards []entity.PlayerWard) error
	PlayerLastSavedMatchID(ctx context.Context, accountID int64) int64
}

type service struct {
	repo        Repository
	openDotaAPI opendota.API
}

func NewService(r Repository, api opendota.API) *service {
	return &service{
		repo:        r,
		openDotaAPI: api,
	}
}

func (s *service) Collect(ctx context.Context) error {
	proPlayers, err := s.openDotaAPI.ProPlayers(ctx)
	if err != nil {
		return err
	}

	for _, p := range proPlayers {
		if err := s.collectPlayer(ctx, p.AccountID); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) CollectPlayer(ctx context.Context, accountID int64) error {
	return s.collectPlayer(ctx, accountID)
}

func (s *service) collectPlayer(ctx context.Context, accountID int64) error {
	logger := logger.LoggerFromContext(ctx).With("account", accountID)

	lastMatchID := s.repo.PlayerLastSavedMatchID(ctx, accountID)
	logger.Infof("Last match id = %d", lastMatchID)

	matches, err := s.openDotaAPI.PlayerAllMatches(ctx, accountID)
	if err != nil {
		logger.Errorf("failed to get all matches: %w", err)
		return err
	}

	logger.Infof("Matches count = %d", len(matches))

	newMatches := make([]int64, 0, len(matches))

	for _, v := range matches {
		if v.MatchID > lastMatchID {
			newMatches = append(newMatches, v.MatchID)
		}
	}

	for _, matchID := range newMatches {
		time.Sleep(100 * time.Millisecond)

		ctxMatch, cancel := context.WithTimeout(ctx, 3*time.Second)

		match, err := s.openDotaAPI.Match(ctxMatch, matchID)
		cancel()

		if err != nil {
			logger.Errorf("failed to get match=%s: %w", matchID, err)
			continue
		}

		logger.Infof("Fetched info about match = %d", matchID)

		wards := make([]entity.PlayerWard, 0, len(newMatches))
		playerStats := match.Player(accountID)

		side := entity.DireSide
		if playerStats.IsRadiant {
			side = entity.RadiantSide
		}

		for _, w := range playerStats.Obs {
			wards = append(wards, entity.PlayerWard{
				AccountID:    accountID,
				MatchID:      matchID,
				PatchVersion: int64(match.Patch),
				Time:         w.Time,
				X:            w.X,
				Y:            w.Y,
				Type:         entity.Observer,
				Side:         side,
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
				Type:         entity.Sentry,
				Side:         side,
			})
		}

		if len(wards) == 0 {
			logger.Info("No wards")
			continue
		}

		if err := s.repo.SaveWards(ctx, wards); err != nil {
			logger.Errorf("failed to save wards: %w", err)
			return err
		}

		logger.Infof("Saved player wards = %d", len(wards))
	}

	return nil
}
