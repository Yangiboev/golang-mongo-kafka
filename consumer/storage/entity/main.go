package entity

// Responses
type GetProductResponse struct {
	Product *Product `json:"product"`
}

type GetAllProductsRequest struct {
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
	Name  string `json:"name"`
}
type GetAllProductsResponse struct {
	Products []*Product `json:"products"`
	Count    int64      `json:"count"`
}
type CreateResponse struct {
	ID string `json:"id"`
}

// Requests
type GetRequest struct {
	ID string `json:"id"`
}

type CreateProductRequest struct {
	Name   string    `json:"name" bson:"name"`
	Price  int64     `json:"price" bson:"price"`
	Photos []*Photos `json:"photos" bson:"photos"`
}
