package service

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/rabbit"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/repository"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/types"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"log"
)

type LoggerService struct {
	repo   *repository.LoggerRepo
	rabbit *rabbit.Client
}

func NewLoggerService(repo *repository.LoggerRepo, r *rabbit.Client) *LoggerService {
	return &LoggerService{
		repo:   repo,
		rabbit: r,
	}
}

func (s *LoggerService) HandleStream(ctx context.Context) error {
	handleAction := func(message amqp.Delivery) error {
		var item types.LogItem

		if err := jsoniter.Unmarshal(message.Body, &item); err != nil {
			log.Println(err)
			if err = message.Ack(true); err != nil {
				return err
			}
			return nil
		}

		if err := s.repo.Add(ctx, item); err != nil {
			log.Println(err)
			return err
		}

		if err := message.Ack(true); err != nil {
			return err
		}

		return nil
	}

	for {
		v, ok := <-s.rabbit.Stream
		if !ok {
			continue
		}
		if err := handleAction(v); err != nil {
			return err
		}
	}
}
