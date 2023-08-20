package storage

import (
	"context"

	"github.com/zardan4/petition-audit-rabbitmq/pkg/core/audit"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuditStorage struct {
	db *mongo.Database
}

func NewAuditStorage(db *mongo.Database) *AuditStorage {
	return &AuditStorage{
		db: db,
	}
}

func (s *AuditStorage) Insert(ctx context.Context, logItem audit.LogItem) error {
	_, err := s.db.Collection(logsCollection).InsertOne(ctx, logItem)
	return err
}
