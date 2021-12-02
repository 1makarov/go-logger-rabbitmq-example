package repository

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoggerRepository struct {
	db *mongo.Collection
}

func newLoggerRepository(db *mongo.Collection) *LoggerRepository {
	return &LoggerRepository{
		db: db,
	}
}

func (r *LoggerRepository) Add(ctx context.Context, v types.LogItem) error {
	_, err := r.db.InsertOne(ctx, v)

	return err
}
