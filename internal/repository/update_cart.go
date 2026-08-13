package repository

import (
	"cart_api/internal/entity"
	"fmt"
)

func (r *sqlRepo) UpdateCartItem(cartId int, itemId int, updatedItem entity.CartItem) (*entity.CartItem, error) {
	query := `WITH checks AS (
		SELECT
			EXISTS (SELECT 1 FROM carts WHERE id = $1) AS cart_exists,
			EXISTS (SELECT 1 FROM carts_items WHERE cart_id = $1 AND id = $2) AS item_exists
	),
	updated AS (
		UPDATE carts_items
		SET product = $3, price = $4
		WHERE cart_id = $1 AND id = $2
		RETURNING id, cart_id, product, price
	)
	SELECT checks.cart_exists, checks.item_exists, updated.id, updated.cart_id, updated.product, updated.price
	FROM checks
	LEFT JOIN updated ON true`

	var cartExists, itemExists bool
	err := r.db.QueryRow(query, cartId, itemId, updatedItem.Product, updatedItem.Price).Scan(&cartExists, &itemExists, &updatedItem.Id, &updatedItem.CartId, &updatedItem.Product, &updatedItem.Price)
	if err != nil {
		return nil, err
	}

	if !cartExists {
		return nil, fmt.Errorf("cart with ID %d does not exist", cartId)
	}
	if !itemExists {
		return nil, fmt.Errorf("item with ID %d does not exist in cart %d", itemId, cartId)
	}

	return &updatedItem, nil
}
