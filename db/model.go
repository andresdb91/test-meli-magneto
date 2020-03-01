package db

import "time"

// DnaCollection indica el nombre de la coleccion donde se almacenan las secuencias de ADN
var DnaCollection = "dna"

// DNA struct para guardar secuencias de ADN y el resultado de su examen
type DNA struct {
	DNA       string
	Result    bool
	Timestamp time.Time
}
