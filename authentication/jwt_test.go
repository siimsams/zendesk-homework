package authentication

import (
	"context"
	"testing"

	"github.com/siimsams/zendesk-homework/test"
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
			token:       test.GenerateValidToken(t),
			expectError: false,
		},
		{
			name:        "invalid token format",
			token:       "invalid-token",
			expectError: true,
		},
		{
			name:        "expired token",
			token:       test.GenerateExpiredToken(t),
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
			ctx:         test.CreateContextWithToken(t, test.GenerateValidToken(t)),
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
			ctx:         test.CreateContextWithToken(t, "invalid-token"),
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
