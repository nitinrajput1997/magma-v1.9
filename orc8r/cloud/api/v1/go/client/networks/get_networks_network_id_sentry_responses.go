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

// GetNetworksNetworkIDSentryReader is a Reader for the GetNetworksNetworkIDSentry structure.
type GetNetworksNetworkIDSentryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDSentryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDSentryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDSentryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDSentryOK creates a GetNetworksNetworkIDSentryOK with default headers values
func NewGetNetworksNetworkIDSentryOK() *GetNetworksNetworkIDSentryOK {
	return &GetNetworksNetworkIDSentryOK{}
}

/*GetNetworksNetworkIDSentryOK handles this case with default header values.

Network-wide Sentry.io configuration
*/
type GetNetworksNetworkIDSentryOK struct {
	Payload *models.NetworkSentryConfig
}

func (o *GetNetworksNetworkIDSentryOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/sentry][%d] getNetworksNetworkIdSentryOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDSentryOK) GetPayload() *models.NetworkSentryConfig {
	return o.Payload
}

func (o *GetNetworksNetworkIDSentryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkSentryConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDSentryDefault creates a GetNetworksNetworkIDSentryDefault with default headers values
func NewGetNetworksNetworkIDSentryDefault(code int) *GetNetworksNetworkIDSentryDefault {
	return &GetNetworksNetworkIDSentryDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDSentryDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDSentryDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID sentry default response
func (o *GetNetworksNetworkIDSentryDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDSentryDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/sentry][%d] GetNetworksNetworkIDSentry default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDSentryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDSentryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
