# Axiom SDK


## Overview

Axiom Public API: A public and stable API for interacting with Axiom services

### Available Operations

* [Query](#query) - Query

## Query

Query

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
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
        Apl: "[dataset_name] | limit 10",
        EndTime: axiom.String("string"),
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

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `format`                                                                             | [operations.Format](../../models/operations/format.md)                               | :heavy_check_mark:                                                                   | N/A                                                                                  |                                                                                      |
| `aplRequestWithOptions`                                                              | [components.APLRequestWithOptions](../../models/components/aplrequestwithoptions.md) | :heavy_check_mark:                                                                   | N/A                                                                                  | {<br/>"apl": "[dataset_name] \| limit 10",<br/>"startTime": "string",<br/>"endTime": "string"<br/>} |
| `nocache`                                                                            | **bool*                                                                              | :heavy_minus_sign:                                                                   | N/A                                                                                  |                                                                                      |
| `saveAsKind`                                                                         | **string*                                                                            | :heavy_minus_sign:                                                                   | N/A                                                                                  |                                                                                      |
| `id`                                                                                 | **string*                                                                            | :heavy_minus_sign:                                                                   | when saveAsKind is true, this parameter indicates the id of the associated dataset   |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |


### Response

**[*operations.QueryAplResponse](../../models/operations/queryaplresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |
