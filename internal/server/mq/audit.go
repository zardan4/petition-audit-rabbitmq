package server

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
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

func (s *AuditServer) ListenAndServe() error {
	q, err := s.ch.QueueDeclare(
		viper.GetString("queues.logs"),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
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
	if err != nil {
		return err
	}

	forever := make(chan bool, 1)

	go func() {
		for m := range msgs {
			s.Log(m)
		}
	}()

	<-forever

	return nil
}

func (s *AuditServer) Log(del amqp.Delivery) error {
	return s.service.Log(del)
}
