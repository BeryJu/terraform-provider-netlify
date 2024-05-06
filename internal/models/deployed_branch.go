// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeployedBranch deployed branch
//
// swagger:model deployedBranch
type DeployedBranch struct {

	// deploy id
	DeployID string `json:"deploy_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// ssl url
	SslURL string `json:"ssl_url,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this deployed branch
func (m *DeployedBranch) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this deployed branch based on context it is used
func (m *DeployedBranch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeployedBranch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeployedBranch) UnmarshalBinary(b []byte) error {
	var res DeployedBranch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
