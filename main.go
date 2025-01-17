package main

import (
	"context"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gufir/money-management/api"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/gapi"
	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db:")
	}

	store := db.NewStore(conn)
	runGRPCServer(config, store)
}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start server:")
	}

}

func runGRPCServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Error().Err(err).Msg("cannot start server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMoneyManagementServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Error().Err(err).Msg("cannot start server")
	}

	log.Printf("start gRPC Server on %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Error().Err(err).Msg("cannot start server")
	}

}
