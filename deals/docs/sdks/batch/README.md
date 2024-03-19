# Batch
(*Batch*)

### Available Operations

* [PostCrmV3ObjectsDealsBatchReadRead](#postcrmv3objectsdealsbatchreadread) - Read a batch of deals by internal ID, or unique property values
* [PostCrmV3ObjectsDealsBatchArchiveArchive](#postcrmv3objectsdealsbatcharchivearchive) - Archive a batch of deals by ID
* [PostCrmV3ObjectsDealsBatchCreateCreate](#postcrmv3objectsdealsbatchcreatecreate) - Create a batch of deals
* [PostCrmV3ObjectsDealsBatchUpdateUpdate](#postcrmv3objectsdealsbatchupdateupdate) - Update a batch of deals

## PostCrmV3ObjectsDealsBatchReadRead

Read a batch of deals by internal ID, or unique property values

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

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `security`                                                                                                                     | [operations.PostCrmV3ObjectsDealsBatchReadReadSecurity](../../models/operations/postcrmv3objectsdealsbatchreadreadsecurity.md) | :heavy_check_mark:                                                                                                             | The security requirements to use for the request.                                                                              |
| `batchReadInputSimplePublicObjectID`                                                                                           | [components.BatchReadInputSimplePublicObjectID](../../models/components/batchreadinputsimplepublicobjectid.md)                 | :heavy_check_mark:                                                                                                             | N/A                                                                                                                            |
| `archived`                                                                                                                     | **bool*                                                                                                                        | :heavy_minus_sign:                                                                                                             | Whether to return only results that have been archived.                                                                        |


### Response

**[*operations.PostCrmV3ObjectsDealsBatchReadReadResponse](../../models/operations/postcrmv3objectsdealsbatchreadreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostCrmV3ObjectsDealsBatchArchiveArchive

Archive a batch of deals by ID

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


    operationSecurity := operations.PostCrmV3ObjectsDealsBatchArchiveArchiveSecurity{
            Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Batch.PostCrmV3ObjectsDealsBatchArchiveArchive(ctx, components.BatchInputSimplePublicObjectID{
        Inputs: []components.SimplePublicObjectID{
            components.SimplePublicObjectID{
                ID: "<id>",
            },
        },
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [components.BatchInputSimplePublicObjectID](../../models/components/batchinputsimplepublicobjectid.md)                                     | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `security`                                                                                                                                 | [operations.PostCrmV3ObjectsDealsBatchArchiveArchiveSecurity](../../models/operations/postcrmv3objectsdealsbatcharchivearchivesecurity.md) | :heavy_check_mark:                                                                                                                         | The security requirements to use for the request.                                                                                          |


### Response

**[*operations.PostCrmV3ObjectsDealsBatchArchiveArchiveResponse](../../models/operations/postcrmv3objectsdealsbatcharchivearchiveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostCrmV3ObjectsDealsBatchCreateCreate

Create a batch of deals

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


    operationSecurity := operations.PostCrmV3ObjectsDealsBatchCreateCreateSecurity{
            Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Batch.PostCrmV3ObjectsDealsBatchCreateCreate(ctx, components.BatchInputSimplePublicObjectInputForCreate{
        Inputs: []components.SimplePublicObjectInputForCreate{
            components.SimplePublicObjectInputForCreate{
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
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.BatchResponseSimplePublicObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [components.BatchInputSimplePublicObjectInputForCreate](../../models/components/batchinputsimplepublicobjectinputforcreate.md)         | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `security`                                                                                                                             | [operations.PostCrmV3ObjectsDealsBatchCreateCreateSecurity](../../models/operations/postcrmv3objectsdealsbatchcreatecreatesecurity.md) | :heavy_check_mark:                                                                                                                     | The security requirements to use for the request.                                                                                      |


### Response

**[*operations.PostCrmV3ObjectsDealsBatchCreateCreateResponse](../../models/operations/postcrmv3objectsdealsbatchcreatecreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostCrmV3ObjectsDealsBatchUpdateUpdate

Update a batch of deals

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


    operationSecurity := operations.PostCrmV3ObjectsDealsBatchUpdateUpdateSecurity{
            Oauth2: deals.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Batch.PostCrmV3ObjectsDealsBatchUpdateUpdate(ctx, components.BatchInputSimplePublicObjectBatchInput{
        Inputs: []components.SimplePublicObjectBatchInput{
            components.SimplePublicObjectBatchInput{
                IDProperty: deals.String("my_unique_property_name"),
                ID: "1",
                Properties: map[string]string{
                    "amount": "1500.00",
                    "dealname": "Custom data integrations",
                    "pipeline": "default",
                    "closedate": "2019-12-07T16:50:06.678Z",
                    "dealstage": "presentationscheduled",
                    "hubspot_owner_id": "910901",
                },
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.BatchResponseSimplePublicObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [components.BatchInputSimplePublicObjectBatchInput](../../models/components/batchinputsimplepublicobjectbatchinput.md)                 | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `security`                                                                                                                             | [operations.PostCrmV3ObjectsDealsBatchUpdateUpdateSecurity](../../models/operations/postcrmv3objectsdealsbatchupdateupdatesecurity.md) | :heavy_check_mark:                                                                                                                     | The security requirements to use for the request.                                                                                      |


### Response

**[*operations.PostCrmV3ObjectsDealsBatchUpdateUpdateResponse](../../models/operations/postcrmv3objectsdealsbatchupdateupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
