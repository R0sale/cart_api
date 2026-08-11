package db

import (
	"cart_api/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConfigureDb(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("couldn't open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("couldn't ping database: %w", err)
	}

	return db, nil
}
