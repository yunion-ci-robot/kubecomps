// Code generated by go-swagger; DO NOT EDIT.

package usergroup

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
)

// NewListUserGroupsParams creates a new ListUserGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListUserGroupsParams() *ListUserGroupsParams {
	return &ListUserGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListUserGroupsParamsWithTimeout creates a new ListUserGroupsParams object
// with the ability to set a timeout on a request.
func NewListUserGroupsParamsWithTimeout(timeout time.Duration) *ListUserGroupsParams {
	return &ListUserGroupsParams{
		timeout: timeout,
	}
}

// NewListUserGroupsParamsWithContext creates a new ListUserGroupsParams object
// with the ability to set a context for a request.
func NewListUserGroupsParamsWithContext(ctx context.Context) *ListUserGroupsParams {
	return &ListUserGroupsParams{
		Context: ctx,
	}
}

// NewListUserGroupsParamsWithHTTPClient creates a new ListUserGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListUserGroupsParamsWithHTTPClient(client *http.Client) *ListUserGroupsParams {
	return &ListUserGroupsParams{
		HTTPClient: client,
	}
}

/*
ListUserGroupsParams contains all the parameters to send to the API endpoint

	for the list user groups operation.

	Typically these are written to a http.Request.
*/
type ListUserGroupsParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* GroupName.

	   group name need to search, fuzzy matches
	*/
	GroupName *string

	/* LdapGroupDn.

	   search with ldap group DN
	*/
	LdapGroupDn *string

	/* Page.

	   The page number

	   Format: int64
	   Default: 1
	*/
	Page *int64

	/* PageSize.

	   The size of per page

	   Format: int64
	   Default: 10
	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUserGroupsParams) WithDefaults() *ListUserGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUserGroupsParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := ListUserGroupsParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list user groups params
func (o *ListUserGroupsParams) WithTimeout(timeout time.Duration) *ListUserGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list user groups params
func (o *ListUserGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list user groups params
func (o *ListUserGroupsParams) WithContext(ctx context.Context) *ListUserGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list user groups params
func (o *ListUserGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list user groups params
func (o *ListUserGroupsParams) WithHTTPClient(client *http.Client) *ListUserGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list user groups params
func (o *ListUserGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list user groups params
func (o *ListUserGroupsParams) WithXRequestID(xRequestID *string) *ListUserGroupsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list user groups params
func (o *ListUserGroupsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithGroupName adds the groupName to the list user groups params
func (o *ListUserGroupsParams) WithGroupName(groupName *string) *ListUserGroupsParams {
	o.SetGroupName(groupName)
	return o
}

// SetGroupName adds the groupName to the list user groups params
func (o *ListUserGroupsParams) SetGroupName(groupName *string) {
	o.GroupName = groupName
}

// WithLdapGroupDn adds the ldapGroupDn to the list user groups params
func (o *ListUserGroupsParams) WithLdapGroupDn(ldapGroupDn *string) *ListUserGroupsParams {
	o.SetLdapGroupDn(ldapGroupDn)
	return o
}

// SetLdapGroupDn adds the ldapGroupDn to the list user groups params
func (o *ListUserGroupsParams) SetLdapGroupDn(ldapGroupDn *string) {
	o.LdapGroupDn = ldapGroupDn
}

// WithPage adds the page to the list user groups params
func (o *ListUserGroupsParams) WithPage(page *int64) *ListUserGroupsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list user groups params
func (o *ListUserGroupsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list user groups params
func (o *ListUserGroupsParams) WithPageSize(pageSize *int64) *ListUserGroupsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list user groups params
func (o *ListUserGroupsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *ListUserGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.GroupName != nil {

		// query param group_name
		var qrGroupName string

		if o.GroupName != nil {
			qrGroupName = *o.GroupName
		}
		qGroupName := qrGroupName
		if qGroupName != "" {

			if err := r.SetQueryParam("group_name", qGroupName); err != nil {
				return err
			}
		}
	}

	if o.LdapGroupDn != nil {

		// query param ldap_group_dn
		var qrLdapGroupDn string

		if o.LdapGroupDn != nil {
			qrLdapGroupDn = *o.LdapGroupDn
		}
		qLdapGroupDn := qrLdapGroupDn
		if qLdapGroupDn != "" {

			if err := r.SetQueryParam("ldap_group_dn", qLdapGroupDn); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
