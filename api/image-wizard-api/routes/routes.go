package routes

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jordanfox1/image-wizard-api/api/image-wizard-api/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋 - from image wizard api")
	})

	// api.Post("/convert", func(c *fiber.Ctx) error {

	// file := decodedData

	// generate new uuid for image name
	// uniqueId := uuid.New()

	// // remove "- from imageName"

	// filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// // extract image extension from original file filename

	// fileExt := strings.Split(file.Filename, ".")[1]

	// // generate image from filename and extension
	// image := fmt.Sprintf("%s.%s", filename, fileExt)

	// // save image to ./images dir
	// err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	// if err != nil {
	// 	log.Println("image save error --> ", err)
	// 	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	// }

	// // generate image url to serve to client using CDN

	// imageUrl := fmt.Sprintf("http://localhost:4000/images/%s", image)

	// // create meta data and send to client

	// data := map[string]interface{}{

	// 	"imageName": image,
	// 	"imageUrl":  imageUrl,
	// 	"header":    file.Header,
	// 	"size":      file.Size,
	// }

	// 	return c.JSON(fiber.Map{"status": 200, "message": "Image conversion successful", "data": decodedData})
	// })
	api.Post("/convert", func(c *fiber.Ctx) error {
		fmt.Println("Incoming request to /convert endpoint")

		input := c.FormValue("image")

		b64data := input[strings.IndexByte(input, ',')+1:]
		// Decode base64-encoded image data
		decodedData, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			log.Println("base64 decoding error --> ", err)
			return c.JSON(fiber.Map{"status": 500, "message": "Base64 decoding error", "dataURL": ""})
		}

		desiredFormat := c.Query("format")

		convertedImage, err := handlers.ConvertImage(decodedData, desiredFormat)
		if err != nil {
			// Return a custom error response with the error message
			fmt.Println("Image conversion failed:", err)
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		// Convert the image bytes to a data URL
		dataURL := fmt.Sprintf("data:image/%s;base64,%s", desiredFormat, base64.StdEncoding.EncodeToString(convertedImage))

		fmt.Println("Image converted successfully")
		return c.JSON(fiber.Map{"status": 200, "message": "Image converted successfully", "dataURL": dataURL})
	})
}
