package gapi

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/token"
	"github.com/gufir/money-management/utils"
)

type Server struct {
	pb.UnimplementedMoneyManagementServer
	config      utils.Config
	store       db.Store
	token       token.Maker
	redisClient *redis.Client
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
		DB:   0,
	})

	// Test Redis connection
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("could not connect to Redis: %v", err)
	}

	server := &Server{
		config:      config,
		store:       store,
		token:       tokenMaker,
		redisClient: redisClient,
	}

	return server, nil
}
