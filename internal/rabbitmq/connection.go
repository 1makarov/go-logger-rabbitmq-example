package rabbitmq

import (
	"github.com/streadway/amqp"
)

func NewConnection(cfg Config) (*amqp.Channel, error) {
	url := cfg.url()

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare(
		cfg.Queue,
		false,
		false,
		false,
		false,
		nil,
	)

	return channel, err
}
