package main

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jordanfox1/img-switch-api/api/img-switch-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 100, // 100MB max request body size
	})

	app.Use(limiter.New(limiter.Config{
		Max:               100,               // 100 requests
		Expiration:        600 * time.Second, // per 10 mins
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Use(cors.New())
	app.Use(logger.New())

	routes.SetupRoutes(app)

	app.Listen(":5000")
}
