package repositories

import (
	"context"

	sb "github.com/somatom98/stylebooking/stylebooking_be"
	"github.com/somatom98/stylebooking/stylebooking_be/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStoreRepository struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewMongoStoreRepository(client *mongo.Client) sb.StoreRepository {
	storesCollection := client.Database("stylebooking").Collection("stores")
	return &MongoStoreRepository{
		collection: storesCollection,
		client:     client,
	}
}

func (r *MongoStoreRepository) GetAll(ctx context.Context) ([]models.Store, error) {
	filter := bson.M{}
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var stores []models.Store
	for cur.Next(ctx) {
		var store models.Store
		if err := cur.Decode(&store); err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return stores, nil
}

func (r *MongoStoreRepository) GetById(ctx context.Context, id string) (models.Store, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Store{}, err
	}
	filter := bson.M{"_id": oid}

	var store models.Store
	err = r.collection.FindOne(ctx, filter).Decode(&store)
	if err != nil {
		return models.Store{}, err
	}
	return store, nil
}

func (r *MongoStoreRepository) Create(ctx context.Context, store models.Store) error {
	_, err := r.collection.InsertOne(ctx, store)
	if err != nil {
		return err
	}
	return nil
}
