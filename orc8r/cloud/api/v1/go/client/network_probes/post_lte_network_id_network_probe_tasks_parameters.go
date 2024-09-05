// Code generated by go-swagger; DO NOT EDIT.

package network_probes

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

// NewPostLTENetworkIDNetworkProbeTasksParams creates a new PostLTENetworkIDNetworkProbeTasksParams object
// with the default values initialized.
func NewPostLTENetworkIDNetworkProbeTasksParams() *PostLTENetworkIDNetworkProbeTasksParams {
	var ()
	return &PostLTENetworkIDNetworkProbeTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLTENetworkIDNetworkProbeTasksParamsWithTimeout creates a new PostLTENetworkIDNetworkProbeTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLTENetworkIDNetworkProbeTasksParamsWithTimeout(timeout time.Duration) *PostLTENetworkIDNetworkProbeTasksParams {
	var ()
	return &PostLTENetworkIDNetworkProbeTasksParams{

		timeout: timeout,
	}
}

// NewPostLTENetworkIDNetworkProbeTasksParamsWithContext creates a new PostLTENetworkIDNetworkProbeTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLTENetworkIDNetworkProbeTasksParamsWithContext(ctx context.Context) *PostLTENetworkIDNetworkProbeTasksParams {
	var ()
	return &PostLTENetworkIDNetworkProbeTasksParams{

		Context: ctx,
	}
}

// NewPostLTENetworkIDNetworkProbeTasksParamsWithHTTPClient creates a new PostLTENetworkIDNetworkProbeTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLTENetworkIDNetworkProbeTasksParamsWithHTTPClient(client *http.Client) *PostLTENetworkIDNetworkProbeTasksParams {
	var ()
	return &PostLTENetworkIDNetworkProbeTasksParams{
		HTTPClient: client,
	}
}

/*PostLTENetworkIDNetworkProbeTasksParams contains all the parameters to send to the API endpoint
for the post LTE network ID network probe tasks operation typically these are written to a http.Request
*/
type PostLTENetworkIDNetworkProbeTasksParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*NetworkProbeTask*/
	NetworkProbeTask *models.NetworkProbeTask

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post LTE network ID network probe tasks params
func (o *PostLTENetworkIDNetworkProbeTasksParams) WithTimeout(timeout time.Duration) *PostLTENetworkIDNetworkProbeTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post LTE network ID network probe tasks params
func (o *PostLTENetworkIDNetworkProbeTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post LTE network ID network probe tasks params
func (o *PostLTENetworkIDNetworkProbeTasksParams) WithContext(ctx context.Context) *PostLTENetworkIDNetworkProbeTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post LTE network ID network probe tasks params
func (o *PostLTENetworkIDNetworkProbeTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post LTE network ID network probe tasks params
func (o *PostLTENetworkIDNetworkProbeTasksParams) WithHTTPClient(client *http.Client) *PostLTENetworkIDNetworkProbeTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post LTE network ID network probe tasks params
func (o *PostLTENetworkIDNetworkProbeTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the post LTE network ID network probe tasks params
func (o *PostLTENetworkIDNetworkProbeTasksParams) WithNetworkID(networkID string) *PostLTENetworkIDNetworkProbeTasksParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post LTE network ID network probe tasks params
func (o *PostLTENetworkIDNetworkProbeTasksParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNetworkProbeTask adds the networkProbeTask to the post LTE network ID network probe tasks params
func (o *PostLTENetworkIDNetworkProbeTasksParams) WithNetworkProbeTask(networkProbeTask *models.NetworkProbeTask) *PostLTENetworkIDNetworkProbeTasksParams {
	o.SetNetworkProbeTask(networkProbeTask)
	return o
}

// SetNetworkProbeTask adds the networkProbeTask to the post LTE network ID network probe tasks params
func (o *PostLTENetworkIDNetworkProbeTasksParams) SetNetworkProbeTask(networkProbeTask *models.NetworkProbeTask) {
	o.NetworkProbeTask = networkProbeTask
}

// WriteToRequest writes these params to a swagger request
func (o *PostLTENetworkIDNetworkProbeTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.NetworkProbeTask != nil {
		if err := r.SetBodyParam(o.NetworkProbeTask); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
