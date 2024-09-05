// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewPostNetworksNetworkIDAlertsSilenceParams creates a new PostNetworksNetworkIDAlertsSilenceParams object
// with the default values initialized.
func NewPostNetworksNetworkIDAlertsSilenceParams() *PostNetworksNetworkIDAlertsSilenceParams {
	var ()
	return &PostNetworksNetworkIDAlertsSilenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworksNetworkIDAlertsSilenceParamsWithTimeout creates a new PostNetworksNetworkIDAlertsSilenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNetworksNetworkIDAlertsSilenceParamsWithTimeout(timeout time.Duration) *PostNetworksNetworkIDAlertsSilenceParams {
	var ()
	return &PostNetworksNetworkIDAlertsSilenceParams{

		timeout: timeout,
	}
}

// NewPostNetworksNetworkIDAlertsSilenceParamsWithContext creates a new PostNetworksNetworkIDAlertsSilenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNetworksNetworkIDAlertsSilenceParamsWithContext(ctx context.Context) *PostNetworksNetworkIDAlertsSilenceParams {
	var ()
	return &PostNetworksNetworkIDAlertsSilenceParams{

		Context: ctx,
	}
}

// NewPostNetworksNetworkIDAlertsSilenceParamsWithHTTPClient creates a new PostNetworksNetworkIDAlertsSilenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNetworksNetworkIDAlertsSilenceParamsWithHTTPClient(client *http.Client) *PostNetworksNetworkIDAlertsSilenceParams {
	var ()
	return &PostNetworksNetworkIDAlertsSilenceParams{
		HTTPClient: client,
	}
}

/*PostNetworksNetworkIDAlertsSilenceParams contains all the parameters to send to the API endpoint
for the post networks network ID alerts silence operation typically these are written to a http.Request
*/
type PostNetworksNetworkIDAlertsSilenceParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Silencer
	  Silencer to be created

	*/
	Silencer *models.AlertSilencer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post networks network ID alerts silence params
func (o *PostNetworksNetworkIDAlertsSilenceParams) WithTimeout(timeout time.Duration) *PostNetworksNetworkIDAlertsSilenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post networks network ID alerts silence params
func (o *PostNetworksNetworkIDAlertsSilenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post networks network ID alerts silence params
func (o *PostNetworksNetworkIDAlertsSilenceParams) WithContext(ctx context.Context) *PostNetworksNetworkIDAlertsSilenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post networks network ID alerts silence params
func (o *PostNetworksNetworkIDAlertsSilenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post networks network ID alerts silence params
func (o *PostNetworksNetworkIDAlertsSilenceParams) WithHTTPClient(client *http.Client) *PostNetworksNetworkIDAlertsSilenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post networks network ID alerts silence params
func (o *PostNetworksNetworkIDAlertsSilenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the post networks network ID alerts silence params
func (o *PostNetworksNetworkIDAlertsSilenceParams) WithNetworkID(networkID string) *PostNetworksNetworkIDAlertsSilenceParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post networks network ID alerts silence params
func (o *PostNetworksNetworkIDAlertsSilenceParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSilencer adds the silencer to the post networks network ID alerts silence params
func (o *PostNetworksNetworkIDAlertsSilenceParams) WithSilencer(silencer *models.AlertSilencer) *PostNetworksNetworkIDAlertsSilenceParams {
	o.SetSilencer(silencer)
	return o
}

// SetSilencer adds the silencer to the post networks network ID alerts silence params
func (o *PostNetworksNetworkIDAlertsSilenceParams) SetSilencer(silencer *models.AlertSilencer) {
	o.Silencer = silencer
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworksNetworkIDAlertsSilenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.Silencer != nil {
		if err := r.SetBodyParam(o.Silencer); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
