package api

import (
	"github.com/gin-gonic/gin"
)

func checkMutant(c *gin.Context) {
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
