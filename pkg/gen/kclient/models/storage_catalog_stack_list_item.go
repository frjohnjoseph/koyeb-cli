// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageCatalogStackListItem A shorter version of a CatalogStack for listing
//
// swagger:model storageCatalogStackListItem
type StorageCatalogStackListItem struct {

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// short description
	ShortDescription string `json:"short_description,omitempty"`

	// source control ref
	SourceControlRef *StorageSCMReference `json:"source_control_ref,omitempty"`

	// status
	Status StorageBaseCatalogStatus `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this storage catalog stack list item
func (m *StorageCatalogStackListItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceControlRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageCatalogStackListItem) validateSourceControlRef(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceControlRef) { // not required
		return nil
	}

	if m.SourceControlRef != nil {
		if err := m.SourceControlRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_control_ref")
			}
			return err
		}
	}

	return nil
}

func (m *StorageCatalogStackListItem) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageCatalogStackListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageCatalogStackListItem) UnmarshalBinary(b []byte) error {
	var res StorageCatalogStackListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
