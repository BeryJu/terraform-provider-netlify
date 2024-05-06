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

// GetSiteDevServerReader is a Reader for the GetSiteDevServer structure.
type GetSiteDevServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSiteDevServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSiteDevServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /sites/{site_id}/dev_servers/{dev_server_id}] getSiteDevServer", response, response.Code())
	}
}

// NewGetSiteDevServerOK creates a GetSiteDevServerOK with default headers values
func NewGetSiteDevServerOK() *GetSiteDevServerOK {
	return &GetSiteDevServerOK{}
}

/*
GetSiteDevServerOK describes a response with status code 200, with default header values.

OK
*/
type GetSiteDevServerOK struct {
	Payload *models.DevServer
}

// IsSuccess returns true when this get site dev server o k response has a 2xx status code
func (o *GetSiteDevServerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get site dev server o k response has a 3xx status code
func (o *GetSiteDevServerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get site dev server o k response has a 4xx status code
func (o *GetSiteDevServerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get site dev server o k response has a 5xx status code
func (o *GetSiteDevServerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get site dev server o k response a status code equal to that given
func (o *GetSiteDevServerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get site dev server o k response
func (o *GetSiteDevServerOK) Code() int {
	return 200
}

func (o *GetSiteDevServerOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/dev_servers/{dev_server_id}][%d] getSiteDevServerOK  %+v", 200, o.Payload)
}

func (o *GetSiteDevServerOK) String() string {
	return fmt.Sprintf("[GET /sites/{site_id}/dev_servers/{dev_server_id}][%d] getSiteDevServerOK  %+v", 200, o.Payload)
}

func (o *GetSiteDevServerOK) GetPayload() *models.DevServer {
	return o.Payload
}

func (o *GetSiteDevServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
