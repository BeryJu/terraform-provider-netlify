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

// NewCancelSiteDeployParams creates a new CancelSiteDeployParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelSiteDeployParams() *CancelSiteDeployParams {
	return &CancelSiteDeployParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelSiteDeployParamsWithTimeout creates a new CancelSiteDeployParams object
// with the ability to set a timeout on a request.
func NewCancelSiteDeployParamsWithTimeout(timeout time.Duration) *CancelSiteDeployParams {
	return &CancelSiteDeployParams{
		timeout: timeout,
	}
}

// NewCancelSiteDeployParamsWithContext creates a new CancelSiteDeployParams object
// with the ability to set a context for a request.
func NewCancelSiteDeployParamsWithContext(ctx context.Context) *CancelSiteDeployParams {
	return &CancelSiteDeployParams{
		Context: ctx,
	}
}

// NewCancelSiteDeployParamsWithHTTPClient creates a new CancelSiteDeployParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelSiteDeployParamsWithHTTPClient(client *http.Client) *CancelSiteDeployParams {
	return &CancelSiteDeployParams{
		HTTPClient: client,
	}
}

/*
CancelSiteDeployParams contains all the parameters to send to the API endpoint

	for the cancel site deploy operation.

	Typically these are written to a http.Request.
*/
type CancelSiteDeployParams struct {

	// DeployID.
	DeployID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel site deploy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelSiteDeployParams) WithDefaults() *CancelSiteDeployParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel site deploy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelSiteDeployParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel site deploy params
func (o *CancelSiteDeployParams) WithTimeout(timeout time.Duration) *CancelSiteDeployParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel site deploy params
func (o *CancelSiteDeployParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel site deploy params
func (o *CancelSiteDeployParams) WithContext(ctx context.Context) *CancelSiteDeployParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel site deploy params
func (o *CancelSiteDeployParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel site deploy params
func (o *CancelSiteDeployParams) WithHTTPClient(client *http.Client) *CancelSiteDeployParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel site deploy params
func (o *CancelSiteDeployParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeployID adds the deployID to the cancel site deploy params
func (o *CancelSiteDeployParams) WithDeployID(deployID string) *CancelSiteDeployParams {
	o.SetDeployID(deployID)
	return o
}

// SetDeployID adds the deployId to the cancel site deploy params
func (o *CancelSiteDeployParams) SetDeployID(deployID string) {
	o.DeployID = deployID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelSiteDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deploy_id
	if err := r.SetPathParam("deploy_id", o.DeployID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
