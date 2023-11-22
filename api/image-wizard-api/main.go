package main

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jordanfox1/image-wizard-api/api/image-wizard-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	routes.SetupRoutes(app)

	app.Listen(":5000")
}
