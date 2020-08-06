package product

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product structure entity
type Product struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty"`
	Description  string             `bson:"description,omitempty"`
	Price        float64            `bson:"price,omitempty"`
	Images       []string           `bson:"images,omitempty"`
	CurrencyCode string             `bson:"currency_code,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty"`
}
