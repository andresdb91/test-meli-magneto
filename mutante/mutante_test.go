package mutante

import (
	"testing"

	"github.com/andresdb91/test-meli-magneto/db/dbtest"
	"github.com/andresdb91/test-meli-magneto/hll"
	"github.com/andresdb91/test-meli-magneto/hll/hlltest"
)

func TestIsMutant(t *testing.T) {
	cases := []struct {
		in     []string
		result bool
	}{
		{
			[]string{
				"ATGCGA",
				"CAGTGC",
				"TTATTT",
				"AGACGG",
				"GCGTCA",
				"TCACTG",
			},
			false,
		},
		{
			[]string{
				"ATGCGA",
				"CAGTGC",
				"TTATGT",
				"AGAAGG",
				"CCCCTA",
				"TCACTG",
			},
			true,
		},
	}

	hll.Client = hlltest.SetupMockRedis()
	dbtest.SetupMockDb()

	for _, c := range cases {
		got := IsMutant(c.in)
		if got != c.result {
			t.Errorf("Mutant check incorrect (%q), got: %t, want: %t", c.in, got, c.result)
		}
	}

	dbtest.CleanupMockDb()
	hlltest.CleanupMockRedis(hll.Client)
}
