package config

import "os"

// create a app level configuration
type Config struct {
	ServerPort string
}

// Load reads values from environment variables
func Load() *Config {
	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

// read from .env with fallback
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}