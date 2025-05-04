package main

import (
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/siimsams/zendesk-homework/authentication"
	"github.com/siimsams/zendesk-homework/env"
	"github.com/siimsams/zendesk-homework/logging"
	scorer "github.com/siimsams/zendesk-homework/proto"
	"github.com/siimsams/zendesk-homework/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load configuration
	config := env.GetConfig()

	// Set up logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logging.SetLogLevel(config.LogLevel)

	// Start TCP listener
	address := ":" + config.Port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen on port " + config.Port)
	}

	// Set up gRPC server with interceptors
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.LoggingUnaryInterceptor,
			authentication.AuthUnaryInterceptor,
		),
	)

	// Register the scorer service
	scorer.RegisterScorerServiceServer(grpcServer, &service.ScorerServer{
		DBPath: config.DbPath,
	})

	// Enable reflection for debugging
	reflection.Register(grpcServer)

	// Start serving
	log.Info().Msg("gRPC server listening on " + address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal().Err(err).Msg("gRPC server failed to start")
	}
}
