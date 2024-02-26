package v1

import "github.com/avelex/kite/internal/entity"

type showPlayerWardsRequest struct {
	Nickname string       `json:"nickname"`
	Patch    string       `json:"patch"`
	Side     entity.Sides `json:"side"`
	Wards    entity.Wards `json:"wards"`
	Time     []int16      `json:"time"`
}
