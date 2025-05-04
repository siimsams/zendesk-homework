package authentication

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestValidateJWT(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		expectError bool
	}{
		{
			name:        "valid token",
			token:       generateValidToken(t),
			expectError: false,
		},
		{
			name:        "invalid token format",
			token:       "invalid-token",
			expectError: true,
		},
		{
			name:        "expired token",
			token:       generateExpiredToken(t),
			expectError: true,
		},
		{
			name:        "empty token",
			token:       "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			claims, err := validateJWT(tt.token)
			if tt.expectError {
				if err == nil {
					t.Error("expected error but got nil")
				}
				if claims != nil {
					t.Error("expected nil claims but got non-nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if claims == nil {
					t.Error("expected non-nil claims but got nil")
				}
			}
		})
	}
}

func TestAuthUnaryInterceptor(t *testing.T) {
	tests := []struct {
		name        string
		ctx         context.Context
		expectError bool
	}{
		{
			name:        "valid token in metadata",
			ctx:         createContextWithToken(t, generateValidToken(t)),
			expectError: false,
		},
		{
			name:        "missing metadata",
			ctx:         context.Background(),
			expectError: true,
		},
		{
			name:        "missing authorization header",
			ctx:         metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{})),
			expectError: true,
		},
		{
			name:        "invalid token",
			ctx:         createContextWithToken(t, "invalid-token"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := func(ctx context.Context, req interface{}) (interface{}, error) {
				return "success", nil
			}

			resp, err := AuthUnaryInterceptor(tt.ctx, nil, &grpc.UnaryServerInfo{}, handler)
			if tt.expectError {
				if err == nil {
					t.Error("expected error but got nil")
				}
				if resp != nil {
					t.Error("expected nil response but got non-nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if resp != "success" {
					t.Errorf("expected 'success' but got %v", resp)
				}
			}
		})
	}
}

func generateValidToken(t *testing.T) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		t.Fatalf("failed to generate valid token: %v", err)
	}
	return "Bearer " + tokenString
}

func generateExpiredToken(t *testing.T) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now().Add(-48 * time.Hour)),
		NotBefore: jwt.NewNumericDate(time.Now().Add(-48 * time.Hour)),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		t.Fatalf("failed to generate expired token: %v", err)
	}
	return "Bearer " + tokenString
}

func createContextWithToken(t *testing.T, token string) context.Context {
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	return metadata.NewIncomingContext(context.Background(), md)
}
