package parsers

import (
	"encoding/json"

	"github.com/Yangiboev/golang-mongodb-kafka/pkg/logger"
	"github.com/Yangiboev/golang-mongodb-kafka/storage/entity"
)

// Parsers - all parsers for available platforms
type Parsers struct {
	logger logger.Logger
}

// NewParsers ...
func NewParsers(logger logger.Logger) *Parsers {
	return &Parsers{
		logger: logger,
	}
}

func (p *Parsers) ParseProduct(data []byte) (*entity.Product, error) {
	product := entity.Product{}

	if err := json.Unmarshal(data, &product); err != nil {
		p.logger.Error("Error while unmarshal byte to proto (product), ", logger.Error(err))
		return nil, err
	}
	return &product, nil
}
