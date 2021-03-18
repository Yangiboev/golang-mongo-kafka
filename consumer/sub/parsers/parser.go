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

func (p *Parsers) Parseproduct(data []byte) (*entity.product, error) {
	product := entity.product{}

	if err := json.Unmarshal(data, &product); err != nil {
		p.logger.Error("Error while unmarshal byte to proto (product), ", logger.Error(err))
		return nil, err
	}
	return &product, nil
}

func (p *Parsers) ParseLog(data []byte) (*entity.Log, error) {
	log := entity.Log{}

	if err := json.Unmarshal(data, &log); err != nil {
		p.logger.Error("Error while unmarshal byte to proto (Log), ", logger.Error(err))
		return nil, err
	}
	return &log, nil
}

func (p *Parsers) ParseUpdateStatus(data []byte) (*entity.UpdateStatus, error) {
	status := entity.UpdateStatus{}

	if err := json.Unmarshal(data, &status); err != nil {
		p.logger.Error("error while unmarshal byte to proto (UpdateStatus)", logger.Error(err))
		return nil, err
	}

	return &status, nil
}
