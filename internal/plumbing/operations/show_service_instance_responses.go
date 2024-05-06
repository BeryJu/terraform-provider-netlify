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

// ShowServiceInstanceReader is a Reader for the ShowServiceInstance structure.
type ShowServiceInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowServiceInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowServiceInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewShowServiceInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowServiceInstanceOK creates a ShowServiceInstanceOK with default headers values
func NewShowServiceInstanceOK() *ShowServiceInstanceOK {
	return &ShowServiceInstanceOK{}
}

/*
ShowServiceInstanceOK describes a response with status code 200, with default header values.

OK
*/
type ShowServiceInstanceOK struct {
	Payload *models.ServiceInstance
}

// IsSuccess returns true when this show service instance o k response has a 2xx status code
func (o *ShowServiceInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this show service instance o k response has a 3xx status code
func (o *ShowServiceInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show service instance o k response has a 4xx status code
func (o *ShowServiceInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this show service instance o k response has a 5xx status code
func (o *ShowServiceInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this show service instance o k response a status code equal to that given
func (o *ShowServiceInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the show service instance o k response
func (o *ShowServiceInstanceOK) Code() int {
	return 200
}

func (o *ShowServiceInstanceOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/services/{addon}/instances/{instance_id}][%d] showServiceInstanceOK  %+v", 200, o.Payload)
}

func (o *ShowServiceInstanceOK) String() string {
	return fmt.Sprintf("[GET /sites/{site_id}/services/{addon}/instances/{instance_id}][%d] showServiceInstanceOK  %+v", 200, o.Payload)
}

func (o *ShowServiceInstanceOK) GetPayload() *models.ServiceInstance {
	return o.Payload
}

func (o *ShowServiceInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowServiceInstanceDefault creates a ShowServiceInstanceDefault with default headers values
func NewShowServiceInstanceDefault(code int) *ShowServiceInstanceDefault {
	return &ShowServiceInstanceDefault{
		_statusCode: code,
	}
}

/*
ShowServiceInstanceDefault describes a response with status code -1, with default header values.

error
*/
type ShowServiceInstanceDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this show service instance default response has a 2xx status code
func (o *ShowServiceInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this show service instance default response has a 3xx status code
func (o *ShowServiceInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this show service instance default response has a 4xx status code
func (o *ShowServiceInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this show service instance default response has a 5xx status code
func (o *ShowServiceInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this show service instance default response a status code equal to that given
func (o *ShowServiceInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the show service instance default response
func (o *ShowServiceInstanceDefault) Code() int {
	return o._statusCode
}

func (o *ShowServiceInstanceDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/services/{addon}/instances/{instance_id}][%d] showServiceInstance default  %+v", o._statusCode, o.Payload)
}

func (o *ShowServiceInstanceDefault) String() string {
	return fmt.Sprintf("[GET /sites/{site_id}/services/{addon}/instances/{instance_id}][%d] showServiceInstance default  %+v", o._statusCode, o.Payload)
}

func (o *ShowServiceInstanceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ShowServiceInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
