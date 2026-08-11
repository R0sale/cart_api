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

func (handler CreateCartHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	cart, err := handler.service.CreateCart()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(cart)
}
