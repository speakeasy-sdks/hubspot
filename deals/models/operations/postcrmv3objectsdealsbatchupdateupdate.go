// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
)

type PostCrmV3ObjectsDealsBatchUpdateUpdateSecurity struct {
	Oauth2      *string `security:"scheme,type=oauth2,name=Authorization"`
	PrivateApps *string `security:"scheme,type=apiKey,subtype=header,name=private-app"`
}

func (o *PostCrmV3ObjectsDealsBatchUpdateUpdateSecurity) GetOauth2() *string {
	if o == nil {
		return nil
	}
	return o.Oauth2
}

func (o *PostCrmV3ObjectsDealsBatchUpdateUpdateSecurity) GetPrivateApps() *string {
	if o == nil {
		return nil
	}
	return o.PrivateApps
}

type PostCrmV3ObjectsDealsBatchUpdateUpdateResponse struct {
	HTTPMeta components.HTTPMetadata
	// successful operation
	BatchResponseSimplePublicObject *components.BatchResponseSimplePublicObject
	// multiple statuses
	BatchResponseSimplePublicObjectWithErrors *components.BatchResponseSimplePublicObjectWithErrors
	Body                                      []byte
}

func (o *PostCrmV3ObjectsDealsBatchUpdateUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostCrmV3ObjectsDealsBatchUpdateUpdateResponse) GetBatchResponseSimplePublicObject() *components.BatchResponseSimplePublicObject {
	if o == nil {
		return nil
	}
	return o.BatchResponseSimplePublicObject
}

func (o *PostCrmV3ObjectsDealsBatchUpdateUpdateResponse) GetBatchResponseSimplePublicObjectWithErrors() *components.BatchResponseSimplePublicObjectWithErrors {
	if o == nil {
		return nil
	}
	return o.BatchResponseSimplePublicObjectWithErrors
}

func (o *PostCrmV3ObjectsDealsBatchUpdateUpdateResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}
