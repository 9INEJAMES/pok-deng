package handlers

import (
	appErr "github.com/9inejames/pok-deng/backend/internal/errors"
	"github.com/9inejames/pok-deng/backend/internal/models"
	svc "github.com/9inejames/pok-deng/backend/internal/services"
	"github.com/9inejames/pok-deng/backend/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type GameHandler struct {
	service svc.GameService
}

func NewGameHandler(service svc.GameService) *GameHandler {
	return &GameHandler{service}
}

func (h *GameHandler) Start(c *fiber.Ctx) error {
	var req models.GameStart

	if err := utils.ParseBody(c, &req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": appErr.ErrInvalidRequest.Error(),
		})
	}

	if req.InitialBalance <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": appErr.ErrInvalidRequest.Error(),
		})
	}

	game, err := h.service.Start(&req)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(game)

}

func (h *GameHandler) Action(c *fiber.Ctx) error {
	id := c.Params("game_id")
	var req models.GameAction

	if err := utils.ParseBody(c, &req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": appErr.ErrInvalidRequest.Error(),
		})
	}

	if req.Action == models.BET && req.Amount <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": appErr.ErrInvalidRequest.Error(),
		})
	}

	game, err := h.service.Action(id, &req)

	if err != nil {
		if err == appErr.ErrSessionNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	if game.State == models.RoundEnd {

		gameRes := &models.GameResponse{
			GameID:            game.ID,
			State:             string(game.State),
			Balance:           game.Balance,
			PlayerHand:        game.PlayerHand,
			DealerHandVisible: game.DealerHand,
			PlayerScore:       &game.PlayerScore,
			DealerScore:       &game.DealerScore,
			Winner:            (*string)(&game.Winner),
		}
		return c.Status(200).JSON(gameRes)
	}

	return c.Status(200).JSON(game)
}
