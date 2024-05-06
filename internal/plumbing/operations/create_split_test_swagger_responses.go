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

// CreateSplitTestReader is a Reader for the CreateSplitTest structure.
type CreateSplitTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSplitTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSplitTestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateSplitTestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSplitTestCreated creates a CreateSplitTestCreated with default headers values
func NewCreateSplitTestCreated() *CreateSplitTestCreated {
	return &CreateSplitTestCreated{}
}

/*
CreateSplitTestCreated describes a response with status code 201, with default header values.

Created
*/
type CreateSplitTestCreated struct {
	Payload *models.SplitTest
}

// IsSuccess returns true when this create split test created response has a 2xx status code
func (o *CreateSplitTestCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create split test created response has a 3xx status code
func (o *CreateSplitTestCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create split test created response has a 4xx status code
func (o *CreateSplitTestCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create split test created response has a 5xx status code
func (o *CreateSplitTestCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create split test created response a status code equal to that given
func (o *CreateSplitTestCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create split test created response
func (o *CreateSplitTestCreated) Code() int {
	return 201
}

func (o *CreateSplitTestCreated) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/traffic_splits][%d] createSplitTestCreated  %+v", 201, o.Payload)
}

func (o *CreateSplitTestCreated) String() string {
	return fmt.Sprintf("[POST /sites/{site_id}/traffic_splits][%d] createSplitTestCreated  %+v", 201, o.Payload)
}

func (o *CreateSplitTestCreated) GetPayload() *models.SplitTest {
	return o.Payload
}

func (o *CreateSplitTestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SplitTest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSplitTestDefault creates a CreateSplitTestDefault with default headers values
func NewCreateSplitTestDefault(code int) *CreateSplitTestDefault {
	return &CreateSplitTestDefault{
		_statusCode: code,
	}
}

/*
CreateSplitTestDefault describes a response with status code -1, with default header values.

error
*/
type CreateSplitTestDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create split test default response has a 2xx status code
func (o *CreateSplitTestDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create split test default response has a 3xx status code
func (o *CreateSplitTestDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create split test default response has a 4xx status code
func (o *CreateSplitTestDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create split test default response has a 5xx status code
func (o *CreateSplitTestDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create split test default response a status code equal to that given
func (o *CreateSplitTestDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create split test default response
func (o *CreateSplitTestDefault) Code() int {
	return o._statusCode
}

func (o *CreateSplitTestDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/traffic_splits][%d] createSplitTest default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSplitTestDefault) String() string {
	return fmt.Sprintf("[POST /sites/{site_id}/traffic_splits][%d] createSplitTest default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSplitTestDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSplitTestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
