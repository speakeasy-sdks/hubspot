// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hubspot-go/deals/internal/utils"
	"github.com/speakeasy-sdks/hubspot-go/deals/models/components"
)

type PostCrmV3ObjectsDealsBatchReadReadSecurity struct {
	Oauth2      *string `security:"scheme,type=oauth2,name=Authorization"`
	PrivateApps *string `security:"scheme,type=apiKey,subtype=header,name=private-app"`
}

func (o *PostCrmV3ObjectsDealsBatchReadReadSecurity) GetOauth2() *string {
	if o == nil {
		return nil
	}
	return o.Oauth2
}

func (o *PostCrmV3ObjectsDealsBatchReadReadSecurity) GetPrivateApps() *string {
	if o == nil {
		return nil
	}
	return o.PrivateApps
}

type PostCrmV3ObjectsDealsBatchReadReadRequest struct {
	BatchReadInputSimplePublicObjectID components.BatchReadInputSimplePublicObjectID `request:"mediaType=application/json"`
	// Whether to return only results that have been archived.
	Archived *bool `default:"false" queryParam:"style=form,explode=true,name=archived"`
}

func (p PostCrmV3ObjectsDealsBatchReadReadRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostCrmV3ObjectsDealsBatchReadReadRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostCrmV3ObjectsDealsBatchReadReadRequest) GetBatchReadInputSimplePublicObjectID() components.BatchReadInputSimplePublicObjectID {
	if o == nil {
		return components.BatchReadInputSimplePublicObjectID{}
	}
	return o.BatchReadInputSimplePublicObjectID
}

func (o *PostCrmV3ObjectsDealsBatchReadReadRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

type PostCrmV3ObjectsDealsBatchReadReadResponse struct {
	HTTPMeta components.HTTPMetadata
	// successful operation
	BatchResponseSimplePublicObject *components.BatchResponseSimplePublicObject
	// multiple statuses
	BatchResponseSimplePublicObjectWithErrors *components.BatchResponseSimplePublicObjectWithErrors
	Body                                      []byte
}

func (o *PostCrmV3ObjectsDealsBatchReadReadResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostCrmV3ObjectsDealsBatchReadReadResponse) GetBatchResponseSimplePublicObject() *components.BatchResponseSimplePublicObject {
	if o == nil {
		return nil
	}
	return o.BatchResponseSimplePublicObject
}

func (o *PostCrmV3ObjectsDealsBatchReadReadResponse) GetBatchResponseSimplePublicObjectWithErrors() *components.BatchResponseSimplePublicObjectWithErrors {
	if o == nil {
		return nil
	}
	return o.BatchResponseSimplePublicObjectWithErrors
}

func (o *PostCrmV3ObjectsDealsBatchReadReadResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}
