// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AssociatedID struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (o *AssociatedID) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AssociatedID) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}