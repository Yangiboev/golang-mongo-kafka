package v1

import (
	"net/http"

	"github.com/Yangiboev/golang-mongodb-kafka/api/models"
	"github.com/Yangiboev/golang-mongodb-kafka/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router /v1/product/{product_id} [get]
// @Summary Get Product
// @Description API for getting a product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product_id path string true "product_id"
// @Success 200 {object} models.GetProductResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetProduct(c *gin.Context) {
	productID := c.Param("product_id")
	_, err := uuid.Parse(productID)
	if err != nil {
		HandleInternalServerError(c, err, "product_id format is invalid format!")
		return
	}
	product, err := h.storage.Product().Get(productID)

	if err != nil {
		HandleInternalServerError(c, err, "error while getting product")
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    product,
		})

}

//@Router /v1/product [post]
//@Summary Create product
//@Description API for creating product
//@Tags product
//@Accept json
//@Produce json
//@Param Product body models.CreateProductRequest  true "product"
// @Success 200 {object} models.CreateResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) CreateProduct(c *gin.Context) {

	var (
		product models.Product
	)
	err := c.ShouldBindJSON(&product)
	if err != nil {
		HandleBadRequest(c, err, "error while binding product to json")
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		HandleInternalServerError(c, err, "error while generating uuid")
		return
	}

	product.ID = id.String()
	resp, err := h.storage.Product().Create(
		&product)

	if err != nil {
		h.log.Error("error while creating product", logger.Error(err))
		HandleInternalServerError(c, err, "Error while creating  product")
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			"data":    resp,
		})
}

// @Router /v1/product [get]
// @Summary Get All Products
// @Description API for getting all Products
// @Tags product
// @Accept  json
// @Produce  json
// @Param name query string false "name"
// @Success 200 {object} models.GetAllProductsResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) GetAllProducts(c *gin.Context) {
	page, err := ParsePageQueryParam(c)
	if err != nil {
		HandleBadRequest(c, err, "Error while parsing page")
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		HandleBadRequest(c, err, "Error while parsing page")
		return
	}

	products, count, err := h.storage.Product().GetAll(page, limit, c.Query("name"))

	if err != nil {
		HandleInternalServerError(c, err, "Error while getting all products")
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"count":   count,
			"data":    products,
		})

}

//@Router /v1/product/{product_id} [put]
//@Summary Update product
//@Description API for creating product
//@Tags product
//@Accept json
//@Produce json
// @Param product_id path string true "product_id"
//@Param Product body models.CreateProductRequest  true "product"
// @Success 200 {object} models.CreateResponse
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalServerError
func (h *handlerV1) UpdateProduct(c *gin.Context) {
	var (
		product   models.Product
		productID string
	)
	productID = c.Param("product_id")
	_, err := uuid.Parse(productID)
	if err != nil {
		HandleInternalServerError(c, err, "product_id format is an invalid format!")
		return
	}

	err = c.ShouldBindJSON(&product)
	if err != nil {
		HandleBadRequest(c, err, "Error while binding product!")
		return
	}
	product.ID = productID
	// resp, err := h.storage.Product().Update(&product)

	if err != nil {
		HandleBadRequest(c, err, "Error while updating product")
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"success": true,
			// "data":    resp,
		})
}
