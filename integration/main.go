package main

import (
	"context"
	"log"
	"testing"

	axiom "../sdks/go"
)

func TestQuery(t *testing.T) {
	a := axiom.New()

	ctx := context.Background()
	tr := true
	_, err := a.Query(ctx, "['test']", "tabular", &tr, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
}
