// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
)

type PostCrmV3ObjectsDealsCreateSecurity struct {
	Oauth2      *string `security:"scheme,type=oauth2,name=Authorization"`
	PrivateApps *string `security:"scheme,type=apiKey,subtype=header,name=private-app"`
}

func (o *PostCrmV3ObjectsDealsCreateSecurity) GetOauth2() *string {
	if o == nil {
		return nil
	}
	return o.Oauth2
}

func (o *PostCrmV3ObjectsDealsCreateSecurity) GetPrivateApps() *string {
	if o == nil {
		return nil
	}
	return o.PrivateApps
}

type PostCrmV3ObjectsDealsCreateResponse struct {
	HTTPMeta components.HTTPMetadata
	// successful operation
	SimplePublicObject *components.SimplePublicObject
	Body               []byte
}

func (o *PostCrmV3ObjectsDealsCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostCrmV3ObjectsDealsCreateResponse) GetSimplePublicObject() *components.SimplePublicObject {
	if o == nil {
		return nil
	}
	return o.SimplePublicObject
}

func (o *PostCrmV3ObjectsDealsCreateResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}
