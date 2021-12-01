package repository

import "go.mongodb.org/mongo-driver/mongo"

const loggerCollection = "logs"

type Repository struct {
	*LoggerRepo
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		LoggerRepo: NewLoggerRepo(db),
	}
}
