package rabbit

import (
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

type Client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	Stream  <-chan amqp.Delivery

	wg *sync.WaitGroup

	queue   string
	connURL string

	done    chan bool
	stopped bool
}

const retryDelay = 5 * time.Second

func Open(cfg Config) (*Client, error) {
	client := &Client{
		queue:   cfg.Queue,
		connURL: cfg.url(),
		wg:      new(sync.WaitGroup),
		done:    make(chan bool),
	}

	client.connect()

	go client.reconnect()

	return client, nil
}

func (c *Client) connect() {
	c.wg.Add(1)
	defer c.wg.Done()

	for {
		if c.stopped {
			return
		}

		var err error

		c.conn, err = amqp.Dial(c.connURL)
		if err != nil {
			log.Println(err)
			time.Sleep(retryDelay)
			continue
		}

		c.channel, err = c.conn.Channel()
		if err != nil {
			log.Println(err)
			time.Sleep(retryDelay)
			continue
		}

		queue, err := c.channel.QueueDeclare(
			c.queue,
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Println(err)
			time.Sleep(retryDelay)
			continue
		}

		c.Stream, err = c.channel.Consume(
			queue.Name,
			"",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Println(err)
			time.Sleep(retryDelay)
			continue
		}

		return
	}
}

func (c *Client) reconnect() {
	c.wg.Add(1)
	defer c.wg.Done()

	graceful := make(chan *amqp.Error, 1)
	errs := c.channel.NotifyClose(graceful)

	for {
		if c.stopped {
			return
		}

		select {
		case <-errs:
			graceful = make(chan *amqp.Error, 1)
			c.connect()
			log.Println("reconnect success")
			errs = c.channel.NotifyClose(graceful)
		case <-c.done:
			return
		}
	}
}

func (c *Client) Close() error {
	close(c.done)
	c.stopped = true

	c.wg.Wait()
	return c.conn.Close()
}
