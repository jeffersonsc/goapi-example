package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewConn crete a new connection with mongodb
func NewConn(ctx context.Context, uri string) (*mongo.Client, error) {
	cCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	return mongo.Connect(cCtx, options.Client().ApplyURI(uri))
}
