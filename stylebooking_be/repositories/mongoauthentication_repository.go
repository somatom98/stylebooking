package repositories

import (
	"context"

	sb "github.com/somatom98/stylebooking/stylebooking_be"
	"github.com/somatom98/stylebooking/stylebooking_be/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAuthenticationRepository struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewMongoAuthenticationRepository(client *mongo.Client) *MongoAuthenticationRepository {
	collection := client.Database("stylebooking").Collection("authentication")
	return &MongoAuthenticationRepository{
		collection: collection,
		client:     client,
	}
}

func (r *MongoAuthenticationRepository) GetByCustomerId(ctx context.Context, customerId string) (models.Authentication, error) {
	var authentication models.Authentication
	filter := bson.M{"customer_id": customerId}
	err := r.collection.FindOne(ctx, filter).Decode(&authentication)
	if err != nil {
		return models.Authentication{}, err
	}
	return authentication, nil
}

func (r *MongoAuthenticationRepository) Create(ctx context.Context, authentication models.Authentication) error {
	_, err := r.collection.InsertOne(ctx, authentication)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoAuthenticationRepository) Delete(ctx context.Context, customerId string) error {
	filter := bson.M{"customer_id": customerId}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return sb.ErrCustomerNotFound{Id: customerId}
	}
	return nil
}

func (r *MongoAuthenticationRepository) Update(ctx context.Context, customerId string, authentication models.Authentication) error {
	filter := bson.M{"customer_id": customerId}
	update := bson.M{"$set": authentication}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return sb.ErrCustomerNotFound{Id: customerId}
	}
	return nil
}
