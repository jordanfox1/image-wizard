package routes

import (
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"

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
		inputImageDataURL := c.FormValue("image")
		inputFileName := c.FormValue("fileName")
		outputFileName := strings.TrimSuffix(inputFileName, filepath.Ext(inputFileName)) + "." + desiredFormat

		convertedImage, err := handlers.ConvertImage(inputImageDataURL, desiredFormat)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
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
