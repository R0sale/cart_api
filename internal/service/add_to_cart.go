package service

import "cart_api/internal/entity"

func (s *cartService) AddItemToCart(cartId int, item entity.NewItemDto) (*entity.CartItem, error) {
	newItem := entity.CartItem{
		CartId:  cartId,
		Product: item.Product,
		Price:   item.Price,
	}

	cartItem, err := s.repository.AddItemToCart(newItem)
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}
