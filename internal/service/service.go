package service

import (
	amqp "github.com/rabbitmq/amqp091-go"
	storage "github.com/zardan4/petition-audit-rabbitmq/internal/storage/mongo"
)

type Audit interface {
	Log(msg amqp.Delivery) error
}

type Service struct {
	Audit
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Audit: NewAuditService(storage),
	}
}
