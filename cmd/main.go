package main

import (
	"cart_api/internal/config"
	"cart_api/internal/db"
	"cart_api/internal/handlers"
	"cart_api/internal/repository"
	"cart_api/internal/service"
	"log"
	"os"
)

func main() {
	configPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		log.Fatalf("couldn't find the env var with the config path")
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("%v", err)
	}

	db, err := db.ConfigureDb(*cfg)
	if err != nil {
		log.Fatalf("problems with db: %v", err)
	}

	defer db.Close()

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers.ConfigureServer(*cfg, service)
}
