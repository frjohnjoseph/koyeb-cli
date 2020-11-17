// Code generated by go-swagger; DO NOT EDIT.

package functions

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

// NewFunctionsCancelFunctionExecutionParams creates a new FunctionsCancelFunctionExecutionParams object
// with the default values initialized.
func NewFunctionsCancelFunctionExecutionParams() *FunctionsCancelFunctionExecutionParams {
	var ()
	return &FunctionsCancelFunctionExecutionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFunctionsCancelFunctionExecutionParamsWithTimeout creates a new FunctionsCancelFunctionExecutionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFunctionsCancelFunctionExecutionParamsWithTimeout(timeout time.Duration) *FunctionsCancelFunctionExecutionParams {
	var ()
	return &FunctionsCancelFunctionExecutionParams{

		timeout: timeout,
	}
}

// NewFunctionsCancelFunctionExecutionParamsWithContext creates a new FunctionsCancelFunctionExecutionParams object
// with the default values initialized, and the ability to set a context for a request
func NewFunctionsCancelFunctionExecutionParamsWithContext(ctx context.Context) *FunctionsCancelFunctionExecutionParams {
	var ()
	return &FunctionsCancelFunctionExecutionParams{

		Context: ctx,
	}
}

// NewFunctionsCancelFunctionExecutionParamsWithHTTPClient creates a new FunctionsCancelFunctionExecutionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFunctionsCancelFunctionExecutionParamsWithHTTPClient(client *http.Client) *FunctionsCancelFunctionExecutionParams {
	var ()
	return &FunctionsCancelFunctionExecutionParams{
		HTTPClient: client,
	}
}

/*FunctionsCancelFunctionExecutionParams contains all the parameters to send to the API endpoint
for the functions cancel function execution operation typically these are written to a http.Request
*/
type FunctionsCancelFunctionExecutionParams struct {

	/*RunID*/
	RunID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the functions cancel function execution params
func (o *FunctionsCancelFunctionExecutionParams) WithTimeout(timeout time.Duration) *FunctionsCancelFunctionExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the functions cancel function execution params
func (o *FunctionsCancelFunctionExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the functions cancel function execution params
func (o *FunctionsCancelFunctionExecutionParams) WithContext(ctx context.Context) *FunctionsCancelFunctionExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the functions cancel function execution params
func (o *FunctionsCancelFunctionExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the functions cancel function execution params
func (o *FunctionsCancelFunctionExecutionParams) WithHTTPClient(client *http.Client) *FunctionsCancelFunctionExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the functions cancel function execution params
func (o *FunctionsCancelFunctionExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRunID adds the runID to the functions cancel function execution params
func (o *FunctionsCancelFunctionExecutionParams) WithRunID(runID string) *FunctionsCancelFunctionExecutionParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the functions cancel function execution params
func (o *FunctionsCancelFunctionExecutionParams) SetRunID(runID string) {
	o.RunID = runID
}

// WithStackID adds the stackID to the functions cancel function execution params
func (o *FunctionsCancelFunctionExecutionParams) WithStackID(stackID string) *FunctionsCancelFunctionExecutionParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the functions cancel function execution params
func (o *FunctionsCancelFunctionExecutionParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *FunctionsCancelFunctionExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param run_id
	if err := r.SetPathParam("run_id", o.RunID); err != nil {
		return err
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
