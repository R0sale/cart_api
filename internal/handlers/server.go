package handlers

import (
	"cart_api/internal/config"
	"fmt"
	"net/http"
)

func ConfigureServer(cfg config.Config, service CartService) {
	mux := http.NewServeMux()

	createCartHandler := NewCreateCartHandler(service)
	mux.Handle("/api/v1/cart", createCartHandler)

	serverAddress := fmt.Sprintf(":%d", cfg.Server.Port)
	http.ListenAndServe(serverAddress, mux)
}
