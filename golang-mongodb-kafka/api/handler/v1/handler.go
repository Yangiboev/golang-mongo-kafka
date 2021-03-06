package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Yangiboev/golang-mongodb-kafka/config"
	"github.com/Yangiboev/golang-mongodb-kafka/pkg/logger"
	"github.com/Yangiboev/golang-mongodb-kafka/storage"
	"github.com/gin-gonic/gin"
)

const (
	//ErrorCodeInvalidURL ...
	ErrorCodeInvalidURL = "INVALID_URL"
	//ErrorCodeInvalidJSON ...
	ErrorCodeInvalidJSON = "INVALID_JSON"
	//ErrorCodeInternal ...
	ErrorCodeInternal = "INTERNAL"
	//ErrorCodeUnauthorized ...
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	//ErrorCodeAlreadyExists ...
	ErrorCodeAlreadyExists = "ALREADY_EXISTS"
	//ErrorCodeNotFound ...
	ErrorCodeNotFound = "NOT_FOUND"
	//ErrorCodeInvalidCode ...
	ErrorCodeInvalidCode = "INVALID_CODE"
	//ErrorBadRequest ...
	ErrorBadRequest = "BAD_REQUEST"
)

type handlerV1 struct {
	storage storage.StorageI
	log     logger.Logger
	cfg     config.Config
}

type HandlerV1Config struct {
	Storage storage.StorageI
	Logger  logger.Logger
	Cfg     config.Config
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		storage: c.Storage,
		log:     c.Logger,
		cfg:     c.Cfg,
	}
}

// Error handler

func HandleBadRequest(c *gin.Context, err error, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"error":   ErrorBadRequest,
		"message": msg,
	})
}
func HandleInternalServerError(c *gin.Context, err error, msg string) {
	fmt.Println(err)
	fmt.Println(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"error":   err,
		"message": msg,
	})
}

func ParsePageQueryParam(c *gin.Context) (int64, error) {
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 0, 10)
	if err != nil {
		return 0, err
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func ParseLimitQueryParam(c *gin.Context) (int64, error) {
	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "10"), 0, 10)
	if err != nil {
		return 0, err
	}
	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}
