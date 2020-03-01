package api

import (
	"fmt"
	"math"
	"net/http"

	"github.com/andresdb91/test-meli-magneto/hll"
	"github.com/andresdb91/test-meli-magneto/mutante"
	"github.com/gin-gonic/gin"
)

func checkMutant(c *gin.Context) {
	type DnaArray struct {
		Dna []string `binding:"required"`
	}
	data := new(DnaArray)
	c.BindJSON(data)
	fmt.Printf("%q\n", data)

	if len(data.Dna) != 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect DNA sequence format",
		})
		return
	}

	validDNA := []rune("ATGC")
	for _, e := range data.Dna {
		if len(e) != 6 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Incorrect DNA sequence format",
			})
			return
		}
		dnaRunes := []rune(e)
		for _, d := range dnaRunes {
			match := false
			for _, v := range validDNA {
				if d == v {
					match = true
					break
				}
			}
			if !match {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Incorrect DNA sequence format",
				})
				return
			}
		}
	}

	result := mutante.IsMutant(data.Dna)

	if result {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusForbidden)
	}
}

func getStats(c *gin.Context) {
	countM, errM := hll.GetCountHLL("mutant")
	countH, errH := hll.GetCountHLL("human")
	ratio := float64(countM) / float64(countH)

	if errM != nil || errH != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	response := gin.H{
		"count_mutant_dna": countM,
		"count_human_dna":  countH,
		"ratio":            math.Round(ratio*10) / 10,
	}
	c.JSON(http.StatusOK, response)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/mutant", checkMutant)
	router.GET("/stats", getStats)

	return router
}

// Run configura e inicia el servidor HTTP Gin
func Run() {
	router := setupRouter()
	router.Run()
}
