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

	var format operations.Format = operations.FormatTabular

	aplRequestWithOptions := components.APLRequestWithOptions{
		Apl:       "[dataset_name] | limit 10",
		EndTime:   axiom.String("string"),
		StartTime: axiom.String("string"),
	}

	var nocache *bool = axiom.Bool(false)

	var saveAsKind *string = axiom.String("<value>")

	var id *string = axiom.String("<value>")

	ctx := context.Background()
	res, err := s.Query(ctx, format, aplRequestWithOptions, nocache, saveAsKind, id)
	if err != nil {
		log.Fatal(err)
	}
	if res.AplResult != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->