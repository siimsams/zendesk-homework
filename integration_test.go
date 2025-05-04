package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/siimsams/zendesk-homework/env"
	"github.com/siimsams/zendesk-homework/observability/logging"
	scorer "github.com/siimsams/zendesk-homework/proto"
	"github.com/siimsams/zendesk-homework/test"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	grpcClient scorer.ScorerServiceClient
	grpcConn   *grpc.ClientConn
)

func TestMain(m *testing.M) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logging.SetLogLevel("debug")

	// Start gRPC server in background
	go func() {
		config := env.Config{
			Port:     "50051",
			DbPath:   "database.db",
			LogLevel: "debug",
		}
		if err := startServer(config); err != nil {
			log.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	// Wait briefly for server to start
	time.Sleep(2 * time.Second)

	// Connect client
	var err error
	grpcConn, err = grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to gRPC server")
	}
	grpcClient = scorer.NewScorerServiceClient(grpcConn)

	code := m.Run()

	grpcConn.Close()
	os.Exit(code)
}

func withAuth(ctx context.Context, t *testing.T) context.Context {
	md := metadata.New(map[string]string{
		"authorization": test.GenerateValidToken(t),
	})
	return metadata.NewOutgoingContext(ctx, md)
}

func TestGetCategoryScores(t *testing.T) {
	ctx := withAuth(context.Background(), t)

	req := &scorer.ScoreRequest{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}

	resp, err := grpcClient.GetCategoryScores(ctx, req)
	if err != nil {
		t.Fatalf("GetCategoryScores failed: %v", err)
	}
	if resp == nil {
		t.Error("Expected non-nil response")
	}
}

func TestGetOverallScore(t *testing.T) {
	ctx := withAuth(context.Background(), t)

	req := &scorer.ScoreRequest{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}

	resp, err := grpcClient.GetOverallScore(ctx, req)
	if err != nil {
		t.Fatalf("GetOverallScore failed: %v", err)
	}
	if resp == nil {
		t.Error("Expected non-nil response")
	}
}

func TestAuthentication(t *testing.T) {
	ctx := context.Background()

	req := &scorer.ScoreRequest{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}

	// Without auth should fail
	_, err := grpcClient.GetCategoryScores(ctx, req)
	if err == nil {
		t.Error("Expected authentication error but got none")
	}

	// With valid auth should succeed
	authCtx := withAuth(ctx, t)

	resp, err := grpcClient.GetCategoryScores(authCtx, req)
	if err != nil {
		t.Fatalf("GetCategoryScores with auth failed: %v", err)
	}
	if resp == nil {
		t.Error("Expected non-nil response with valid auth")
	}
}
