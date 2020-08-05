package product

import (
	"context"
	"time"

	"github.com/jeffersonsc/natureapi"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const collectionName = "products"

// MongoRepository Instance mongo repository
type MongoRepository struct {
	ctx    context.Context
	client *mongo.Client
}

// GetCollectionName return as collection name
func GetCollectionName() string {
	return collectionName
}

// NewMongoRepository Instance new mongo repository
func NewMongoRepository(ctx context.Context, client *mongo.Client) Repository {
	return &MongoRepository{
		ctx:    ctx,
		client: client,
	}
}

// FindAll products in databse
func (r *MongoRepository) FindAll(filter map[string]interface{}) ([]*Product, error) {
	coll := r.client.Database(natureapi.DBNAME).Collection(GetCollectionName())

	ctx, cancel := context.WithTimeout(r.ctx, time.Second*10)
	defer cancel()

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	products := []*Product{}
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// Find product by id
func (r *MongoRepository) Find(id string) (*Product, error) {
	coll := r.client.Database(natureapi.DBNAME).Collection(GetCollectionName())

	ctx, cancel := context.WithTimeout(r.ctx, time.Second*10)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}
	product := Product{}
	err = coll.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Create a new product or update product has created
func (r *MongoRepository) Create(product *Product) error {
	coll := r.client.Database(natureapi.DBNAME).Collection(GetCollectionName())

	ctx, cancel := context.WithTimeout(r.ctx, time.Second*10)
	defer cancel()

	result, err := coll.InsertOne(ctx, product)

	if err != nil {
		return err
	}

	product.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}

// Update a project in database
func (r *MongoRepository) Update(product *Product) error {
	coll := r.client.Database(natureapi.DBNAME).Collection(GetCollectionName())

	ctx, cancel := context.WithTimeout(r.ctx, time.Second*10)
	defer cancel()

	filter := bson.M{"_id": bson.M{"$eq": product.ID}}
	_, err := coll.UpdateOne(ctx, filter, bson.M{"$set": product})
	if err != nil {
		return err
	}

	return nil
}

// Delete project database WARN!!!
func (r *MongoRepository) Delete(product *Product) error {
	coll := r.client.Database(natureapi.DBNAME).Collection(GetCollectionName())

	ctx, cancel := context.WithTimeout(r.ctx, time.Second*10)
	defer cancel()

	filter := bson.M{"_id": bson.M{"$eq": product.ID}}
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
