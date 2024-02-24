package entity

type ProPlayer struct {
	AccountID int64  `json:"account_id"`
	Name      string `json:"name"`
}

type PlayerMatchOverview struct {
	MatchID      int64 `json:"match_id"`
	PlayerSlot   int   `json:"player_slot"`
	RadiantWin   bool  `json:"radiant_win"`
	Duration     int   `json:"duration"`
	GameMode     int   `json:"game_mode"`
	LobbyType    int   `json:"lobby_type"`
	HeroID       int   `json:"hero_id"`
	StartTime    int   `json:"start_time"`
	Version      int   `json:"version"`
	Kills        int   `json:"kills"`
	Deaths       int   `json:"deaths"`
	Assists      int   `json:"assists"`
	AverageRank  int   `json:"average_rank"`
	LeaverStatus int   `json:"leaver_status"`
	PartySize    int   `json:"party_size"`
}
type PlayerWardsView struct {
	Profile Profile `json:"profile"`
	Dire    *Side   `json:"dire,omitempty"`
	Radiant *Side   `json:"radiant,omitempty"`
}

type Profile struct {
	AccountID   int64  `json:"account_id"`
	Nickname    string `json:"nickname"`
	MatchPlayed int64  `json:"match_played"`
	Observers   int64  `json:"observers"`
	Sentry      int64  `json:"sentry"`
}
type Side struct {
	Obs []WardView `json:"obs,omitempty"`
	Sen []WardView `json:"sen,omitempty"`
}
