package main

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/config"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/db/mongo"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/logger"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/pkg/signaler"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/rabbitmq"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/repository"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {
	cfg := config.Init()
	ctx := context.Background()

	rabbitConn, err := rabbitmq.NewConnection(cfg.Rabbit)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer rabbitConn.Close()

	mongoClient, err := mongo.Open(ctx, cfg.DB)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer mongoClient.Disconnect(ctx)

	db := mongoClient.Database(cfg.DB.Name)
	rabbit := rabbitmq.New(rabbitConn)

	channel := make(chan amqp.Delivery)
	if err = rabbit.Consume(cfg.Rabbit.Queue, channel); err != nil {
		logrus.Errorln(err)
		return
	}

	repo := repository.New(db)
	services := service.New(repo)
	loggers := logger.New(channel, services.LoggerService)

	go func() {
		if err = loggers.Add(); err != nil {
			logrus.Errorln(err)
			signaler.Signal()
		}
	}()

	logrus.Infoln("logger started")

	signaler.Wait()
}
