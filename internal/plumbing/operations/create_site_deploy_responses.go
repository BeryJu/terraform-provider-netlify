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

// CreateSiteDeployReader is a Reader for the CreateSiteDeploy structure.
type CreateSiteDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSiteDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSiteDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateSiteDeployDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSiteDeployOK creates a CreateSiteDeployOK with default headers values
func NewCreateSiteDeployOK() *CreateSiteDeployOK {
	return &CreateSiteDeployOK{}
}

/*
CreateSiteDeployOK describes a response with status code 200, with default header values.

OK
*/
type CreateSiteDeployOK struct {
	Payload *models.Deploy
}

// IsSuccess returns true when this create site deploy o k response has a 2xx status code
func (o *CreateSiteDeployOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create site deploy o k response has a 3xx status code
func (o *CreateSiteDeployOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create site deploy o k response has a 4xx status code
func (o *CreateSiteDeployOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create site deploy o k response has a 5xx status code
func (o *CreateSiteDeployOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create site deploy o k response a status code equal to that given
func (o *CreateSiteDeployOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create site deploy o k response
func (o *CreateSiteDeployOK) Code() int {
	return 200
}

func (o *CreateSiteDeployOK) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/deploys][%d] createSiteDeployOK  %+v", 200, o.Payload)
}

func (o *CreateSiteDeployOK) String() string {
	return fmt.Sprintf("[POST /sites/{site_id}/deploys][%d] createSiteDeployOK  %+v", 200, o.Payload)
}

func (o *CreateSiteDeployOK) GetPayload() *models.Deploy {
	return o.Payload
}

func (o *CreateSiteDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deploy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteDeployDefault creates a CreateSiteDeployDefault with default headers values
func NewCreateSiteDeployDefault(code int) *CreateSiteDeployDefault {
	return &CreateSiteDeployDefault{
		_statusCode: code,
	}
}

/*
CreateSiteDeployDefault describes a response with status code -1, with default header values.

error
*/
type CreateSiteDeployDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create site deploy default response has a 2xx status code
func (o *CreateSiteDeployDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create site deploy default response has a 3xx status code
func (o *CreateSiteDeployDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create site deploy default response has a 4xx status code
func (o *CreateSiteDeployDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create site deploy default response has a 5xx status code
func (o *CreateSiteDeployDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create site deploy default response a status code equal to that given
func (o *CreateSiteDeployDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create site deploy default response
func (o *CreateSiteDeployDefault) Code() int {
	return o._statusCode
}

func (o *CreateSiteDeployDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/deploys][%d] createSiteDeploy default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSiteDeployDefault) String() string {
	return fmt.Sprintf("[POST /sites/{site_id}/deploys][%d] createSiteDeploy default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSiteDeployDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSiteDeployDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
