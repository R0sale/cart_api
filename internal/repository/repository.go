package repository

import "database/sql"

type sqlRepo struct {
	db *sql.DB
}

func NewRepository(currentDb *sql.DB) *sqlRepo {
	return &sqlRepo{
		db: currentDb,
	}
}
