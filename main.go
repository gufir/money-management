package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/gufir/money-management/api"
	db "github.com/gufir/money-management/db/sqlc"
	_ "github.com/gufir/money-management/doc/statik"
	"github.com/gufir/money-management/gapi"
	"github.com/gufir/money-management/mail"
	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/utils"
	"github.com/gufir/money-management/worker"
	"github.com/jackc/pgx/v5/pgxpool"
)

var interruptSignal = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignal...)
	defer stop()

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db:")
	}

	store := db.NewStore(conn)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}
	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	waitGroup, ctx := errgroup.WithContext(ctx)

	go RunTaskProcessor(ctx, waitGroup, config, redisOpt, store)
	go runGatewayServer(ctx, waitGroup, config, store, taskDistributor)
	runGRPCServer(ctx, waitGroup, config, store, taskDistributor)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
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

func RunTaskProcessor(ctx context.Context, waitGroup *errgroup.Group, config utils.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("start task processor")

	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start task processor")
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("shutting down task processor")
		taskProcessor.Shutdown()
		log.Info().Msg("task processor stopped")
		return nil
	})
}

func runGRPCServer(ctx context.Context, waitGroup *errgroup.Group, config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Error().Err(err).Msg("cannot start server")
	}

	// grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer()
	pb.RegisterMoneyManagementServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Error().Err(err).Msg("cannot start server")
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start gRPC server on %s", listener.Addr().String())
		err = grpcServer.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error().Err(err).Msg("gRPC server failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("shutting down gRPC server")

		grpcServer.GracefulStop()
		log.Info().Msg("gRPC server stopped")

		return nil
	})

	// log.Printf("start gRPC Server on %s", listener.Addr().String())

	// err = grpcServer.Serve(listener)
	// if err != nil {
	// 	log.Error().Err(err).Msg("cannot start server")
	// }

}

func runGatewayServer(ctx context.Context, waitGroup *errgroup.Group, config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},

		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	err = pb.RegisterMoneyManagementHandlerServer(ctx, grpcMux, server)

	if err != nil {
		log.Fatal().Msg("cannot register service")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Msg("cannot create statik file system")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	// listener, err := net.Listen("tcp", config.HTTPServerAddress)
	// if err != nil {
	// 	log.Fatal().Msgf("cannot listen to gRPC server on %s: %v", config.GRPCServerAddress, err)
	// }

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   config.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(mux)
	httpServer := &http.Server{
		Handler: handler,
		Addr:    config.HTTPServerAddress,
	}

	// log.Info().Msgf("start HTTP gateway server on %s", listener.Addr().String())
	// handler := gapi.HttpLogger(corsHandler)

	// err = http.Serve(listener, handler)
	// if err != nil {
	// 	log.Fatal().Msg("cannot start server")
	// }

	waitGroup.Go(func() error {
		if err != nil {
			log.Fatal().Msgf("cannot listen to gRPC server on %s: %v", config.GRPCServerAddress, err)
		}
		log.Info().Msgf("start HTTP gateway server on %s", httpServer.Addr)
		err = httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Error().Err(err).Msg("HTTP gateway server failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("shutting down HTTP gateway server")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error().Msg("HTTP gateway server failed to shutdown")
		}

		log.Info().Msg("HTTP gateway server stopped")
		return nil
	})

}
