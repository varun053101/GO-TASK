package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/varun053101/GO-TASK/internal/logger"
	"go.uber.org/zap"
)

func main() {
	// initialize logger
	logger.Init()
	defer logger.Log.Sync()

	// create fiber app
	app := fiber.New()

	// basic health endpoint to check server status
	app.Get("/health", func(c *fiber.Ctx) error {
		logger.Log.Info("health endpoint called")
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// start the server
	err := app.Listen(":8080")
	if err != nil {
		logger.Log.Fatal("server failed to start", zap.Error(err))
	}
}
