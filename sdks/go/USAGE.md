<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"axiom"
	"axiom/models/components"
	"axiom/models/operations"
	"context"
	"log"
)

func main() {
	s := axiom.New(
		axiom.WithSecurity("<YOUR_AUTH_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Query(ctx, operations.QueryAplRequest{
		Format: operations.FormatTabular,
		APLRequestWithOptions: components.APLRequestWithOptions{
			Apl:       "[dataset_name] | limit 10",
			EndTime:   axiom.String("string"),
			StartTime: axiom.String("string"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AplResult != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->