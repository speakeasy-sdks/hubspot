// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Paging struct {
	Next *NextPage     `json:"next,omitempty"`
	Prev *PreviousPage `json:"prev,omitempty"`
}

func (o *Paging) GetNext() *NextPage {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *Paging) GetPrev() *PreviousPage {
	if o == nil {
		return nil
	}
	return o.Prev
}