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

// EnableHookReader is a Reader for the EnableHook structure.
type EnableHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEnableHookDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnableHookOK creates a EnableHookOK with default headers values
func NewEnableHookOK() *EnableHookOK {
	return &EnableHookOK{}
}

/*
EnableHookOK describes a response with status code 200, with default header values.

OK
*/
type EnableHookOK struct {
	Payload *models.Hook
}

// IsSuccess returns true when this enable hook o k response has a 2xx status code
func (o *EnableHookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enable hook o k response has a 3xx status code
func (o *EnableHookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable hook o k response has a 4xx status code
func (o *EnableHookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable hook o k response has a 5xx status code
func (o *EnableHookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enable hook o k response a status code equal to that given
func (o *EnableHookOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the enable hook o k response
func (o *EnableHookOK) Code() int {
	return 200
}

func (o *EnableHookOK) Error() string {
	return fmt.Sprintf("[POST /hooks/{hook_id}/enable][%d] enableHookOK  %+v", 200, o.Payload)
}

func (o *EnableHookOK) String() string {
	return fmt.Sprintf("[POST /hooks/{hook_id}/enable][%d] enableHookOK  %+v", 200, o.Payload)
}

func (o *EnableHookOK) GetPayload() *models.Hook {
	return o.Payload
}

func (o *EnableHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableHookDefault creates a EnableHookDefault with default headers values
func NewEnableHookDefault(code int) *EnableHookDefault {
	return &EnableHookDefault{
		_statusCode: code,
	}
}

/*
EnableHookDefault describes a response with status code -1, with default header values.

error
*/
type EnableHookDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this enable hook default response has a 2xx status code
func (o *EnableHookDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enable hook default response has a 3xx status code
func (o *EnableHookDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enable hook default response has a 4xx status code
func (o *EnableHookDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enable hook default response has a 5xx status code
func (o *EnableHookDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enable hook default response a status code equal to that given
func (o *EnableHookDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the enable hook default response
func (o *EnableHookDefault) Code() int {
	return o._statusCode
}

func (o *EnableHookDefault) Error() string {
	return fmt.Sprintf("[POST /hooks/{hook_id}/enable][%d] enableHook default  %+v", o._statusCode, o.Payload)
}

func (o *EnableHookDefault) String() string {
	return fmt.Sprintf("[POST /hooks/{hook_id}/enable][%d] enableHook default  %+v", o._statusCode, o.Payload)
}

func (o *EnableHookDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnableHookDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
