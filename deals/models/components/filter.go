// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Operator - null
type Operator string

const (
	OperatorEq               Operator = "EQ"
	OperatorNeq              Operator = "NEQ"
	OperatorLt               Operator = "LT"
	OperatorLte              Operator = "LTE"
	OperatorGt               Operator = "GT"
	OperatorGte              Operator = "GTE"
	OperatorBetween          Operator = "BETWEEN"
	OperatorIn               Operator = "IN"
	OperatorNotIn            Operator = "NOT_IN"
	OperatorHasProperty      Operator = "HAS_PROPERTY"
	OperatorNotHasProperty   Operator = "NOT_HAS_PROPERTY"
	OperatorContainsToken    Operator = "CONTAINS_TOKEN"
	OperatorNotContainsToken Operator = "NOT_CONTAINS_TOKEN"
)

func (e Operator) ToPointer() *Operator {
	return &e
}

func (e *Operator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EQ":
		fallthrough
	case "NEQ":
		fallthrough
	case "LT":
		fallthrough
	case "LTE":
		fallthrough
	case "GT":
		fallthrough
	case "GTE":
		fallthrough
	case "BETWEEN":
		fallthrough
	case "IN":
		fallthrough
	case "NOT_IN":
		fallthrough
	case "HAS_PROPERTY":
		fallthrough
	case "NOT_HAS_PROPERTY":
		fallthrough
	case "CONTAINS_TOKEN":
		fallthrough
	case "NOT_CONTAINS_TOKEN":
		*e = Operator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Operator: %v", v)
	}
}

type Filter struct {
	HighValue    *string  `json:"highValue,omitempty"`
	PropertyName string   `json:"propertyName"`
	Values       []string `json:"values,omitempty"`
	Value        *string  `json:"value,omitempty"`
	// null
	Operator Operator `json:"operator"`
}

func (o *Filter) GetHighValue() *string {
	if o == nil {
		return nil
	}
	return o.HighValue
}

func (o *Filter) GetPropertyName() string {
	if o == nil {
		return ""
	}
	return o.PropertyName
}

func (o *Filter) GetValues() []string {
	if o == nil {
		return nil
	}
	return o.Values
}

func (o *Filter) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *Filter) GetOperator() Operator {
	if o == nil {
		return Operator("")
	}
	return o.Operator
}
