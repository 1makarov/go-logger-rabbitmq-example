package logger

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/services"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/types"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
)

type Logger struct {
	channel <-chan amqp.Delivery
	logger  services.Logger
}

func New(channel <-chan amqp.Delivery, logger services.Logger) *Logger {
	return &Logger{
		channel: channel,
		logger:  logger,
	}
}

func (l *Logger) Add(ctx context.Context) error {
	for v := range l.channel {
		var item types.LogItem

		if err := jsoniter.Unmarshal(v.Body, &item); err != nil {
			return err
		}

		if err := l.logger.Add(ctx, item); err != nil {
			return err
		}

		if err := v.Ack(true); err != nil {
			return err
		}
	}

	return nil
}
