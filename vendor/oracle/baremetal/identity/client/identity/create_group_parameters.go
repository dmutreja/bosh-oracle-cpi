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

	"oracle/baremetal/identity/models"
)

// NewCreateGroupParams creates a new CreateGroupParams object
// with the default values initialized.
func NewCreateGroupParams() *CreateGroupParams {
	var ()
	return &CreateGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGroupParamsWithTimeout creates a new CreateGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateGroupParamsWithTimeout(timeout time.Duration) *CreateGroupParams {
	var ()
	return &CreateGroupParams{

		timeout: timeout,
	}
}

// NewCreateGroupParamsWithContext creates a new CreateGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateGroupParamsWithContext(ctx context.Context) *CreateGroupParams {
	var ()
	return &CreateGroupParams{

		Context: ctx,
	}
}

// NewCreateGroupParamsWithHTTPClient creates a new CreateGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateGroupParamsWithHTTPClient(client *http.Client) *CreateGroupParams {
	var ()
	return &CreateGroupParams{
		HTTPClient: client,
	}
}

/*CreateGroupParams contains all the parameters to send to the API endpoint
for the create group operation typically these are written to a http.Request
*/
type CreateGroupParams struct {

	/*CreateGroupDetails
	  Request object for creating a new group.

	*/
	CreateGroupDetails *models.CreateGroupDetails
	/*OpcRetryToken
	  A token that uniquely identifies a request so it can be retried in case of a timeout or
	server error without risk of executing that same action again. Retry tokens expire after 24
	hours, but can be invalidated before then due to conflicting operations (e.g., if a resource
	has been deleted and purged from the system, then a retry of the original creation request
	may be rejected).


	*/
	OpcRetryToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create group params
func (o *CreateGroupParams) WithTimeout(timeout time.Duration) *CreateGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create group params
func (o *CreateGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create group params
func (o *CreateGroupParams) WithContext(ctx context.Context) *CreateGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create group params
func (o *CreateGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create group params
func (o *CreateGroupParams) WithHTTPClient(client *http.Client) *CreateGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create group params
func (o *CreateGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateGroupDetails adds the createGroupDetails to the create group params
func (o *CreateGroupParams) WithCreateGroupDetails(createGroupDetails *models.CreateGroupDetails) *CreateGroupParams {
	o.SetCreateGroupDetails(createGroupDetails)
	return o
}

// SetCreateGroupDetails adds the createGroupDetails to the create group params
func (o *CreateGroupParams) SetCreateGroupDetails(createGroupDetails *models.CreateGroupDetails) {
	o.CreateGroupDetails = createGroupDetails
}

// WithOpcRetryToken adds the opcRetryToken to the create group params
func (o *CreateGroupParams) WithOpcRetryToken(opcRetryToken *string) *CreateGroupParams {
	o.SetOpcRetryToken(opcRetryToken)
	return o
}

// SetOpcRetryToken adds the opcRetryToken to the create group params
func (o *CreateGroupParams) SetOpcRetryToken(opcRetryToken *string) {
	o.OpcRetryToken = opcRetryToken
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateGroupDetails == nil {
		o.CreateGroupDetails = new(models.CreateGroupDetails)
	}

	if err := r.SetBodyParam(o.CreateGroupDetails); err != nil {
		return err
	}

	if o.OpcRetryToken != nil {

		// header param opc-retry-token
		if err := r.SetHeaderParam("opc-retry-token", *o.OpcRetryToken); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}