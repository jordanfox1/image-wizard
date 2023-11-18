package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋! from API!!")
	})
}
