package repo

import "github.com/Yangiboev/golang-mongodb-kafka/storage/entity"

type ProductStorageI interface {
	Create(product *entity.Product) (string, error)
	Get(id string) (*entity.Product, error)
	GetAll(page, limit int64, name string) ([]*entity.Product, int64, error)
}
