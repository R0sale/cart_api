package service

import "cart_api/internal/entity"

func (s service) ViewCart(cartId int) (*entity.Cart, error) {
	cart, err := s.repository.GetCartById(cartId)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
