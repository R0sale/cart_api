package service

func (s *cartService) RemoveFromCart(cartId int, itemId int) error {
	return s.repository.RemoveItemFromCart(cartId, itemId)
}
