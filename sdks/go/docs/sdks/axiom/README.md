# Axiom SDK


## Overview

Axiom API: 100% of your data for every possible need: o11y, security, analytics, and new insights.


### Available Operations

* [GetAnnotation](#getannotation) - Get a single annotation.
* [UpdateAnnotation](#updateannotation) - Update an annotation.
* [DeleteAnnotation](#deleteannotation) - Delete an annotation.

## GetAnnotation

Get a single annotation by id.

### Example Usage

```go
package main

import(
	"axiom-go/models/components"
	axiomgo "axiom-go"
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | The id of the annotation.                             |


### Response

**[*operations.GetAnnotationResponse](../../models/operations/getannotationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateAnnotation

Update an annotation.

### Example Usage

```go
package main

import(
	"axiom-go/models/components"
	axiomgo "axiom-go"
	"axiom-go/types"
	"context"
	"log"
)

func main() {
    s := axiomgo.New(
        axiomgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    baseAnnotation := components.BaseAnnotation{
        Title: axiomgo.String("Production deployment"),
        Description: axiomgo.String("A production deployment happened."),
        Time: types.MustNewTimeFromString("2024-04-19T15:00:00Z"),
        EndTime: types.MustNewTimeFromString("2024-04-19T16:00:00Z"),
        Type: axiomgo.String("deploy"),
        URL: axiomgo.String("https://deployments.example.com/42"),
        Datasets: []string{
            "o",
            "n",
            "e",
            "-",
            "d",
            "a",
            "t",
            "a",
            "s",
            "e",
            "t",
            ",",
            "a",
            "n",
            "o",
            "t",
            "h",
            "e",
            "r",
            "-",
            "d",
            "a",
            "t",
            "a",
            "s",
            "e",
            "t",
        },
    }

    ctx := context.Background()
    res, err := s.UpdateAnnotation(ctx, id, baseAnnotation)
    if err != nil {
        log.Fatal(err)
    }
    if res.Annotation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The id of the annotation.                                              |
| `baseAnnotation`                                                       | [components.BaseAnnotation](../../models/components/baseannotation.md) | :heavy_check_mark:                                                     | The fields to update.                                                  |


### Response

**[*operations.UpdateAnnotationResponse](../../models/operations/updateannotationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteAnnotation

Delete an annotation by id.

### Example Usage

```go
package main

import(
	"axiom-go/models/components"
	axiomgo "axiom-go"
	"context"
	"log"
)

func main() {
    s := axiomgo.New(
        axiomgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.DeleteAnnotation(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | The id of the annotation.                             |


### Response

**[*operations.DeleteAnnotationResponse](../../models/operations/deleteannotationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
