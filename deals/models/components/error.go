// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Error struct {
	// A specific category that contains more specific detail about the error
	SubCategory *string
	// Context about the error condition
	Context map[string][]string
	// A unique identifier for the request. Include this value with any error reports or support tickets
	CorrelationID string
	// A map of link names to associated URIs containing documentation about the error or recommended remediation steps
	Links map[string]string
	// A human readable message describing the error along with remediation steps where appropriate
	Message string
	// The error category
	Category string
	// further information about the error
	Errors []ErrorDetail
}

func (o *Error) GetSubCategory() *string {
	if o == nil {
		return nil
	}
	return o.SubCategory
}

func (o *Error) GetContext() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Context
}

func (o *Error) GetCorrelationID() string {
	if o == nil {
		return ""
	}
	return o.CorrelationID
}

func (o *Error) GetLinks() map[string]string {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *Error) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *Error) GetCategory() string {
	if o == nil {
		return ""
	}
	return o.Category
}

func (o *Error) GetErrors() []ErrorDetail {
	if o == nil {
		return nil
	}
	return o.Errors
}
