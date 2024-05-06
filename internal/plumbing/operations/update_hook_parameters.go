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

// NewUpdateHookParams creates a new UpdateHookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateHookParams() *UpdateHookParams {
	return &UpdateHookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHookParamsWithTimeout creates a new UpdateHookParams object
// with the ability to set a timeout on a request.
func NewUpdateHookParamsWithTimeout(timeout time.Duration) *UpdateHookParams {
	return &UpdateHookParams{
		timeout: timeout,
	}
}

// NewUpdateHookParamsWithContext creates a new UpdateHookParams object
// with the ability to set a context for a request.
func NewUpdateHookParamsWithContext(ctx context.Context) *UpdateHookParams {
	return &UpdateHookParams{
		Context: ctx,
	}
}

// NewUpdateHookParamsWithHTTPClient creates a new UpdateHookParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateHookParamsWithHTTPClient(client *http.Client) *UpdateHookParams {
	return &UpdateHookParams{
		HTTPClient: client,
	}
}

/*
UpdateHookParams contains all the parameters to send to the API endpoint

	for the update hook operation.

	Typically these are written to a http.Request.
*/
type UpdateHookParams struct {

	// Hook.
	Hook *models.Hook

	// HookID.
	HookID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHookParams) WithDefaults() *UpdateHookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update hook params
func (o *UpdateHookParams) WithTimeout(timeout time.Duration) *UpdateHookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update hook params
func (o *UpdateHookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update hook params
func (o *UpdateHookParams) WithContext(ctx context.Context) *UpdateHookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update hook params
func (o *UpdateHookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update hook params
func (o *UpdateHookParams) WithHTTPClient(client *http.Client) *UpdateHookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update hook params
func (o *UpdateHookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHook adds the hook to the update hook params
func (o *UpdateHookParams) WithHook(hook *models.Hook) *UpdateHookParams {
	o.SetHook(hook)
	return o
}

// SetHook adds the hook to the update hook params
func (o *UpdateHookParams) SetHook(hook *models.Hook) {
	o.Hook = hook
}

// WithHookID adds the hookID to the update hook params
func (o *UpdateHookParams) WithHookID(hookID string) *UpdateHookParams {
	o.SetHookID(hookID)
	return o
}

// SetHookID adds the hookId to the update hook params
func (o *UpdateHookParams) SetHookID(hookID string) {
	o.HookID = hookID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Hook != nil {
		if err := r.SetBodyParam(o.Hook); err != nil {
			return err
		}
	}

	// path param hook_id
	if err := r.SetPathParam("hook_id", o.HookID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
