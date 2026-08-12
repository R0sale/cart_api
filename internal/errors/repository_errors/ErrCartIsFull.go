package repositoryerrors

import "errors"

var (
	ErrCartIsFull = errors.New("cart has more than 5 items, cannot add more")
)
