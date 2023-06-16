// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Storage storage
//
// swagger:model Storage
type Storage struct {

	// Free volume size.
	Free uint64 `json:"free,omitempty"`

	// Total volume size.
	Total uint64 `json:"total,omitempty"`
}

// Validate validates this storage
func (m *Storage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage based on context it is used
func (m *Storage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Storage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Storage) UnmarshalBinary(b []byte) error {
	var res Storage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
