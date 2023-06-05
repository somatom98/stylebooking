package repositories

import (
	"context"
	"errors"

	sb "github.com/somatom98/stylebooking/stylebooking_be"
	"github.com/somatom98/stylebooking/stylebooking_be/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCustomerRepository struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewMongoCustomerRepository(client *mongo.Client) sb.CustomerRepository {
	storesCollection := client.Database("stylebooking").Collection("customers")
	return &MongoCustomerRepository{
		collection: storesCollection,
		client:     client,
	}
}

func (r *MongoCustomerRepository) GetAll(ctx context.Context) ([]models.Customer, error) {
	filter := bson.M{}
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var customers []models.Customer
	for cur.Next(ctx) {
		var customer models.Customer
		if err := cur.Decode(&customer); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *MongoCustomerRepository) GetById(ctx context.Context, id string) (models.Customer, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Customer{}, err
	}
	filter := bson.M{"_id": oid}

	var customer models.Customer
	err = r.collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func (r *MongoCustomerRepository) GetByEmail(ctx context.Context, email string) (models.Customer, error) {
	filter := bson.M{"email": email}

	var customer models.Customer
	err := r.collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func (r *MongoCustomerRepository) Create(ctx context.Context, customer models.Customer) (string, error) {
	result, err := r.collection.InsertOne(ctx, customer)
	if err != nil {
		return "", err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("error converting oid")
	}
	return oid.Hex(), nil
}

func (r *MongoCustomerRepository) Update(ctx context.Context, id string, customer models.Customer) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": customer}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return sb.ErrCustomerNotFound{Id: id}
	}
	return nil
}

func (r *MongoCustomerRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return sb.ErrCustomerNotFound{Id: id}
	}
	return nil
}
