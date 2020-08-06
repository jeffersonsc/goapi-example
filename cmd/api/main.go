package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/jeffersonsc/natureapi/api"
)

func main() {
	port := os.Getenv("PORT")
	ctx := context.TODO()

	api := api.NewServer(ctx)
	log.Println("Server running on port", port)

	log.Fatal(http.ListenAndServe(net.JoinHostPort("", port), api.Router))
}
