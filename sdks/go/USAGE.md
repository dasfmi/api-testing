<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	axiomgo "axiom-go"
	"axiom-go/models/components"
	"context"
	"log"
)

func main() {
	s := axiomgo.New(
		axiomgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	var id string = "<value>"

	ctx := context.Background()
	res, err := s.GetAnnotation(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	if res.Annotation != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->