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

	"github.com/netlify/terraform-provider-netlify/internal/models"
)

// NewUpdateSiteBuildLogParams creates a new UpdateSiteBuildLogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSiteBuildLogParams() *UpdateSiteBuildLogParams {
	return &UpdateSiteBuildLogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSiteBuildLogParamsWithTimeout creates a new UpdateSiteBuildLogParams object
// with the ability to set a timeout on a request.
func NewUpdateSiteBuildLogParamsWithTimeout(timeout time.Duration) *UpdateSiteBuildLogParams {
	return &UpdateSiteBuildLogParams{
		timeout: timeout,
	}
}

// NewUpdateSiteBuildLogParamsWithContext creates a new UpdateSiteBuildLogParams object
// with the ability to set a context for a request.
func NewUpdateSiteBuildLogParamsWithContext(ctx context.Context) *UpdateSiteBuildLogParams {
	return &UpdateSiteBuildLogParams{
		Context: ctx,
	}
}

// NewUpdateSiteBuildLogParamsWithHTTPClient creates a new UpdateSiteBuildLogParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSiteBuildLogParamsWithHTTPClient(client *http.Client) *UpdateSiteBuildLogParams {
	return &UpdateSiteBuildLogParams{
		HTTPClient: client,
	}
}

/*
UpdateSiteBuildLogParams contains all the parameters to send to the API endpoint

	for the update site build log operation.

	Typically these are written to a http.Request.
*/
type UpdateSiteBuildLogParams struct {

	// BuildID.
	BuildID string

	// Msg.
	Msg *models.BuildLogMsg

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update site build log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSiteBuildLogParams) WithDefaults() *UpdateSiteBuildLogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update site build log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSiteBuildLogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update site build log params
func (o *UpdateSiteBuildLogParams) WithTimeout(timeout time.Duration) *UpdateSiteBuildLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update site build log params
func (o *UpdateSiteBuildLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update site build log params
func (o *UpdateSiteBuildLogParams) WithContext(ctx context.Context) *UpdateSiteBuildLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update site build log params
func (o *UpdateSiteBuildLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update site build log params
func (o *UpdateSiteBuildLogParams) WithHTTPClient(client *http.Client) *UpdateSiteBuildLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update site build log params
func (o *UpdateSiteBuildLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildID adds the buildID to the update site build log params
func (o *UpdateSiteBuildLogParams) WithBuildID(buildID string) *UpdateSiteBuildLogParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the update site build log params
func (o *UpdateSiteBuildLogParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithMsg adds the msg to the update site build log params
func (o *UpdateSiteBuildLogParams) WithMsg(msg *models.BuildLogMsg) *UpdateSiteBuildLogParams {
	o.SetMsg(msg)
	return o
}

// SetMsg adds the msg to the update site build log params
func (o *UpdateSiteBuildLogParams) SetMsg(msg *models.BuildLogMsg) {
	o.Msg = msg
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSiteBuildLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param build_id
	if err := r.SetPathParam("build_id", o.BuildID); err != nil {
		return err
	}
	if o.Msg != nil {
		if err := r.SetBodyParam(o.Msg); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
