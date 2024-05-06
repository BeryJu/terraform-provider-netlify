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

// NewCreateTicketParams creates a new CreateTicketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTicketParams() *CreateTicketParams {
	return &CreateTicketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTicketParamsWithTimeout creates a new CreateTicketParams object
// with the ability to set a timeout on a request.
func NewCreateTicketParamsWithTimeout(timeout time.Duration) *CreateTicketParams {
	return &CreateTicketParams{
		timeout: timeout,
	}
}

// NewCreateTicketParamsWithContext creates a new CreateTicketParams object
// with the ability to set a context for a request.
func NewCreateTicketParamsWithContext(ctx context.Context) *CreateTicketParams {
	return &CreateTicketParams{
		Context: ctx,
	}
}

// NewCreateTicketParamsWithHTTPClient creates a new CreateTicketParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTicketParamsWithHTTPClient(client *http.Client) *CreateTicketParams {
	return &CreateTicketParams{
		HTTPClient: client,
	}
}

/*
CreateTicketParams contains all the parameters to send to the API endpoint

	for the create ticket operation.

	Typically these are written to a http.Request.
*/
type CreateTicketParams struct {

	// ClientID.
	ClientID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create ticket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTicketParams) WithDefaults() *CreateTicketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create ticket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTicketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create ticket params
func (o *CreateTicketParams) WithTimeout(timeout time.Duration) *CreateTicketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create ticket params
func (o *CreateTicketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create ticket params
func (o *CreateTicketParams) WithContext(ctx context.Context) *CreateTicketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create ticket params
func (o *CreateTicketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create ticket params
func (o *CreateTicketParams) WithHTTPClient(client *http.Client) *CreateTicketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create ticket params
func (o *CreateTicketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the create ticket params
func (o *CreateTicketParams) WithClientID(clientID string) *CreateTicketParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the create ticket params
func (o *CreateTicketParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTicketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param client_id
	qrClientID := o.ClientID
	qClientID := qrClientID
	if qClientID != "" {

		if err := r.SetQueryParam("client_id", qClientID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
