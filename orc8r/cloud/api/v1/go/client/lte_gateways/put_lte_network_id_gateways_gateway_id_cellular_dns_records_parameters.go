// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams creates a new PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams object
// with the default values initialized.
func NewPutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams() *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithTimeout creates a new PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithContext creates a new PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithContext(ctx context.Context) *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithHTTPClient creates a new PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams contains all the parameters to send to the API endpoint
for the put LTE network ID gateways gateway ID cellular DNS records operation typically these are written to a http.Request
*/
type PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Records
	  Custom DNS records for the gateway

	*/
	Records []*models.DNSConfigRecord

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithContext(ctx context.Context) *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithGatewayID(gatewayID string) *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithNetworkID(networkID string) *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRecords adds the records to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WithRecords(records []*models.DNSConfigRecord) *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams {
	o.SetRecords(records)
	return o
}

// SetRecords adds the records to the put LTE network ID gateways gateway ID cellular DNS records params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) SetRecords(records []*models.DNSConfigRecord) {
	o.Records = records
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.Records != nil {
		if err := r.SetBodyParam(o.Records); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
