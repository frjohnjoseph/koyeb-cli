// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// NewUpdateStore2Params creates a new UpdateStore2Params object
// with the default values initialized.
func NewUpdateStore2Params() *UpdateStore2Params {
	var ()
	return &UpdateStore2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateStore2ParamsWithTimeout creates a new UpdateStore2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateStore2ParamsWithTimeout(timeout time.Duration) *UpdateStore2Params {
	var ()
	return &UpdateStore2Params{

		timeout: timeout,
	}
}

// NewUpdateStore2ParamsWithContext creates a new UpdateStore2Params object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateStore2ParamsWithContext(ctx context.Context) *UpdateStore2Params {
	var ()
	return &UpdateStore2Params{

		Context: ctx,
	}
}

// NewUpdateStore2ParamsWithHTTPClient creates a new UpdateStore2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateStore2ParamsWithHTTPClient(client *http.Client) *UpdateStore2Params {
	var ()
	return &UpdateStore2Params{
		HTTPClient: client,
	}
}

/*UpdateStore2Params contains all the parameters to send to the API endpoint
for the update store2 operation typically these are written to a http.Request
*/
type UpdateStore2Params struct {

	/*Body*/
	Body *models.StorageStore
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update store2 params
func (o *UpdateStore2Params) WithTimeout(timeout time.Duration) *UpdateStore2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update store2 params
func (o *UpdateStore2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update store2 params
func (o *UpdateStore2Params) WithContext(ctx context.Context) *UpdateStore2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update store2 params
func (o *UpdateStore2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update store2 params
func (o *UpdateStore2Params) WithHTTPClient(client *http.Client) *UpdateStore2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update store2 params
func (o *UpdateStore2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update store2 params
func (o *UpdateStore2Params) WithBody(body *models.StorageStore) *UpdateStore2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update store2 params
func (o *UpdateStore2Params) SetBody(body *models.StorageStore) {
	o.Body = body
}

// WithID adds the id to the update store2 params
func (o *UpdateStore2Params) WithID(id string) *UpdateStore2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the update store2 params
func (o *UpdateStore2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateStore2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
