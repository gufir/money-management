package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/gufir/money-management/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Routes
	router.POST("/users", server.CreateUser)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Helper for generating error responses.
func errorResponse(message string) gin.H {
	return gin.H{"error": message}
}
