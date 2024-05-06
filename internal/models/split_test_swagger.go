// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SplitTest split test
//
// swagger:model splitTest
type SplitTest struct {

	// active
	Active bool `json:"active,omitempty"`

	// branches
	Branches []interface{} `json:"branches"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// site id
	SiteID string `json:"site_id,omitempty"`

	// unpublished at
	UnpublishedAt string `json:"unpublished_at,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this split test
func (m *SplitTest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this split test based on context it is used
func (m *SplitTest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SplitTest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SplitTest) UnmarshalBinary(b []byte) error {
	var res SplitTest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
