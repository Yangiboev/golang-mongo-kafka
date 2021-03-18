package handlers

// MessageHandlerI is interface contains methods which MsgEventHandler has to implement
type productHandlerI interface {
	Handle(topic string, key []byte, data []byte) error
	OnTopic(topic string, key string, message *entity.product) error
	parse(data []byte) (*entity.product, error)
}
