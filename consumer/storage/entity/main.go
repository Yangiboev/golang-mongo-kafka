package entity

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

type CreateResponse struct {
	ID string `json:"id"`
}

type CreateProductRequest struct {
	ID     string    `json:"id" bson:"id"`
	Name   string    `json:"name" bson:"name"`
	Price  int64     `json:"price" bson:"price"`
	Photos []*Photos `json:"photos" bson:"photos"`
}
