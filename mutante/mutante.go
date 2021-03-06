package mutante

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/andresdb91/test-meli-magneto/db"
	"github.com/andresdb91/test-meli-magneto/hll"
)

// IsMutant verifica si una cadena de ADN corresponde a un mutante
func IsMutant(dna []string) bool {
	var v [6]int
	var h, i, j, r, l, lh int
	var dr, dl [5]int
	var coin int

	dnaString := strings.Join(dna[:], "")
	dnaVect := []rune(dnaString)

	fmt.Printf("\nNew DNA sequence: %q\n\n", dna)

	exists, result := checkDNA(dnaString)
	if exists {
		return result
	}

	for x := 0; x < len(dnaVect); x++ {
		i = x % 6
		j = x / 6
		r = i - j + 2
		l = i + j - 3

		if i > j {
			lh = i + j
		} else {
			lh = i - j
		}

		if i == 0 {
			h = 0
			fmt.Printf("New row: %d\n", j)
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

		if (0 <= l) && (l <= 4) && (((lh)*6-x)/5+dl[l] >= 3) {
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
			saveDNA(dnaString, true)
			return true
		}
	}

	fmt.Printf("--------------------------------\nResult: Human\nSequences: %d\n\n", coin)
	saveDNA(dnaString, false)
	return false
}

func checkDNA(dna string) (exists bool, result bool) {
	exists, result, _ = db.Find(dna)
	return
}

func saveDNA(dna string, result bool) {
	var set string
	if result {
		set = "mutant"
	} else {
		set = "human"
	}

	hll.AddToHLL(set, dna)

	dnaObj := db.DNA{
		DNA:       dna,
		Result:    result,
		Timestamp: time.Now(),
	}

	db.Save(dnaObj)
}
