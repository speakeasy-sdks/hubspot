// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type PublicGdprDeleteInput struct {
	IDProperty *string `json:"idProperty,omitempty"`
	ObjectID   string  `json:"objectId"`
}

func (o *PublicGdprDeleteInput) GetIDProperty() *string {
	if o == nil {
		return nil
	}
	return o.IDProperty
}

func (o *PublicGdprDeleteInput) GetObjectID() string {
	if o == nil {
		return ""
	}
	return o.ObjectID
}