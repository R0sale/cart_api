package handlers

import (
	"cart_api/internal/entity"
	"encoding/json"
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

	if newItem.Product == "" {
		http.Error(w, "product name is required", http.StatusBadRequest)
		return
	} else if newItem.Price < 0 {
		http.Error(w, "price must be a non-negative value", http.StatusBadRequest)
		return
	}

	cartId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "couldn't parse the cart id from the url", http.StatusBadRequest)
		return
	}

	item, err := handler.service.AddItemToCart(cartId, newItem)
	if err != nil {
		errorsCheck(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}
