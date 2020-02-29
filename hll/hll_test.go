package hll

import "testing"

func TestRedisSetup(t *testing.T) {
	err := SetupHLL()

	if err != nil {
		t.Fatalf("Error al conectar a Redis: %s", err)
	}
}
