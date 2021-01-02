package main

import (
	"log"

	"github.com/inctnce/quizard-api/internal/server"
	"github.com/joho/godotenv"
)

var ()

func main() {
	godotenv.Load()
	config := server.NewConfig()

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
