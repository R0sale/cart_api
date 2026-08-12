package repository

import (
	"cart_api/internal/entity"
	repositoryerrors "cart_api/internal/errors/repository_errors"
)

func (r *sqlRepo) AddItemToCart(item entity.CartItem) (*entity.CartItem, error) {
	query := `WITH checks AS (
		SELECT
			EXISTS (SELECT 1 FROM carts WHERE id = $1) AS cart_exists) AS cart_exists,
			(SELECT COUNT(1) FROM carts_items WHERE cart_id = $1) < 5 AS items_full
	)
	inserted AS (
		INSERT INTO carts_items (cart_id, product, price)
		SELECT $1, $2, $3
		FROM checks
		WHERE cart_exists AND items_full
		RETURNING id
	)
		SELECT checks.cart_exists, checks.items_full, inserted.id
		FROM checks
		LEFT JOIN inserted ON true`

	var cartExists, items_full bool
	err := r.db.QueryRow(query, item.CartId, item.Product, item.Price).Scan(&cartExists, &items_full, &item.Id)
	if err != nil {
		return nil, err
	}

	switch {
	case !cartExists:
		return nil, repositoryerrors.ErrCartNotFound
	case !items_full:
		return nil, repositoryerrors.ErrCartIsFull
	}

	return &item, nil
}
