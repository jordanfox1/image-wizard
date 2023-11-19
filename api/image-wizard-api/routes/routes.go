package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jordanfox1/image-wizard-api/api/image-wizard-api/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋 - from image wizard api")
	})

	api.Post("/convert", func(c *fiber.Ctx) error {
		desiredFormat := c.Query("format")
		// TODO: handle improper desired format

		convertedImage, err := handlers.ConvertImage(c.Body(), desiredFormat)
		if err != nil {
			return err
		}

		return c.Send(convertedImage)
	})
}
