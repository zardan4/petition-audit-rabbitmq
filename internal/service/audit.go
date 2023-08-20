package service

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	storage "github.com/zardan4/petition-audit-rabbitmq/internal/storage/mongo"
	"github.com/zardan4/petition-audit-rabbitmq/pkg/core/audit"
)

type AuditService struct {
	storage *storage.Storage
}

func NewAuditService(storage *storage.Storage) *AuditService {
	return &AuditService{
		storage: storage,
	}
}

func (s *AuditService) Log(msg amqp.Delivery) error {
	var item audit.LogItem

	err := json.Unmarshal(msg.Body, &item)

	if err != nil {
		return err
	}

	return s.storage.Insert(context.TODO(), item)
}
