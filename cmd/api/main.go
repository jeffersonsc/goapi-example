package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/jeffersonsc/natureapi/api"
	"github.com/jeffersonsc/natureapi/api/handler"
	"github.com/jeffersonsc/natureapi/internal/db"
	"github.com/jeffersonsc/natureapi/pkg/product"
)

func main() {
	log := log.New(os.Stdout, "[main]", 0)

	port := os.Getenv("PORT")
	ctx := context.TODO()

	dbConn, err := db.NewMongoConn(ctx, os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal(err)
	}

	redisConn, err := db.NewRedisConn(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	productRepo := product.NewMongoRepository(ctx, dbConn)
	productCache := product.NewStorageCache(redisConn)

	api := api.NewServer(ctx)

	// Setup routes from here
	handler.NewProductHandler(productRepo, productCache, api.Router)

	log.Println("Server running on port", port)

	log.Fatal(http.ListenAndServe(net.JoinHostPort("", port), api.Router))
}
