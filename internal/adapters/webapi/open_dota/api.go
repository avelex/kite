package opendota

import (
	"context"

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
)

type API struct {
	key string
}

func New(key string) *API {
	return &API{
		key: key,
	}
}

func (a *API) PlayerAllMatches(ctx context.Context, accountID int64) ([]entity.PlayerMatchOverview, error) {
	var matches []entity.PlayerMatchOverview

	err := requests.URL(_API_BASE).
		Pathf(_PLAYER_MATCHES_PATH, accountID).
		Param(_API_KEY, a.key).
		ToJSON(&matches).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return matches, nil
}

func (a *API) ProPlayers(ctx context.Context) ([]int64, error) {
	type proPlayer struct {
		AccountID int64 `json:"account_id"`
	}

	var p []proPlayer

	err := requests.URL(_API_BASE).
		Path(_PRO_PLAYERS_PATH).
		Param(_API_KEY, a.key).
		ToJSON(&p).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]int64, 0, len(p))
	for _, pp := range p {
		out = append(out, pp.AccountID)
	}

	return out, nil
}

func (a *API) Match(ctx context.Context, matchID int64) (entity.FullMatch, error) {
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
