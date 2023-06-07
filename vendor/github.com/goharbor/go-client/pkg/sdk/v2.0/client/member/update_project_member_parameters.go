// Code generated by go-swagger; DO NOT EDIT.

package member

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
	"github.com/go-openapi/swag"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// NewUpdateProjectMemberParams creates a new UpdateProjectMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProjectMemberParams() *UpdateProjectMemberParams {
	return &UpdateProjectMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProjectMemberParamsWithTimeout creates a new UpdateProjectMemberParams object
// with the ability to set a timeout on a request.
func NewUpdateProjectMemberParamsWithTimeout(timeout time.Duration) *UpdateProjectMemberParams {
	return &UpdateProjectMemberParams{
		timeout: timeout,
	}
}

// NewUpdateProjectMemberParamsWithContext creates a new UpdateProjectMemberParams object
// with the ability to set a context for a request.
func NewUpdateProjectMemberParamsWithContext(ctx context.Context) *UpdateProjectMemberParams {
	return &UpdateProjectMemberParams{
		Context: ctx,
	}
}

// NewUpdateProjectMemberParamsWithHTTPClient creates a new UpdateProjectMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProjectMemberParamsWithHTTPClient(client *http.Client) *UpdateProjectMemberParams {
	return &UpdateProjectMemberParams{
		HTTPClient: client,
	}
}

/*
UpdateProjectMemberParams contains all the parameters to send to the API endpoint

	for the update project member operation.

	Typically these are written to a http.Request.
*/
type UpdateProjectMemberParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Mid.

	   Member ID.

	   Format: int64
	*/
	Mid int64

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string

	// Role.
	Role *models.RoleRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update project member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectMemberParams) WithDefaults() *UpdateProjectMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update project member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectMemberParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := UpdateProjectMemberParams{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update project member params
func (o *UpdateProjectMemberParams) WithTimeout(timeout time.Duration) *UpdateProjectMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update project member params
func (o *UpdateProjectMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update project member params
func (o *UpdateProjectMemberParams) WithContext(ctx context.Context) *UpdateProjectMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update project member params
func (o *UpdateProjectMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update project member params
func (o *UpdateProjectMemberParams) WithHTTPClient(client *http.Client) *UpdateProjectMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update project member params
func (o *UpdateProjectMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the update project member params
func (o *UpdateProjectMemberParams) WithXIsResourceName(xIsResourceName *bool) *UpdateProjectMemberParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the update project member params
func (o *UpdateProjectMemberParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the update project member params
func (o *UpdateProjectMemberParams) WithXRequestID(xRequestID *string) *UpdateProjectMemberParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update project member params
func (o *UpdateProjectMemberParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithMid adds the mid to the update project member params
func (o *UpdateProjectMemberParams) WithMid(mid int64) *UpdateProjectMemberParams {
	o.SetMid(mid)
	return o
}

// SetMid adds the mid to the update project member params
func (o *UpdateProjectMemberParams) SetMid(mid int64) {
	o.Mid = mid
}

// WithProjectNameOrID adds the projectNameOrID to the update project member params
func (o *UpdateProjectMemberParams) WithProjectNameOrID(projectNameOrID string) *UpdateProjectMemberParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the update project member params
func (o *UpdateProjectMemberParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WithRole adds the role to the update project member params
func (o *UpdateProjectMemberParams) WithRole(role *models.RoleRequest) *UpdateProjectMemberParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the update project member params
func (o *UpdateProjectMemberParams) SetRole(role *models.RoleRequest) {
	o.Role = role
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProjectMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param mid
	if err := r.SetPathParam("mid", swag.FormatInt64(o.Mid)); err != nil {
		return err
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}
	if o.Role != nil {
		if err := r.SetBodyParam(o.Role); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
