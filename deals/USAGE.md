<!-- Start SDK Example Usage [usage] -->
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
	security := operations.PostCrmV3ObjectsDealsBatchReadReadSecurity{
		Oauth2: deals.String("<YOUR_OAUTH2_HERE>"),
	}

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
	ctx := context.Background()
	res, err := s.Batch.PostCrmV3ObjectsDealsBatchReadRead(ctx, security, batchReadInputSimplePublicObjectID, archived)
	if err != nil {
		log.Fatal(err)
	}
	if res.BatchResponseSimplePublicObject != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->