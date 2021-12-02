package repository

import "go.mongodb.org/mongo-driver/mongo"

const loggerCollection = "logs"

type Repository struct {
	*LoggerRepository
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		LoggerRepository: newLoggerRepository(db.Collection(loggerCollection)),
	}
}
