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
<!-- End SDK Example Usage [usage] -->