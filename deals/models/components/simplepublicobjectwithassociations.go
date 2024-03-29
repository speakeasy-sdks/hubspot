// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/hubspot-go/deals/internal/utils"
	"time"
)

type SimplePublicObjectWithAssociations struct {
	Associations          map[string]CollectionResponseAssociatedID `json:"associations,omitempty"`
	CreatedAt             time.Time                                 `json:"createdAt"`
	Archived              *bool                                     `json:"archived,omitempty"`
	ArchivedAt            *time.Time                                `json:"archivedAt,omitempty"`
	PropertiesWithHistory map[string][]ValueWithTimestamp           `json:"propertiesWithHistory,omitempty"`
	ID                    string                                    `json:"id"`
	Properties            map[string]string                         `json:"properties"`
	UpdatedAt             time.Time                                 `json:"updatedAt"`
}

func (s SimplePublicObjectWithAssociations) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SimplePublicObjectWithAssociations) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SimplePublicObjectWithAssociations) GetAssociations() map[string]CollectionResponseAssociatedID {
	if o == nil {
		return nil
	}
	return o.Associations
}

func (o *SimplePublicObjectWithAssociations) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *SimplePublicObjectWithAssociations) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *SimplePublicObjectWithAssociations) GetArchivedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *SimplePublicObjectWithAssociations) GetPropertiesWithHistory() map[string][]ValueWithTimestamp {
	if o == nil {
		return nil
	}
	return o.PropertiesWithHistory
}

func (o *SimplePublicObjectWithAssociations) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SimplePublicObjectWithAssociations) GetProperties() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Properties
}

func (o *SimplePublicObjectWithAssociations) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
