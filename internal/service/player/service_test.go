package player

import (
	"fmt"
	"testing"

	"github.com/avelex/kite/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	some := []uint16{1, 2, 3}
	convert := [2]uint16(some)
	fmt.Printf("convert: %v\n", convert)
}

func TestCalculate(t *testing.T) {
	r := require.New(t)

	testCases := []struct {
		name  string
		wards []entity.PlayerWard
		side  string
		wtype entity.WardType
		want  []entity.WardView
	}{
		{
			name: "One Spot / One Ward / One Side",
			wards: []entity.PlayerWard{
				{
					X:    1,
					Y:    1,
					Side: entity.DireSide,
					Type: entity.Sentry,
				},
				{
					X:    2,
					Y:    2,
					Side: entity.DireSide,
					Type: entity.Sentry,
				},
			},
			side:  entity.DireSide,
			wtype: entity.Sentry,
			want: []entity.WardView{
				{
					X:    0,
					Y:    0,
					Rate: 10000,
				},
			},
		},
		{
			name: "One Spot / Different Ward / One Side",
			wards: []entity.PlayerWard{
				{
					X:    1,
					Y:    1,
					Side: entity.DireSide,
					Type: entity.Sentry,
				},
				{
					X:    2,
					Y:    2,
					Side: entity.DireSide,
					Type: entity.Observer,
				},
			},
			side:  entity.DireSide,
			wtype: entity.Sentry,
			want: []entity.WardView{
				{
					X:    0,
					Y:    0,
					Rate: 10000,
				},
			},
		},
		{
			name: "One Spot / One Ward / Different Side",
			wards: []entity.PlayerWard{
				{
					X:    1,
					Y:    1,
					Side: entity.DireSide,
					Type: entity.Sentry,
				},
				{
					X:    2,
					Y:    2,
					Side: entity.RadiantSide,
					Type: entity.Sentry,
				},
			},
			side:  entity.DireSide,
			wtype: entity.Sentry,
			want: []entity.WardView{
				{
					X:    0,
					Y:    0,
					Rate: 10000,
				},
			},
		},
		{
			name: "Different Spot / One Ward / One Side",
			wards: []entity.PlayerWard{
				{
					X:    1,
					Y:    1,
					Side: entity.DireSide,
					Type: entity.Sentry,
				},
				{
					X:    6,
					Y:    6,
					Side: entity.DireSide,
					Type: entity.Sentry,
				},
			},
			side:  entity.DireSide,
			wtype: entity.Sentry,
			want: []entity.WardView{
				{
					X:    0,
					Y:    0,
					Rate: 5000,
				},
				{
					X:    5,
					Y:    5,
					Rate: 5000,
				},
			},
		},
		{
			name: "Different Spot / Different Ward / One Side",
			wards: []entity.PlayerWard{
				{
					X:    1,
					Y:    1,
					Side: entity.DireSide,
					Type: entity.Sentry,
				},
				{
					X:    6,
					Y:    6,
					Side: entity.DireSide,
					Type: entity.Observer,
				},
			},
			side:  entity.DireSide,
			wtype: entity.Sentry,
			want: []entity.WardView{
				{
					X:    0,
					Y:    0,
					Rate: 10000,
				},
			},
		},
		{
			name: "Different Spot / Different Ward / Deferent Side",
			wards: []entity.PlayerWard{
				{
					X:    1,
					Y:    1,
					Side: entity.DireSide,
					Type: entity.Sentry,
				},
				{
					X:    6,
					Y:    6,
					Side: entity.RadiantSide,
					Type: entity.Observer,
				},
			},
			side:  entity.DireSide,
			wtype: entity.Sentry,
			want: []entity.WardView{
				{
					X:    0,
					Y:    0,
					Rate: 10000,
				},
			},
		},
		{
			name:  "Empty",
			wards: []entity.PlayerWard{},
			side:  entity.DireSide,
			wtype: entity.Sentry,
			want:  []entity.WardView{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := calculateWardsFreqBySides(tC.wards, tC.side, tC.wtype)
			r.ElementsMatch(tC.want, got)
		})
	}
}
