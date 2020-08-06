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
	server := &Server{
		ctx: ctx,
	}
	router := mux.NewRouter()

	router.HandleFunc("/health", handler.Health)

	server.Router = router

	return server
}
