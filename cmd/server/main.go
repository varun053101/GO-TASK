package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/varun053101/GO-TASK/config"
	"github.com/varun053101/GO-TASK/internal/handler"
	"github.com/varun053101/GO-TASK/internal/logger"
	"github.com/varun053101/GO-TASK/internal/middleware"
	"github.com/varun053101/GO-TASK/internal/repository"
	"go.uber.org/zap"
)

func main() {

	_ = godotenv.Load()

	// initialize logger
	logger.Init()
	defer logger.Log.Sync()

	// load config
	cfg := config.Load()

	// connect to database
	if err := repository.Connect(cfg); err != nil {
		logger.Log.Fatal("failed to connect to database", zap.Error(err))
	}

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

	userRepo := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepo)
	// Create a User
	app.Post("/users", userHandler.CreateUser)
	// Get user by user id
	app.Get("/users/:id", userHandler.GetUserByID)

	// start the server
	err := app.Listen(":" + cfg.ServerPort)
	if err != nil {
		logger.Log.Fatal("server failed to start", zap.Error(err))
	}
}
