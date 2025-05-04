package main

import (
	"log"
	"net"

	"github.com/siimsams/zendesk-homework/authentication"
	"github.com/siimsams/zendesk-homework/env"
	scorer "github.com/siimsams/zendesk-homework/proto"
	"github.com/siimsams/zendesk-homework/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := env.GetConfig()

	lis, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authentication.AuthUnaryInterceptor))
	scorer.RegisterScorerServiceServer(grpcServer, &service.ScorerServer{
		DBPath: config.DbPath,
	})
	reflection.Register(grpcServer)
	log.Printf("gRPC server listening on :%s", config.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
