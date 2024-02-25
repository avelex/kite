package mock

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/avelex/kite/internal/entity"
)

type api struct {
}

func New() *api {
	return &api{}
}

func (a *api) PlayerAllMatches(_ context.Context, accountID int64) ([]entity.PlayerMatchOverview, error) {
	var out []entity.PlayerMatchOverview

	if err := json.NewDecoder(strings.NewReader(matches)).Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}

func (a *api) ProPlayers(_ context.Context) ([]entity.ProPlayer, error) {
	var out []entity.ProPlayer

	if err := json.NewDecoder(strings.NewReader(proPlayers)).Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}

func (a *api) Patches(_ context.Context) ([]entity.Patch, error) {
	var out []entity.Patch

	if err := json.NewDecoder(strings.NewReader(patches)).Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}

func (a *api) Match(_ context.Context, matchID int64) (entity.FullMatch, error) {
	var out entity.FullMatch

	if err := json.NewDecoder(strings.NewReader(match)).Decode(&out); err != nil {
		return entity.FullMatch{}, err
	}

	return out, nil
}

var proPlayers = `[{
	"account_id": 16497807,
	"steamid": "76561197976763535",
	"avatar": "https://avatars.steamstatic.com/a0aca11d96d24ee6796bc8017fe7d988ac69006d.jpg",
	"avatarmedium": "https://avatars.steamstatic.com/a0aca11d96d24ee6796bc8017fe7d988ac69006d_medium.jpg",
	"avatarfull": "https://avatars.steamstatic.com/a0aca11d96d24ee6796bc8017fe7d988ac69006d_full.jpg",
	"profileurl": "https://steamcommunity.com/id/to_Ofu/",
	"personaname": "tOfu",
	"last_login": "2020-12-05T02:26:53.827Z",
	"full_history_time": "2024-02-01T13:54:35.779Z",
	"cheese": 0,
	"fh_unavailable": true,
	"loccountrycode": "DE",
	"last_match_time": "2024-02-14T12:20:21.000Z",
	"plus": true,
	"name": "tOfu",
	"country_code": "de",
	"fantasy_role": 2,
	"team_id": 8599101,
	"team_name": "Gaimin Gladiators",
	"team_tag": "GG",
	"is_locked": true,
	"is_pro": true,
	"locked_until": null
	}
]`
var patches string = `[{"name":"6.70","date":"2010-12-24T00:00:00Z","id":0},{"name":"6.71","date":"2011-01-21T00:00:00Z","id":1},{"name":"6.72","date":"2011-04-27T00:00:00Z","id":2},{"name":"6.73","date":"2011-12-24T00:00:00Z","id":3},{"name":"6.74","date":"2012-03-10T00:00:00Z","id":4},{"name":"6.75","date":"2012-09-30T00:00:00Z","id":5},{"name":"6.76","date":"2012-10-21T00:00:00Z","id":6},{"name":"6.77","date":"2012-12-15T00:00:00Z","id":7},{"name":"6.78","date":"2013-05-30T00:00:00Z","id":8},{"name":"6.79","date":"2013-11-24T00:00:00Z","id":9},{"name":"6.80","date":"2014-01-27T00:00:00Z","id":10},{"name":"6.81","date":"2014-04-29T00:00:00Z","id":11},{"name":"6.82","date":"2014-09-24T00:00:00Z","id":12},{"name":"6.83","date":"2014-12-17T00:00:00Z","id":13},{"name":"6.84","date":"2015-04-30T21:00:00Z","id":14},{"name":"6.85","date":"2015-09-24T20:00:00Z","id":15},{"name":"6.86","date":"2015-12-16T20:00:00Z","id":16},{"name":"6.87","date":"2016-04-26T01:00:00Z","id":17},{"name":"6.88","date":"2016-06-12T08:00:00Z","id":18},{"name":"7.00","date":"2016-12-13T00:00:00Z","id":19},{"name":"7.01","date":"2016-12-21T03:00:00Z","id":20},{"name":"7.02","date":"2017-02-09T04:00:00Z","id":21},{"name":"7.03","date":"2017-03-16T00:00:00Z","id":22},{"name":"7.04","date":"2017-03-23T18:00:00Z","id":23},{"name":"7.05","date":"2017-04-09T22:00:00Z","id":24},{"name":"7.06","date":"2017-05-15T15:00:00Z","id":25},{"name":"7.07","date":"2017-10-31T23:00:00Z","id":26},{"name":"7.08","date":"2018-02-01T00:00:00Z","id":27},{"name":"7.09","date":"2018-02-15T00:00:00.000Z","id":28},{"name":"7.10","date":"2018-03-01T00:00:00.000Z","id":29},{"name":"7.11","date":"2018-03-15T00:00:00.000Z","id":30},{"name":"7.12","date":"2018-03-29T00:00:00.000Z","id":31},{"name":"7.13","date":"2018-04-12T00:00:00.000Z","id":32},{"name":"7.14","date":"2018-04-26T00:00:00.000Z","id":33},{"name":"7.15","date":"2018-05-10T00:00:00.000Z","id":34},{"name":"7.16","date":"2018-05-27T00:00:00.000Z","id":35},{"name":"7.17","date":"2018-06-10T00:00:00.000Z","id":36},{"name":"7.18","date":"2018-06-24T00:00:00.000Z","id":37},{"name":"7.19","date":"2018-07-30T00:00:00.000Z","id":38},{"name":"7.20","date":"2018-11-19T18:00:00.000Z","id":39},{"name":"7.21","date":"2019-01-29T18:00:00.000Z","id":40},{"name":"7.22","date":"2019-05-25T00:00:00.000Z","id":41},{"name":"7.23","date":"2019-11-26T18:00:00.000Z","id":42},{"name":"7.24","date":"2020-01-27T00:00:00.000Z","id":43},{"name":"7.25","date":"2020-03-17T18:00:00.000Z","id":44},{"name":"7.26","date":"2020-04-18T00:00:00.000Z","id":45},{"name":"7.27","date":"2020-06-29T00:00:00.000Z","id":46},{"name":"7.28","date":"2020-12-18T00:00:00.000Z","id":47},{"name":"7.29","date":"2021-04-10T00:00:00.000Z","id":48},{"name":"7.30","date":"2021-08-18T02:53:21.000Z","id":49},{"name":"7.31","date":"2022-02-23T23:46:14.907Z","id":50},{"name":"7.32","date":"2022-08-24T02:16:32.682Z","id":51},{"name":"7.33","date":"2023-04-21T01:22:56.804Z","id":52},{"name":"7.34","date":"2023-08-09T00:11:15.281Z","id":53},{"name":"7.35","date":"2023-12-14T16:07:43.429Z","id":54}]`
var matches = `[
	{
	"match_id": 7586292999,
	"player_slot": 131,
	"radiant_win": true,
	"duration": 3007,
	"game_mode": 2,
	"lobby_type": 1,
	"hero_id": 53,
	"start_time": 1707913221,
	"version": 21,
	"kills": 1,
	"deaths": 8,
	"assists": 18,
	"average_rank": 81,
	"leaver_status": 0,
	"party_size": 10
	}
]`

//go:nolint
var match = `{
    "version": 21,
    "match_id": 7586292999,
    "players": [
        {
            "obs_log": [
                {
                    "time": 334,
                    "type": "obs_log",
                    "key": "[104, 140]",
                    "slot": 8,
                    "x": 104,
                    "y": 140,
                    "z": 128,
                    "entityleft": false,
                    "ehandle": 10275063,
                    "player_slot": 131
                },
                {
                    "time": 661,
                    "type": "obs_log",
                    "key": "[144, 120]",
                    "slot": 8,
                    "x": 144,
                    "y": 120,
                    "z": 132,
                    "entityleft": false,
                    "ehandle": 4606446,
                    "player_slot": 131
                },
                {
                    "time": 898,
                    "type": "obs_log",
                    "key": "[110, 162]",
                    "slot": 8,
                    "x": 110,
                    "y": 162,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 15665001,
                    "player_slot": 131
                },
                {
                    "time": 1064,
                    "type": "obs_log",
                    "key": "[146, 88]",
                    "slot": 8,
                    "x": 146,
                    "y": 88,
                    "z": 128,
                    "entityleft": false,
                    "ehandle": 14894744,
                    "player_slot": 131
                },
                {
                    "time": 1600,
                    "type": "obs_log",
                    "key": "[114, 180]",
                    "slot": 8,
                    "x": 114,
                    "y": 180,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 2065949,
                    "player_slot": 131
                },
                {
                    "time": 1796,
                    "type": "obs_log",
                    "key": "[146, 70]",
                    "slot": 8,
                    "x": 146,
                    "y": 70,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 706413,
                    "player_slot": 131
                },
                {
                    "time": 2013,
                    "type": "obs_log",
                    "key": "[146, 70]",
                    "slot": 8,
                    "x": 146,
                    "y": 70,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 15566869,
                    "player_slot": 131
                },
                {
                    "time": 2267,
                    "type": "obs_log",
                    "key": "[158, 90]",
                    "slot": 8,
                    "x": 158,
                    "y": 90,
                    "z": 132,
                    "entityleft": false,
                    "ehandle": 9865790,
                    "player_slot": 131
                },
                {
                    "time": 2394,
                    "type": "obs_log",
                    "key": "[144, 120]",
                    "slot": 8,
                    "x": 144,
                    "y": 120,
                    "z": 132,
                    "entityleft": false,
                    "ehandle": 5736991,
                    "player_slot": 131
                },
                {
                    "time": 2566,
                    "type": "obs_log",
                    "key": "[146, 70]",
                    "slot": 8,
                    "x": 146,
                    "y": 70,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 3590375,
                    "player_slot": 131
                },
                {
                    "time": 2849,
                    "type": "obs_log",
                    "key": "[186, 116]",
                    "slot": 8,
                    "x": 186,
                    "y": 116,
                    "z": 132,
                    "entityleft": false,
                    "ehandle": 11815796,
                    "player_slot": 131
                }
            ],
            "sen_log": [
                {
                    "time": 30,
                    "type": "sen_log",
                    "key": "[160, 88]",
                    "slot": 8,
                    "x": 160,
                    "y": 88,
                    "z": 128,
                    "entityleft": false,
                    "ehandle": 9520684,
                    "player_slot": 131
                },
                {
                    "time": 353,
                    "type": "sen_log",
                    "key": "[126, 114]",
                    "slot": 8,
                    "x": 126,
                    "y": 114,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 10799494,
                    "player_slot": 131
                },
                {
                    "time": 430,
                    "type": "sen_log",
                    "key": "[60, 132]",
                    "slot": 8,
                    "x": 60,
                    "y": 132,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 16255104,
                    "player_slot": 131
                },
                {
                    "time": 663,
                    "type": "sen_log",
                    "key": "[148, 122]",
                    "slot": 8,
                    "x": 148,
                    "y": 122,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 5736465,
                    "player_slot": 131
                },
                {
                    "time": 825,
                    "type": "sen_log",
                    "key": "[100, 156]",
                    "slot": 8,
                    "x": 100,
                    "y": 156,
                    "z": 128,
                    "entityleft": false,
                    "ehandle": 5489427,
                    "player_slot": 131
                },
                {
                    "time": 1052,
                    "type": "sen_log",
                    "key": "[160, 96]",
                    "slot": 8,
                    "x": 160,
                    "y": 96,
                    "z": 128,
                    "entityleft": false,
                    "ehandle": 5948262,
                    "player_slot": 131
                },
                {
                    "time": 1233,
                    "type": "sen_log",
                    "key": "[116, 132]",
                    "slot": 8,
                    "x": 116,
                    "y": 132,
                    "z": 128,
                    "entityleft": false,
                    "ehandle": 10225310,
                    "player_slot": 131
                },
                {
                    "time": 1550,
                    "type": "sen_log",
                    "key": "[116, 154]",
                    "slot": 8,
                    "x": 116,
                    "y": 154,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 16236810,
                    "player_slot": 131
                },
                {
                    "time": 1710,
                    "type": "sen_log",
                    "key": "[146, 156]",
                    "slot": 8,
                    "x": 146,
                    "y": 156,
                    "z": 128,
                    "entityleft": false,
                    "ehandle": 1065434,
                    "player_slot": 131
                },
                {
                    "time": 1728,
                    "type": "sen_log",
                    "key": "[142, 128]",
                    "slot": 8,
                    "x": 142,
                    "y": 128,
                    "z": 128,
                    "entityleft": false,
                    "ehandle": 10421072,
                    "player_slot": 131
                },
                {
                    "time": 1796,
                    "type": "sen_log",
                    "key": "[150, 72]",
                    "slot": 8,
                    "x": 150,
                    "y": 72,
                    "z": 128,
                    "entityleft": false,
                    "ehandle": 12484916,
                    "player_slot": 131
                },
                {
                    "time": 2011,
                    "type": "sen_log",
                    "key": "[146, 70]",
                    "slot": 8,
                    "x": 146,
                    "y": 70,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 6572461,
                    "player_slot": 131
                },
                {
                    "time": 2163,
                    "type": "sen_log",
                    "key": "[124, 162]",
                    "slot": 8,
                    "x": 124,
                    "y": 162,
                    "z": 128,
                    "entityleft": false,
                    "ehandle": 9013445,
                    "player_slot": 131
                },
                {
                    "time": 2395,
                    "type": "sen_log",
                    "key": "[144, 120]",
                    "slot": 8,
                    "x": 144,
                    "y": 120,
                    "z": 132,
                    "entityleft": false,
                    "ehandle": 11241002,
                    "player_slot": 131
                },
                {
                    "time": 2402,
                    "type": "sen_log",
                    "key": "[156, 130]",
                    "slot": 8,
                    "x": 156,
                    "y": 130,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 7948910,
                    "player_slot": 131
                },
                {
                    "time": 2567,
                    "type": "sen_log",
                    "key": "[146, 70]",
                    "slot": 8,
                    "x": 146,
                    "y": 70,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 6785423,
                    "player_slot": 131
                },
                {
                    "time": 2734,
                    "type": "sen_log",
                    "key": "[172, 148]",
                    "slot": 8,
                    "x": 172,
                    "y": 148,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 8880448,
                    "player_slot": 131
                },
                {
                    "time": 2882,
                    "type": "sen_log",
                    "key": "[156, 176]",
                    "slot": 8,
                    "x": 156,
                    "y": 176,
                    "z": 130,
                    "entityleft": false,
                    "ehandle": 5046778,
                    "player_slot": 131
                }
            ],
            "account_id": 16497807,
            "name": "tOfu"
        }
    ]
}`
