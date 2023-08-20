package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type MongoConfig struct {
	ConnectionLine string
	Database       string
}

type RabbitMQConfig struct {
	Password string
	User     string
}

type Config struct {
	DB MongoConfig
	MQ RabbitMQConfig
}

func New() (*Config, error) {
	godotenv.Load()

	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("rabbitmq", &cfg.MQ); err != nil {
		return nil, err
	}

	return cfg, nil
}
