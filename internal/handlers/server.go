package handlers

import (
	"cart_api/internal/config"
	"fmt"
	"log"
	"net/http"
)

func ConfigureServer(cfg config.Config, service cartService) {
	mux := http.NewServeMux()

	createCartHandler := NewCreateCartHandler(service)
	viewCartHandler := NewViewCartHandler(service)
	addItemHandler := NewAddItemHandler(service)
	removeFromCartHandler := NewRemoveFromCartHandler(service)
	updateCartHandler := NewUpdateCartHandler(service)
	mux.Handle("POST /api/v1/carts", createCartHandler)
	mux.Handle("POST /api/v1/carts/{id}/items", addItemHandler)
	mux.Handle("GET /api/v1/carts/{id}", viewCartHandler)
	mux.Handle("DELETE /api/v1/carts/{id}/items/{itemId}", removeFromCartHandler)
	mux.Handle("PUT /api/v1/carts/{id}/items/{itemId}", updateCartHandler)

	serverAddress := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(serverAddress, mux))
}
