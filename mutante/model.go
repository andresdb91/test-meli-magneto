package mutante

import "time"

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
	Ratio       float32
	Timestamp   time.Time
}
