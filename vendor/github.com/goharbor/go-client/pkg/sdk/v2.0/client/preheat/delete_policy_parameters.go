// Code generated by go-swagger; DO NOT EDIT.

package preheat

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

// NewDeletePolicyParams creates a new DeletePolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePolicyParams() *DeletePolicyParams {
	return &DeletePolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePolicyParamsWithTimeout creates a new DeletePolicyParams object
// with the ability to set a timeout on a request.
func NewDeletePolicyParamsWithTimeout(timeout time.Duration) *DeletePolicyParams {
	return &DeletePolicyParams{
		timeout: timeout,
	}
}

// NewDeletePolicyParamsWithContext creates a new DeletePolicyParams object
// with the ability to set a context for a request.
func NewDeletePolicyParamsWithContext(ctx context.Context) *DeletePolicyParams {
	return &DeletePolicyParams{
		Context: ctx,
	}
}

// NewDeletePolicyParamsWithHTTPClient creates a new DeletePolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePolicyParamsWithHTTPClient(client *http.Client) *DeletePolicyParams {
	return &DeletePolicyParams{
		HTTPClient: client,
	}
}

/*
DeletePolicyParams contains all the parameters to send to the API endpoint

	for the delete policy operation.

	Typically these are written to a http.Request.
*/
type DeletePolicyParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* PreheatPolicyName.

	   Preheat Policy Name
	*/
	PreheatPolicyName string

	/* ProjectName.

	   The name of the project
	*/
	ProjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePolicyParams) WithDefaults() *DeletePolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete policy params
func (o *DeletePolicyParams) WithTimeout(timeout time.Duration) *DeletePolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete policy params
func (o *DeletePolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete policy params
func (o *DeletePolicyParams) WithContext(ctx context.Context) *DeletePolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete policy params
func (o *DeletePolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete policy params
func (o *DeletePolicyParams) WithHTTPClient(client *http.Client) *DeletePolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete policy params
func (o *DeletePolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the delete policy params
func (o *DeletePolicyParams) WithXRequestID(xRequestID *string) *DeletePolicyParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete policy params
func (o *DeletePolicyParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPreheatPolicyName adds the preheatPolicyName to the delete policy params
func (o *DeletePolicyParams) WithPreheatPolicyName(preheatPolicyName string) *DeletePolicyParams {
	o.SetPreheatPolicyName(preheatPolicyName)
	return o
}

// SetPreheatPolicyName adds the preheatPolicyName to the delete policy params
func (o *DeletePolicyParams) SetPreheatPolicyName(preheatPolicyName string) {
	o.PreheatPolicyName = preheatPolicyName
}

// WithProjectName adds the projectName to the delete policy params
func (o *DeletePolicyParams) WithProjectName(projectName string) *DeletePolicyParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the delete policy params
func (o *DeletePolicyParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param preheat_policy_name
	if err := r.SetPathParam("preheat_policy_name", o.PreheatPolicyName); err != nil {
		return err
	}

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
