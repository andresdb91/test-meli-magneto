package hll

import (
	"fmt"
	"math"
	"testing"

	"github.com/go-redis/redis/v7"
)

var redisHLLTest = 2

func TestHLLAddHuman(t *testing.T) {

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       redisHLLTest,
	})

	tolerance := 0.01

	cases := []struct {
		in    string
		set   string
		wantH int64
		wantM int64
	}{
		{
			in:    "humano_1",
			set:   "human",
			wantH: 1,
			wantM: 0,
		},
		{
			in:    "humano_2",
			set:   "human",
			wantH: 2,
			wantM: 0,
		},
		{
			in:    "humano_1",
			set:   "human",
			wantH: 2,
			wantM: 0,
		},
		{
			in:    "mutante_1",
			set:   "mutant",
			wantH: 2,
			wantM: 1,
		},
	}

	for _, c := range cases {
		var e float64
		AddToHLL(c.set, c.in)

		countH := GetCountHLL("human")
		e = math.Abs(float64(countH-c.wantH)) / float64(c.wantH)
		if e > tolerance {
			t.Errorf("Returned human count doesn't match, got: %d, want: %d +/- %4.3f\n", countH, c.wantH, tolerance)
		}

		countM := GetCountHLL("mutant")
		e = math.Abs(float64(countM-c.wantM)) / float64(c.wantM)
		if e > tolerance {
			t.Errorf("Returned mutant count doesn't match, got: %d, want: %d +/- %4.3f\n", countM, c.wantM, tolerance)
		}

		fmt.Printf("human: %d\nmutant: %d\n", countH, countM)
	}

	client.FlushDB()
}
