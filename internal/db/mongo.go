package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewMongoConn crete a new connection with mongodb
func NewMongoConn(ctx context.Context, uri string) (*mongo.Client, error) {
	cCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	client, err := mongo.Connect(cCtx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(cCtx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}
