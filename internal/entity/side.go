package entity

const (
	DireSide    = "dire"
	RadiantSide = "radiant"
)

type Sides map[string]struct{}

func (s Sides) HasRadiant() bool {
	_, ok := s[RadiantSide]
	return ok
}

func (s Sides) HasDire() bool {
	_, ok := s[DireSide]
	return ok
}

func (s Sides) Slice() []string {
	out := make([]string, 0, len(s))
	for k := range s {
		out = append(out, k)
	}
	return out
}
