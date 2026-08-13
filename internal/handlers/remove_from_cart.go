package handlers

import (
	repositoryerrors "cart_api/internal/errors/repository_errors"
	"errors"
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
		if errors.Is(err, repositoryerrors.ErrCartNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if errors.Is(err, repositoryerrors.ErrItemNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
