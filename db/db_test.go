package db

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestSave(t *testing.T) {
	Client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	DbName = "mutantdb_test"

	dna := DNA{
		DNA:       "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		Result:    true,
		Timestamp: time.Now(),
	}

	err := Save(dna)
	if err != nil {
		t.Errorf("Error when storing document: %v", err)
	}

	Client.Database("mutantdb_test").Drop(context.TODO())
}

func TestSaveAndFind(t *testing.T) {
	Client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	DbName = "mutantdb_test"

	dna := DNA{
		DNA:       "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		Result:    true,
		Timestamp: time.Now(),
	}

	err := Save(dna)
	if err != nil {
		t.Errorf("Error when storing document: %v", err)
	}

	cases := []struct {
		dna    string
		exists bool
		result bool
	}{
		{
			dna:    "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			exists: true,
			result: true,
		},
		{
			dna:    "AAAAAAAAAAAAAAAAAAAAAAAAAATAAAAAAAAA",
			exists: false,
			result: false,
		},
	}

	for _, c := range cases {
		exists, result, err := Find(c.dna)
		if err != nil {
			t.Fatalf("Error while retrieving from database: %v", err)
		}

		if exists != c.exists {
			if c.exists {
				t.Errorf("DNA sample should exist, got: %v, want: %v", exists, c.exists)
			} else {
				t.Errorf("DNA sample should not exist, got: %v, want: %v", exists, c.exists)
			}
		} else if c.exists && result != c.result {
			t.Errorf("Expected result doesn't match, got %v, want %v", result, c.result)
		}
	}

	Client.Database("mutantdb_test").Drop(context.TODO())
}
