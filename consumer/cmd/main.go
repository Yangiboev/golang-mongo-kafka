package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Yangiboev/golang-mongodb-kafka/config"
	"github.com/Yangiboev/golang-mongodb-kafka/pkg/logger"
	"github.com/Yangiboev/golang-mongodb-kafka/storage/entity"
	"github.com/Yangiboev/golang-mongodb-kafka/sub/handlers"
	"github.com/Yangiboev/golang-mongodb-kafka/sub/kafka"
	"github.com/Yangiboev/golang-mongodb-kafka/sub/parsers"
	"github.com/Yangiboev/golang-mongodb-kafka/sub/topics"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "mongo_consumer")
	defer func() { _ = logger.Cleanup(log) }()

	log.Info("main: mongoConfig",
		logger.String("host", cfg.MongoHost),
		logger.Int("port", cfg.MongoPort))
	logger.String("user", cfg.MongoUser)

	credential := options.Credential{
		Username: cfg.MongoUser,
		Password: cfg.MongoPassword,
	}

	mongoString := fmt.Sprintf("mongodb://%s:%d", cfg.MongoHost, cfg.MongoPort)

	conn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoString).SetAuth(credential))

	if err != nil {
		log.Error("error connect mongo", logger.Error(err))
		return
	}

	db := conn.Database("mongo_consumer")

	log.Info("connected", logger.Any("db", db.Name()))

	lis, err := net.Listen("http", cfg.HTTPort)

	if err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return
	}

	ps := parsers.NewParsers(log)
	handler := handlers.EventHandler(&handlers.EventHandlerArgs{
		Logger:  log,
		Parsers: ps,
		DB:      db,
		Product: []entity.Product{},
	})

	consumer, err := kafka.NewConsumer(&cfg, log, handler)

	if err != nil {
		log.Error("error while consuming", logger.Error(err))
		return
	}

	go func() {
		if err := consumer.Subscribe([]string{topics.ProductInfoTopic}); err != nil {
			log.Error("error", logger.Error(err))
		}
	}()

	s := grpc.NewServer()
	reflection.Register(s)

	log.Info("main: server running", logger.Error(err))
	if err := s.Serve(lis); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
	}
}
