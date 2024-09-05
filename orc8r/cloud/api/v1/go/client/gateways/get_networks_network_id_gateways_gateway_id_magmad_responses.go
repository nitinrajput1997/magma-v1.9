// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDGatewaysGatewayIDMagmadReader is a Reader for the GetNetworksNetworkIDGatewaysGatewayIDMagmad structure.
type GetNetworksNetworkIDGatewaysGatewayIDMagmadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDGatewaysGatewayIDMagmadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDMagmadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDMagmadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDMagmadOK creates a GetNetworksNetworkIDGatewaysGatewayIDMagmadOK with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDMagmadOK() *GetNetworksNetworkIDGatewaysGatewayIDMagmadOK {
	return &GetNetworksNetworkIDGatewaysGatewayIDMagmadOK{}
}

/*GetNetworksNetworkIDGatewaysGatewayIDMagmadOK handles this case with default header values.

Magmad agent configuration
*/
type GetNetworksNetworkIDGatewaysGatewayIDMagmadOK struct {
	Payload *models.MagmadGatewayConfigs
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDMagmadOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/magmad][%d] getNetworksNetworkIdGatewaysGatewayIdMagmadOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDMagmadOK) GetPayload() *models.MagmadGatewayConfigs {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDMagmadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MagmadGatewayConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDGatewaysGatewayIDMagmadDefault creates a GetNetworksNetworkIDGatewaysGatewayIDMagmadDefault with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDMagmadDefault(code int) *GetNetworksNetworkIDGatewaysGatewayIDMagmadDefault {
	return &GetNetworksNetworkIDGatewaysGatewayIDMagmadDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDGatewaysGatewayIDMagmadDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDGatewaysGatewayIDMagmadDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID gateways gateway ID magmad default response
func (o *GetNetworksNetworkIDGatewaysGatewayIDMagmadDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDMagmadDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/magmad][%d] GetNetworksNetworkIDGatewaysGatewayIDMagmad default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDMagmadDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDMagmadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
