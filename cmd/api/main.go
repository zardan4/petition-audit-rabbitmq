package main

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/zardan4/petition-audit-rabbitmq/internal/config"
	server "github.com/zardan4/petition-audit-rabbitmq/internal/server/mq"
	"github.com/zardan4/petition-audit-rabbitmq/internal/service"
	storage "github.com/zardan4/petition-audit-rabbitmq/internal/storage/mongo"
	errhand "github.com/zardan4/petition-audit-rabbitmq/pkg/core/error"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		logrus.Fatal(err)
	}

	// mq
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@localhost:15672/", cfg.MQ.User, cfg.MQ.Password))
	errhand.FailOnError(err, "Error while initializing amqp connection")

	defer conn.Close()

	// create channel
	ch, err := conn.Channel()
	errhand.FailOnError(err, "Failed to open a channel")

	defer ch.Close()

	// init db connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.ApplyURI(cfg.DB.ConnectionLine)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logrus.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.Fatal(err)
	}

	// levels
	db := client.Database(cfg.DB.Database)
	storage := storage.NewStorage(db)
	service := service.NewService(storage)

	auditSrv := server.NewAuditServer(service, ch)
	srv := server.NewServer(auditSrv)

	fmt.Printf("Server started at %s", time.Now())

	// serve all mqs
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
