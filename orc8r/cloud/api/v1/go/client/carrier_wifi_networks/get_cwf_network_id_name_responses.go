// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetCwfNetworkIDNameReader is a Reader for the GetCwfNetworkIDName structure.
type GetCwfNetworkIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDNameOK creates a GetCwfNetworkIDNameOK with default headers values
func NewGetCwfNetworkIDNameOK() *GetCwfNetworkIDNameOK {
	return &GetCwfNetworkIDNameOK{}
}

/*GetCwfNetworkIDNameOK handles this case with default header values.

Name of the network
*/
type GetCwfNetworkIDNameOK struct {
	Payload models.NetworkName
}

func (o *GetCwfNetworkIDNameOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/name][%d] getCwfNetworkIdNameOK  %+v", 200, o.Payload)
}

func (o *GetCwfNetworkIDNameOK) GetPayload() models.NetworkName {
	return o.Payload
}

func (o *GetCwfNetworkIDNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDNameDefault creates a GetCwfNetworkIDNameDefault with default headers values
func NewGetCwfNetworkIDNameDefault(code int) *GetCwfNetworkIDNameDefault {
	return &GetCwfNetworkIDNameDefault{
		_statusCode: code,
	}
}

/*GetCwfNetworkIDNameDefault handles this case with default header values.

Unexpected Error
*/
type GetCwfNetworkIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID name default response
func (o *GetCwfNetworkIDNameDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDNameDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/name][%d] GetCwfNetworkIDName default  %+v", o._statusCode, o.Payload)
}

func (o *GetCwfNetworkIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
