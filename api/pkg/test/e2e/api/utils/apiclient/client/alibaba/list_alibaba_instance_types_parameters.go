// Code generated by go-swagger; DO NOT EDIT.

package alibaba

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
)

// NewListAlibabaInstanceTypesParams creates a new ListAlibabaInstanceTypesParams object
// with the default values initialized.
func NewListAlibabaInstanceTypesParams() *ListAlibabaInstanceTypesParams {
	var ()
	return &ListAlibabaInstanceTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAlibabaInstanceTypesParamsWithTimeout creates a new ListAlibabaInstanceTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAlibabaInstanceTypesParamsWithTimeout(timeout time.Duration) *ListAlibabaInstanceTypesParams {
	var ()
	return &ListAlibabaInstanceTypesParams{

		timeout: timeout,
	}
}

// NewListAlibabaInstanceTypesParamsWithContext creates a new ListAlibabaInstanceTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAlibabaInstanceTypesParamsWithContext(ctx context.Context) *ListAlibabaInstanceTypesParams {
	var ()
	return &ListAlibabaInstanceTypesParams{

		Context: ctx,
	}
}

// NewListAlibabaInstanceTypesParamsWithHTTPClient creates a new ListAlibabaInstanceTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAlibabaInstanceTypesParamsWithHTTPClient(client *http.Client) *ListAlibabaInstanceTypesParams {
	var ()
	return &ListAlibabaInstanceTypesParams{
		HTTPClient: client,
	}
}

/*ListAlibabaInstanceTypesParams contains all the parameters to send to the API endpoint
for the list alibaba instance types operation typically these are written to a http.Request
*/
type ListAlibabaInstanceTypesParams struct {

	/*AccessKeyID*/
	AccessKeyID *string
	/*AccessKeySecret*/
	AccessKeySecret *string
	/*Credential*/
	Credential *string
	/*Region*/
	Region *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) WithTimeout(timeout time.Duration) *ListAlibabaInstanceTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) WithContext(ctx context.Context) *ListAlibabaInstanceTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) WithHTTPClient(client *http.Client) *ListAlibabaInstanceTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKeyID adds the accessKeyID to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) WithAccessKeyID(accessKeyID *string) *ListAlibabaInstanceTypesParams {
	o.SetAccessKeyID(accessKeyID)
	return o
}

// SetAccessKeyID adds the accessKeyId to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) SetAccessKeyID(accessKeyID *string) {
	o.AccessKeyID = accessKeyID
}

// WithAccessKeySecret adds the accessKeySecret to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) WithAccessKeySecret(accessKeySecret *string) *ListAlibabaInstanceTypesParams {
	o.SetAccessKeySecret(accessKeySecret)
	return o
}

// SetAccessKeySecret adds the accessKeySecret to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) SetAccessKeySecret(accessKeySecret *string) {
	o.AccessKeySecret = accessKeySecret
}

// WithCredential adds the credential to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) WithCredential(credential *string) *ListAlibabaInstanceTypesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithRegion adds the region to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) WithRegion(region *string) *ListAlibabaInstanceTypesParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the list alibaba instance types params
func (o *ListAlibabaInstanceTypesParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *ListAlibabaInstanceTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessKeyID != nil {

		// header param AccessKeyID
		if err := r.SetHeaderParam("AccessKeyID", *o.AccessKeyID); err != nil {
			return err
		}

	}

	if o.AccessKeySecret != nil {

		// header param AccessKeySecret
		if err := r.SetHeaderParam("AccessKeySecret", *o.AccessKeySecret); err != nil {
			return err
		}

	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}

	}

	if o.Region != nil {

		// header param Region
		if err := r.SetHeaderParam("Region", *o.Region); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
