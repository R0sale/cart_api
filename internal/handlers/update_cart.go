package handlers

import (
	"cart_api/internal/entity"
	"encoding/json"
	"net/http"
	"strconv"
)

type UpdateCartHandler struct {
	service cartService
}

func NewUpdateCartHandler(service cartService) *UpdateCartHandler {
	return &UpdateCartHandler{
		service: service,
	}
}

func (h *UpdateCartHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cartId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "couldn't parse the cart id from the url", http.StatusBadRequest)
		return
	}

	itemId, err := strconv.Atoi(r.PathValue("itemId"))
	if err != nil {
		http.Error(w, "couldn't parse the item id from the url", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var updatedItem entity.NewItemDto
	if err := decoder.Decode(&updatedItem); err != nil {
		http.Error(w, "couldn't decode the body from request in update_cart.go", http.StatusBadRequest)
		return
	}

	if updatedItem.Product == "" {
		http.Error(w, "product name is required", http.StatusBadRequest)
		return
	}

	if updatedItem.Price < 0 {
		http.Error(w, "price must be a non-negative value", http.StatusBadRequest)
		return
	}

	item, err := h.service.UpdateCartItem(cartId, itemId, updatedItem)
	if err != nil {
		errorsCheck(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}
