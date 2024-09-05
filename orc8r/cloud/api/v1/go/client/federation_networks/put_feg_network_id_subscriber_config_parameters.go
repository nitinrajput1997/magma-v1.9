// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

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

// NewPutFegNetworkIDSubscriberConfigParams creates a new PutFegNetworkIDSubscriberConfigParams object
// with the default values initialized.
func NewPutFegNetworkIDSubscriberConfigParams() *PutFegNetworkIDSubscriberConfigParams {
	var ()
	return &PutFegNetworkIDSubscriberConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutFegNetworkIDSubscriberConfigParamsWithTimeout creates a new PutFegNetworkIDSubscriberConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutFegNetworkIDSubscriberConfigParamsWithTimeout(timeout time.Duration) *PutFegNetworkIDSubscriberConfigParams {
	var ()
	return &PutFegNetworkIDSubscriberConfigParams{

		timeout: timeout,
	}
}

// NewPutFegNetworkIDSubscriberConfigParamsWithContext creates a new PutFegNetworkIDSubscriberConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutFegNetworkIDSubscriberConfigParamsWithContext(ctx context.Context) *PutFegNetworkIDSubscriberConfigParams {
	var ()
	return &PutFegNetworkIDSubscriberConfigParams{

		Context: ctx,
	}
}

// NewPutFegNetworkIDSubscriberConfigParamsWithHTTPClient creates a new PutFegNetworkIDSubscriberConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutFegNetworkIDSubscriberConfigParamsWithHTTPClient(client *http.Client) *PutFegNetworkIDSubscriberConfigParams {
	var ()
	return &PutFegNetworkIDSubscriberConfigParams{
		HTTPClient: client,
	}
}

/*PutFegNetworkIDSubscriberConfigParams contains all the parameters to send to the API endpoint
for the put feg network ID subscriber config operation typically these are written to a http.Request
*/
type PutFegNetworkIDSubscriberConfigParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Record
	  Subscriber Config for the Network

	*/
	Record *models.NetworkSubscriberConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put feg network ID subscriber config params
func (o *PutFegNetworkIDSubscriberConfigParams) WithTimeout(timeout time.Duration) *PutFegNetworkIDSubscriberConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put feg network ID subscriber config params
func (o *PutFegNetworkIDSubscriberConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put feg network ID subscriber config params
func (o *PutFegNetworkIDSubscriberConfigParams) WithContext(ctx context.Context) *PutFegNetworkIDSubscriberConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put feg network ID subscriber config params
func (o *PutFegNetworkIDSubscriberConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put feg network ID subscriber config params
func (o *PutFegNetworkIDSubscriberConfigParams) WithHTTPClient(client *http.Client) *PutFegNetworkIDSubscriberConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put feg network ID subscriber config params
func (o *PutFegNetworkIDSubscriberConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the put feg network ID subscriber config params
func (o *PutFegNetworkIDSubscriberConfigParams) WithNetworkID(networkID string) *PutFegNetworkIDSubscriberConfigParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put feg network ID subscriber config params
func (o *PutFegNetworkIDSubscriberConfigParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRecord adds the record to the put feg network ID subscriber config params
func (o *PutFegNetworkIDSubscriberConfigParams) WithRecord(record *models.NetworkSubscriberConfig) *PutFegNetworkIDSubscriberConfigParams {
	o.SetRecord(record)
	return o
}

// SetRecord adds the record to the put feg network ID subscriber config params
func (o *PutFegNetworkIDSubscriberConfigParams) SetRecord(record *models.NetworkSubscriberConfig) {
	o.Record = record
}

// WriteToRequest writes these params to a swagger request
func (o *PutFegNetworkIDSubscriberConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.Record != nil {
		if err := r.SetBodyParam(o.Record); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
