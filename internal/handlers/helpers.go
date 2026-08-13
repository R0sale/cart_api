package handlers

import (
	repositoryerrors "cart_api/internal/errors/repository_errors"
	"errors"
	"net/http"
)

func errorsCheck(err error, w http.ResponseWriter) {
	if errors.Is(err, repositoryerrors.ErrCartIsFull) {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	} else if errors.Is(err, repositoryerrors.ErrCartNotFound) {
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
