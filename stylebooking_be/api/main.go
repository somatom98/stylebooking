package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// set up router
	router := gin.Default()

	// set up routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"products": []string{"cut", "nails", "wax"},
		})
	})

	// set up port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// run server
	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}
