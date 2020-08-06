package api

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/jeffersonsc/natureapi/api/handler"
)

// Server make as http server
type Server struct {
	ctx    context.Context
	Router *mux.Router
}

// NewServer create a new server
func NewServer(ctx context.Context) *Server {
	server := &Server{}
	router := mux.NewRouter()

	router.HandleFunc("/health", handler.Health)

	// Setup routes from here
	// router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

	// }).Methods(http.MethodGet)

	server.Router = router

	return server
}
