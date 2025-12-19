package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/varun053101/GO-TASK/config"
	"github.com/varun053101/GO-TASK/internal/logger"
	"github.com/varun053101/GO-TASK/internal/middleware"
	"go.uber.org/zap"
)

func main() {
	// initialize logger
	logger.Init()
	defer logger.Log.Sync()

	// create fiber app
	app := fiber.New()
	app.Use(middleware.RequestLogger())

	// basic health endpoint to check server status
	app.Get("/health", func(c *fiber.Ctx) error {
		logger.Log.Info("health endpoint called")
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// start the server
	cfg := config.Load()
	err := app.Listen(":" + cfg.ServerPort)
	if err != nil {
		logger.Log.Fatal("server failed to start", zap.Error(err))
	}
}
