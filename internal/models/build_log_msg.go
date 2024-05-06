// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildLogMsg build log msg
//
// swagger:model buildLogMsg
type BuildLogMsg struct {

	// error
	Error bool `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// section
	// Enum: [initializing building deploying cleanup postprocessing]
	Section string `json:"section,omitempty"`
}

// Validate validates this build log msg
func (m *BuildLogMsg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var buildLogMsgTypeSectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["initializing","building","deploying","cleanup","postprocessing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buildLogMsgTypeSectionPropEnum = append(buildLogMsgTypeSectionPropEnum, v)
	}
}

const (

	// BuildLogMsgSectionInitializing captures enum value "initializing"
	BuildLogMsgSectionInitializing string = "initializing"

	// BuildLogMsgSectionBuilding captures enum value "building"
	BuildLogMsgSectionBuilding string = "building"

	// BuildLogMsgSectionDeploying captures enum value "deploying"
	BuildLogMsgSectionDeploying string = "deploying"

	// BuildLogMsgSectionCleanup captures enum value "cleanup"
	BuildLogMsgSectionCleanup string = "cleanup"

	// BuildLogMsgSectionPostprocessing captures enum value "postprocessing"
	BuildLogMsgSectionPostprocessing string = "postprocessing"
)

// prop value enum
func (m *BuildLogMsg) validateSectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, buildLogMsgTypeSectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BuildLogMsg) validateSection(formats strfmt.Registry) error {
	if swag.IsZero(m.Section) { // not required
		return nil
	}

	// value enum
	if err := m.validateSectionEnum("section", "body", m.Section); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this build log msg based on context it is used
func (m *BuildLogMsg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BuildLogMsg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildLogMsg) UnmarshalBinary(b []byte) error {
	var res BuildLogMsg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
