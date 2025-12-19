package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/varun053101/GO-TASK/config"
	"github.com/varun053101/GO-TASK/internal/logger"
	"github.com/varun053101/GO-TASK/internal/middleware"
	"github.com/varun053101/GO-TASK/internal/repository"
	"github.com/varun053101/GO-TASK/internal/handler"
	"github.com/varun053101/GO-TASK/internal/routes"
)

func main() {
	_ = godotenv.Load()

	logger.Init()
	defer logger.Log.Sync()

	cfg := config.Load()
	repository.Connect(cfg)

	app := fiber.New()
	app.Use(middleware.RequestLogger())

	userRepo := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepo)

	routes.RegisterHealthRoutes(app)
	routes.RegisterUserRoutes(app, userHandler)

	app.Listen(":" + cfg.ServerPort)
}