package repository

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type Logger interface {
	Add(ctx context.Context, v types.LogItem) error
}

type Repository struct {
	Logger Logger
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		Logger: newLoggerRepository(db.Collection(loggerCollection)),
	}
}
