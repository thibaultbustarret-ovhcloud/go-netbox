// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDcimPowerFeedsReadParams creates a new DcimPowerFeedsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerFeedsReadParams() *DcimPowerFeedsReadParams {
	return &DcimPowerFeedsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerFeedsReadParamsWithTimeout creates a new DcimPowerFeedsReadParams object
// with the ability to set a timeout on a request.
func NewDcimPowerFeedsReadParamsWithTimeout(timeout time.Duration) *DcimPowerFeedsReadParams {
	return &DcimPowerFeedsReadParams{
		timeout: timeout,
	}
}

// NewDcimPowerFeedsReadParamsWithContext creates a new DcimPowerFeedsReadParams object
// with the ability to set a context for a request.
func NewDcimPowerFeedsReadParamsWithContext(ctx context.Context) *DcimPowerFeedsReadParams {
	return &DcimPowerFeedsReadParams{
		Context: ctx,
	}
}

// NewDcimPowerFeedsReadParamsWithHTTPClient creates a new DcimPowerFeedsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerFeedsReadParamsWithHTTPClient(client *http.Client) *DcimPowerFeedsReadParams {
	return &DcimPowerFeedsReadParams{
		HTTPClient: client,
	}
}

/* DcimPowerFeedsReadParams contains all the parameters to send to the API endpoint
   for the dcim power feeds read operation.

   Typically these are written to a http.Request.
*/
type DcimPowerFeedsReadParams struct {

	/* ID.

	   A unique integer value identifying this power feed.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power feeds read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsReadParams) WithDefaults() *DcimPowerFeedsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power feeds read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power feeds read params
func (o *DcimPowerFeedsReadParams) WithTimeout(timeout time.Duration) *DcimPowerFeedsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power feeds read params
func (o *DcimPowerFeedsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power feeds read params
func (o *DcimPowerFeedsReadParams) WithContext(ctx context.Context) *DcimPowerFeedsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power feeds read params
func (o *DcimPowerFeedsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power feeds read params
func (o *DcimPowerFeedsReadParams) WithHTTPClient(client *http.Client) *DcimPowerFeedsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power feeds read params
func (o *DcimPowerFeedsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power feeds read params
func (o *DcimPowerFeedsReadParams) WithID(id int64) *DcimPowerFeedsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power feeds read params
func (o *DcimPowerFeedsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerFeedsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
