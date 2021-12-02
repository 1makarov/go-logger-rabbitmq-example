package rabbitmq

import "github.com/streadway/amqp"

type RabbitMQ struct {
	channel *amqp.Channel
}

func New(channel *amqp.Channel) *RabbitMQ {
	return &RabbitMQ{
		channel: channel,
	}
}

func (rabbitmq *RabbitMQ) Consume(queue string, channel chan amqp.Delivery) error {
	stream, err := rabbitmq.channel.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for m := range stream {
			channel <- m
		}
		close(channel)
	}()

	return nil
}
