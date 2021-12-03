package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Open(ctx context.Context, cfg Config) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(cfg.url())

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return client, client.Ping(ctx, nil)
}
