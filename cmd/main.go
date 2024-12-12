package main

import (
	"coding-chelleng/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Get("/metrics", monitor.New())
	app.Use(recover.New())

	config.Initialized(app)
	app.Listen(":3000")
}
