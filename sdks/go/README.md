# axiom-go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/advanced-setup/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get axiom
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Axiom SDK](docs/sdks/axiom/README.md)

* [Query](docs/sdks/axiom/README.md#query) - Query

### [Datasets](docs/sdks/datasets/README.md)

* [GetDatasets](docs/sdks/datasets/README.md#getdatasets) - List Datasets
* [CreateDataset](docs/sdks/datasets/README.md#createdataset) - Create a dataset
* [GetDataset](docs/sdks/datasets/README.md#getdataset) - Retrieve dataset by ID
* [UpdateDataset](docs/sdks/datasets/README.md#updatedataset) - Update dataset
* [DeleteDataset](docs/sdks/datasets/README.md#deletedataset) - Delete dataset
* [IngestIntoDatasetJSON](docs/sdks/datasets/README.md#ingestintodatasetjson) - Ingest
* [IngestIntoDatasetRaw](docs/sdks/datasets/README.md#ingestintodatasetraw) - Ingest
* [QueryDataset](docs/sdks/datasets/README.md#querydataset) - Query (Legacy)
* [TrimDataset](docs/sdks/datasets/README.md#trimdataset) - Trim dataset

### [Users](docs/sdks/users/README.md)

* [GetCurrentUser](docs/sdks/users/README.md#getcurrentuser) - Get current user
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ForbiddenError | 401                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

### Example

```go
package main

import (
	"axiom"
	"axiom/models/components"
	"axiom/models/operations"
	"axiom/models/sdkerrors"
	"context"
	"errors"
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

		var e *sdkerrors.ForbiddenError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.axiom.co/v1/` | None |

#### Example

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
		axiom.WithServerIndex(0),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		axiom.WithServerURL("https://api.axiom.co/v1/"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `Auth`       | oauth2       | OAuth2 token |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
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
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `RetryConfig` object to the call by using the `WithRetries` option:
```go
package main

import (
	"axiom"
	"axiom/internal/utils"
	"axiom/models/components"
	"axiom/models/operations"
	"context"
	"log"
	"models/operations"
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
	res, err := s.Query(ctx, format, aplRequestWithOptions, nocache, saveAsKind, id, operations.WithRetries(
		utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.AplResult != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"axiom"
	"axiom/internal/utils"
	"axiom/models/components"
	"axiom/models/operations"
	"context"
	"log"
)

func main() {
	s := axiom.New(
		axiom.WithRetryConfig(
			utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
