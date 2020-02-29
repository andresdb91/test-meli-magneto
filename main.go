package main

import (
	"github.com/andresdb91/test-meli-magneto/api"
	"github.com/andresdb91/test-meli-magneto/db"
	"github.com/andresdb91/test-meli-magneto/hll"
)

// func main() {
// 	mutante.IsMutant(os.Args[1:])
// }

func main() {
	db.SetupDB()
	hll.SetupHLL()
	api.Run()
}
