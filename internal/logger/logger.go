package logger

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/service"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/types"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
)

type Logger struct {
	channel <-chan amqp.Delivery
	logger  *service.LoggerService
}

func New(channel <-chan amqp.Delivery, logger *service.LoggerService) *Logger {
	return &Logger{
		channel: channel,
		logger:  logger,
	}
}

func (l *Logger) Add() error {
	for v := range l.channel {
		var item types.LogItem

		if err := jsoniter.Unmarshal(v.Body, &item); err != nil {
			return err
		}

		if err := l.logger.Add(context.Background(), item); err != nil {
			return err
		}

		if err := v.Ack(true); err != nil {
			return err
		}
	}

	return nil
}
