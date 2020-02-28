package main

import (
	"github.com/andresdb91/test-meli-magneto/api"
)

// func main() {
// 	mutante.IsMutant(os.Args[1:])
// }

func main() {
	api.Run(":8080")
}
