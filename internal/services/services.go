package services

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/repository"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/types"
)

type Logger interface {
	Add(ctx context.Context, item types.LogItem) error
}

type Services struct {
	Logger Logger
}

func New(repo *repository.Repository) *Services {
	return &Services{
		Logger: NewLoggerService(repo.Logger),
	}
}
