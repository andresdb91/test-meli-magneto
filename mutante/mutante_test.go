package mutante

import "testing"

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

	for _, c := range cases {
		got := IsMutant(c.in)
		if got != c.result {
			t.Errorf("isMutant(%q) == %t, deberia ser %t", c.in, got, c.result)
		}
	}
}