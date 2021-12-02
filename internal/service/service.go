package service

import (
	"github.com/1makarov/go-logger-rabbitmq-example/internal/repository"
)

type Service struct {
	*LoggerService
}

func New(repo *repository.Repository) *Service {
	return &Service{
		LoggerService: NewLoggerService(repo.LoggerRepository),
	}
}
