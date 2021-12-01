package repository

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoggerRepo struct {
	db *mongo.Collection
}

func NewLoggerRepo(db *mongo.Database) *LoggerRepo {
	return &LoggerRepo{
		db: db.Collection(loggerCollection),
	}
}

func (r *LoggerRepo) Add(ctx context.Context, v types.LogItem) error {
	_, err := r.db.InsertOne(ctx, v)

	return err
}
