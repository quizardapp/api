package server

import "os"

type Config struct {
	Port     string `toml:"PORT"`
	LogLevel string `toml:"LOG_LEVEL"`
	DBUrl    string `toml:"DB_URL"`
}

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
	}
}
