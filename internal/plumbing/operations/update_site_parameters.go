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

// NewUpdateSiteParams creates a new UpdateSiteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSiteParams() *UpdateSiteParams {
	return &UpdateSiteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSiteParamsWithTimeout creates a new UpdateSiteParams object
// with the ability to set a timeout on a request.
func NewUpdateSiteParamsWithTimeout(timeout time.Duration) *UpdateSiteParams {
	return &UpdateSiteParams{
		timeout: timeout,
	}
}

// NewUpdateSiteParamsWithContext creates a new UpdateSiteParams object
// with the ability to set a context for a request.
func NewUpdateSiteParamsWithContext(ctx context.Context) *UpdateSiteParams {
	return &UpdateSiteParams{
		Context: ctx,
	}
}

// NewUpdateSiteParamsWithHTTPClient creates a new UpdateSiteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSiteParamsWithHTTPClient(client *http.Client) *UpdateSiteParams {
	return &UpdateSiteParams{
		HTTPClient: client,
	}
}

/*
UpdateSiteParams contains all the parameters to send to the API endpoint

	for the update site operation.

	Typically these are written to a http.Request.
*/
type UpdateSiteParams struct {

	// Site.
	Site *models.SiteSetup

	// SiteID.
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update site params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSiteParams) WithDefaults() *UpdateSiteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update site params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSiteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update site params
func (o *UpdateSiteParams) WithTimeout(timeout time.Duration) *UpdateSiteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update site params
func (o *UpdateSiteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update site params
func (o *UpdateSiteParams) WithContext(ctx context.Context) *UpdateSiteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update site params
func (o *UpdateSiteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update site params
func (o *UpdateSiteParams) WithHTTPClient(client *http.Client) *UpdateSiteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update site params
func (o *UpdateSiteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSite adds the site to the update site params
func (o *UpdateSiteParams) WithSite(site *models.SiteSetup) *UpdateSiteParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the update site params
func (o *UpdateSiteParams) SetSite(site *models.SiteSetup) {
	o.Site = site
}

// WithSiteID adds the siteID to the update site params
func (o *UpdateSiteParams) WithSiteID(siteID string) *UpdateSiteParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the update site params
func (o *UpdateSiteParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Site != nil {
		if err := r.SetBodyParam(o.Site); err != nil {
			return err
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
