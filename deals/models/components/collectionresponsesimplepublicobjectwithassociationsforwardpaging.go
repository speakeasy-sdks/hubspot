// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type CollectionResponseSimplePublicObjectWithAssociationsForwardPaging struct {
	Paging  *ForwardPaging                       `json:"paging,omitempty"`
	Results []SimplePublicObjectWithAssociations `json:"results"`
}

func (o *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) GetPaging() *ForwardPaging {
	if o == nil {
		return nil
	}
	return o.Paging
}

func (o *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) GetResults() []SimplePublicObjectWithAssociations {
	if o == nil {
		return []SimplePublicObjectWithAssociations{}
	}
	return o.Results
}
