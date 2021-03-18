package storage

import (
	"github.com/Yangiboev/golang-mongodb-kafka/storage/mongo"
	"github.com/Yangiboev/golang-mongodb-kafka/storage/repo"
	db "go.mongodb.org/mongo-driver/mongo"
)

type Pg interface {
	Product() repo.ProductStorageI
}

type pg struct {
	productRepo repo.ProductStorageI
}

func NewPg(db *db.Database) Pg {
	return &pg{
		productRepo: mongo.NewProductRepo(db),
	}
}

func (p pg) product() repo.productStorageI {
	return p.productRepo
}
