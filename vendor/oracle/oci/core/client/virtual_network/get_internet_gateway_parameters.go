package virtual_network

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

// NewGetInternetGatewayParams creates a new GetInternetGatewayParams object
// with the default values initialized.
func NewGetInternetGatewayParams() *GetInternetGatewayParams {
	var ()
	return &GetInternetGatewayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInternetGatewayParamsWithTimeout creates a new GetInternetGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInternetGatewayParamsWithTimeout(timeout time.Duration) *GetInternetGatewayParams {
	var ()
	return &GetInternetGatewayParams{

		timeout: timeout,
	}
}

// NewGetInternetGatewayParamsWithContext creates a new GetInternetGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInternetGatewayParamsWithContext(ctx context.Context) *GetInternetGatewayParams {
	var ()
	return &GetInternetGatewayParams{

		Context: ctx,
	}
}

// NewGetInternetGatewayParamsWithHTTPClient creates a new GetInternetGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInternetGatewayParamsWithHTTPClient(client *http.Client) *GetInternetGatewayParams {
	var ()
	return &GetInternetGatewayParams{
		HTTPClient: client,
	}
}

/*GetInternetGatewayParams contains all the parameters to send to the API endpoint
for the get internet gateway operation typically these are written to a http.Request
*/
type GetInternetGatewayParams struct {

	/*IgID
	  The OCID of the Internet Gateway.

	*/
	IgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get internet gateway params
func (o *GetInternetGatewayParams) WithTimeout(timeout time.Duration) *GetInternetGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get internet gateway params
func (o *GetInternetGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get internet gateway params
func (o *GetInternetGatewayParams) WithContext(ctx context.Context) *GetInternetGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get internet gateway params
func (o *GetInternetGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get internet gateway params
func (o *GetInternetGatewayParams) WithHTTPClient(client *http.Client) *GetInternetGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get internet gateway params
func (o *GetInternetGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIgID adds the igID to the get internet gateway params
func (o *GetInternetGatewayParams) WithIgID(igID string) *GetInternetGatewayParams {
	o.SetIgID(igID)
	return o
}

// SetIgID adds the igId to the get internet gateway params
func (o *GetInternetGatewayParams) SetIgID(igID string) {
	o.IgID = igID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInternetGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param igId
	if err := r.SetPathParam("igId", o.IgID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
