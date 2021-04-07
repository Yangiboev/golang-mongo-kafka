package handlers

import "github.com/Yangiboev/golang-mongodb-kafka/storage/entity"

// MessageHandlerI is interface contains methods which MsgEventHandler has to implement
type ProductEventHandlerI interface {
	Handle(topic string, key []byte, data []byte) error
	OnTopic(topic string, key string, message *entity.Product) error
	Parse(data []byte) (*entity.Product, error)
}
