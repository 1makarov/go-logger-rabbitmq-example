package service

import (
	amqp "github.com/1makarov/go-logger-rabbitmq-example/internal/rabbit"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/repository"
)

type Service struct {
	*LoggerService
}

func New(repo *repository.Repository, rabbit *amqp.Client) *Service {
	return &Service{
		LoggerService: NewLoggerService(repo.LoggerRepo, rabbit),
	}
}
