# Datasets
(*Datasets*)

### Available Operations

* [GetDatasets](#getdatasets) - List Datasets
* [CreateDataset](#createdataset) - Create a dataset
* [GetDataset](#getdataset) - Retrieve dataset by ID
* [UpdateDataset](#updatedataset) - Update dataset
* [DeleteDataset](#deletedataset) - Delete dataset
* [IngestIntoDatasetJSON](#ingestintodatasetjson) - Ingest
* [IngestIntoDatasetRaw](#ingestintodatasetraw) - Ingest
* [QueryDataset](#querydataset) - Query (Legacy)
* [TrimDataset](#trimdataset) - Trim dataset

## GetDatasets

Retrieves a list of datasets for the authenticated user.

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"log"
)

func main() {
    s := axiom.New(
        axiom.WithSecurity("<YOUR_AUTH_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Datasets.GetDatasets(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.DatasetSpecs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetDatasetsResponse](../../models/operations/getdatasetsresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## CreateDataset

Create a dataset

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"log"
)

func main() {
    s := axiom.New(
        axiom.WithSecurity("<YOUR_AUTH_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Datasets.CreateDataset(ctx, components.CreateDataset{
        Description: axiom.String("string"),
        Name: "string",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatasetSpec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.CreateDataset](../../models/components/createdataset.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |


### Response

**[*operations.CreateDatasetResponse](../../models/operations/createdatasetresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## GetDataset

Retrieve dataset by ID

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"log"
)

func main() {
    s := axiom.New(
        axiom.WithSecurity("<YOUR_AUTH_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.Datasets.GetDataset(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.DatasetSpec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetDatasetResponse](../../models/operations/getdatasetresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.NotFoundError  | 404                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## UpdateDataset

Update dataset

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"log"
)

func main() {
    s := axiom.New(
        axiom.WithSecurity("<YOUR_AUTH_HERE>"),
    )


    var id string = "<value>"

    updateDataset := components.UpdateDataset{
        Description: axiom.String("string"),
    }

    ctx := context.Background()
    res, err := s.Datasets.UpdateDataset(ctx, id, updateDataset)
    if err != nil {
        log.Fatal(err)
    }
    if res.DatasetSpec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |                                                                      |
| `id`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |
| `updateDataset`                                                      | [components.UpdateDataset](../../models/components/updatedataset.md) | :heavy_check_mark:                                                   | N/A                                                                  | {<br/>"description": "string"<br/>}                                  |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |                                                                      |


### Response

**[*operations.UpdateDatasetResponse](../../models/operations/updatedatasetresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## DeleteDataset

Delete dataset

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"log"
)

func main() {
    s := axiom.New(
        axiom.WithSecurity("<YOUR_AUTH_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.Datasets.DeleteDataset(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.DeleteDatasetResponse](../../models/operations/deletedatasetresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## IngestIntoDatasetJSON

Ingest

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


    var id string = "<value>"

    requestBody := []operations.RequestBody{
        operations.RequestBody{},
        operations.RequestBody{},
    }

    var timestampField *string = axiom.String("<value>")

    var timestampFormat *string = axiom.String("<value>")

    var csvDelimiter *string = axiom.String("<value>")

    xAxiomCSVFields := []string{
        "<value>",
    }

    var xAxiomEventLabels *string = axiom.String("<value>")

    ctx := context.Background()
    res, err := s.Datasets.IngestIntoDatasetJSON(ctx, id, requestBody, timestampField, timestampFormat, csvDelimiter, xAxiomCSVFields, xAxiomEventLabels)
    if err != nil {
        log.Fatal(err)
    }
    if res.IngestStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    | Example                                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |                                                                                                                                                                |
| `id`                                                                                                                                                           | *string*                                                                                                                                                       | :heavy_check_mark:                                                                                                                                             | Unique ID of the dataset where you want to send data.                                                                                                          |                                                                                                                                                                |
| `requestBody`                                                                                                                                                  | [][operations.RequestBody](../../models/operations/requestbody.md)                                                                                             | :heavy_check_mark:                                                                                                                                             | Data you want to send to Axiom in a supported format.                                                                                                          | [<br/>{<br/>"message": "Hello, World!",<br/>"foo": "bar"<br/>},<br/>{<br/>"bar": "foz"<br/>}<br/>]                                                             |
| `timestampField`                                                                                                                                               | **string*                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                             | The name of the field to use as the timestamp. If not specified, the timestamp will be the time the event was received by Axiom.                               |                                                                                                                                                                |
| `timestampFormat`                                                                                                                                              | **string*                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                             | The date-time format of the timestamp field. The reference time is `Mon Jan 2 15:04:05 -0700 MST 2006`, as specified in https://pkg.go.dev/time/?tab=doc#Parse |                                                                                                                                                                |
| `csvDelimiter`                                                                                                                                                 | **string*                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                             | The delimiter to use when parsing CSV data. If not specified, the default delimiter is `,`.                                                                    |                                                                                                                                                                |
| `xAxiomCSVFields`                                                                                                                                              | []*string*                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                             | A list of optional comma separated fields to use for CSV ingestion. If not specified, the first line of the CSV will be used as the field names.               |                                                                                                                                                                |
| `xAxiomEventLabels`                                                                                                                                            | **string*                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                             | An optional JSON encoded object with additional labels to add to all events in the request                                                                     |                                                                                                                                                                |
| `opts`                                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |                                                                                                                                                                |


### Response

**[*operations.IngestIntoDatasetJSONResponse](../../models/operations/ingestintodatasetjsonresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## IngestIntoDatasetRaw

Ingest

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"log"
)

func main() {
    s := axiom.New(
        axiom.WithSecurity("<YOUR_AUTH_HERE>"),
    )


    var id string = "<value>"

    var requestBody []byte = []byte("0xfdbAd8082B")

    var timestampField *string = axiom.String("<value>")

    var timestampFormat *string = axiom.String("<value>")

    var csvDelimiter *string = axiom.String("<value>")

    xAxiomCSVFields := []string{
        "<value>",
    }

    var xAxiomEventLabels *string = axiom.String("<value>")

    ctx := context.Background()
    res, err := s.Datasets.IngestIntoDatasetRaw(ctx, id, requestBody, timestampField, timestampFormat, csvDelimiter, xAxiomCSVFields, xAxiomEventLabels)
    if err != nil {
        log.Fatal(err)
    }
    if res.IngestStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `id`                                                                                                                                                           | *string*                                                                                                                                                       | :heavy_check_mark:                                                                                                                                             | Unique ID of the dataset where you want to send data.                                                                                                          |
| `requestBody`                                                                                                                                                  | *[]byte*                                                                                                                                                       | :heavy_check_mark:                                                                                                                                             | Data you want to send to Axiom in a supported format.                                                                                                          |
| `timestampField`                                                                                                                                               | **string*                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                             | The name of the field to use as the timestamp. If not specified, the timestamp will be the time the event was received by Axiom.                               |
| `timestampFormat`                                                                                                                                              | **string*                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                             | The date-time format of the timestamp field. The reference time is `Mon Jan 2 15:04:05 -0700 MST 2006`, as specified in https://pkg.go.dev/time/?tab=doc#Parse |
| `csvDelimiter`                                                                                                                                                 | **string*                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                             | The delimiter to use when parsing CSV data. If not specified, the default delimiter is `,`.                                                                    |
| `xAxiomCSVFields`                                                                                                                                              | []*string*                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                             | A list of optional comma separated fields to use for CSV ingestion. If not specified, the first line of the CSV will be used as the field names.               |
| `xAxiomEventLabels`                                                                                                                                            | **string*                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                             | An optional JSON encoded object with additional labels to add to all events in the request                                                                     |
| `opts`                                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                             | The options for this request.                                                                                                                                  |


### Response

**[*operations.IngestIntoDatasetRawResponse](../../models/operations/ingestintodatasetrawresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## QueryDataset

Query (Legacy)

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"log"
)

func main() {
    s := axiom.New(
        axiom.WithSecurity("<YOUR_AUTH_HERE>"),
    )


    var id string = "<value>"

    queryRequestWithOptions := components.QueryRequestWithOptions{
        Aggregations: []components.Aggregation{
            components.Aggregation{
                Field: "<value>",
                Op: components.OpMin,
            },
        },
        ContinuationToken: axiom.String("string"),
        Cursor: axiom.String("string"),
        EndTime: "string",
        Filter: &components.Filter{
            Field: "<value>",
            Op: components.FilterOpExists,
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
    }

    var saveAsKind *string = axiom.String("<value>")

    var streamingDuration *string = axiom.String("<value>")

    var nocache *bool = axiom.Bool(false)

    ctx := context.Background()
    res, err := s.Datasets.QueryDataset(ctx, id, queryRequestWithOptions, saveAsKind, streamingDuration, nocache)
    if err != nil {
        log.Fatal(err)
    }
    if res.Result != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                   | Type                                                                                                                                                                                                                                                                                                                                                                        | Required                                                                                                                                                                                                                                                                                                                                                                    | Description                                                                                                                                                                                                                                                                                                                                                                 | Example                                                                                                                                                                                                                                                                                                                                                                     |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                          | The context to use for the request.                                                                                                                                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                                                                                                             |
| `id`                                                                                                                                                                                                                                                                                                                                                                        | *string*                                                                                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                          | N/A                                                                                                                                                                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                                                                                                             |
| `queryRequestWithOptions`                                                                                                                                                                                                                                                                                                                                                   | [components.QueryRequestWithOptions](../../models/components/queryrequestwithoptions.md)                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                          | N/A                                                                                                                                                                                                                                                                                                                                                                         | {<br/>"aggregations": [],<br/>"continuationToken": "string",<br/>"cursor": "string",<br/>"endTime": "string",<br/>"filter": {},<br/>"groupBy": [<br/>"string"<br/>],<br/>"includeCursor": true,<br/>"limit": 10,<br/>"order": [<br/>{<br/>"desc": true,<br/>"field": "string"<br/>}<br/>],<br/>"project": [<br/>{<br/>"alias": "string",<br/>"field": "string"<br/>}<br/>],<br/>"queryOptions": {<br/>"displayNull": "0"<br/>},<br/>"resolution": "string",<br/>"startTime": "string"<br/>} |
| `saveAsKind`                                                                                                                                                                                                                                                                                                                                                                | **string*                                                                                                                                                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                          | N/A                                                                                                                                                                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                                                                                                             |
| `streamingDuration`                                                                                                                                                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                          | N/A                                                                                                                                                                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                                                                                                             |
| `nocache`                                                                                                                                                                                                                                                                                                                                                                   | **bool*                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                          | N/A                                                                                                                                                                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                                                                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                          | The options for this request.                                                                                                                                                                                                                                                                                                                                               |                                                                                                                                                                                                                                                                                                                                                                             |


### Response

**[*operations.QueryDatasetResponse](../../models/operations/querydatasetresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## TrimDataset

Trim dataset

### Example Usage

```go
package main

import(
	"axiom/models/components"
	"axiom"
	"context"
	"log"
)

func main() {
    s := axiom.New(
        axiom.WithSecurity("<YOUR_AUTH_HERE>"),
    )


    var id string = "<value>"

    trimOptions := components.TrimOptions{
        MaxDuration: "1h",
    }

    ctx := context.Background()
    res, err := s.Datasets.TrimDataset(ctx, id, trimOptions)
    if err != nil {
        log.Fatal(err)
    }
    if res.TrimResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |                                                                  |
| `id`                                                             | *string*                                                         | :heavy_check_mark:                                               | N/A                                                              |                                                                  |
| `trimOptions`                                                    | [components.TrimOptions](../../models/components/trimoptions.md) | :heavy_check_mark:                                               | N/A                                                              | {<br/>"maxDuration": "1h"<br/>}                                  |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |                                                                  |


### Response

**[*operations.TrimDatasetResponse](../../models/operations/trimdatasetresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |
