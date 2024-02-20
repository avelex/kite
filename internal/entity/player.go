package entity

type PlayerWardsView struct {
	Nickname  string         `json:"nickname"`
	AccountID int64          `json:"account_id"`
	Dire      map[int][]Ward `json:"dire"`
	Radiant   map[int][]Ward `json:"radiant"`
}
