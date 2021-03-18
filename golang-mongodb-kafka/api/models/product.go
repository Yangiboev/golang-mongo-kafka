package models

import "time"

type Photos struct {
	ID string `json:"id" bson:"id"`
}

type Product struct {
	ID        string    `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Price     int64     `json:"price" bson:"price"`
	Photos    []*Photos `json:"photos" bson:"photos"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type CreateProductRequest struct {
	ID     string    `json:"id" bson:"id"`
	Name   string    `json:"name" bson:"name"`
	Price  int64     `json:"price" bson:"price"`
	Photos []*Photos `json:"photos" bson:"photos"`
}

type GetProductResponse struct {
	Product *Product `json:"product" bson:"product"`
}
type GetAllProductsRequest struct {
	Name  string `json:"name" bson:"name"`
	Page  int64  `json:"page" bson:"page"`
	Limit int64  `json:"limit" bson:"limit"`
}
type GetAllProductsResponse struct {
	Products []*Product `json:"products" bson:"products"`
	Count    int64      `json:"count" bson:"count"`
}
