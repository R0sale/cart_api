package repository

import "database/sql"

type repository struct {
	db *sql.DB
}

func NewRepository(currentDb *sql.DB) *repository {
	return &repository{
		db: currentDb,
	}
}
