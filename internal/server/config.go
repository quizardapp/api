package server

import (
	"os"

	"github.com/inctnce/quizard-api/internal/store"
)

// Config of the server
type Config struct {
	Port     string `toml:"PORT"`
	LogLevel string `toml:"LOG_LEVEL"`
	Store    *store.Config
}

// NewConfig returns new Config instance
func NewConfig() *Config {

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "debug"
	}

	return &Config{
		Port:     port,
		LogLevel: logLevel,
		Store:    store.NewConfig(),
	}
}
