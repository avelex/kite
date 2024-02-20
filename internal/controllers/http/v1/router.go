package v1

import (
	"context"

	"github.com/avelex/kite/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type PlayersUsecases interface {
	PlayerWards(ctx context.Context, nickname string) (entity.PlayerWardsView, error)
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
		v1.Post("/player", h.showPlayerWards)
	}
}

func (h *handler) status(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func (h *handler) showPlayerWards(ctx *fiber.Ctx) error {
	var in showPlayerInfoRequest
	if err := ctx.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	view, err := h.pu.PlayerWards(ctx.Context(), in.Nickname)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(view)
}
