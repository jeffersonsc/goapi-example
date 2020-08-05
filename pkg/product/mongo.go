package product

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
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
func (r *MongoRepository) FindAll(filter map[string]interface{}) ([]*Product, error) { return nil, nil }

// Find product by id
func (r *MongoRepository) Find(id string) (*Product, error) { return nil, nil }

// Create a new product or update product has created
func (r *MongoRepository) Create(product *Product) error { return nil }

// Update a project in database
func (r *MongoRepository) Update(product *Product) error { return nil }

// Delete project database WARN!!!
func (r *MongoRepository) Delete(product *Product) error { return nil }
