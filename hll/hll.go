package hll

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
)

// Client representa el cliente Redis
var Client *redis.Client

var redisHLLProd = 1

// SetupHLL configura la conexion a Redis para utilizar como HLL
func SetupHLL() {
	host := os.Getenv("REDIS_SERVER_ADDR")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("REDIS_SERVER_PORT")
	if port == "" {
		port = "6379"
	}
	passwd := os.Getenv("REDIS_CREDS_PWD")

	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: passwd,
		DB:       redisHLLProd,
	})
}

// AddToHLL agrega un valor al set del HLL
func AddToHLL(set string, value string) {
	Client.PFAdd(set, value)
}

// GetCountHLL obtiene la cardinalidad del set
func GetCountHLL(set string) int64 {
	count, err := Client.PFCount(set).Result()
	if err != nil {
		panic(err)
	}
	return count
}
