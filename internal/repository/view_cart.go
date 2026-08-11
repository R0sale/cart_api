package repository

import "cart_api/internal/entity"

func (r repository) GetCartById(cartId int) (*entity.Cart, error) {
	query := `SELECT id, cart_id, product, price FROM carts_items
			WHERE cart_id = $1`

	var cartItems []entity.CartItem
	rows, err := r.db.Query(query, cartId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.CartItem
		if err := rows.Scan(&item.Id, &item.CartId, &item.Product, &item.Price); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if cartItems == nil {
		cartItems = []entity.CartItem{}
	}

	return &entity.Cart{Id: cartId, Items: cartItems}, nil
}
