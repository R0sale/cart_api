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

func (handler ViewCartHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid cart ID", http.StatusBadRequest)
		return
	}

	cart, err := handler.service.ViewCart(id)
	if err != nil {
		errorsCheck(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}
