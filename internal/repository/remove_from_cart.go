package repository

import repositoryerrors "cart_api/internal/errors/repository_errors"

func (r *sqlRepo) RemoveItemFromCart(cartId int, itemId int) error {
	query := `WITH checks AS (
		SELECT 
			EXISTS (SELECT 1 FROM carts WHERE id = $1) AS cart_exists
	),
	deleted AS (
	 	DELETE FROM carts_items WHERE cart_id = $1 AND id = $2
		RETURNING id
	)
		SELECT checks.cart_exists, COALESCE(deleted.id, -1) AS deleted_id
		FROM checks
		LEFT JOIN deleted ON true`

	var cartExists bool
	err := r.db.QueryRow(query, cartId, itemId).Scan(&cartExists, &itemId)
	if err != nil {
		return err
	}

	switch {
	case !cartExists:
		return repositoryerrors.ErrCartNotFound
	case itemId == -1:
		return repositoryerrors.ErrItemNotFound
	}

	return nil
}
