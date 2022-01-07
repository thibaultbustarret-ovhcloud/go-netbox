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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewDcimPowerPortsBulkUpdateParams creates a new DcimPowerPortsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortsBulkUpdateParams() *DcimPowerPortsBulkUpdateParams {
	return &DcimPowerPortsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortsBulkUpdateParamsWithTimeout creates a new DcimPowerPortsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortsBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerPortsBulkUpdateParams {
	return &DcimPowerPortsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortsBulkUpdateParamsWithContext creates a new DcimPowerPortsBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerPortsBulkUpdateParamsWithContext(ctx context.Context) *DcimPowerPortsBulkUpdateParams {
	return &DcimPowerPortsBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerPortsBulkUpdateParamsWithHTTPClient creates a new DcimPowerPortsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortsBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerPortsBulkUpdateParams {
	return &DcimPowerPortsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the dcim power ports bulk update operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortsBulkUpdateParams struct {

	// Data.
	Data *models.WritablePowerPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power ports bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsBulkUpdateParams) WithDefaults() *DcimPowerPortsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power ports bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power ports bulk update params
func (o *DcimPowerPortsBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerPortsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power ports bulk update params
func (o *DcimPowerPortsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power ports bulk update params
func (o *DcimPowerPortsBulkUpdateParams) WithContext(ctx context.Context) *DcimPowerPortsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power ports bulk update params
func (o *DcimPowerPortsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power ports bulk update params
func (o *DcimPowerPortsBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerPortsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power ports bulk update params
func (o *DcimPowerPortsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power ports bulk update params
func (o *DcimPowerPortsBulkUpdateParams) WithData(data *models.WritablePowerPort) *DcimPowerPortsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power ports bulk update params
func (o *DcimPowerPortsBulkUpdateParams) SetData(data *models.WritablePowerPort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
