package handlers

import "cart_api/internal/entity"

type cartService interface {
	CreateCart() (*entity.Cart, error)
	ViewCart(cartId int) (*entity.Cart, error)
	AddItemToCart(cartId int, item entity.NewItemDto) (*entity.CartItem, error)
	RemoveFromCart(cartId int, itemId int) error
}
