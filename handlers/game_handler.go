package handlers

import "github.com/gofiber/fiber/v2"

type GameHandler interface {
	CreateGame(c *fiber.Ctx) error
	GetGames(c *fiber.Ctx) error
	GetGame(c *fiber.Ctx) error
	UpdateGame(c *fiber.Ctx) error
	DeleteGame(c *fiber.Ctx) error
}
