package server

import (
	"github.com/spf13/viper"
	"github.com/zardan4/petition-audit-rabbitmq/internal/config"
	errhand "github.com/zardan4/petition-audit-rabbitmq/pkg/error"
)

type Server struct {
	auditServer *AuditServer
}

func NewServer(auditServer *AuditServer) *Server {
	return &Server{
		auditServer: auditServer,
	}
}

func (s *Server) ListenAndServe(qNames config.MQNames) error {
	var err error

	forever := make(chan bool)

	logChan, err := s.auditServer.CreateQueueConsumer(viper.GetString("queues.logs"))
	if err != nil {
		return err
	}

	go func() {
		for log := range logChan {
			err := s.auditServer.Log(log)
			errhand.ErrorOnError(err, "")
		}
	}()

	<-forever

	return nil
}
