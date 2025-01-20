package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/token"
	"github.com/gufir/money-management/utils"
)

type Server struct {
	config utils.Config
	store  db.Store
	token  token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	server := &Server{
		config: config,
		store:  store,
		token:  tokenMaker,
	}

	server.SetupRouter()

	return server, nil
}

func (server *Server) SetupRouter() {
	router := gin.Default()
	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.LoginUser)
	router.POST("/users/renew", server.renewAccessToken)

	// AuthRoutes := router.Group("/").Use(AuthMiddleware(server.token))

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Helper for generating error responses.
func errorResponse(message string) gin.H {
	return gin.H{"error": message}
}
