# Search
(*Search*)

### Available Operations

* [PostCrmV3ObjectsDealsSearchDoSearch](#postcrmv3objectsdealssearchdosearch)

## PostCrmV3ObjectsDealsSearchDoSearch

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"context"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
	"log"
)

func main() {
    s := deals.New()


    operationSecurity := operations.PostCrmV3ObjectsDealsSearchDoSearchSecurity{
            Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Search.PostCrmV3ObjectsDealsSearchDoSearch(ctx, components.PublicObjectSearchRequest{
        Limit: 378622,
        After: "<value>",
        Sorts: []string{
            "<value>",
        },
        Properties: []string{
            "<value>",
        },
        FilterGroups: []components.FilterGroup{
            components.FilterGroup{
                Filters: []components.Filter{
                    components.Filter{
                        PropertyName: "<value>",
                        Operator: components.OperatorGte,
                    },
                },
            },
        },
    }, operationSecurity)
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


### Response

**[*operations.PostCrmV3ObjectsDealsSearchDoSearchResponse](../../models/operations/postcrmv3objectsdealssearchdosearchresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
