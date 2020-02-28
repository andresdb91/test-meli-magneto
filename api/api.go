package api

import (
	"fmt"
	"net/http"

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
	result := mutante.IsMutant(data.Dna)

	if result {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusForbidden)
	}
}

func getStats() {}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/mutant", checkMutant)

	return router
}

// Run configura e inicia el servidor HTTP Gin
func Run(routerParam string) {
	router := setupRouter()
	router.Run(routerParam)
}
