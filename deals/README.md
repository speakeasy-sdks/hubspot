# openapi

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
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/hubspot/deals
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"log"
)

func main() {
	s := deals.New()

	batchReadInputSimplePublicObjectID := components.BatchReadInputSimplePublicObjectID{
		PropertiesWithHistory: []string{
			"<value>",
		},
		Inputs: []components.SimplePublicObjectID{
			components.SimplePublicObjectID{
				ID: "<id>",
			},
		},
		Properties: []string{
			"<value>",
		},
	}

	var archived *bool = deals.Bool(false)

	operationSecurity := operations.PostCrmV3ObjectsDealsBatchReadReadSecurity{
		Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
	}

	ctx := context.Background()
	res, err := s.Batch.PostCrmV3ObjectsDealsBatchReadRead(ctx, operationSecurity, batchReadInputSimplePublicObjectID, archived)
	if err != nil {
		log.Fatal(err)
	}
	if res.BatchResponseSimplePublicObject != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Batch](docs/sdks/batch/README.md)

* [PostCrmV3ObjectsDealsBatchReadRead](docs/sdks/batch/README.md#postcrmv3objectsdealsbatchreadread) - Read a batch of deals by internal ID, or unique property values
* [PostCrmV3ObjectsDealsBatchArchiveArchive](docs/sdks/batch/README.md#postcrmv3objectsdealsbatcharchivearchive) - Archive a batch of deals by ID
* [PostCrmV3ObjectsDealsBatchCreateCreate](docs/sdks/batch/README.md#postcrmv3objectsdealsbatchcreatecreate) - Create a batch of deals
* [PostCrmV3ObjectsDealsBatchUpdateUpdate](docs/sdks/batch/README.md#postcrmv3objectsdealsbatchupdateupdate) - Update a batch of deals

### [Basic](docs/sdks/basic/README.md)

* [GetCrmV3ObjectsDealsDealIDGetByID](docs/sdks/basic/README.md#getcrmv3objectsdealsdealidgetbyid) - Read
* [DeleteCrmV3ObjectsDealsDealIDArchive](docs/sdks/basic/README.md#deletecrmv3objectsdealsdealidarchive) - Archive
* [PatchCrmV3ObjectsDealsDealIDUpdate](docs/sdks/basic/README.md#patchcrmv3objectsdealsdealidupdate) - Update
* [GetCrmV3ObjectsDealsGetPage](docs/sdks/basic/README.md#getcrmv3objectsdealsgetpage) - List
* [PostCrmV3ObjectsDealsCreate](docs/sdks/basic/README.md#postcrmv3objectsdealscreate) - Create

### [PublicObject](docs/sdks/publicobject/README.md)

* [PostCrmV3ObjectsDealsMergeMerge](docs/sdks/publicobject/README.md#postcrmv3objectsdealsmergemerge) - Merge two deals with same type

### [Gdpr](docs/sdks/gdpr/README.md)

* [PostCrmV3ObjectsDealsGdprDeletePurge](docs/sdks/gdpr/README.md#postcrmv3objectsdealsgdprdeletepurge) - GDPR DELETE

### [Search](docs/sdks/search/README.md)

* [PostCrmV3ObjectsDealsSearchDoSearch](docs/sdks/search/README.md#postcrmv3objectsdealssearchdosearch)
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/sdkerrors"
	"log"
)

func main() {
	s := deals.New()

	batchReadInputSimplePublicObjectID := components.BatchReadInputSimplePublicObjectID{
		PropertiesWithHistory: []string{
			"<value>",
		},
		Inputs: []components.SimplePublicObjectID{
			components.SimplePublicObjectID{
				ID: "<id>",
			},
		},
		Properties: []string{
			"<value>",
		},
	}

	var archived *bool = deals.Bool(false)

	operationSecurity := operations.PostCrmV3ObjectsDealsBatchReadReadSecurity{
		Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
	}

	ctx := context.Background()
	res, err := s.Batch.PostCrmV3ObjectsDealsBatchReadRead(ctx, operationSecurity, batchReadInputSimplePublicObjectID, archived)
	if err != nil {

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
| 0 | `https://api.hubapi.com` | None |

#### Example

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"log"
)

func main() {
	s := deals.New(
		deals.WithServerIndex(0),
	)

	batchReadInputSimplePublicObjectID := components.BatchReadInputSimplePublicObjectID{
		PropertiesWithHistory: []string{
			"<value>",
		},
		Inputs: []components.SimplePublicObjectID{
			components.SimplePublicObjectID{
				ID: "<id>",
			},
		},
		Properties: []string{
			"<value>",
		},
	}

	var archived *bool = deals.Bool(false)

	operationSecurity := operations.PostCrmV3ObjectsDealsBatchReadReadSecurity{
		Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
	}

	ctx := context.Background()
	res, err := s.Batch.PostCrmV3ObjectsDealsBatchReadRead(ctx, operationSecurity, batchReadInputSimplePublicObjectID, archived)
	if err != nil {
		log.Fatal(err)
	}
	if res.BatchResponseSimplePublicObject != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"log"
)

func main() {
	s := deals.New(
		deals.WithServerURL("https://api.hubapi.com"),
	)

	batchReadInputSimplePublicObjectID := components.BatchReadInputSimplePublicObjectID{
		PropertiesWithHistory: []string{
			"<value>",
		},
		Inputs: []components.SimplePublicObjectID{
			components.SimplePublicObjectID{
				ID: "<id>",
			},
		},
		Properties: []string{
			"<value>",
		},
	}

	var archived *bool = deals.Bool(false)

	operationSecurity := operations.PostCrmV3ObjectsDealsBatchReadReadSecurity{
		Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
	}

	ctx := context.Background()
	res, err := s.Batch.PostCrmV3ObjectsDealsBatchReadRead(ctx, operationSecurity, batchReadInputSimplePublicObjectID, archived)
	if err != nil {
		log.Fatal(err)
	}
	if res.BatchResponseSimplePublicObject != nil {
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
| `Oauth2`     | oauth2       | OAuth2 token |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:


### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"log"
)

func main() {
	s := deals.New()

	batchReadInputSimplePublicObjectID := components.BatchReadInputSimplePublicObjectID{
		PropertiesWithHistory: []string{
			"<value>",
		},
		Inputs: []components.SimplePublicObjectID{
			components.SimplePublicObjectID{
				ID: "<id>",
			},
		},
		Properties: []string{
			"<value>",
		},
	}

	var archived *bool = deals.Bool(false)

	operationSecurity := operations.PostCrmV3ObjectsDealsBatchReadReadSecurity{
		Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
	}

	ctx := context.Background()
	res, err := s.Batch.PostCrmV3ObjectsDealsBatchReadRead(ctx, operationSecurity, batchReadInputSimplePublicObjectID, archived)
	if err != nil {
		log.Fatal(err)
	}
	if res.BatchResponseSimplePublicObject != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

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
