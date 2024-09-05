// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

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

// NewPutLTENetworkIDCellularParams creates a new PutLTENetworkIDCellularParams object
// with the default values initialized.
func NewPutLTENetworkIDCellularParams() *PutLTENetworkIDCellularParams {
	var ()
	return &PutLTENetworkIDCellularParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDCellularParamsWithTimeout creates a new PutLTENetworkIDCellularParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDCellularParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDCellularParams {
	var ()
	return &PutLTENetworkIDCellularParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDCellularParamsWithContext creates a new PutLTENetworkIDCellularParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDCellularParamsWithContext(ctx context.Context) *PutLTENetworkIDCellularParams {
	var ()
	return &PutLTENetworkIDCellularParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDCellularParamsWithHTTPClient creates a new PutLTENetworkIDCellularParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDCellularParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDCellularParams {
	var ()
	return &PutLTENetworkIDCellularParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDCellularParams contains all the parameters to send to the API endpoint
for the put LTE network ID cellular operation typically these are written to a http.Request
*/
type PutLTENetworkIDCellularParams struct {

	/*Config
	  New cellular configuration for the network

	*/
	Config *models.NetworkCellularConfigs
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put LTE network ID cellular params
func (o *PutLTENetworkIDCellularParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDCellularParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID cellular params
func (o *PutLTENetworkIDCellularParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID cellular params
func (o *PutLTENetworkIDCellularParams) WithContext(ctx context.Context) *PutLTENetworkIDCellularParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID cellular params
func (o *PutLTENetworkIDCellularParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID cellular params
func (o *PutLTENetworkIDCellularParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDCellularParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID cellular params
func (o *PutLTENetworkIDCellularParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the put LTE network ID cellular params
func (o *PutLTENetworkIDCellularParams) WithConfig(config *models.NetworkCellularConfigs) *PutLTENetworkIDCellularParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the put LTE network ID cellular params
func (o *PutLTENetworkIDCellularParams) SetConfig(config *models.NetworkCellularConfigs) {
	o.Config = config
}

// WithNetworkID adds the networkID to the put LTE network ID cellular params
func (o *PutLTENetworkIDCellularParams) WithNetworkID(networkID string) *PutLTENetworkIDCellularParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID cellular params
func (o *PutLTENetworkIDCellularParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDCellularParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
