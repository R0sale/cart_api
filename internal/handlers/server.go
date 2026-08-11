package handlers

import (
	"cart_api/internal/config"
	"fmt"
	"net/http"
)

func ConfigureServer(cfg config.Config, service cartService) {
	mux := http.NewServeMux()

	createCartHandler := NewCreateCartHandler(service)
	viewCartHandler := NewViewCartHandler(service)
	mux.Handle("POST /api/v1/cart", createCartHandler)
	mux.Handle("GET /api/v1/cart/{id}", viewCartHandler)

	serverAddress := fmt.Sprintf(":%d", cfg.Server.Port)
	http.ListenAndServe(serverAddress, mux)
}
