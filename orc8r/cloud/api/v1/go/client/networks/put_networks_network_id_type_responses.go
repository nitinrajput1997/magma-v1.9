// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutNetworksNetworkIDTypeReader is a Reader for the PutNetworksNetworkIDType structure.
type PutNetworksNetworkIDTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDTypeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDTypeNoContent creates a PutNetworksNetworkIDTypeNoContent with default headers values
func NewPutNetworksNetworkIDTypeNoContent() *PutNetworksNetworkIDTypeNoContent {
	return &PutNetworksNetworkIDTypeNoContent{}
}

/*PutNetworksNetworkIDTypeNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDTypeNoContent struct {
}

func (o *PutNetworksNetworkIDTypeNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/type][%d] putNetworksNetworkIdTypeNoContent ", 204)
}

func (o *PutNetworksNetworkIDTypeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDTypeDefault creates a PutNetworksNetworkIDTypeDefault with default headers values
func NewPutNetworksNetworkIDTypeDefault(code int) *PutNetworksNetworkIDTypeDefault {
	return &PutNetworksNetworkIDTypeDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDTypeDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDTypeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID type default response
func (o *PutNetworksNetworkIDTypeDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDTypeDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/type][%d] PutNetworksNetworkIDType default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDTypeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
