// Code generated by go-swagger; DO NOT EDIT.

package admin

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

	"github.com/ory/hydra-client-go/models"
)

// NewAcceptLoginRequestParams creates a new AcceptLoginRequestParams object
// with the default values initialized.
func NewAcceptLoginRequestParams() *AcceptLoginRequestParams {
	var ()
	return &AcceptLoginRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptLoginRequestParamsWithTimeout creates a new AcceptLoginRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAcceptLoginRequestParamsWithTimeout(timeout time.Duration) *AcceptLoginRequestParams {
	var ()
	return &AcceptLoginRequestParams{

		timeout: timeout,
	}
}

// NewAcceptLoginRequestParamsWithContext creates a new AcceptLoginRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewAcceptLoginRequestParamsWithContext(ctx context.Context) *AcceptLoginRequestParams {
	var ()
	return &AcceptLoginRequestParams{

		Context: ctx,
	}
}

// NewAcceptLoginRequestParamsWithHTTPClient creates a new AcceptLoginRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAcceptLoginRequestParamsWithHTTPClient(client *http.Client) *AcceptLoginRequestParams {
	var ()
	return &AcceptLoginRequestParams{
		HTTPClient: client,
	}
}

/*AcceptLoginRequestParams contains all the parameters to send to the API endpoint
for the accept login request operation typically these are written to a http.Request
*/
type AcceptLoginRequestParams struct {

	/*Body*/
	Body *models.AcceptLoginRequest
	/*LoginChallenge*/
	LoginChallenge string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the accept login request params
func (o *AcceptLoginRequestParams) WithTimeout(timeout time.Duration) *AcceptLoginRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept login request params
func (o *AcceptLoginRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept login request params
func (o *AcceptLoginRequestParams) WithContext(ctx context.Context) *AcceptLoginRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept login request params
func (o *AcceptLoginRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept login request params
func (o *AcceptLoginRequestParams) WithHTTPClient(client *http.Client) *AcceptLoginRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept login request params
func (o *AcceptLoginRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the accept login request params
func (o *AcceptLoginRequestParams) WithBody(body *models.AcceptLoginRequest) *AcceptLoginRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the accept login request params
func (o *AcceptLoginRequestParams) SetBody(body *models.AcceptLoginRequest) {
	o.Body = body
}

// WithLoginChallenge adds the loginChallenge to the accept login request params
func (o *AcceptLoginRequestParams) WithLoginChallenge(loginChallenge string) *AcceptLoginRequestParams {
	o.SetLoginChallenge(loginChallenge)
	return o
}

// SetLoginChallenge adds the loginChallenge to the accept login request params
func (o *AcceptLoginRequestParams) SetLoginChallenge(loginChallenge string) {
	o.LoginChallenge = loginChallenge
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptLoginRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param login_challenge
	qrLoginChallenge := o.LoginChallenge
	qLoginChallenge := qrLoginChallenge
	if qLoginChallenge != "" {
		if err := r.SetQueryParam("login_challenge", qLoginChallenge); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
