package repository

import (
	"cart_api/internal/entity"
	"errors"
)

func (r *sqlRepo) AddItemToCart(item entity.CartItem) (*entity.CartItem, error) {
	query := `INSERT INTO carts_items (cart_id, product, price) SELECT $1, $2, $3
			WHERE EXISTS(SELECT 1 FROM carts WHERE id = $1) AND
			(SELECT COUNT(1) FROM carts_items WHERE cart_id = $1) < 5
			RETURNING id`

	var newItemId int
	err := r.db.QueryRow(query, item.CartId, item.Product, item.Price).Scan(&newItemId)
	if err != nil {
		return nil, errors.New("couldn't proceed with adding the item to the cart with id %d. Any of two rules is violated: cart already has 5 items, cart doesn't exists")
	}

	return &entity.CartItem{
		Id:      newItemId,
		CartId:  item.CartId,
		Product: item.Product,
		Price:   item.Price,
	}, nil
}
