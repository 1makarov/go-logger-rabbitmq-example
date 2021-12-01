package main

import (
	"context"
	"github.com/1makarov/go-logger-rabbitmq-example/config"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/db/mongo"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/pkg/signaler"
	amqp "github.com/1makarov/go-logger-rabbitmq-example/internal/rabbit"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/repository"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Init()
	ctx := context.Background()

	rabbit, err := amqp.Open(cfg.Rabbit)
	if err != nil {
		logrus.Fatalln(err)
	}

	db, err := mongo.Open(ctx, cfg.DB)
	if err != nil {
		logrus.Fatalln(err)
	}

	repo := repository.New(db.Connect())
	services := service.New(repo, rabbit)

	go func() {
		if err = services.HandleStream(ctx); err != nil {
			logrus.Errorln(err)
			signaler.Signal()
		}
	}()

	logrus.Infoln("logger started")

	signaler.Wait()

	if err = rabbit.Close(); err != nil {
		logrus.Errorln(err)
	}

	if err = db.Close(ctx); err != nil {
		logrus.Errorln(err)
	}
}
