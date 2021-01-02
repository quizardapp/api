package store

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	DBUrl string
}

func NewConfig() *Config {

	DBUrl := os.Getenv("CLEARDB_DATABASE_URL")
	if DBUrl == "" {
		logrus.Error("cannot connect to the database")
	}

	return &Config{
		DBUrl: DBUrl,
	}
}
