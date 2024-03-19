# Gdpr
(*Gdpr*)

### Available Operations

* [PostCrmV3ObjectsDealsGdprDeletePurge](#postcrmv3objectsdealsgdprdeletepurge) - GDPR DELETE

## PostCrmV3ObjectsDealsGdprDeletePurge

Permanently delete a contact and all associated content to follow GDPR. Use optional property 'idProperty' set to 'email' to identify contact by email address. If email address is not found, the email address will be added to a blocklist and prevent it from being used in the future.

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


    operationSecurity := operations.PostCrmV3ObjectsDealsGdprDeletePurgeSecurity{
            Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Gdpr.PostCrmV3ObjectsDealsGdprDeletePurge(ctx, components.PublicGdprDeleteInput{
        ObjectID: "<value>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [components.PublicGdprDeleteInput](../../models/components/publicgdprdeleteinput.md)                                               | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `security`                                                                                                                         | [operations.PostCrmV3ObjectsDealsGdprDeletePurgeSecurity](../../models/operations/postcrmv3objectsdealsgdprdeletepurgesecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |


### Response

**[*operations.PostCrmV3ObjectsDealsGdprDeletePurgeResponse](../../models/operations/postcrmv3objectsdealsgdprdeletepurgeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
