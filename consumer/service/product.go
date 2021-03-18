package service

import (
	"context"

	"github.com/Yangiboev/golang-mongodb-kafka/pkg/helper"
	"github.com/Yangiboev/golang-mongodb-kafka/pkg/logger"
	"github.com/Yangiboev/golang-mongodb-kafka/storage"
	"github.com/Yangiboev/golang-mongodb-kafka/storage/entity"
	db "go.mongodb.org/mongo-driver/mongo"
)

type productService struct {
	logger  logger.Logger
	storage storage.StorageI
}

func NewProductService(db *db.Database, log logger.Logger) *productService {
	return &productService{
		storage: storage.NewProductStorage(db),
		logger:  log,
	}
}

func (ps *productService) Create(ctx context.Context, req *entity.Product) (*entity.CreateResponse, error) {
	ID, err := ps.storage.Product().Create(req)

	if err != nil {
		ps.logger.Error("error while creating product", logger.Error(err))
		return nil, helper.HandleError(ps.logger, err, "error while creating product", req)
	}
	return &entity.CreateResponse{
		ID: ID,
	}, nil
}

func (ps *productService) Get(ctx context.Context, req *entity.GetRequest) (*entity.GetProductResponse, error) {

	response, err := ps.storage.Product().Get(req.ID)

	if err != nil {
		return nil, helper.HandleError(ps.logger, err, "error while getting product", req)
	}
	return &entity.GetProductResponse{
		Product: response,
	}, nil
}
func (ps *productService) GetAll(ctx context.Context, req *entity.GetAllProductsRequest) (*entity.GetAllProductsResponse, error) {

	response, count, err := ps.storage.Product().GetAll(req.Page, req.Limit, req.Name)

	if err != nil {
		return nil, helper.HandleError(ps.logger, err, "error while getting product", req)
	}
	return &entity.GetAllProductsResponse{
		Count:    count,
		Products: response,
	}, nil
}
