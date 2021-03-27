package main

import (
	"context"
	"fmt"

	"github.com/Yangiboev/golang-mongodb-kafka/api"
	"github.com/Yangiboev/golang-mongodb-kafka/config"
	"github.com/Yangiboev/golang-mongodb-kafka/pkg/logger"
	"github.com/Yangiboev/golang-mongodb-kafka/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	log  logger.Logger
	strg storage.StorageI
	cfg  config.Config
)

func initDependencies() {
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "mongo-golang")

	log.Info("main: MONGO",
		logger.String("Host", cfg.MongoHost),
		logger.Int("Port", cfg.MongoPort),
		logger.String("Database", cfg.MongoDatabase),
	)

	credential := options.Credential{
		Username: cfg.MongoUser,
		Password: cfg.MongoPassword,
	}
	mongoString := fmt.Sprintf("mongodb://%s:%d", cfg.MongoHost, cfg.MongoPort)

	mongoConn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoString).SetAuth(credential))

	if err != nil {
		log.Error("error to connect to mongo database", logger.Error(err))
	}
	connDB := mongoConn.Database(cfg.MongoDatabase)

	log.Info("Connected to MongoDB", logger.Any("database: ", connDB.Name()))
	strg = storage.NewProductStorage(connDB)
}

func main() {
	initDependencies()
	server := api.New(api.RouterOptions{
		Config:  cfg,
		Log:     log,
		Storage: strg,
	})

	err := server.Run(cfg.Port)

	if err != nil {
		log.Error("Something went wrong", logger.Error(err))
		panic(err)
	}

}
