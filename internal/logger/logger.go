package logger

import "go.uber.org/zap"

// global logger instance used across the app
var Log *zap.Logger

// initialize the zap logger and will be called once at startup
func Init() {
	Log, _ = zap.NewProduction()
}

