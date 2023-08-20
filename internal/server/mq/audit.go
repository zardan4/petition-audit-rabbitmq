package server

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zardan4/petition-audit-rabbitmq/internal/service"
)

type AuditServer struct {
	ch      *amqp.Channel
	service *service.Service
}

func NewAuditServer(service *service.Service, ch *amqp.Channel) *AuditServer {
	return &AuditServer{
		service: service,
		ch:      ch,
	}
}

func (s *AuditServer) CreateQueueConsumer(qName string) (<-chan amqp.Delivery, error) {
	q, err := s.ch.QueueDeclare(
		qName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	msgs, err := s.ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	return msgs, err
}

func (s *AuditServer) Log(del amqp.Delivery) error {
	return s.service.Log(del)
}
