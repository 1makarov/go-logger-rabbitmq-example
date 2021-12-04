package repository

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/types"
	"github.com/jmoiron/sqlx"
)

type Logger interface {
	Add(ctx context.Context, v types.LogItem) error
}

type Repository struct {
	Logger Logger
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Logger: newLoggerRepository(db),
	}
}
