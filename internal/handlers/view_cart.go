package handlers

import (
	repositoryerrors "cart_api/internal/errors/repository_errors"
	"encoding/json"
	"errors"
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
		if errors.Is(err, repositoryerrors.ErrCartNotFound) {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(cart)
}
