# Annotations
(*Annotations*)

## Overview

Annotations are used to mark important events in your data.

### Available Operations

* [GetAnnotations](#getannotations) - Get annotations
* [CreateAnnotation](#createannotation) - Create a new annotation.

## GetAnnotations

Get all annotations and filter by datasets or timerange.

### Example Usage

```go
package main

import(
	"axiom-go/models/components"
	axiomgo "axiom-go"
	"time"
	"axiom-go/types"
	"context"
	"log"
)

func main() {
    s := axiomgo.New(
        axiomgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    datasets := []string{
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
    }

    var start *time.Time = types.MustNewTimeFromString("2024-04-19T15:00:00Z")

    var end *time.Time = types.MustNewTimeFromString("2024-04-19T16:00:00Z")

    ctx := context.Background()
    res, err := s.Annotations.GetAnnotations(ctx, datasets, start, end)
    if err != nil {
        log.Fatal(err)
    }
    if res.Annotations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |                                                                       |
| `datasets`                                                            | []*string*                                                            | :heavy_minus_sign:                                                    | The datasets to filter by.                                            | one-dataset,another-dataset                                           |
| `start`                                                               | [*time.Time](https://pkg.go.dev/time#Time)                            | :heavy_minus_sign:                                                    | If set, will filter to events after this date. Should be in RFC3339.  | 2024-04-19T15:00:00Z                                                  |
| `end`                                                                 | [*time.Time](https://pkg.go.dev/time#Time)                            | :heavy_minus_sign:                                                    | If set, will filter to events before this date. Should be in RFC3339. | 2024-04-19T16:00:00Z                                                  |


### Response

**[*operations.GetAnnotationsResponse](../../models/operations/getannotationsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateAnnotation

Create a new annotation.

### Example Usage

```go
package main

import(
	"axiom-go/models/components"
	axiomgo "axiom-go"
	"context"
	"axiom-go/types"
	"log"
)

func main() {
    s := axiomgo.New(
        axiomgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Annotations.CreateAnnotation(ctx, components.NewAnnotation{
        Title: axiomgo.String("Production deployment"),
        Description: axiomgo.String("A production deployment happened."),
        Time: types.MustNewTimeFromString("2024-04-19T15:00:00Z"),
        EndTime: types.MustNewTimeFromString("2024-04-19T16:00:00Z"),
        Type: "deploy",
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
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Annotation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.NewAnnotation](../../models/components/newannotation.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.CreateAnnotationResponse](../../models/operations/createannotationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
