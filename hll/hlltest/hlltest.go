package hlltest

import (
	"github.com/go-redis/redis/v7"
)

var redisHLLTest = 2

// SetupMockRedis prepara una instancia de Redis de prueba
func SetupMockRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       redisHLLTest,
	})
	return client
}

// CleanupMockRedis limpia la instancia de prueba
func CleanupMockRedis(rc *redis.Client) {
	rc.FlushDB()
}
