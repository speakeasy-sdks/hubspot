# PublicObject
(*PublicObject*)

### Available Operations

* [PostCrmV3ObjectsDealsMergeMerge](#postcrmv3objectsdealsmergemerge) - Merge two deals with same type

## PostCrmV3ObjectsDealsMergeMerge

Merge two deals with same type

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"context"
	"log"
)

func main() {
    s := deals.New()
    request := components.PublicMergeInput{
        ObjectIDToMerge: "<value>",
        PrimaryObjectID: "<value>",
    }

    security := operations.PostCrmV3ObjectsDealsMergeMergeSecurity{
            Oauth2: deals.String("<YOUR_OAUTH2_HERE>"),
        }
    ctx := context.Background()
    res, err := s.PublicObject.PostCrmV3ObjectsDealsMergeMerge(ctx, request, security)
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


### Response

**[*operations.PostCrmV3ObjectsDealsMergeMergeResponse](../../models/operations/postcrmv3objectsdealsmergemergeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
