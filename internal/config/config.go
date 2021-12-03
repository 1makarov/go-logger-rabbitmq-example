package config

import (
	"github.com/1makarov/go-logger-rabbitmq-example/internal/mongodb"
	"github.com/1makarov/go-logger-rabbitmq-example/internal/rabbitmq"
	"os"
)

type Config struct {
	DB     mongodb.Config
	Rabbit rabbitmq.Config
}

func Init() *Config {
	var cfg Config

	setFromEnv(&cfg)

	return &cfg
}

func setFromEnv(cfg *Config) {
	cfg.DB.Name = os.Getenv("DB_NAME")
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.Port = os.Getenv("DB_PORT")

	cfg.Rabbit.Queue = os.Getenv("RABBIT_QUEUE")
	cfg.Rabbit.User = os.Getenv("RABBIT_USER")
	cfg.Rabbit.Host = os.Getenv("RABBIT_HOST")
	cfg.Rabbit.Password = os.Getenv("RABBIT_PASSWORD")
	cfg.Rabbit.Port = os.Getenv("RABBIT_PORT")
}
