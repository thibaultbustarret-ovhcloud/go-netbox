// Code generated by go-swagger; DO NOT EDIT.

package extras

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

// NewExtrasImageAttachmentsBulkUpdateParams creates a new ExtrasImageAttachmentsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasImageAttachmentsBulkUpdateParams() *ExtrasImageAttachmentsBulkUpdateParams {
	return &ExtrasImageAttachmentsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsBulkUpdateParamsWithTimeout creates a new ExtrasImageAttachmentsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasImageAttachmentsBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsBulkUpdateParams {
	return &ExtrasImageAttachmentsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsBulkUpdateParamsWithContext creates a new ExtrasImageAttachmentsBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasImageAttachmentsBulkUpdateParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsBulkUpdateParams {
	return &ExtrasImageAttachmentsBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasImageAttachmentsBulkUpdateParamsWithHTTPClient creates a new ExtrasImageAttachmentsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasImageAttachmentsBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsBulkUpdateParams {
	return &ExtrasImageAttachmentsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasImageAttachmentsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras image attachments bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasImageAttachmentsBulkUpdateParams struct {

	// Data.
	Data *models.ImageAttachment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras image attachments bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsBulkUpdateParams) WithDefaults() *ExtrasImageAttachmentsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras image attachments bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras image attachments bulk update params
func (o *ExtrasImageAttachmentsBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments bulk update params
func (o *ExtrasImageAttachmentsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments bulk update params
func (o *ExtrasImageAttachmentsBulkUpdateParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments bulk update params
func (o *ExtrasImageAttachmentsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments bulk update params
func (o *ExtrasImageAttachmentsBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments bulk update params
func (o *ExtrasImageAttachmentsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras image attachments bulk update params
func (o *ExtrasImageAttachmentsBulkUpdateParams) WithData(data *models.ImageAttachment) *ExtrasImageAttachmentsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras image attachments bulk update params
func (o *ExtrasImageAttachmentsBulkUpdateParams) SetData(data *models.ImageAttachment) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
