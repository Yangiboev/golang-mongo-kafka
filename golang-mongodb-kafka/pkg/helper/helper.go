package helper

import (
	"github.com/Yangiboev/golang-mongodb-kafka/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(log logger.Logger, err error, message string, req interface{}) error {

	if err == mongo.ErrNoDocuments {
		log.Error(message+", Not Found", logger.Error(err), logger.Any("req", req))
		return status.Error(codes.NotFound, "Not Found")
	} else if err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(codes.Internal, message)
	}
	return nil
}
