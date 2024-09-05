// Code generated by go-swagger; DO NOT EDIT.

package metrics

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

// NewPostNetworksNetworkIDMetricsPushParams creates a new PostNetworksNetworkIDMetricsPushParams object
// with the default values initialized.
func NewPostNetworksNetworkIDMetricsPushParams() *PostNetworksNetworkIDMetricsPushParams {
	var ()
	return &PostNetworksNetworkIDMetricsPushParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworksNetworkIDMetricsPushParamsWithTimeout creates a new PostNetworksNetworkIDMetricsPushParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNetworksNetworkIDMetricsPushParamsWithTimeout(timeout time.Duration) *PostNetworksNetworkIDMetricsPushParams {
	var ()
	return &PostNetworksNetworkIDMetricsPushParams{

		timeout: timeout,
	}
}

// NewPostNetworksNetworkIDMetricsPushParamsWithContext creates a new PostNetworksNetworkIDMetricsPushParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNetworksNetworkIDMetricsPushParamsWithContext(ctx context.Context) *PostNetworksNetworkIDMetricsPushParams {
	var ()
	return &PostNetworksNetworkIDMetricsPushParams{

		Context: ctx,
	}
}

// NewPostNetworksNetworkIDMetricsPushParamsWithHTTPClient creates a new PostNetworksNetworkIDMetricsPushParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNetworksNetworkIDMetricsPushParamsWithHTTPClient(client *http.Client) *PostNetworksNetworkIDMetricsPushParams {
	var ()
	return &PostNetworksNetworkIDMetricsPushParams{
		HTTPClient: client,
	}
}

/*PostNetworksNetworkIDMetricsPushParams contains all the parameters to send to the API endpoint
for the post networks network ID metrics push operation typically these are written to a http.Request
*/
type PostNetworksNetworkIDMetricsPushParams struct {

	/*Metrics
	  Metrics to be submitted

	*/
	Metrics []*models.PushedMetric
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post networks network ID metrics push params
func (o *PostNetworksNetworkIDMetricsPushParams) WithTimeout(timeout time.Duration) *PostNetworksNetworkIDMetricsPushParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post networks network ID metrics push params
func (o *PostNetworksNetworkIDMetricsPushParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post networks network ID metrics push params
func (o *PostNetworksNetworkIDMetricsPushParams) WithContext(ctx context.Context) *PostNetworksNetworkIDMetricsPushParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post networks network ID metrics push params
func (o *PostNetworksNetworkIDMetricsPushParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post networks network ID metrics push params
func (o *PostNetworksNetworkIDMetricsPushParams) WithHTTPClient(client *http.Client) *PostNetworksNetworkIDMetricsPushParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post networks network ID metrics push params
func (o *PostNetworksNetworkIDMetricsPushParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetrics adds the metrics to the post networks network ID metrics push params
func (o *PostNetworksNetworkIDMetricsPushParams) WithMetrics(metrics []*models.PushedMetric) *PostNetworksNetworkIDMetricsPushParams {
	o.SetMetrics(metrics)
	return o
}

// SetMetrics adds the metrics to the post networks network ID metrics push params
func (o *PostNetworksNetworkIDMetricsPushParams) SetMetrics(metrics []*models.PushedMetric) {
	o.Metrics = metrics
}

// WithNetworkID adds the networkID to the post networks network ID metrics push params
func (o *PostNetworksNetworkIDMetricsPushParams) WithNetworkID(networkID string) *PostNetworksNetworkIDMetricsPushParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post networks network ID metrics push params
func (o *PostNetworksNetworkIDMetricsPushParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworksNetworkIDMetricsPushParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Metrics != nil {
		if err := r.SetBodyParam(o.Metrics); err != nil {
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
