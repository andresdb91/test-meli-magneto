package db

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client representa el cliente a la base de datos MongoDB
var Client *mongo.Client

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

	dbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, passwd, server, port)
	var err error
	Client, err = mongo.Connect(nil, options.Client().ApplyURI(dbURI))
	if err != nil {
		fmt.Printf("Error al conectar a mongodb: %v\n", err)
	}
}
