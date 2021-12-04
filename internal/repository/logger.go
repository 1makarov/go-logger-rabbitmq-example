package repository

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/types"
	"github.com/jmoiron/sqlx"
)

type LoggerRepository struct {
	db *sqlx.DB
}

func newLoggerRepository(db *sqlx.DB) *LoggerRepository {
	return &LoggerRepository{
		db: db,
	}
}

func (r *LoggerRepository) Add(ctx context.Context, v types.LogItem) error {
	_, err := r.db.ExecContext(ctx, `

	insert into logs values ($1, $2, $3, $4)
	
	`, v.Entity, v.Action, v.UserID, v.Timestamp)

	return err
}
