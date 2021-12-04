package postgres

import (
	"github.com/jmoiron/sqlx"
)

func Open(cfg Config) (*sqlx.DB, error) {
	return sqlx.Connect("pgx", cfg.string())
}
