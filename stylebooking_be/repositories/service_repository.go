package repositories

import (
	"context"

	sb "github.com/somatom98/stylebooking/stylebooking_be"
	"github.com/somatom98/stylebooking/stylebooking_be/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoServiceRepository struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewMongoServiceRepository(client *mongo.Client) sb.ServiceRepository {
	servicesCollection := client.Database("stylebooking").Collection("services")
	return &MongoServiceRepository{
		collection: servicesCollection,
		client:     client,
	}
}

func (r *MongoServiceRepository) GetAll(context context.Context) ([]models.Service, error) {
	filter := bson.M{}
	cur, err := r.collection.Find(context, filter)
	if err != nil {
		panic(err)
	}
	defer cur.Close(context)

	var services []models.Service
	for cur.Next(context) {
		var service models.Service
		if err := cur.Decode(&service); err != nil {
			panic(err)
		}
		services = append(services, service)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	return services, nil
}
