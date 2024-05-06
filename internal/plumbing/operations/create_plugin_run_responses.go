// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/terraform-provider-netlify/internal/models"
)

// CreatePluginRunReader is a Reader for the CreatePluginRun structure.
type CreatePluginRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePluginRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePluginRunCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreatePluginRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePluginRunCreated creates a CreatePluginRunCreated with default headers values
func NewCreatePluginRunCreated() *CreatePluginRunCreated {
	return &CreatePluginRunCreated{}
}

/*
CreatePluginRunCreated describes a response with status code 201, with default header values.

CREATED
*/
type CreatePluginRunCreated struct {
	Payload *models.PluginRun
}

// IsSuccess returns true when this create plugin run created response has a 2xx status code
func (o *CreatePluginRunCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create plugin run created response has a 3xx status code
func (o *CreatePluginRunCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create plugin run created response has a 4xx status code
func (o *CreatePluginRunCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create plugin run created response has a 5xx status code
func (o *CreatePluginRunCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create plugin run created response a status code equal to that given
func (o *CreatePluginRunCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create plugin run created response
func (o *CreatePluginRunCreated) Code() int {
	return 201
}

func (o *CreatePluginRunCreated) Error() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/plugin_runs][%d] createPluginRunCreated  %+v", 201, o.Payload)
}

func (o *CreatePluginRunCreated) String() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/plugin_runs][%d] createPluginRunCreated  %+v", 201, o.Payload)
}

func (o *CreatePluginRunCreated) GetPayload() *models.PluginRun {
	return o.Payload
}

func (o *CreatePluginRunCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PluginRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePluginRunDefault creates a CreatePluginRunDefault with default headers values
func NewCreatePluginRunDefault(code int) *CreatePluginRunDefault {
	return &CreatePluginRunDefault{
		_statusCode: code,
	}
}

/*
CreatePluginRunDefault describes a response with status code -1, with default header values.

error
*/
type CreatePluginRunDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create plugin run default response has a 2xx status code
func (o *CreatePluginRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create plugin run default response has a 3xx status code
func (o *CreatePluginRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create plugin run default response has a 4xx status code
func (o *CreatePluginRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create plugin run default response has a 5xx status code
func (o *CreatePluginRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create plugin run default response a status code equal to that given
func (o *CreatePluginRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create plugin run default response
func (o *CreatePluginRunDefault) Code() int {
	return o._statusCode
}

func (o *CreatePluginRunDefault) Error() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/plugin_runs][%d] createPluginRun default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePluginRunDefault) String() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/plugin_runs][%d] createPluginRun default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePluginRunDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePluginRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
