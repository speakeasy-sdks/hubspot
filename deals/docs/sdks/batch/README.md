# Batch
(*Batch*)

## Overview

### Available Operations

* [PostCrmV3ObjectsDealsBatchReadRead](#postcrmv3objectsdealsbatchreadread) - Read a batch of deals by internal ID, or unique property values
* [PostCrmV3ObjectsDealsBatchArchiveArchive](#postcrmv3objectsdealsbatcharchivearchive) - Archive a batch of deals by ID
* [PostCrmV3ObjectsDealsBatchCreateCreate](#postcrmv3objectsdealsbatchcreatecreate) - Create a batch of deals
* [PostCrmV3ObjectsDealsBatchUpdateUpdate](#postcrmv3objectsdealsbatchupdateupdate) - Update a batch of deals

## PostCrmV3ObjectsDealsBatchReadRead

Read a batch of deals by internal ID, or unique property values

### Example Usage

<!-- UsageSnippet language="go" operationID="post-/crm/v3/objects/deals/batch/read_read" method="post" path="/crm/v3/objects/deals/batch/read" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-sdks/hubspot-go/deals"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/operations"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := deals.New()

    res, err := s.Batch.PostCrmV3ObjectsDealsBatchReadRead(ctx, operations.PostCrmV3ObjectsDealsBatchReadReadSecurity{
        Oauth2: deals.Pointer("<YOUR_OAUTH2_HERE>"),
    }, components.BatchReadInputSimplePublicObjectID{
        PropertiesWithHistory: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
        },
        Inputs: []components.SimplePublicObjectID{},
        Properties: []string{
            "<value 1>",
            "<value 2>",
        },
    }, deals.Pointer(false))
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
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.PostCrmV3ObjectsDealsBatchReadReadResponse](../../models/operations/postcrmv3objectsdealsbatchreadreadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PostCrmV3ObjectsDealsBatchArchiveArchive

Archive a batch of deals by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="post-/crm/v3/objects/deals/batch/archive_archive" method="post" path="/crm/v3/objects/deals/batch/archive" -->
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

    res, err := s.Batch.PostCrmV3ObjectsDealsBatchArchiveArchive(ctx, components.BatchInputSimplePublicObjectID{
        Inputs: []components.SimplePublicObjectID{},
    }, operations.PostCrmV3ObjectsDealsBatchArchiveArchiveSecurity{
        Oauth2: deals.Pointer("<YOUR_OAUTH2_HERE>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Body != nil {
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
| `opts`                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.PostCrmV3ObjectsDealsBatchArchiveArchiveResponse](../../models/operations/postcrmv3objectsdealsbatcharchivearchiveresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PostCrmV3ObjectsDealsBatchCreateCreate

Create a batch of deals

### Example Usage

<!-- UsageSnippet language="go" operationID="post-/crm/v3/objects/deals/batch/create_create" method="post" path="/crm/v3/objects/deals/batch/create" -->
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
    }, operations.PostCrmV3ObjectsDealsBatchCreateCreateSecurity{
        Oauth2: deals.Pointer("<YOUR_OAUTH2_HERE>"),
    })
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
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.PostCrmV3ObjectsDealsBatchCreateCreateResponse](../../models/operations/postcrmv3objectsdealsbatchcreatecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PostCrmV3ObjectsDealsBatchUpdateUpdate

Update a batch of deals

### Example Usage

<!-- UsageSnippet language="go" operationID="post-/crm/v3/objects/deals/batch/update_update" method="post" path="/crm/v3/objects/deals/batch/update" -->
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

    res, err := s.Batch.PostCrmV3ObjectsDealsBatchUpdateUpdate(ctx, components.BatchInputSimplePublicObjectBatchInput{
        Inputs: []components.SimplePublicObjectBatchInput{
            components.SimplePublicObjectBatchInput{
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
    }, operations.PostCrmV3ObjectsDealsBatchUpdateUpdateSecurity{
        Oauth2: deals.Pointer("<YOUR_OAUTH2_HERE>"),
    })
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
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.PostCrmV3ObjectsDealsBatchUpdateUpdateResponse](../../models/operations/postcrmv3objectsdealsbatchupdateupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |