package container

import (
	"gamestone/config"

	"github.com/gofiber/fiber/v2"
)

type AppContainer struct {
	App    *fiber.App
	Config config.AppConfig
}
