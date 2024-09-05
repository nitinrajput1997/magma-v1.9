// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDGatewaysGatewayIDStatusReader is a Reader for the GetLTENetworkIDGatewaysGatewayIDStatus structure.
type GetLTENetworkIDGatewaysGatewayIDStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDGatewaysGatewayIDStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDStatusOK creates a GetLTENetworkIDGatewaysGatewayIDStatusOK with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDStatusOK() *GetLTENetworkIDGatewaysGatewayIDStatusOK {
	return &GetLTENetworkIDGatewaysGatewayIDStatusOK{}
}

/*GetLTENetworkIDGatewaysGatewayIDStatusOK handles this case with default header values.

The status of the gateway
*/
type GetLTENetworkIDGatewaysGatewayIDStatusOK struct {
	Payload *models.GatewayStatus
}

func (o *GetLTENetworkIDGatewaysGatewayIDStatusOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/status][%d] getLteNetworkIdGatewaysGatewayIdStatusOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDStatusOK) GetPayload() *models.GatewayStatus {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDStatusDefault creates a GetLTENetworkIDGatewaysGatewayIDStatusDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDStatusDefault(code int) *GetLTENetworkIDGatewaysGatewayIDStatusDefault {
	return &GetLTENetworkIDGatewaysGatewayIDStatusDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDStatusDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID status default response
func (o *GetLTENetworkIDGatewaysGatewayIDStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDStatusDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/status][%d] GetLTENetworkIDGatewaysGatewayIDStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
