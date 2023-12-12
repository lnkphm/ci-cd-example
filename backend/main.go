package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello world"})
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "pass"})
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", getHello)
	router.GET("/health", getHealth)
	return router
}

func main() {
	router := setupRouter()
	router.Run("127.0.0.1:8080")
}
