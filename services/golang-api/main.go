package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(struct {
			Message string
		}{
			Message: os.Getenv("MESSAGE"),
		})
	})

	_ = app.Listen(":8000")
}
