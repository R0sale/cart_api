package handlers

import (
	"cart_api/internal/entity"
	repositoryerrors "cart_api/internal/errors/repository_errors"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type AddItemHandler struct {
	service cartService
}

func NewAddItemHandler(service cartService) *AddItemHandler {
	return &AddItemHandler{
		service: service,
	}
}

func (handler *AddItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var newItem entity.NewItemDto

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&newItem); err != nil {
		http.Error(w, "couldn't decode the body from request in add_to_cart.go", http.StatusBadRequest)
		return
	}

	cartId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "couldn't parse the cart id from the url", http.StatusBadRequest)
		return
	}

	item, err := handler.service.AddItemToCart(cartId, newItem)
	if err != nil {
		if errors.Is(err, repositoryerrors.ErrCartIsFull) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		} else if errors.Is(err, repositoryerrors.ErrCartNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}
