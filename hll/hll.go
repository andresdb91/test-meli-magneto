package hll

import (
	"github.com/go-redis/redis/v7"
)

// Client representa el cliente Redis
var Client *redis.Client

// SetupHLL configura la conexion a Redis para utilizar como HLL
func SetupHLL() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
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
