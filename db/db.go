package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client representa el cliente a la base de datos MongoDB
var Client *mongo.Client

// DbName indica el nombre de la base de datos
var DbName = "mutantdb"

// SetupDB configura la conexion con el servidor MongoDB
func SetupDB() {
	user := os.Getenv("MONGODB_CREDS_USER")
	if user == "" {
		user = "mutant"
	}
	passwd := os.Getenv("MONGODB_CREDS_PWD")
	if passwd == "" {
		passwd = "mutant"
	}
	server := os.Getenv("MONGODB_SERVER_ADDR")
	if server == "" {
		server = "127.0.0.1"
	}
	port := os.Getenv("MONGODB_SERVER_PORT")
	if port == "" {
		port = "27017"
	}
	database := os.Getenv("MONGODB_DATABASE")
	if database == "" {
		database = "mutantdb"
	}

	dbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, passwd, server, port, database)
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		fmt.Printf("Error al conectar a mongodb: %v\n", err)
		panic(err)
	}

	col := Client.Database(DbName).Collection(DnaCollection)
	mod := mongo.IndexModel{
		Keys: bson.M{
			"dna": 1,
		},
		Options: nil,
	}
	col.Indexes().CreateOne(context.TODO(), mod)
}

// Find busca un documento en la base de datos
func Find(dna string) (exists bool, result bool, err error) {
	dnaCol := Client.Database(DbName).Collection(DnaCollection)

	var dnaObj DNA
	filter := bson.D{{
		"dna",
		dna,
	}}
	findOpts := options.Find()
	findOpts.SetLimit(1)

	cur, err := dnaCol.Find(context.TODO(), filter, findOpts)
	if err != nil {
		fmt.Printf("Error when fetching results: %v\n", err)
		return false, false, err
	}

	exists = cur.Next(context.TODO())
	if exists {
		cur.Decode(&dnaObj)
	}

	return exists, dnaObj.Result, nil
}

// Save guarda un documento en la base de datos
func Save(document DNA) error {
	dnaCol := Client.Database(DbName).Collection(DnaCollection)

	res, err := dnaCol.InsertOne(context.TODO(), document)

	if err != nil {
		fmt.Printf("Error while storing DNA: %v\n", err)
		return err
	}

	fmt.Printf("Inserted document: %v\n", res)
	return err

}
