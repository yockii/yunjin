package controller

import (
	"github.com/gofiber/fiber/v3"
)

type Controller interface {
	InitializeRouter(router fiber.Router)
}

var controllers []Controller

func InitializeRouter(app *fiber.App) {
	router := app.Group("/api/v1")

	for _, controller := range controllers {
		controller.InitializeRouter(router)
	}
}
