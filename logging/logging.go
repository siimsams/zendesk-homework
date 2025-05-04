package logging

import (
	"context"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func LoggingUnaryInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	duration := time.Since(start)

	client := ""
	if p, ok := peer.FromContext(ctx); ok {
		client = p.Addr.String()
	}

	logEvent := log.Info().
		Str("method", info.FullMethod).
		Str("client", client).
		Dur("duration", duration)

	if err != nil {
		logEvent.Err(err).
			Str("status", status.Code(err).String()).
			Msg("gRPC request failed")
	} else {
		logEvent.Str("status", "OK").
			Msg("gRPC request handled")
	}

	return resp, err
}

func SetLogLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
