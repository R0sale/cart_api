package repository

import (
	"cart_api/internal/entity"
	"fmt"
)

func (r *sqlRepo) AddCart() (*entity.Cart, error) {
	query := `INSERT INTO carts
			DEFAULT VALUES
			RETURNING id`

	var cart entity.Cart

	err := r.db.QueryRow(query).Scan(&cart.Id)
	if err != nil {
		return nil, fmt.Errorf("couldn't insert the cart to the db")
	}

	return &cart, nil
}
