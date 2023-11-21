package routes

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
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
		fmt.Println("Incoming request to /convert endpoint")

		desiredFormat := c.Query("format")
		inputImage := c.FormValue("image")
		inputFileName := c.FormValue("fileName")
		outputFileName := strings.TrimSuffix(inputFileName, filepath.Ext(inputFileName)) + "." + desiredFormat

		b64data := inputImage[strings.IndexByte(inputImage, ',')+1:]
		decodedData, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			log.Println("base64 decoding error --> ", err)
			return c.JSON(fiber.Map{"status": 500, "message": "Base64 decoding error", "dataURL": ""})
		}

		convertedImage, err := handlers.ConvertImage(decodedData, desiredFormat)
		if err != nil {
			// Return a custom error response with the error message
			fmt.Println("Image conversion failed:", err)
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		// Convert the image bytes to a data URL
		dataURL := fmt.Sprintf("data:image/%s;base64,%s", desiredFormat, base64.StdEncoding.EncodeToString(convertedImage))

		fmt.Println("Image converted successfully")
		res := c.JSON(fiber.Map{"status": 200, "message": "Image converted successfully", "dataURL": dataURL, "fileName": outputFileName})
		return res
	})
}
