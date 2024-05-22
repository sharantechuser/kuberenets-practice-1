package config

import (
	"fmt"
	"os"
)

const (
	MongoDatabase = "gotodoservicetest"
)

type Config struct {
	MongoDatabase string
	MongoDbPort   string
}

var config Config

func init() {

	config = Config{
		MongoDatabase: os.Getenv("MONGO_DATABASE"),
		MongoDbPort:   os.Getenv("MONGO_PORT"),
	}
	fmt.Println("Configs: %v", config)

}

func GetConfig() Config {
	return config
}
