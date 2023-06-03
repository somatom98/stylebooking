package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	// set up mongo client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://user:password@free.qko2zsp.mongodb.net/"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	servicesCollection := client.Database("stylebooking").Collection("services")

	// set up router
	router := gin.Default()

	// set up routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/products", func(c *gin.Context) {
		filter := bson.M{}
		cur, err := servicesCollection.Find(context.Background(), filter)
		if err != nil {
			panic(err)
		}
		defer cur.Close(context.Background())

		var services []map[string]interface{}
		for cur.Next(context.Background()) {
			var service map[string]interface{}
			if err := cur.Decode(&service); err != nil {
				panic(err)
			}
			services = append(services, service)
		}
		if err := cur.Err(); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, services)
	})

	// set up port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// run server
	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}
