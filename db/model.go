package db

import "time"

// DnaCollection indica el nombre de la coleccion donde se almacenan las secuencias de ADN
var DnaCollection = "dna"

// StatsCollection indica la coleccion donde se consolidan las estadisticas
var StatsCollection = "stats"

// DNA struct para guardar secuencias de ADN y el resultado de su examen
type DNA struct {
	DNA       string
	Result    bool
	Timestamp time.Time
}

// Stats struct para consolidar estadisticas periodicamente
type Stats struct {
	CountMutant int
	CountHuman  int
	Timestamp   time.Time
}
