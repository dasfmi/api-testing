package main

import (
	"context"
	"log"

	axiom "../sdks/go"
)

func main() {
	ax := New()

	ctx := context.Background()
	_, err := ax.Query(ctx, operations.QueryAplRequest{
		ID:                    "datasetName",
		APLRequestWithOptions: components.APLRequestWithOptions{},
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = ax.IngestJSON(ctx, operations.IngestIntoDatasetJSONRequest{})
	if err != nil {
		log.Fatal(err)
	}
	_, err = ax.IngestRaw(ctx, "dataset", []byte{}, nil, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	ax.Datasets.Create(ctx, components.CreateDataset{
		Name: "name",
	})
	ax.Datasets.List(ctx)
	ax.Datasets.Update(ctx, "name", components.UpdateDataset{
		Description: axiom.String("dataset description"),
	})
	ax.Datasets.Get(ctx, "name")
	ax.Datasets.Trim(ctx, "name", components.TrimOptions{})
	ax.Datasets.Delete(ctx, "name")
}
