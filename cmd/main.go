package main

import (
	"cart_api/internal/config"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	DB *sql.DB
}

func (app *App) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi there we are done")
}

func main() {
	configPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		log.Fatalf("couldn't find the env var with the config path")
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("%v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("couldn't open database: %v", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("couldn't ping database: %v", err)
	}

	myApp := &App{
		DB: db,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", myApp.Handle)

	serverAddress := fmt.Sprintf(":%d", cfg.Server.Port)
	http.ListenAndServe(serverAddress, mux)
}
