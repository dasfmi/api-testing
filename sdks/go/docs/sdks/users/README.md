# Users
(*Users*)

### Available Operations

* [GetCurrentUser](#getcurrentuser) - Get current user

## GetCurrentUser

Get current user

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
    res, err := s.Users.GetCurrentUser(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
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

**[*operations.GetCurrentUserResponse](../../models/operations/getcurrentuserresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |
