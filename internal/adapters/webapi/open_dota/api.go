package opendota

import (
	"context"
	"slices"

	"github.com/avelex/kite/internal/entity"
	"github.com/carlmjohnson/requests"
)

const _API_BASE = "https://api.opendota.com"
const _API_KEY = "api_key"

const (
	_PLAYER_INFO_PATH    = "/api/players/%d"
	_PLAYER_MATCHES_PATH = "/api/players/%d/matches"
	_MATCHES_INFO_PATH   = "/api/matches/%d"
	_PRO_PLAYERS_PATH    = "/api/proPlayers"
	_PATCH_PATH          = "/api/constants/patch"
)

type API interface {
	PlayerAllMatches(ctx context.Context, accountID int64) ([]entity.PlayerMatchOverview, error)
	ProPlayers(ctx context.Context) ([]entity.ProPlayer, error)
	Patches(ctx context.Context) ([]entity.Patch, error)
	Match(ctx context.Context, matchID int64) (entity.FullMatch, error)
}

type api struct {
	key string
}

func New(key string) *api {
	return &api{
		key: key,
	}
}

func (a *api) PlayerAllMatches(ctx context.Context, accountID int64) ([]entity.PlayerMatchOverview, error) {
	var matches []entity.PlayerMatchOverview

	err := requests.URL(_API_BASE).
		Pathf(_PLAYER_MATCHES_PATH, accountID).
		Param(_API_KEY, a.key).
		ToJSON(&matches).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	slices.Reverse(matches)

	return matches, nil
}

func (a *api) ProPlayers(ctx context.Context) ([]entity.ProPlayer, error) {
	var list []entity.ProPlayer

	err := requests.URL(_API_BASE).
		Path(_PRO_PLAYERS_PATH).
		Param(_API_KEY, a.key).
		ToJSON(&list).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a *api) Match(ctx context.Context, matchID int64) (entity.FullMatch, error) {
	var m entity.FullMatch

	err := requests.URL(_API_BASE).
		Pathf(_MATCHES_INFO_PATH, matchID).
		Param(_API_KEY, a.key).
		ToJSON(&m).
		Fetch(ctx)
	if err != nil {
		return entity.FullMatch{}, err
	}

	return m, nil
}

func (a *api) Patches(ctx context.Context) ([]entity.Patch, error) {
	var list []entity.Patch

	err := requests.URL(_API_BASE).
		Path(_PATCH_PATH).
		Param(_API_KEY, a.key).
		ToJSON(&list).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}
