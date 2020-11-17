// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewOrganizationGetOrganizationParams creates a new OrganizationGetOrganizationParams object
// with the default values initialized.
func NewOrganizationGetOrganizationParams() *OrganizationGetOrganizationParams {
	var ()
	return &OrganizationGetOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationGetOrganizationParamsWithTimeout creates a new OrganizationGetOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrganizationGetOrganizationParamsWithTimeout(timeout time.Duration) *OrganizationGetOrganizationParams {
	var ()
	return &OrganizationGetOrganizationParams{

		timeout: timeout,
	}
}

// NewOrganizationGetOrganizationParamsWithContext creates a new OrganizationGetOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrganizationGetOrganizationParamsWithContext(ctx context.Context) *OrganizationGetOrganizationParams {
	var ()
	return &OrganizationGetOrganizationParams{

		Context: ctx,
	}
}

// NewOrganizationGetOrganizationParamsWithHTTPClient creates a new OrganizationGetOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrganizationGetOrganizationParamsWithHTTPClient(client *http.Client) *OrganizationGetOrganizationParams {
	var ()
	return &OrganizationGetOrganizationParams{
		HTTPClient: client,
	}
}

/*OrganizationGetOrganizationParams contains all the parameters to send to the API endpoint
for the organization get organization operation typically these are written to a http.Request
*/
type OrganizationGetOrganizationParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the organization get organization params
func (o *OrganizationGetOrganizationParams) WithTimeout(timeout time.Duration) *OrganizationGetOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organization get organization params
func (o *OrganizationGetOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organization get organization params
func (o *OrganizationGetOrganizationParams) WithContext(ctx context.Context) *OrganizationGetOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organization get organization params
func (o *OrganizationGetOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organization get organization params
func (o *OrganizationGetOrganizationParams) WithHTTPClient(client *http.Client) *OrganizationGetOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organization get organization params
func (o *OrganizationGetOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the organization get organization params
func (o *OrganizationGetOrganizationParams) WithID(id string) *OrganizationGetOrganizationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organization get organization params
func (o *OrganizationGetOrganizationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationGetOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
