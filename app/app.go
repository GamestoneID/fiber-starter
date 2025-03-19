package app

import (
	"gamestone/config"
	"gamestone/container"
	"gamestone/handlers"
	"gamestone/routes"

	"github.com/gofiber/fiber/v2"
)

// NewFiberApp initializes the Fiber application with all necessary routes
func NewFiberApp(
	appConfig config.AppConfig,
	gameHandler handlers.GameHandler,
) (*container.AppContainer, error) {
	app := fiber.New(
		fiber.Config{
			AppName: appConfig.AppName,
		},
	)

	// API Versioning
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Register routes
	routes.RegisterGameRoutes(v1, gameHandler)

	return &container.AppContainer{
		App:    app,
		Config: appConfig,
	}, nil
}
