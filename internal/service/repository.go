package service

import "cart_api/internal/entity"

type repository interface {
	AddCart() (*entity.Cart, error)
	GetCartById(cartId int) (*entity.Cart, error)
}
