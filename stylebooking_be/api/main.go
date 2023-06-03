package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	stylebooking "github.com/somatom98/stylebooking/stylebooking_be"
	"github.com/somatom98/stylebooking/stylebooking_be/config"
	"github.com/somatom98/stylebooking/stylebooking_be/models"
	"github.com/somatom98/stylebooking/stylebooking_be/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var router *gin.Engine
var mongoClient *mongo.Client
var serviceRepository stylebooking.ServiceRepository

func init() {
	c := config.GetConfig()

	// set up mongo client
	var err error
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(c.Mongo.ConnectionString))
	if err != nil {
		panic(err)
	}
	if err := mongoClient.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}

	// set up repositories
	serviceRepository = repositories.NewMongoServiceRepository(mongoClient)

	// set up router
	router = gin.Default()
}

func main() {
	// set up routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/services", func(c *gin.Context) {
		services, err := serviceRepository.GetAll(context.Background())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, services)
	})

	router.GET("/services/:id", func(c *gin.Context) {
		id := c.Param("id")
		service, err := serviceRepository.GetById(context.Background(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, service)
	})

	router.POST("/services", func(c *gin.Context) {
		var service models.Service
		if err := c.ShouldBindJSON(&service); err != nil {
			panic(err)
		}

		if err := serviceRepository.Create(context.Background(), service); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, service)
	})

	// set up port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// run server
	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}
