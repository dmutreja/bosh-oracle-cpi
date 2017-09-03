package identity

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

// NewGetIdentityProviderParams creates a new GetIdentityProviderParams object
// with the default values initialized.
func NewGetIdentityProviderParams() *GetIdentityProviderParams {
	var ()
	return &GetIdentityProviderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityProviderParamsWithTimeout creates a new GetIdentityProviderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIdentityProviderParamsWithTimeout(timeout time.Duration) *GetIdentityProviderParams {
	var ()
	return &GetIdentityProviderParams{

		timeout: timeout,
	}
}

// NewGetIdentityProviderParamsWithContext creates a new GetIdentityProviderParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIdentityProviderParamsWithContext(ctx context.Context) *GetIdentityProviderParams {
	var ()
	return &GetIdentityProviderParams{

		Context: ctx,
	}
}

// NewGetIdentityProviderParamsWithHTTPClient creates a new GetIdentityProviderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIdentityProviderParamsWithHTTPClient(client *http.Client) *GetIdentityProviderParams {
	var ()
	return &GetIdentityProviderParams{
		HTTPClient: client,
	}
}

/*GetIdentityProviderParams contains all the parameters to send to the API endpoint
for the get identity provider operation typically these are written to a http.Request
*/
type GetIdentityProviderParams struct {

	/*IdentityProviderID
	  The OCID of the identity provider.

	*/
	IdentityProviderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get identity provider params
func (o *GetIdentityProviderParams) WithTimeout(timeout time.Duration) *GetIdentityProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identity provider params
func (o *GetIdentityProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identity provider params
func (o *GetIdentityProviderParams) WithContext(ctx context.Context) *GetIdentityProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identity provider params
func (o *GetIdentityProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identity provider params
func (o *GetIdentityProviderParams) WithHTTPClient(client *http.Client) *GetIdentityProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identity provider params
func (o *GetIdentityProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentityProviderID adds the identityProviderID to the get identity provider params
func (o *GetIdentityProviderParams) WithIdentityProviderID(identityProviderID string) *GetIdentityProviderParams {
	o.SetIdentityProviderID(identityProviderID)
	return o
}

// SetIdentityProviderID adds the identityProviderId to the get identity provider params
func (o *GetIdentityProviderParams) SetIdentityProviderID(identityProviderID string) {
	o.IdentityProviderID = identityProviderID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param identityProviderId
	if err := r.SetPathParam("identityProviderId", o.IdentityProviderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}