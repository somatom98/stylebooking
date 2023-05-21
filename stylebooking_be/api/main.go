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
		c.JSON(http.StatusOK, []map[string]string{
			{
				"name":        "cut",
				"description": "cut hair",
				"price":       "10",
			},
			{
				"name":        "nails",
				"description": "cut nails",
				"price":       "5",
			},
			{
				"name":        "wax",
				"description": "wax hair",
				"price":       "15",
			},
			{
				"name":        "shave",
				"description": "shave hair",
				"price":       "10",
			},
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
