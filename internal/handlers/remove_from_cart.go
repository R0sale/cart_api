package handlers

import (
	"net/http"
	"strconv"
)

type RemoveFromCartHandler struct {
	service cartService
}

func NewRemoveFromCartHandler(service cartService) *RemoveFromCartHandler {
	return &RemoveFromCartHandler{
		service: service,
	}
}

func (h *RemoveFromCartHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.RemoveFromCart(cartId, itemId)
	if err != nil {
		errorsCheck(err, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
