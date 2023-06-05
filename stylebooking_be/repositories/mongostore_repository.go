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

func (r *MongoStoreRepository) AddService(ctx context.Context, storeId string, service models.Service) error {
	oid, err := primitive.ObjectIDFromHex(storeId)
	if err != nil {
		return err
	}

	service.ID = primitive.NewObjectID()

	filter := bson.M{"_id": oid}
	update := bson.M{"$push": bson.M{"services": service}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return sb.ErrStoreNotFound{Id: storeId}
	}
	return nil
}

func (r *MongoStoreRepository) UpdateService(ctx context.Context, storeId string, serviceId string, service models.Service) error {
	oid, err := primitive.ObjectIDFromHex(storeId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid, "services._id": serviceId}
	update := bson.M{"$set": bson.M{"services.$": service}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return sb.ErrServiceNotFound{Id: storeId, StoreId: serviceId}
	}
	return nil
}

func (r *MongoStoreRepository) DeleteService(ctx context.Context, storeId string, serviceId string) error {
	oid, err := primitive.ObjectIDFromHex(storeId)
	if err != nil {
		return err
	}
	soid, err := primitive.ObjectIDFromHex(serviceId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$pull": bson.M{"services": bson.M{"_id": soid}}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return sb.ErrServiceNotFound{Id: storeId, StoreId: serviceId}
	}
	return nil
}
