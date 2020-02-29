package mutante

import (
	"testing"

	"github.com/andresdb91/test-meli-magneto/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupMockDb() {
	db.Client, _ = mongo.Connect(nil, options.Client().ApplyURI("mongodb://localhost:27017"))
	db.DbName = "mutantdb_test"
}

func cleanupMockDb() {
	db.Client.Database("mutantdb_test").Collection(db.DnaCollection).Drop(nil)
}

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

	setupMockDb()

	for _, c := range cases {
		got := IsMutant(c.in)
		if got != c.result {
			t.Errorf("Mutant check incorrect (%q), got: %t, want: %t", c.in, got, c.result)
		}
	}

	cleanupMockDb()
}
