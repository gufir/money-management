package api

import (
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/utils"
	"github.com/labstack/echo/v4"
)

type Server struct {
	config utils.Config
	store  db.Store
	router *echo.Echo
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
		router: echo.New(),
	}

	server.SetupRouter()
	return server, nil
}

func (server *Server) SetupRouter() {
	server.router.POST("/create_users", server.CreateUser)
}

func (server *Server) Start(address string) error {
	return server.router.Start(address)
}

func ErrorResponse(message string) map[string]string {
	return map[string]string{"error": message}
}
