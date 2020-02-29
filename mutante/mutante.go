package mutante

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/andresdb91/test-meli-magneto/db"
)

var dbName = "mutantdb"
var dnaCollection = "dna"
var statsCollection = "stats"

// IsMutant verifica si una cadena de ADN corresponde a un mutante
func IsMutant(dna []string) bool {
	var v [6]int
	var h, i, j, r, l int
	var dr, dl [5]int
	var coin int

	dnaString := strings.Join(dna[:], "")
	dnaVect := []rune(dnaString)

	fmt.Printf("\nNew DNA sequence: %q\n", dna)

	for x := 0; x < len(dnaVect); x++ {
		i = x % 6
		j = x / 6
		r = i - j + 2
		l = i + j - 3

		if i == 0 {
			h = 0
			fmt.Printf("\nNew row: %d\n", j)
		}

		if 6-i+h >= 4 {
			fmt.Printf("Comparing: %q == %q\n", dnaVect[x], dnaVect[x+1])
			if dnaVect[x] == dnaVect[x+1] {
				if h < 2 {
					h++
					fmt.Printf("Horizontal matches %d\n", h)
				} else {
					coin++
					h = 0
					fmt.Printf("Sequence found, total: %d\n", coin)
				}
			} else {
				h = 0
			}
		}

		if 6-j+v[i] >= 4 {
			fmt.Printf("Comparing: %q == %q\n", dnaVect[x], dnaVect[x+6])
			if dnaVect[x] == dnaVect[x+6] {
				if v[i] < 2 {
					v[i]++
					fmt.Printf("Vertical matches %d\n", v[i])
				} else {
					coin++
					v[i] = 0
					fmt.Printf("Sequence found, total: %d\n", coin)
				}
			} else {
				v[i] = 0
			}
		}

		if (0 <= r) && (r <= 4) && ((35-6*int(math.Abs(float64(i-j)))-x)/7+dr[r] >= 3) {
			fmt.Printf("Comparing: %q == %q\n", dnaVect[x], dnaVect[x+7])
			if dnaVect[x] == dnaVect[x+7] {
				if dr[r] < 2 {
					dr[r]++
					fmt.Printf("Diagonal (right-down) matches %d\n", dr[r])
				} else {
					coin++
					dr[r] = 0
					fmt.Printf("Sequence found, total: %d\n", coin)
				}
			} else {
				dr[r] = 0
			}
		}

		if (0 <= l) && (l <= 4) && (((i+j)*6-x)/5+dl[l] >= 3) {
			fmt.Printf("Comparing: %q == %q\n", dnaVect[x], dnaVect[x+5])
			if dnaVect[x] == dnaVect[x+5] {
				if dl[l] < 2 {
					dl[l]++
					fmt.Printf("Diagonal (left-down) matches %d\n", dl[l])
				} else {
					coin++
					dl[l] = 0
					fmt.Printf("Sequence found, total: %d\n", coin)
				}
			} else {
				dl[l] = 0
			}
		}

		if coin > 1 {
			fmt.Printf("--------------------------------\nResult: Mutant\nSequences: %d\n\n", coin)
			SaveDNA(dnaString, true)
			return true
		}
	}

	fmt.Printf("--------------------------------\nResult: Human\nSequences: %d\n\n", coin)
	SaveDNA(dnaString, false)
	return false
}

// CheckDNA verifica si la cadena de ADN recibida ya esta en la base de datos
// Si la cadena existe también retorna el resultado del examen
func CheckDNA(dna string) (exists bool, result bool) {
	// dnaCol := db.Client.Database(dbName).Collection(dnaCollection)
	return false, false
}

// SaveDNA almacena una secuencia de ADN y el resultado de su examen en la base de datos
func SaveDNA(dna string, result bool) {
	dnaCol := db.Client.Database(dbName).Collection(dnaCollection)

	dnaObj := DNA{
		DNA:       dna,
		Result:    result,
		Timestamp: time.Now(),
	}

	res, err := dnaCol.InsertOne(nil, dnaObj)

	if err != nil {
		fmt.Printf("Error while storing DNA: %v", err)
	} else {
		fmt.Printf("Inserted document: %v", res)
	}
}
