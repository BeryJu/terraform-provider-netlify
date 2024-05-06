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

// ListDeployKeysReader is a Reader for the ListDeployKeys structure.
type ListDeployKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDeployKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDeployKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDeployKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDeployKeysOK creates a ListDeployKeysOK with default headers values
func NewListDeployKeysOK() *ListDeployKeysOK {
	return &ListDeployKeysOK{}
}

/*
ListDeployKeysOK describes a response with status code 200, with default header values.

OK
*/
type ListDeployKeysOK struct {
	Payload []*models.DeployKey
}

// IsSuccess returns true when this list deploy keys o k response has a 2xx status code
func (o *ListDeployKeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list deploy keys o k response has a 3xx status code
func (o *ListDeployKeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list deploy keys o k response has a 4xx status code
func (o *ListDeployKeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list deploy keys o k response has a 5xx status code
func (o *ListDeployKeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list deploy keys o k response a status code equal to that given
func (o *ListDeployKeysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list deploy keys o k response
func (o *ListDeployKeysOK) Code() int {
	return 200
}

func (o *ListDeployKeysOK) Error() string {
	return fmt.Sprintf("[GET /deploy_keys][%d] listDeployKeysOK  %+v", 200, o.Payload)
}

func (o *ListDeployKeysOK) String() string {
	return fmt.Sprintf("[GET /deploy_keys][%d] listDeployKeysOK  %+v", 200, o.Payload)
}

func (o *ListDeployKeysOK) GetPayload() []*models.DeployKey {
	return o.Payload
}

func (o *ListDeployKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDeployKeysDefault creates a ListDeployKeysDefault with default headers values
func NewListDeployKeysDefault(code int) *ListDeployKeysDefault {
	return &ListDeployKeysDefault{
		_statusCode: code,
	}
}

/*
ListDeployKeysDefault describes a response with status code -1, with default header values.

error
*/
type ListDeployKeysDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list deploy keys default response has a 2xx status code
func (o *ListDeployKeysDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list deploy keys default response has a 3xx status code
func (o *ListDeployKeysDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list deploy keys default response has a 4xx status code
func (o *ListDeployKeysDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list deploy keys default response has a 5xx status code
func (o *ListDeployKeysDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list deploy keys default response a status code equal to that given
func (o *ListDeployKeysDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list deploy keys default response
func (o *ListDeployKeysDefault) Code() int {
	return o._statusCode
}

func (o *ListDeployKeysDefault) Error() string {
	return fmt.Sprintf("[GET /deploy_keys][%d] listDeployKeys default  %+v", o._statusCode, o.Payload)
}

func (o *ListDeployKeysDefault) String() string {
	return fmt.Sprintf("[GET /deploy_keys][%d] listDeployKeys default  %+v", o._statusCode, o.Payload)
}

func (o *ListDeployKeysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDeployKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
