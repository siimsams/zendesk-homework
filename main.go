package main

import (
	"log"
	"net"

	scorer "github.com/siimsams/zendesk-homework/proto"
	"github.com/siimsams/zendesk-homework/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	scorer.RegisterScorerServiceServer(grpcServer, &service.ScorerServer{
		DBPath: "database.db",
	})
	reflection.Register(grpcServer)
	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
