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

// PopulateMockRedis carga datos ficticios en Redis para realizar pruebas
func PopulateMockRedis(rc *redis.Client) (int64, int64) {
	mockData := []struct {
		set   string
		value string
	}{
		{
			set:   "human",
			value: "a",
		},
		{
			set:   "human",
			value: "b",
		},
		{
			set:   "human",
			value: "c",
		},
		{
			set:   "human",
			value: "d",
		},
		{
			set:   "mutant",
			value: "x",
		},
		{
			set:   "mutant",
			value: "y",
		},
		{
			set:   "mutant",
			value: "z",
		},
	}

	var countH, countM int64

	for _, d := range mockData {
		rc.PFAdd(d.set, d.value)
		if d.set == "human" {
			countH++
		} else {
			countM++
		}
	}

	return countH, countM
}
