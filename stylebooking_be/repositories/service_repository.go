package repositories

import (
	"context"

	sb "github.com/somatom98/stylebooking/stylebooking_be"
	"github.com/somatom98/stylebooking/stylebooking_be/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *MongoServiceRepository) GetAll(ctx context.Context) ([]models.Service, error) {
	filter := bson.M{}
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var services []models.Service
	for cur.Next(ctx) {
		var service models.Service
		if err := cur.Decode(&service); err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return services, nil
}

func (r *MongoServiceRepository) GetById(ctx context.Context, id string) (models.Service, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Service{}, err
	}
	filter := bson.M{"_id": oid}

	var service models.Service
	err = r.collection.FindOne(ctx, filter).Decode(&service)
	if err != nil {
		return models.Service{}, err
	}

	return service, nil
}

func (r *MongoServiceRepository) Create(ctx context.Context, service models.Service) error {
	_, err := r.collection.InsertOne(ctx, service)
	if err != nil {
		return err
	}
	return nil
}
