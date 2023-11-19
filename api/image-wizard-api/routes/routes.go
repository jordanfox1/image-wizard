package routes

import (
	"net/http"

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

		convertedImage, err := handlers.ConvertImage(c.Body(), desiredFormat)
		if err != nil {
			// Return a custom error response with the error message
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		return c.Send(convertedImage)
	})
}
