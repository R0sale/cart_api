package service

import "cart_api/internal/entity"

func (s *cartService) UpdateCartItem(cartId int, itemId int, updatedItem entity.NewItemDto) (*entity.CartItem, error) {
	cartItem := entity.CartItem{
		CartId:  cartId,
		Product: updatedItem.Product,
		Price:   updatedItem.Price,
	}

	item, err := s.repository.UpdateCartItem(cartId, itemId, cartItem)
	if err != nil {
		return nil, err
	}

	return item, nil
}
