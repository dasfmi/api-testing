# Datasets
(*Datasets*)

### Available Operations

* [List](#list) - List Datasets
* [Create](#create) - Create a dataset
* [Get](#get) - Retrieve dataset by ID
* [Update](#update) - Update dataset
* [Delete](#delete) - Delete dataset
* [Trim](#trim) - Trim dataset

## List

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
    res, err := s.Datasets.List(ctx)
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

## Create

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
    res, err := s.Datasets.Create(ctx, components.CreateDataset{
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

## Get

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
    res, err := s.Datasets.Get(ctx, id)
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

## Update

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
    res, err := s.Datasets.Update(ctx, id, updateDataset)
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

## Delete

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
    res, err := s.Datasets.Delete(ctx, id)
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

## Trim

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
    res, err := s.Datasets.Trim(ctx, id, trimOptions)
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
