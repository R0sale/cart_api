package service

import "cart_api/internal/entity"

func (s service) CreateCart() (*entity.Cart, error) {
	cart, err := s.repository.AddCart()
	if err != nil {
		return nil, err
	}

	return cart, nil
}
