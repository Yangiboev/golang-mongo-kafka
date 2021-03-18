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
	storage storage.Pg
	product []entity.product
}

type EventHandlerArgs struct {
	Logger  logger.Logger
	Parsers *parsers.Parsers
	DB      *mongo.Database
	product []entity.product
}

func NewproductEventHandler(args *EventHandlerArgs) *EventHandler {
	return &EventHandler{
		logger:  args.Logger,
		parsers: args.Parsers,
		storage: storage.NewPg(args.DB),
		product: []entity.product{},
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
	case topics.productInfoTopic:
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
	case topics.productKey:
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

	case topics.LogKey:
		m.logger.Info("data", logger.Any("", string(data)))
		log, err := m.parsers.ParseLog(data)
		if err != nil {
			m.logger.Error("Could not parse the message with Log: err >", logger.Error(err))
			return err
		}
		err = m.InsertLog(log)

		if err != nil {
			m.logger.Error("Error while creating log", logger.Error(err))
			return err
		}
		m.logger.Info("message created", logger.Any("log", log))
		break
	case topics.UpdateKey:
		status, err := m.parsers.ParseUpdateStatus(data)
		m.logger.Info("status", logger.Any("status", status))
		if err != nil {
			m.logger.Error("Could not parse the message with UpdateStatus: err >", logger.Error(err))
			return err
		}
		err = m.UpdateStatus(status.MessageID, status.Status)

		if err != nil {
			m.logger.Error("Error while updating status", logger.Error(err))
			return err
		}
		break
	}
	return nil
}

func (m *EventHandler) Createproduct(product *entity.product) error {
	m.logger.Info("product", logger.Any("product", product))
	_, err := m.storage.product().Create(product)

	//bazaga yozguncha ko'p nagruzka tushib cpuni ko'p ishlatib yubormasligi uchun
	<-time.After(10 * time.Millisecond)

	return err
}

func (m *EventHandler) InsertLog(log *entity.Log) error {
	m.logger.Info("log", logger.Any("log", log))
	err := m.storage.product().InsertLog(log)

	//bazaga yozguncha ko'p nagruzka tushib cpuni ko'p ishlatib yubormasligi uchun
	<-time.After(10 * time.Millisecond)

	return err
}

func (m *EventHandler) UpdateStatus(messageID, status string) error {
	m.logger.Info("update status", logger.String("status", status))
	err := m.storage.product().UpdateStatus(messageID, status)

	return err
}
