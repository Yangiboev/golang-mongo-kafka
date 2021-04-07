package config

import (
	"os"

	"github.com/spf13/cast"
)

var (
	CollectionName = "product"
)

// Config ...
type Config struct {
	Port          string
	Environment   string // develop, staging, production
	MongoHost     string
	MongoPort     int
	MongoDatabase string
	MongoPassword string
	MongoUser     string
	LogLevel      string
	KafkaHost     string
	KafkaPort     int
	KafkaTopic    string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.MongoHost = cast.ToString(getOrReturnDefault("MONGO_HOST", "localhost"))
	c.MongoPort = cast.ToInt(getOrReturnDefault("MONGO_PORT", 27017))
	c.MongoDatabase = cast.ToString(getOrReturnDefault("MONGO_DATABASE", "mongo_consumer"))
	c.MongoUser = cast.ToString(getOrReturnDefault("MONGO_USER", "dell"))
	c.MongoPassword = cast.ToString(getOrReturnDefault("MONGO_PASSWORD", "icon"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.Port = cast.ToString(getOrReturnDefault("RPC_PORT", ":8081"))
	c.KafkaHost = cast.ToString(getOrReturnDefault("KAFKA_HOST", "localhost"))
	c.KafkaPort = cast.ToInt(getOrReturnDefault("KAFKA_PORT", 9092))
	c.KafkaTopic = cast.ToString(getOrReturnDefault("KAFKA_TOPIC", "msg.receive"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
