package main

import (
	"fmt"

	"github.com/andresdb91/test-meli-magneto/api"
	"github.com/andresdb91/test-meli-magneto/db"
)

// func main() {
// 	mutante.IsMutant(os.Args[1:])
// }

func main() {
	fmt.Println("asd")
	db.SetupDB()
	api.Run()
}
