package main

import (
	"github.com/andresdb91/test-meli-magneto/api"
	"github.com/andresdb91/test-meli-magneto/db"
)

// func main() {
// 	mutante.IsMutant(os.Args[1:])
// }

func main() {
	db.SetupDB()
	api.Run()
}
