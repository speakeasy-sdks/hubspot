# Search
(*Search*)

## Overview

### Available Operations

* [PostCrmV3ObjectsDealsSearchDoSearch](#postcrmv3objectsdealssearchdosearch)

## PostCrmV3ObjectsDealsSearchDoSearch

### Example Usage

<!-- UsageSnippet language="go" operationID="post-/crm/v3/objects/deals/search_doSearch" method="post" path="/crm/v3/objects/deals/search" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := deals.New()

    res, err := s.Search.PostCrmV3ObjectsDealsSearchDoSearch(ctx, components.PublicObjectSearchRequest{
        Limit: 322945,
        After: "<value>",
        Sorts: []string{
            "<value 1>",
        },
        Properties: []string{
            "<value 1>",
        },
        FilterGroups: []components.FilterGroup{},
    }, operations.PostCrmV3ObjectsDealsSearchDoSearchSecurity{
        Oauth2: deals.Pointer("<YOUR_OAUTH2_HERE>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionResponseWithTotalSimplePublicObjectForwardPaging != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [components.PublicObjectSearchRequest](../../models/components/publicobjectsearchrequest.md)                                     | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `security`                                                                                                                       | [operations.PostCrmV3ObjectsDealsSearchDoSearchSecurity](../../models/operations/postcrmv3objectsdealssearchdosearchsecurity.md) | :heavy_check_mark:                                                                                                               | The security requirements to use for the request.                                                                                |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.PostCrmV3ObjectsDealsSearchDoSearchResponse](../../models/operations/postcrmv3objectsdealssearchdosearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |