package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	stylebooking "github.com/somatom98/stylebooking/stylebooking_be"
	"github.com/somatom98/stylebooking/stylebooking_be/config"
	"github.com/somatom98/stylebooking/stylebooking_be/controllers"
	"github.com/somatom98/stylebooking/stylebooking_be/models"
	"github.com/somatom98/stylebooking/stylebooking_be/repositories"
	"github.com/somatom98/stylebooking/stylebooking_be/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var conf config.Configuration
var router *gin.Engine
var mongoClient *mongo.Client
var serviceRepository stylebooking.ServiceRepository
var storeRepository stylebooking.StoreRepository
var storeService stylebooking.StoreService
var storeController *controllers.StoreController

func init() {
	conf = config.GetConfig()

	// set up mongo client
	var err error
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(conf.Mongo.ConnectionString))
	if err != nil {
		panic(err)
	}
	if err := mongoClient.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}

	// set up repositories
	serviceRepository = repositories.NewMongoServiceRepository(mongoClient)
	storeRepository = repositories.NewMongoStoreRepository(mongoClient)

	// set up services
	storeService = services.NewStoreService(storeRepository)

	// set up controllers
	storeController = controllers.NewStoreController(storeService)

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

	router.GET("/stores", storeController.GetAll)

	router.GET("/stores/:id", storeController.GetById)

	router.POST("/stores", storeController.Create)

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

	// run server
	log.Fatal(router.Run(conf.App.Addr))
}
