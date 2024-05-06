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

// DeleteDeployKeyReader is a Reader for the DeleteDeployKey structure.
type DeleteDeployKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeployKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDeployKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteDeployKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDeployKeyNoContent creates a DeleteDeployKeyNoContent with default headers values
func NewDeleteDeployKeyNoContent() *DeleteDeployKeyNoContent {
	return &DeleteDeployKeyNoContent{}
}

/*
DeleteDeployKeyNoContent describes a response with status code 204, with default header values.

Not Content
*/
type DeleteDeployKeyNoContent struct {
}

// IsSuccess returns true when this delete deploy key no content response has a 2xx status code
func (o *DeleteDeployKeyNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete deploy key no content response has a 3xx status code
func (o *DeleteDeployKeyNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete deploy key no content response has a 4xx status code
func (o *DeleteDeployKeyNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete deploy key no content response has a 5xx status code
func (o *DeleteDeployKeyNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete deploy key no content response a status code equal to that given
func (o *DeleteDeployKeyNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete deploy key no content response
func (o *DeleteDeployKeyNoContent) Code() int {
	return 204
}

func (o *DeleteDeployKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /deploy_keys/{key_id}][%d] deleteDeployKeyNoContent ", 204)
}

func (o *DeleteDeployKeyNoContent) String() string {
	return fmt.Sprintf("[DELETE /deploy_keys/{key_id}][%d] deleteDeployKeyNoContent ", 204)
}

func (o *DeleteDeployKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeployKeyDefault creates a DeleteDeployKeyDefault with default headers values
func NewDeleteDeployKeyDefault(code int) *DeleteDeployKeyDefault {
	return &DeleteDeployKeyDefault{
		_statusCode: code,
	}
}

/*
DeleteDeployKeyDefault describes a response with status code -1, with default header values.

error
*/
type DeleteDeployKeyDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete deploy key default response has a 2xx status code
func (o *DeleteDeployKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete deploy key default response has a 3xx status code
func (o *DeleteDeployKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete deploy key default response has a 4xx status code
func (o *DeleteDeployKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete deploy key default response has a 5xx status code
func (o *DeleteDeployKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete deploy key default response a status code equal to that given
func (o *DeleteDeployKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete deploy key default response
func (o *DeleteDeployKeyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDeployKeyDefault) Error() string {
	return fmt.Sprintf("[DELETE /deploy_keys/{key_id}][%d] deleteDeployKey default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDeployKeyDefault) String() string {
	return fmt.Sprintf("[DELETE /deploy_keys/{key_id}][%d] deleteDeployKey default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDeployKeyDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDeployKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
