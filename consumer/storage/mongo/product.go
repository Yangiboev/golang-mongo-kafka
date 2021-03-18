package mongo

import (
	"context"

	"github.com/Yangiboev/golang-mongodb-kafka/config"
	"github.com/Yangiboev/golang-mongodb-kafka/storage/entity"
	"github.com/Yangiboev/golang-mongodb-kafka/storage/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepo struct {
	collection *mongo.Collection
}

func NewProductRepo(db *mongo.Database) repo.ProductStorageI {
	return &productRepo{
		collection: db.Collection(config.CollectionName)}
}

func (pr *productRepo) Create(product *entity.Product) (string, error) {
	_, err := pr.collection.InsertOne(
		context.Background(),
		product,
	)

	if err != nil {
		return "", err
	}

	return product.ID, nil
}

func (pr *productRepo) Get(id string) (*entity.Product, error) {
	var product entity.Product
	response := pr.collection.FindOne(
		context.Background(),
		bson.M{
			"id": id,
		})

	err := response.Decode(&product)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pr *productRepo) GetAll(page, limit int64, name string) ([]*entity.Product, int64, error) {

	var (
		products []*entity.Product
		filter   = bson.D{}
	)

	if name != "" {
		filter = append(filter, bson.E{Key: "name", Value: name})
	}

	opts := options.Find()

	skip := (page - 1) * limit
	opts.SetLimit(limit)
	opts.SetSkip(skip)
	opts.SetSort(bson.M{
		"created_at": -1,
	})

	count, err := pr.collection.CountDocuments(context.Background(), filter)

	if err != nil {
		return nil, 0, err
	}

	rows, err := pr.collection.Find(
		context.Background(),
		filter,
		opts,
	)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next(context.Background()) {
		var product entity.Product

		err := rows.Decode(&product)

		if err != nil {
			return nil, 0, err
		}

		products = append(products, &product)

	}

	return products, count, nil
}
