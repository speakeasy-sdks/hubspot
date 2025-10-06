# PublicObject
(*PublicObject*)

## Overview

### Available Operations

* [PostCrmV3ObjectsDealsMergeMerge](#postcrmv3objectsdealsmergemerge) - Merge two deals with same type

## PostCrmV3ObjectsDealsMergeMerge

Merge two deals with same type

### Example Usage

<!-- UsageSnippet language="go" operationID="post-/crm/v3/objects/deals/merge_merge" method="post" path="/crm/v3/objects/deals/merge" -->
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

    res, err := s.PublicObject.PostCrmV3ObjectsDealsMergeMerge(ctx, components.PublicMergeInput{
        ObjectIDToMerge: "<value>",
        PrimaryObjectID: "<id>",
    }, operations.PostCrmV3ObjectsDealsMergeMergeSecurity{
        Oauth2: deals.Pointer("<YOUR_OAUTH2_HERE>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SimplePublicObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [components.PublicMergeInput](../../models/components/publicmergeinput.md)                                               | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.PostCrmV3ObjectsDealsMergeMergeSecurity](../../models/operations/postcrmv3objectsdealsmergemergesecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.PostCrmV3ObjectsDealsMergeMergeResponse](../../models/operations/postcrmv3objectsdealsmergemergeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |