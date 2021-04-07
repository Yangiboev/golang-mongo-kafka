package handlers

import (
	"time"

	"github.com/Yangiboev/golang-mongodb-kafka/pkg/logger"
	"github.com/Yangiboev/golang-mongodb-kafka/storage"
	"github.com/Yangiboev/golang-mongodb-kafka/storage/entity"
	"github.com/Yangiboev/golang-mongodb-kafka/sub/parsers"
	"github.com/Yangiboev/golang-mongodb-kafka/sub/topics"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventHandler struct {
	logger  logger.Logger
	parsers *parsers.Parsers
	storage storage.StorageI
	product []*entity.Product
}

type EventHandlerArgs struct {
	Logger  logger.Logger
	Parsers *parsers.Parsers
	DB      *mongo.Database
	Product []*entity.Product
}

func NewProductEventHandler(args *EventHandlerArgs) *EventHandler {
	return &EventHandler{
		logger:  args.Logger,
		parsers: args.Parsers,
		storage: storage.NewProductStorage(args.DB),
		product: []*entity.Product{},
	}
}

func (m *EventHandler) Handle(topic string, key []byte, data []byte) error {
	messageKey := string(key)

	if err := m.OnTopic(topic, messageKey, data); err != nil {
		m.logger.Error("Error while handling event in OnTopic function: ", logger.Error(err))
		return err
	}

	return nil
}

// Publish(topic, key, message)
func (m *EventHandler) OnTopic(topic string, key string, data []byte) (err error) {
	switch topic {
	case topics.ProductInfoTopic:
		err := m.OnKey(key, data)

		if err != nil {
			return err
		}
		break
	}
	return nil
}

func (m *EventHandler) OnKey(key string, data []byte) error {
	switch key {
	case topics.ProductKey:
		m.logger.Info("data", logger.Any("", string(data)))
		product, err := m.parsers.Parseproduct(data)
		if err != nil {
			m.logger.Error("Could not parse the message with product: err >", logger.Error(err))
			return err
		}
		err = m.Createproduct(product)

		if err != nil {
			m.logger.Error("Error while creating product", logger.Error(err))
			return err
		}
		m.logger.Info("message created", logger.Any("product", product))
		break

	}
	return nil
}

func (m *EventHandler) Createproduct(product *entity.Product) error {
	m.logger.Info("product", logger.Any("product", product))
	_, err := m.storage.Product().Create(product)

	<-time.After(10 * time.Millisecond)

	return err
}
