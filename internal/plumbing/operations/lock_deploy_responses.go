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

// LockDeployReader is a Reader for the LockDeploy structure.
type LockDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LockDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLockDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLockDeployDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLockDeployOK creates a LockDeployOK with default headers values
func NewLockDeployOK() *LockDeployOK {
	return &LockDeployOK{}
}

/*
LockDeployOK describes a response with status code 200, with default header values.

OK
*/
type LockDeployOK struct {
	Payload *models.Deploy
}

// IsSuccess returns true when this lock deploy o k response has a 2xx status code
func (o *LockDeployOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lock deploy o k response has a 3xx status code
func (o *LockDeployOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lock deploy o k response has a 4xx status code
func (o *LockDeployOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lock deploy o k response has a 5xx status code
func (o *LockDeployOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lock deploy o k response a status code equal to that given
func (o *LockDeployOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lock deploy o k response
func (o *LockDeployOK) Code() int {
	return 200
}

func (o *LockDeployOK) Error() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/lock][%d] lockDeployOK  %+v", 200, o.Payload)
}

func (o *LockDeployOK) String() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/lock][%d] lockDeployOK  %+v", 200, o.Payload)
}

func (o *LockDeployOK) GetPayload() *models.Deploy {
	return o.Payload
}

func (o *LockDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deploy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLockDeployDefault creates a LockDeployDefault with default headers values
func NewLockDeployDefault(code int) *LockDeployDefault {
	return &LockDeployDefault{
		_statusCode: code,
	}
}

/*
LockDeployDefault describes a response with status code -1, with default header values.

error
*/
type LockDeployDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this lock deploy default response has a 2xx status code
func (o *LockDeployDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lock deploy default response has a 3xx status code
func (o *LockDeployDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lock deploy default response has a 4xx status code
func (o *LockDeployDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lock deploy default response has a 5xx status code
func (o *LockDeployDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lock deploy default response a status code equal to that given
func (o *LockDeployDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lock deploy default response
func (o *LockDeployDefault) Code() int {
	return o._statusCode
}

func (o *LockDeployDefault) Error() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/lock][%d] lockDeploy default  %+v", o._statusCode, o.Payload)
}

func (o *LockDeployDefault) String() string {
	return fmt.Sprintf("[POST /deploys/{deploy_id}/lock][%d] lockDeploy default  %+v", o._statusCode, o.Payload)
}

func (o *LockDeployDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *LockDeployDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
