// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewListSiteSubmissionsParams creates a new ListSiteSubmissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSiteSubmissionsParams() *ListSiteSubmissionsParams {
	return &ListSiteSubmissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSiteSubmissionsParamsWithTimeout creates a new ListSiteSubmissionsParams object
// with the ability to set a timeout on a request.
func NewListSiteSubmissionsParamsWithTimeout(timeout time.Duration) *ListSiteSubmissionsParams {
	return &ListSiteSubmissionsParams{
		timeout: timeout,
	}
}

// NewListSiteSubmissionsParamsWithContext creates a new ListSiteSubmissionsParams object
// with the ability to set a context for a request.
func NewListSiteSubmissionsParamsWithContext(ctx context.Context) *ListSiteSubmissionsParams {
	return &ListSiteSubmissionsParams{
		Context: ctx,
	}
}

// NewListSiteSubmissionsParamsWithHTTPClient creates a new ListSiteSubmissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSiteSubmissionsParamsWithHTTPClient(client *http.Client) *ListSiteSubmissionsParams {
	return &ListSiteSubmissionsParams{
		HTTPClient: client,
	}
}

/*
ListSiteSubmissionsParams contains all the parameters to send to the API endpoint

	for the list site submissions operation.

	Typically these are written to a http.Request.
*/
type ListSiteSubmissionsParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	// SiteID.
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list site submissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSiteSubmissionsParams) WithDefaults() *ListSiteSubmissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list site submissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSiteSubmissionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list site submissions params
func (o *ListSiteSubmissionsParams) WithTimeout(timeout time.Duration) *ListSiteSubmissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list site submissions params
func (o *ListSiteSubmissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list site submissions params
func (o *ListSiteSubmissionsParams) WithContext(ctx context.Context) *ListSiteSubmissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list site submissions params
func (o *ListSiteSubmissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list site submissions params
func (o *ListSiteSubmissionsParams) WithHTTPClient(client *http.Client) *ListSiteSubmissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list site submissions params
func (o *ListSiteSubmissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the list site submissions params
func (o *ListSiteSubmissionsParams) WithPage(page *int32) *ListSiteSubmissionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list site submissions params
func (o *ListSiteSubmissionsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the list site submissions params
func (o *ListSiteSubmissionsParams) WithPerPage(perPage *int32) *ListSiteSubmissionsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the list site submissions params
func (o *ListSiteSubmissionsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithSiteID adds the siteID to the list site submissions params
func (o *ListSiteSubmissionsParams) WithSiteID(siteID string) *ListSiteSubmissionsParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the list site submissions params
func (o *ListSiteSubmissionsParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *ListSiteSubmissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
