package routes

import (
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jordanfox1/img-switch-api/api/img-switch-api/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋 - from Img Switch api")
	})

	api.Post("/convert", func(c *fiber.Ctx) error {
		desiredFormat := c.Query("format")
		inputImageDataURL := c.FormValue("image")
		inputFileName := c.FormValue("fileName")
		outputFileName := strings.TrimSuffix(inputFileName, filepath.Ext(inputFileName)) + "." + desiredFormat

		convertedImage, err := handlers.ConvertImage(inputImageDataURL, desiredFormat)
		if err != nil {
			if strings.Contains(err.Error(), "input and output formats cannot be the same") {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"status":  400,
					"message": "Bad Request: input and output formats cannot be the same",
					"error":   err.Error(),
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  500,
				"message": "Failed to convert image",
				"error":   err.Error(),
			})
		}

		dataURL := fmt.Sprintf("data:image/%s;base64,%s", desiredFormat, base64.StdEncoding.EncodeToString(convertedImage))

		return c.JSON(fiber.Map{
			"status":   200,
			"message":  "Image converted successfully",
			"dataURL":  dataURL,
			"fileName": outputFileName,
		})
	})
}
