package router

import (
	"github.com/gofiber/fiber/v2"
)

func ConfigRoute(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})
}