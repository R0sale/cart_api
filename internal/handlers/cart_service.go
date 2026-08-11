package handlers

import "cart_api/internal/entity"

type CartService interface {
	CreateCart() (*entity.Cart, error)
}
