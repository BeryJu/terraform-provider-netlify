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
)

// NewDeleteSiteSnippetParams creates a new DeleteSiteSnippetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSiteSnippetParams() *DeleteSiteSnippetParams {
	return &DeleteSiteSnippetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSiteSnippetParamsWithTimeout creates a new DeleteSiteSnippetParams object
// with the ability to set a timeout on a request.
func NewDeleteSiteSnippetParamsWithTimeout(timeout time.Duration) *DeleteSiteSnippetParams {
	return &DeleteSiteSnippetParams{
		timeout: timeout,
	}
}

// NewDeleteSiteSnippetParamsWithContext creates a new DeleteSiteSnippetParams object
// with the ability to set a context for a request.
func NewDeleteSiteSnippetParamsWithContext(ctx context.Context) *DeleteSiteSnippetParams {
	return &DeleteSiteSnippetParams{
		Context: ctx,
	}
}

// NewDeleteSiteSnippetParamsWithHTTPClient creates a new DeleteSiteSnippetParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSiteSnippetParamsWithHTTPClient(client *http.Client) *DeleteSiteSnippetParams {
	return &DeleteSiteSnippetParams{
		HTTPClient: client,
	}
}

/*
DeleteSiteSnippetParams contains all the parameters to send to the API endpoint

	for the delete site snippet operation.

	Typically these are written to a http.Request.
*/
type DeleteSiteSnippetParams struct {

	// SiteID.
	SiteID string

	// SnippetID.
	SnippetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete site snippet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSiteSnippetParams) WithDefaults() *DeleteSiteSnippetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete site snippet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSiteSnippetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete site snippet params
func (o *DeleteSiteSnippetParams) WithTimeout(timeout time.Duration) *DeleteSiteSnippetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete site snippet params
func (o *DeleteSiteSnippetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete site snippet params
func (o *DeleteSiteSnippetParams) WithContext(ctx context.Context) *DeleteSiteSnippetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete site snippet params
func (o *DeleteSiteSnippetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete site snippet params
func (o *DeleteSiteSnippetParams) WithHTTPClient(client *http.Client) *DeleteSiteSnippetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete site snippet params
func (o *DeleteSiteSnippetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the delete site snippet params
func (o *DeleteSiteSnippetParams) WithSiteID(siteID string) *DeleteSiteSnippetParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete site snippet params
func (o *DeleteSiteSnippetParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithSnippetID adds the snippetID to the delete site snippet params
func (o *DeleteSiteSnippetParams) WithSnippetID(snippetID string) *DeleteSiteSnippetParams {
	o.SetSnippetID(snippetID)
	return o
}

// SetSnippetID adds the snippetId to the delete site snippet params
func (o *DeleteSiteSnippetParams) SetSnippetID(snippetID string) {
	o.SnippetID = snippetID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSiteSnippetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	// path param snippet_id
	if err := r.SetPathParam("snippet_id", o.SnippetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
