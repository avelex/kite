package entity

type WardType string

const (
	Observer WardType = "obs"
	Sentry   WardType = "sentry"
)

type PlayerWard struct {
	AccountID    int64
	MatchID      int64
	PatchVersion int64
	Time         int64
	X, Y         int
	IsRadiant    bool
	Type         WardType
}

type Ward struct {
	Time      int64  `json:"time"`
	Type      string `json:"type"`
	Key       string `json:"key"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Z         int    `json:"z"`
	IsRadiant bool   `json:"is_radiant"`
}
