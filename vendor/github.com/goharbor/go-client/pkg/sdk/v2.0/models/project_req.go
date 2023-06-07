// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProjectReq project req
//
// swagger:model ProjectReq
type ProjectReq struct {

	// The CVE allowlist of the project.
	CVEAllowlist *CVEAllowlist `json:"cve_allowlist,omitempty"`

	// The metadata of the project.
	Metadata *ProjectMetadata `json:"metadata,omitempty"`

	// The name of the project.
	// Max Length: 255
	ProjectName string `json:"project_name,omitempty"`

	// deprecated, reserved for project creation in replication
	Public *bool `json:"public,omitempty"`

	// The ID of referenced registry when creating the proxy cache project
	RegistryID *int64 `json:"registry_id,omitempty"`

	// The storage quota of the project.
	StorageLimit *int64 `json:"storage_limit,omitempty"`
}

// Validate validates this project req
func (m *ProjectReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCVEAllowlist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectReq) validateCVEAllowlist(formats strfmt.Registry) error {
	if swag.IsZero(m.CVEAllowlist) { // not required
		return nil
	}

	if m.CVEAllowlist != nil {
		if err := m.CVEAllowlist.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cve_allowlist")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cve_allowlist")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectReq) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectReq) validateProjectName(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectName) { // not required
		return nil
	}

	if err := validate.MaxLength("project_name", "body", m.ProjectName, 255); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this project req based on the context it is used
func (m *ProjectReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCVEAllowlist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectReq) contextValidateCVEAllowlist(ctx context.Context, formats strfmt.Registry) error {

	if m.CVEAllowlist != nil {
		if err := m.CVEAllowlist.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cve_allowlist")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cve_allowlist")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectReq) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectReq) UnmarshalBinary(b []byte) error {
	var res ProjectReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
