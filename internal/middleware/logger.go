package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/varun053101/GO-TASK/internal/logger"
	"go.uber.org/zap"
)

// logs basic details for every incoming request
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// continue handling the request
		err := c.Next()

		// log request details after response is sent
		logger.Log.Info(
			"request completed",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", time.Duration(time.Since(start).Milliseconds())),
		)

		return err
	}
}
