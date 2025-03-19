package handlers

import (
	"gamestone/dto"
	"gamestone/errors"
	"gamestone/services"
	"gamestone/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type GameHandlerImpl struct {
	service  services.GameService
	validate *validator.Validate
}

// CreateGame implements GameHandler.
func (v *GameHandlerImpl) CreateGame(c *fiber.Ctx) error {
	var game dto.CreateGameRequest

	if err := c.BodyParser(&game); err != nil {
		response, code := utils.HandleErrors(errors.NewHttpError(fiber.StatusBadRequest, "Invalid request body"))
		return c.Status(code).JSON(response)
	}

	if err := v.validate.Struct(game); err != nil {
		response, code := utils.HandleErrors(err)
		return c.Status(code).JSON(response)
	}

	gameCreated, err := v.service.CreateGame(&game)
	if err != nil {
		response, code := utils.HandleErrors(err)
		return c.Status(code).JSON(response)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    gameCreated,
		"message": "Game created successfully",
	})
}

// DeleteGame implements GameHandler.
func (v *GameHandlerImpl) DeleteGame(c *fiber.Ctx) error {
	slug := c.Params("slug")

	if err := v.service.DeleteGame(slug); err != nil {
		response, code := utils.HandleErrors(err)
		return c.Status(code).JSON(response)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetGame implements GameHandler.
func (v *GameHandlerImpl) GetGame(c *fiber.Ctx) error {
	slug := c.Params("slug")

	game, err := v.service.GetGameBySlug(slug)

	if err != nil {
		response, code := utils.HandleErrors(err)
		return c.Status(code).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    game,
		"message": "Game retrieved successfully",
	})
}

// GetGames implements GameHandler.
func (v *GameHandlerImpl) GetGames(c *fiber.Ctx) error {
	isActiveFilter := utils.GetOptionalBoolQuery(c, "is_active")
	limitFilter := utils.GetOptionalIntQuery(c, "limit")
	pageFilter := utils.GetOptionalIntQuery(c, "page")

	response, err := v.service.GetAllGames(&dto.FilterGame{
		IsActive: isActiveFilter,
		Limit:    limitFilter,
		Page:     pageFilter,
	})

	if err != nil {
		response, code := utils.HandleErrors(err)
		return c.Status(code).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":     response.Data,
		"metadata": response.Metadata,
		"message":  "Games retrieved successfully",
	})
}

// UpdateGame implements GameHandler.
func (v *GameHandlerImpl) UpdateGame(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var game dto.UpdateGameRequest

	if err := c.BodyParser(&game); err != nil {
		response, code := utils.HandleErrors(errors.NewHttpError(fiber.StatusBadRequest, "Invalid request body"))
		return c.Status(code).JSON(response)
	}

	if err := v.validate.Struct(game); err != nil {
		response, code := utils.HandleErrors(err)
		return c.Status(code).JSON(response)
	}

	gameUpdated, err := v.service.UpdateGame(slug, &game)
	if err != nil {
		response, code := utils.HandleErrors(err)
		return c.Status(code).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    gameUpdated,
		"message": "Game updated successfully",
	})
}

func NewGameHandler(service services.GameService, validate *validator.Validate) GameHandler {
	return &GameHandlerImpl{
		service:  service,
		validate: validate,
	}
}
