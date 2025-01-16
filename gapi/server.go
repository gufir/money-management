package gapi

import (
	"fmt"

	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/token"
	"github.com/gufir/money-management/utils"
)

type Server struct {
	pb.UnimplementedMoneyManagementServer
	config utils.Config
	store  db.Store
	token  token.Maker
}

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

	return server, nil
}
