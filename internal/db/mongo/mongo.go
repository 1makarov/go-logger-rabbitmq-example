package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client

	cfg Config
}

func Open(ctx context.Context, cfg Config) (*DB, error) {
	opts := options.Client().ApplyURI(cfg.url())

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &DB{client: client, cfg: cfg}, client.Ping(ctx, nil)
}

func (db *DB) Connect() *mongo.Database {
	return db.client.Database(db.cfg.Name)
}

func (db *DB) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}
