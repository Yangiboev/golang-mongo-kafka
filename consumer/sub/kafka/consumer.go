package kafka

import (
	"fmt"
	"time"

	"github.com/Yangiboev/golang-mongodb-kafka/config"
	"github.com/Yangiboev/golang-mongodb-kafka/pkg/logger"
	"github.com/Yangiboev/golang-mongodb-kafka/sub/handlers"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	kafkaReader         *kafka.Consumer
	logger              logger.Logger
	productEventHandler *handlers.EventHandler
}

func NewConsumer(cfg *config.Config, logger logger.Logger, productHandler *handlers.EventHandler) (*Consumer, error) {
	connString := fmt.Sprintf("%s:%d", cfg.KafkaHost, cfg.KafkaPort)

	productReader, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        connString,
		"group.id":                 "product",
		"auto.offset.reset":        "latest",
		"allow.auto.create.topics": true,
	})

	if err != nil {
		return nil, err
	}

	err = productReader.SubscribeTopics([]string{"product"}, nil)

	if err != nil {
		return nil, err
	}
	return &Consumer{
		kafkaReader:         productReader,
		logger:              logger,
		productEventHandler: productHandler,
	}, nil
}

func (c *Consumer) Subscribe(topics []string) error {
	c.logger.Info(fmt.Sprintf(">>> Kafka Consumer is started for %v", topics))
	err := c.kafkaReader.SubscribeTopics(topics, nil)

	if err != nil {
		return err
	}

	for {
		msg, err := c.kafkaReader.ReadMessage(-1)
		fmt.Println(msg)
		if err != nil {
			c.logger.Error(fmt.Sprintf("Error while consuming message %v", msg), logger.Error(err))
		}

		err = c.productEventHandler.Handle(*msg.TopicPartition.Topic, msg.Key, msg.Value)

		if err != nil {
			c.logger.Error("Error while handling product events", logger.Error(err))
		}

		c.logger.Info(fmt.Sprintf(">>> Consumed message at topic/partition %v", msg.TopicPartition))
		<-time.After(10 * time.Millisecond)
	}
}
