package test

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/siimsams/zendesk-homework/env"
	"google.golang.org/grpc/metadata"
)

func GenerateValidToken(t *testing.T) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})

	tokenString, err := token.SignedString(env.JwtSecretBytes())
	if err != nil {
		t.Fatalf("failed to generate valid token: %v", err)
	}
	return "Bearer " + tokenString
}

func GenerateExpiredToken(t *testing.T) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now().Add(-48 * time.Hour)),
		NotBefore: jwt.NewNumericDate(time.Now().Add(-48 * time.Hour)),
	})

	tokenString, err := token.SignedString(env.JwtSecretBytes())
	if err != nil {
		t.Fatalf("failed to generate expired token: %v", err)
	}
	return "Bearer " + tokenString
}

func CreateContextWithToken(t *testing.T, token string) context.Context {
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	return metadata.NewIncomingContext(context.Background(), md)
}
