package main

import (
	"os"

	"github.com/andresdb91/test-meli-magneto/mutante"
)

func main() {
	mutante.IsMutant(os.Args[1:])
}
