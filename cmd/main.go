package main

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/config"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/logger"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/postgres"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/rabbitmq"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/repository"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/services"
	"github.com/1makarov/go-logger-rabbitmq-example/pkg/signaler"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func main() {
	cfg := config.Init()
	ctx := context.Background()

	rabbitConn, err := rabbitmq.NewConnection(cfg.Rabbit)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer func() {
		if err = rabbitConn.Close(); err != nil {
			logrus.Errorln(err)
		}
	}()

	db, err := postgres.Open(cfg.DB)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer func() {
		if err = db.Close(); err != nil {
			logrus.Errorln(err)
		}
	}()

	rabbit := rabbitmq.New(rabbitConn)

	channel := make(chan amqp.Delivery)
	if err = rabbit.Consume(cfg.Rabbit.Queue, channel); err != nil {
		logrus.Errorln(err)
		return
	}

	repo := repository.New(db)
	service := services.New(repo)
	loggers := logger.New(channel, service.Logger)

	go func() {
		if err = loggers.Add(ctx); err != nil {
			logrus.Errorln(err)
			signaler.Signal()
		}
	}()

	logrus.Infoln("started")

	signaler.Wait()
}
