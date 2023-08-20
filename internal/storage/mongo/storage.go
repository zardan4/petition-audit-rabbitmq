package storage

import (
	"context"

	"github.com/zardan4/petition-audit-rabbitmq/pkg/core/audit"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	logsCollection = "logs"
)

type Audit interface {
	Insert(ctx context.Context, logItem audit.LogItem) error
}

type Storage struct {
	Audit
}

func NewStorage(db *mongo.Database) *Storage {
	return &Storage{
		Audit: NewAuditStorage(db),
	}
}
