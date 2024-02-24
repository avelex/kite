package v1

import "github.com/avelex/kite/internal/entity"

type showPlayerWardsRequest struct {
	Nickname string       `json:"nickname"`
	Patch    string       `json:"patch,omitempty"`
	Side     entity.Sides `json:"side,omitempty"`
	Wards    entity.Wards `json:"wards,omitempty"`
	Time     []uint16     `json:"time,omitempty"`
}
