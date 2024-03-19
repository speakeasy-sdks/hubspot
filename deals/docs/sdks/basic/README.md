# Basic
(*Basic*)

### Available Operations

* [GetCrmV3ObjectsDealsDealIDGetByID](#getcrmv3objectsdealsdealidgetbyid) - Read
* [DeleteCrmV3ObjectsDealsDealIDArchive](#deletecrmv3objectsdealsdealidarchive) - Archive
* [PatchCrmV3ObjectsDealsDealIDUpdate](#patchcrmv3objectsdealsdealidupdate) - Update
* [GetCrmV3ObjectsDealsGetPage](#getcrmv3objectsdealsgetpage) - List
* [PostCrmV3ObjectsDealsCreate](#postcrmv3objectsdealscreate) - Create

## GetCrmV3ObjectsDealsDealIDGetByID

Read an Object identified by `{dealId}`. `{dealId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param.  Control what is returned via the `properties` query param.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"context"
	"log"
)

func main() {
    s := deals.New()


    operationSecurity := operations.GetCrmV3ObjectsDealsDealIDGetByIDSecurity{
            Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Basic.GetCrmV3ObjectsDealsDealIDGetByID(ctx, operations.GetCrmV3ObjectsDealsDealIDGetByIDRequest{
        DealID: "<value>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.SimplePublicObjectWithAssociations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetCrmV3ObjectsDealsDealIDGetByIDRequest](../../models/operations/getcrmv3objectsdealsdealidgetbyidrequest.md)   | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `security`                                                                                                                   | [operations.GetCrmV3ObjectsDealsDealIDGetByIDSecurity](../../models/operations/getcrmv3objectsdealsdealidgetbyidsecurity.md) | :heavy_check_mark:                                                                                                           | The security requirements to use for the request.                                                                            |


### Response

**[*operations.GetCrmV3ObjectsDealsDealIDGetByIDResponse](../../models/operations/getcrmv3objectsdealsdealidgetbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteCrmV3ObjectsDealsDealIDArchive

Move an Object identified by `{dealId}` to the recycling bin.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"context"
	"log"
)

func main() {
    s := deals.New()


    var dealID string = "<value>"

    operationSecurity := operations.DeleteCrmV3ObjectsDealsDealIDArchiveSecurity{
            Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Basic.DeleteCrmV3ObjectsDealsDealIDArchive(ctx, operationSecurity, dealID)
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
| `security`                                                                                                                         | [operations.DeleteCrmV3ObjectsDealsDealIDArchiveSecurity](../../models/operations/deletecrmv3objectsdealsdealidarchivesecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |
| `dealID`                                                                                                                           | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |


### Response

**[*operations.DeleteCrmV3ObjectsDealsDealIDArchiveResponse](../../models/operations/deletecrmv3objectsdealsdealidarchiveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchCrmV3ObjectsDealsDealIDUpdate

Perform a partial update of an Object identified by `{dealId}`. `{dealId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param. Provided property values will be overwritten. Read-only and non-existent properties will be ignored. Properties values can be cleared by passing an empty string.

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


    var dealID string = "<value>"

    simplePublicObjectInput := components.SimplePublicObjectInput{
        Properties: map[string]string{
            "amount": "1500.00",
            "dealname": "Custom data integrations",
            "pipeline": "default",
            "closedate": "2019-12-07T16:50:06.678Z",
            "dealstage": "presentationscheduled",
            "hubspot_owner_id": "910901",
        },
    }

    var idProperty *string = deals.String("<value>")

    operationSecurity := operations.PatchCrmV3ObjectsDealsDealIDUpdateSecurity{
            Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Basic.PatchCrmV3ObjectsDealsDealIDUpdate(ctx, operationSecurity, dealID, simplePublicObjectInput, idProperty)
    if err != nil {
        log.Fatal(err)
    }
    if res.SimplePublicObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                                                                                                                                                                                                                    | Required                                                                                                                                                                                                                                                                                                                                                | Description                                                                                                                                                                                                                                                                                                                                             | Example                                                                                                                                                                                                                                                                                                                                                 |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                      | The context to use for the request.                                                                                                                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                                                                                                         |
| `security`                                                                                                                                                                                                                                                                                                                                              | [operations.PatchCrmV3ObjectsDealsDealIDUpdateSecurity](../../models/operations/patchcrmv3objectsdealsdealidupdatesecurity.md)                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                      | The security requirements to use for the request.                                                                                                                                                                                                                                                                                                       |                                                                                                                                                                                                                                                                                                                                                         |
| `dealID`                                                                                                                                                                                                                                                                                                                                                | *string*                                                                                                                                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                      | N/A                                                                                                                                                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                                                                                                         |
| `simplePublicObjectInput`                                                                                                                                                                                                                                                                                                                               | [components.SimplePublicObjectInput](../../models/components/simplepublicobjectinput.md)                                                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                      | N/A                                                                                                                                                                                                                                                                                                                                                     | {<br/>"properties": {<br/>"amount": "1500.00",<br/>"dealname": "Custom data integrations",<br/>"pipeline": "default",<br/>"closedate": "2019-12-07T16:50:06.678Z",<br/>"dealstage": "presentationscheduled",<br/>"hubspot_owner_id": "910901"<br/>},<br/>"associations": [<br/>{<br/>"to": {<br/>"id": "101"<br/>},<br/>"types": [<br/>{<br/>"associationCategory": "HUBSPOT_DEFINED",<br/>"associationTypeId": 2<br/>}<br/>]<br/>}<br/>]<br/>} |
| `idProperty`                                                                                                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                      | The name of a property whose values are unique for this object type                                                                                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                                                                                                         |


### Response

**[*operations.PatchCrmV3ObjectsDealsDealIDUpdateResponse](../../models/operations/patchcrmv3objectsdealsdealidupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCrmV3ObjectsDealsGetPage

Read a page of deals. Control what is returned via the `properties` query param.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"context"
	"log"
)

func main() {
    s := deals.New()


    operationSecurity := operations.GetCrmV3ObjectsDealsGetPageSecurity{
            Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Basic.GetCrmV3ObjectsDealsGetPage(ctx, operations.GetCrmV3ObjectsDealsGetPageRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionResponseSimplePublicObjectWithAssociationsForwardPaging != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetCrmV3ObjectsDealsGetPageRequest](../../models/operations/getcrmv3objectsdealsgetpagerequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.GetCrmV3ObjectsDealsGetPageSecurity](../../models/operations/getcrmv3objectsdealsgetpagesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.GetCrmV3ObjectsDealsGetPageResponse](../../models/operations/getcrmv3objectsdealsgetpageresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostCrmV3ObjectsDealsCreate

Create a deal with the given properties and return a copy of the object, including the ID. Documentation and examples for creating standard deals is provided.

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


    operationSecurity := operations.PostCrmV3ObjectsDealsCreateSecurity{
            Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Basic.PostCrmV3ObjectsDealsCreate(ctx, components.SimplePublicObjectInputForCreate{
        Associations: []components.PublicAssociationsForObject{
            components.PublicAssociationsForObject{
                Types: []components.AssociationSpec{
                    components.AssociationSpec{
                        AssociationCategory: components.AssociationCategoryHubspotDefined,
                        AssociationTypeID: 2,
                    },
                },
                To: components.PublicObjectID{
                    ID: "101",
                },
            },
        },
        Properties: map[string]string{
            "amount": "1500.00",
            "dealname": "Custom data integrations",
            "pipeline": "default",
            "closedate": "2019-12-07T16:50:06.678Z",
            "dealstage": "presentationscheduled",
            "hubspot_owner_id": "910901",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.SimplePublicObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [components.SimplePublicObjectInputForCreate](../../models/components/simplepublicobjectinputforcreate.md)       | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.PostCrmV3ObjectsDealsCreateSecurity](../../models/operations/postcrmv3objectsdealscreatesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.PostCrmV3ObjectsDealsCreateResponse](../../models/operations/postcrmv3objectsdealscreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
