package entity

type FullMatch struct {
	MatchID               int64             `json:"match_id"`
	BarracksStatusDire    int               `json:"barracks_status_dire"`
	BarracksStatusRadiant int               `json:"barracks_status_radiant"`
	Cluster               int               `json:"cluster"`
	DireScore             int               `json:"dire_score"`
	Duration              int               `json:"duration"`
	Engine                int               `json:"engine"`
	FirstBloodTime        int               `json:"first_blood_time"`
	GameMode              int               `json:"game_mode"`
	HumanPlayers          int               `json:"human_players"`
	Leagueid              int               `json:"leagueid"`
	LobbyType             int               `json:"lobby_type"`
	MatchSeqNum           int               `json:"match_seq_num"`
	NegativeVotes         int               `json:"negative_votes"`
	PositiveVotes         int               `json:"positive_votes"`
	RadiantGoldAdv        []int             `json:"radiant_gold_adv"`
	RadiantScore          int               `json:"radiant_score"`
	RadiantWin            bool              `json:"radiant_win"`
	RadiantXpAdv          []int             `json:"radiant_xp_adv"`
	StartTime             int               `json:"start_time"`
	TowerStatusDire       int               `json:"tower_status_dire"`
	TowerStatusRadiant    int               `json:"tower_status_radiant"`
	Version               int               `json:"version"`
	ReplaySalt            int               `json:"replay_salt"`
	SeriesID              int               `json:"series_id"`
	SeriesType            int               `json:"series_type"`
	Skill                 int               `json:"skill"`
	Players               []FullMatchPlayer `json:"players"`
	Patch                 int               `json:"patch"`
	Region                int               `json:"region"`
	ReplayURL             string            `json:"replay_url"`
}

type FullMatchPlayer struct {
	MatchID       int64  `json:"match_id"`
	PlayerSlot    int    `json:"player_slot"`
	AccountID     int64  `json:"account_id"`
	PersonaName   string `json:"personaname"`
	Name          string `json:"name"`
	StartTime     int    `json:"start_time"`
	CampsStacked  int    `json:"camps_stacked"`
	CreepsStacked int    `json:"creeps_stacked"`
	HeroID        int    `json:"hero_id"`
	IsRadiant     bool   `json:"isRadiant"`

	Obs           []FullMatchWard `json:"obs_log"`
	ObsPlaced     int             `json:"obs_placed"`
	ObserverKills int             `json:"observer_kills"`
	ObserverUses  int             `json:"observer_uses"`

	Sen         []FullMatchWard `json:"sen_log"`
	SenPlaced   int             `json:"sen_placed"`
	SentryKills int             `json:"sentry_kills"`
	SentryUses  int             `json:"sentry_uses"`
}

type FullMatchWard struct {
	Time int64  `json:"time"`
	Type string `json:"type"`
	Key  string `json:"key"`
	X    uint8  `json:"x"`
	Y    uint8  `json:"y"`
	Z    uint8  `json:"z"`
}

func (m FullMatch) Player(accountID int64) FullMatchPlayer {
	for _, v := range m.Players {
		if v.AccountID == accountID {
			return v
		}
	}
	return FullMatchPlayer{}
}
