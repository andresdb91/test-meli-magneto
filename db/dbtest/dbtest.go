package dbtest

import (
	"github.com/andresdb91/test-meli-magneto/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SetupMockDb prepara una base de datos de prueba
func SetupMockDb() {
	db.Client, _ = mongo.Connect(nil, options.Client().ApplyURI("mongodb://localhost:27017"))
	db.DbName = "mutantdb_test"
}

// CleanupMockDb limpia la base de datos de prueba
func CleanupMockDb() {
	db.Client.Database("mutantdb_test").Drop(nil)
}
