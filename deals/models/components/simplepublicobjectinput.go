// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type SimplePublicObjectInput struct {
	Properties map[string]string `json:"properties"`
}

func (o *SimplePublicObjectInput) GetProperties() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Properties
}
