package handlers

import (
	"encoding/json"
	"net/http"
)

type CreateCartHandler struct {
	service CartService
}

func NewCreateCartHandler(service CartService) *CreateCartHandler {
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

	json.NewEncoder(writer).Encode(cart)
}
