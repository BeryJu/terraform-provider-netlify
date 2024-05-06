// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Purge purge
//
// swagger:model purge
type Purge struct {

	// cache tags
	CacheTags []string `json:"cache_tags"`

	// site id
	SiteID string `json:"site_id,omitempty"`

	// site slug
	SiteSlug string `json:"site_slug,omitempty"`
}

// Validate validates this purge
func (m *Purge) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this purge based on context it is used
func (m *Purge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Purge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Purge) UnmarshalBinary(b []byte) error {
	var res Purge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
