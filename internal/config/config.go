package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type MongoConfig struct {
	ConnectionLine string
	Database       string
}

type MQNames struct {
	MQLogName string
}

type RabbitMQConfig struct {
	Password string
	User     string

	MQNames MQNames
}

type Config struct {
	DB MongoConfig
	MQ RabbitMQConfig
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("rabbitmq", &cfg.MQ); err != nil {
		return nil, err
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName("main")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg.MQ.MQNames.MQLogName = viper.GetString("queues.logs")

	return cfg, nil
}
