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

// NewGetSiteBuildParams creates a new GetSiteBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSiteBuildParams() *GetSiteBuildParams {
	return &GetSiteBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSiteBuildParamsWithTimeout creates a new GetSiteBuildParams object
// with the ability to set a timeout on a request.
func NewGetSiteBuildParamsWithTimeout(timeout time.Duration) *GetSiteBuildParams {
	return &GetSiteBuildParams{
		timeout: timeout,
	}
}

// NewGetSiteBuildParamsWithContext creates a new GetSiteBuildParams object
// with the ability to set a context for a request.
func NewGetSiteBuildParamsWithContext(ctx context.Context) *GetSiteBuildParams {
	return &GetSiteBuildParams{
		Context: ctx,
	}
}

// NewGetSiteBuildParamsWithHTTPClient creates a new GetSiteBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSiteBuildParamsWithHTTPClient(client *http.Client) *GetSiteBuildParams {
	return &GetSiteBuildParams{
		HTTPClient: client,
	}
}

/*
GetSiteBuildParams contains all the parameters to send to the API endpoint

	for the get site build operation.

	Typically these are written to a http.Request.
*/
type GetSiteBuildParams struct {

	// BuildID.
	BuildID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get site build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSiteBuildParams) WithDefaults() *GetSiteBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get site build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSiteBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get site build params
func (o *GetSiteBuildParams) WithTimeout(timeout time.Duration) *GetSiteBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get site build params
func (o *GetSiteBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get site build params
func (o *GetSiteBuildParams) WithContext(ctx context.Context) *GetSiteBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get site build params
func (o *GetSiteBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get site build params
func (o *GetSiteBuildParams) WithHTTPClient(client *http.Client) *GetSiteBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get site build params
func (o *GetSiteBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildID adds the buildID to the get site build params
func (o *GetSiteBuildParams) WithBuildID(buildID string) *GetSiteBuildParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the get site build params
func (o *GetSiteBuildParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSiteBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param build_id
	if err := r.SetPathParam("build_id", o.BuildID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
