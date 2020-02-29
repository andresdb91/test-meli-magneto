package hll

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

// var Client

// SetupHLL configura la conexion a Redis para utilizar como HLL
func SetupHLL() error {
	Client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)

	return err
}
