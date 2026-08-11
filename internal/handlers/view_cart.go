package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ViewCartHandler struct {
	service cartService
}

func NewViewCartHandler(service cartService) *ViewCartHandler {
	return &ViewCartHandler{
		service: service,
	}
}

func (handler ViewCartHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.PathValue("id"))
	if err != nil {
		http.Error(writer, "Invalid cart ID", http.StatusBadRequest)
		return
	}

	cart, err := handler.service.ViewCart(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(cart)
}
