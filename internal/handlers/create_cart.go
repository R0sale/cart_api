package handlers

import (
	"encoding/json"
	"net/http"
)

type CreateCartHandler struct {
	service cartService
}

func NewCreateCartHandler(service cartService) *CreateCartHandler {
	return &CreateCartHandler{
		service: service,
	}
}

func (handler CreateCartHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cart, err := handler.service.CreateCart()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}
