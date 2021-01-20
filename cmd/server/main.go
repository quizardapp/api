package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	apiserver "github.com/quizardapp/auth-api/pkg/server"
)

func main() {
	godotenv.Load("config.env")

	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
