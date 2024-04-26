package main

import (
	"context"
	"log"

	"github.com/axiomhq/axiom/public-api/tests/api"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target api --clean ../openapi.yaml

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := TestAnnotations(ctx); err != nil {
		log.Fatalf("TestAnnotations failed: %v", err)
	}
}

// TokenSecuritySource implements api.SecuritySource.
type TokenSecuritySource struct {
	token string
}

// NewTokenSecuritySource creates a new TokenSecuritySource.
func NewTokenSecuritySource(token string) api.SecuritySource {
	return &TokenSecuritySource{token: token}
}

// BearerAuth implements api.SecuritySource.
func (t *TokenSecuritySource) BearerAuth(ctx context.Context, operationName string) (api.BearerAuth, error) {
	return api.BearerAuth{Token: t.token}, nil
}
