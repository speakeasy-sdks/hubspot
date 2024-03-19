// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
)

type PostCrmV3ObjectsDealsSearchDoSearchSecurity struct {
	Oauth2      *string `security:"scheme,type=oauth2,name=Authorization"`
	PrivateApps *string `security:"scheme,type=apiKey,subtype=header,name=private-app"`
}

func (o *PostCrmV3ObjectsDealsSearchDoSearchSecurity) GetOauth2() *string {
	if o == nil {
		return nil
	}
	return o.Oauth2
}

func (o *PostCrmV3ObjectsDealsSearchDoSearchSecurity) GetPrivateApps() *string {
	if o == nil {
		return nil
	}
	return o.PrivateApps
}

type PostCrmV3ObjectsDealsSearchDoSearchResponse struct {
	HTTPMeta components.HTTPMetadata
	// successful operation
	CollectionResponseWithTotalSimplePublicObjectForwardPaging *components.CollectionResponseWithTotalSimplePublicObjectForwardPaging
	Body                                                       []byte
}

func (o *PostCrmV3ObjectsDealsSearchDoSearchResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostCrmV3ObjectsDealsSearchDoSearchResponse) GetCollectionResponseWithTotalSimplePublicObjectForwardPaging() *components.CollectionResponseWithTotalSimplePublicObjectForwardPaging {
	if o == nil {
		return nil
	}
	return o.CollectionResponseWithTotalSimplePublicObjectForwardPaging
}

func (o *PostCrmV3ObjectsDealsSearchDoSearchResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}