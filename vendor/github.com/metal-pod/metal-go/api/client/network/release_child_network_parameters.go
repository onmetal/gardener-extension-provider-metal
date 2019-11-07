// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewReleaseChildNetworkParams creates a new ReleaseChildNetworkParams object
// with the default values initialized.
func NewReleaseChildNetworkParams() *ReleaseChildNetworkParams {
	var ()
	return &ReleaseChildNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReleaseChildNetworkParamsWithTimeout creates a new ReleaseChildNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReleaseChildNetworkParamsWithTimeout(timeout time.Duration) *ReleaseChildNetworkParams {
	var ()
	return &ReleaseChildNetworkParams{

		timeout: timeout,
	}
}

// NewReleaseChildNetworkParamsWithContext creates a new ReleaseChildNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewReleaseChildNetworkParamsWithContext(ctx context.Context) *ReleaseChildNetworkParams {
	var ()
	return &ReleaseChildNetworkParams{

		Context: ctx,
	}
}

// NewReleaseChildNetworkParamsWithHTTPClient creates a new ReleaseChildNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReleaseChildNetworkParamsWithHTTPClient(client *http.Client) *ReleaseChildNetworkParams {
	var ()
	return &ReleaseChildNetworkParams{
		HTTPClient: client,
	}
}

/*ReleaseChildNetworkParams contains all the parameters to send to the API endpoint
for the release child network operation typically these are written to a http.Request
*/
type ReleaseChildNetworkParams struct {

	/*ID
	  identifier of the network

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the release child network params
func (o *ReleaseChildNetworkParams) WithTimeout(timeout time.Duration) *ReleaseChildNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the release child network params
func (o *ReleaseChildNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the release child network params
func (o *ReleaseChildNetworkParams) WithContext(ctx context.Context) *ReleaseChildNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the release child network params
func (o *ReleaseChildNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the release child network params
func (o *ReleaseChildNetworkParams) WithHTTPClient(client *http.Client) *ReleaseChildNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the release child network params
func (o *ReleaseChildNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the release child network params
func (o *ReleaseChildNetworkParams) WithID(id string) *ReleaseChildNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the release child network params
func (o *ReleaseChildNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReleaseChildNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}