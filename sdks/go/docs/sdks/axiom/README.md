# Axiom SDK


## Overview

Axiom Public API: A public and stable API for interacting with Axiom services

### Available Operations

* [Query](#query) - Query
* [IngestJSON](#ingestjson) - Ingest
* [IngestRaw](#ingestraw) - Ingest
* [QueryLegacyDollar](#querylegacydollar) - Query (Legacy)

## Query

Query

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"axiom/models/operations"
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
            Apl: "[dataset_name] | limit 10",
            EndTime: axiom.String("string"),
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

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.QueryAplRequest](../../models/operations/queryaplrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |


### Response

**[*operations.QueryAplResponse](../../models/operations/queryaplresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## IngestJSON

Ingest

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"axiom/models/operations"
	"log"
)

func main() {
    s := axiom.New(
        axiom.WithSecurity("<YOUR_AUTH_HERE>"),
    )

    ctx := context.Background()
    res, err := s.IngestJSON(ctx, operations.IngestIntoDatasetJSONRequest{
        ID: "<id>",
        RequestBody: []operations.RequestBody{
            operations.RequestBody{},
            operations.RequestBody{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IngestStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.IngestIntoDatasetJSONRequest](../../models/operations/ingestintodatasetjsonrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.IngestIntoDatasetJSONResponse](../../models/operations/ingestintodatasetjsonresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## IngestRaw

Ingest

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"axiom/models/operations"
	"log"
)

func main() {
    s := axiom.New(
        axiom.WithSecurity("<YOUR_AUTH_HERE>"),
    )

    ctx := context.Background()
    res, err := s.IngestRaw(ctx, operations.IngestIntoDatasetRawRequest{
        ID: "<id>",
        RequestBody: []byte("0xC34eBEb2EB"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IngestStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.IngestIntoDatasetRawRequest](../../models/operations/ingestintodatasetrawrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.IngestIntoDatasetRawResponse](../../models/operations/ingestintodatasetrawresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## QueryLegacyDollar

Query (Legacy)

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"axiom/models/operations"
	"log"
)

func main() {
    s := axiom.New(
        axiom.WithSecurity("<YOUR_AUTH_HERE>"),
    )

    ctx := context.Background()
    res, err := s.QueryLegacyDollar(ctx, operations.QueryDatasetRequest{
        ID: "<id>",
        QueryRequestWithOptions: components.QueryRequestWithOptions{
            Aggregations: []components.Aggregation{
                components.Aggregation{
                    Field: "<value>",
                    Op: components.OpAvg,
                },
            },
            ContinuationToken: axiom.String("string"),
            Cursor: axiom.String("string"),
            EndTime: "string",
            Filter: &components.Filter{
                Field: "<value>",
                Op: components.FilterOpEndsWith,
            },
            GroupBy: []string{
                "string",
            },
            IncludeCursor: axiom.Bool(true),
            Limit: axiom.Int64(10),
            Order: []components.Order{
                components.Order{
                    Desc: true,
                    Field: "string",
                },
            },
            Project: []components.Projection{
                components.Projection{
                    Alias: axiom.String("string"),
                    Field: "string",
                },
            },
            QueryOptions: &components.QueryOptions{
                DisplayNull: axiom.String("0"),
            },
            Resolution: "string",
            StartTime: "string",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Result != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.QueryDatasetRequest](../../models/operations/querydatasetrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.QueryDatasetResponse](../../models/operations/querydatasetresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |
