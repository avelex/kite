package entity

type WardType string

const (
	Observer WardType = "observer"
	Sentry   WardType = "sentry"
)

type Wards map[WardType]struct{}

func (w Wards) HasObserver() bool {
	_, ok := w[Observer]
	return ok
}

func (w Wards) HasSentry() bool {
	_, ok := w[Sentry]
	return ok
}

func (w Wards) Slice() []string {
	out := make([]string, 0, len(w))
	for k := range w {
		out = append(out, string(k))
	}
	return out
}

type PlayerWard struct {
	ID           uint
	AccountID    int64 `gorm:"index"`
	MatchID      int64
	PatchVersion int64
	Time         int64
	X, Y         uint8
	Side         string
	Type         WardType
}

type WardView struct {
	X    uint8  `json:"x"`
	Y    uint8  `json:"y"`
	Rate uint16 `json:"rate"`
}
