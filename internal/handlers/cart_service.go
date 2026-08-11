package handlers

import "cart_api/internal/entity"

type cartService interface {
	CreateCart() (*entity.Cart, error)
	ViewCart(cartId int) (*entity.Cart, error)
}
