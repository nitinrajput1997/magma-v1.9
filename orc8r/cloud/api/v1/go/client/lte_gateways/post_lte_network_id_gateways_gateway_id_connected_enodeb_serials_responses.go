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

// PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsReader is a Reader for the PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerials structure.
type PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsNoContent creates a PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsNoContent with default headers values
func NewPostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsNoContent() *PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsNoContent {
	return &PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsNoContent{}
}

/*PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsNoContent handles this case with default header values.

Success
*/
type PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsNoContent struct {
}

func (o *PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsNoContent) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/gateways/{gateway_id}/connected_enodeb_serials][%d] postLteNetworkIdGatewaysGatewayIdConnectedEnodebSerialsNoContent ", 204)
}

func (o *PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault creates a PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault with default headers values
func NewPostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault(code int) *PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault {
	return &PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault{
		_statusCode: code,
	}
}

/*PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault handles this case with default header values.

Unexpected Error
*/
type PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post LTE network ID gateways gateway ID connected ENODEB serials default response
func (o *PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault) Code() int {
	return o._statusCode
}

func (o *PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/gateways/{gateway_id}/connected_enodeb_serials][%d] PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerials default  %+v", o._statusCode, o.Payload)
}

func (o *PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
