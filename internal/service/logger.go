package service

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/repository"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/types"
)

type LoggerService struct {
	repo *repository.LoggerRepository
}

func NewLoggerService(repo *repository.LoggerRepository) *LoggerService {
	return &LoggerService{
		repo: repo,
	}
}

func (l *LoggerService) Add(ctx context.Context, item types.LogItem) error {
	return l.repo.Add(ctx, item)
}
