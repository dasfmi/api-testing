package main

import (
	"context"
	"fmt"
	"os"

	"github.com/axiomhq/axiom/public-api/tests/api"
)

func TestAnnotations(ctx context.Context) error {
	serverURL := os.Getenv("AXIOM_URL")
	token := os.Getenv("AXIOM_TOKEN")
	if serverURL == "" || token == "" {
		return fmt.Errorf("AXIOM_URL and AXIOM_TOKEN must be set")
	}

	client, err := api.NewClient(serverURL, NewTokenSecuritySource(token))
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	res, err := client.GetAnnotations(ctx, api.GetAnnotationsParams{})
	if err != nil {
		return fmt.Errorf("failed to get annotations: %w", err)
	}
	annotations, ok := res.(*api.GetAnnotationsOKApplicationJSON)
	if !ok {
		return fmt.Errorf("unexpected response type: %T", res)
	}

	fmt.Println(annotations)

	return nil
}
