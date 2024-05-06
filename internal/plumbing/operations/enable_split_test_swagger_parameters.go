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

// NewEnableSplitTestParams creates a new EnableSplitTestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableSplitTestParams() *EnableSplitTestParams {
	return &EnableSplitTestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableSplitTestParamsWithTimeout creates a new EnableSplitTestParams object
// with the ability to set a timeout on a request.
func NewEnableSplitTestParamsWithTimeout(timeout time.Duration) *EnableSplitTestParams {
	return &EnableSplitTestParams{
		timeout: timeout,
	}
}

// NewEnableSplitTestParamsWithContext creates a new EnableSplitTestParams object
// with the ability to set a context for a request.
func NewEnableSplitTestParamsWithContext(ctx context.Context) *EnableSplitTestParams {
	return &EnableSplitTestParams{
		Context: ctx,
	}
}

// NewEnableSplitTestParamsWithHTTPClient creates a new EnableSplitTestParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableSplitTestParamsWithHTTPClient(client *http.Client) *EnableSplitTestParams {
	return &EnableSplitTestParams{
		HTTPClient: client,
	}
}

/*
EnableSplitTestParams contains all the parameters to send to the API endpoint

	for the enable split test operation.

	Typically these are written to a http.Request.
*/
type EnableSplitTestParams struct {

	// SiteID.
	SiteID string

	// SplitTestID.
	SplitTestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enable split test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableSplitTestParams) WithDefaults() *EnableSplitTestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable split test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableSplitTestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enable split test params
func (o *EnableSplitTestParams) WithTimeout(timeout time.Duration) *EnableSplitTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable split test params
func (o *EnableSplitTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable split test params
func (o *EnableSplitTestParams) WithContext(ctx context.Context) *EnableSplitTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable split test params
func (o *EnableSplitTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable split test params
func (o *EnableSplitTestParams) WithHTTPClient(client *http.Client) *EnableSplitTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable split test params
func (o *EnableSplitTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the enable split test params
func (o *EnableSplitTestParams) WithSiteID(siteID string) *EnableSplitTestParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the enable split test params
func (o *EnableSplitTestParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithSplitTestID adds the splitTestID to the enable split test params
func (o *EnableSplitTestParams) WithSplitTestID(splitTestID string) *EnableSplitTestParams {
	o.SetSplitTestID(splitTestID)
	return o
}

// SetSplitTestID adds the splitTestId to the enable split test params
func (o *EnableSplitTestParams) SetSplitTestID(splitTestID string) {
	o.SplitTestID = splitTestID
}

// WriteToRequest writes these params to a swagger request
func (o *EnableSplitTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	// path param split_test_id
	if err := r.SetPathParam("split_test_id", o.SplitTestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
