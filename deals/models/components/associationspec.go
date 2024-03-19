// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type AssociationCategory string

const (
	AssociationCategoryHubspotDefined    AssociationCategory = "HUBSPOT_DEFINED"
	AssociationCategoryUserDefined       AssociationCategory = "USER_DEFINED"
	AssociationCategoryIntegratorDefined AssociationCategory = "INTEGRATOR_DEFINED"
)

func (e AssociationCategory) ToPointer() *AssociationCategory {
	return &e
}

func (e *AssociationCategory) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HUBSPOT_DEFINED":
		fallthrough
	case "USER_DEFINED":
		fallthrough
	case "INTEGRATOR_DEFINED":
		*e = AssociationCategory(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssociationCategory: %v", v)
	}
}

type AssociationSpec struct {
	AssociationCategory AssociationCategory `json:"associationCategory"`
	AssociationTypeID   int                 `json:"associationTypeId"`
}

func (o *AssociationSpec) GetAssociationCategory() AssociationCategory {
	if o == nil {
		return AssociationCategory("")
	}
	return o.AssociationCategory
}

func (o *AssociationSpec) GetAssociationTypeID() int {
	if o == nil {
		return 0
	}
	return o.AssociationTypeID
}