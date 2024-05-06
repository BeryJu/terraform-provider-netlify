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

// NewUpdateServiceInstanceParams creates a new UpdateServiceInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateServiceInstanceParams() *UpdateServiceInstanceParams {
	return &UpdateServiceInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceInstanceParamsWithTimeout creates a new UpdateServiceInstanceParams object
// with the ability to set a timeout on a request.
func NewUpdateServiceInstanceParamsWithTimeout(timeout time.Duration) *UpdateServiceInstanceParams {
	return &UpdateServiceInstanceParams{
		timeout: timeout,
	}
}

// NewUpdateServiceInstanceParamsWithContext creates a new UpdateServiceInstanceParams object
// with the ability to set a context for a request.
func NewUpdateServiceInstanceParamsWithContext(ctx context.Context) *UpdateServiceInstanceParams {
	return &UpdateServiceInstanceParams{
		Context: ctx,
	}
}

// NewUpdateServiceInstanceParamsWithHTTPClient creates a new UpdateServiceInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateServiceInstanceParamsWithHTTPClient(client *http.Client) *UpdateServiceInstanceParams {
	return &UpdateServiceInstanceParams{
		HTTPClient: client,
	}
}

/*
UpdateServiceInstanceParams contains all the parameters to send to the API endpoint

	for the update service instance operation.

	Typically these are written to a http.Request.
*/
type UpdateServiceInstanceParams struct {

	// Addon.
	Addon string

	// Config.
	Config interface{}

	// InstanceID.
	InstanceID string

	// SiteID.
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update service instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServiceInstanceParams) WithDefaults() *UpdateServiceInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update service instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServiceInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update service instance params
func (o *UpdateServiceInstanceParams) WithTimeout(timeout time.Duration) *UpdateServiceInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service instance params
func (o *UpdateServiceInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service instance params
func (o *UpdateServiceInstanceParams) WithContext(ctx context.Context) *UpdateServiceInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service instance params
func (o *UpdateServiceInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service instance params
func (o *UpdateServiceInstanceParams) WithHTTPClient(client *http.Client) *UpdateServiceInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service instance params
func (o *UpdateServiceInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddon adds the addon to the update service instance params
func (o *UpdateServiceInstanceParams) WithAddon(addon string) *UpdateServiceInstanceParams {
	o.SetAddon(addon)
	return o
}

// SetAddon adds the addon to the update service instance params
func (o *UpdateServiceInstanceParams) SetAddon(addon string) {
	o.Addon = addon
}

// WithConfig adds the config to the update service instance params
func (o *UpdateServiceInstanceParams) WithConfig(config interface{}) *UpdateServiceInstanceParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the update service instance params
func (o *UpdateServiceInstanceParams) SetConfig(config interface{}) {
	o.Config = config
}

// WithInstanceID adds the instanceID to the update service instance params
func (o *UpdateServiceInstanceParams) WithInstanceID(instanceID string) *UpdateServiceInstanceParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the update service instance params
func (o *UpdateServiceInstanceParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithSiteID adds the siteID to the update service instance params
func (o *UpdateServiceInstanceParams) WithSiteID(siteID string) *UpdateServiceInstanceParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the update service instance params
func (o *UpdateServiceInstanceParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addon
	if err := r.SetPathParam("addon", o.Addon); err != nil {
		return err
	}
	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
	}

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
		return err
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
