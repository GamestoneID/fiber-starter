package routes

import (
	"gamestone/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterGameRoutes(router fiber.Router, handler handlers.GameHandler) {
	game := router.Group("/games")
	game.Get("/", handler.GetGames)           // Digunakan pada dashboard gamestone
	game.Post("/", handler.CreateGame)        // Kemungkinan tidak digunakan, data voca product hanya bisa diambil dari sync
	game.Get("/:slug", handler.GetGame)       // Kemungkinan tidak digunakan, sync products langsung lewat repository
	game.Put("/:slug", handler.UpdateGame)    // Kemungkinan tidak digunakan, data voca product hanya bisa diambil dari sync
	game.Delete("/:slug", handler.DeleteGame) // Kemungkinan tidak digunakan, data voca product hanya bisa diambil dari sync
}
