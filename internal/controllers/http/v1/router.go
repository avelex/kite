package v1

import (
	"context"
	"strings"

	"github.com/avelex/kite/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type PlayersUsecases interface {
	PlayerWards(ctx context.Context, nickname, patch string, sides entity.Sides, wards entity.Wards, times []int16) (entity.PlayerWardsView, error)
}

type handler struct {
	pu PlayersUsecases
}

func New(pu PlayersUsecases) *handler {
	return &handler{
		pu: pu,
	}
}

func (h *handler) RegisterRoutes(r fiber.Router) {
	v1 := r.Group("/v1")
	{
		v1.Get("/status", h.status)
		v1.Get("/hello", h.hello)
		v1.Post("/player", h.showPlayerWards)
	}
}

func (h *handler) status(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func (h *handler) hello(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello")
}

func (h *handler) showPlayerWards(ctx *fiber.Ctx) error {
	var in showPlayerWardsRequest
	if err := ctx.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	if strings.TrimSpace(in.Nickname) == "" {
		return fiber.ErrBadRequest
	}

	view, err := h.pu.PlayerWards(ctx.Context(), in.Nickname, in.Patch, in.Side, in.Wards, in.Time)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(view)
}
