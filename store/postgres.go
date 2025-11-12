package store

import (
	"database/sql"
	"go-boilerplate/config"
)

var postgresDB *sql.DB

func InitPostgres(cfg *config.DBConfig) (*sql.DB, error) {
	if postgresDB != nil {
		return postgresDB, nil
	}

	db, err := sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	postgresDB = db
	return postgresDB, nil
}
