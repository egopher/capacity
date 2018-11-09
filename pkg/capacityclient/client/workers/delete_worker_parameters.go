// Code generated by go-swagger; DO NOT EDIT.

package workers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewDeleteWorkerParams creates a new DeleteWorkerParams object
// with the default values initialized.
func NewDeleteWorkerParams() *DeleteWorkerParams {
	var ()
	return &DeleteWorkerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkerParamsWithTimeout creates a new DeleteWorkerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWorkerParamsWithTimeout(timeout time.Duration) *DeleteWorkerParams {
	var ()
	return &DeleteWorkerParams{

		timeout: timeout,
	}
}

// NewDeleteWorkerParamsWithContext creates a new DeleteWorkerParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWorkerParamsWithContext(ctx context.Context) *DeleteWorkerParams {
	var ()
	return &DeleteWorkerParams{

		Context: ctx,
	}
}

// NewDeleteWorkerParamsWithHTTPClient creates a new DeleteWorkerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWorkerParamsWithHTTPClient(client *http.Client) *DeleteWorkerParams {
	var ()
	return &DeleteWorkerParams{
		HTTPClient: client,
	}
}

/*DeleteWorkerParams contains all the parameters to send to the API endpoint
for the delete worker operation typically these are written to a http.Request
*/
type DeleteWorkerParams struct {

	/*MachineID*/
	MachineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete worker params
func (o *DeleteWorkerParams) WithTimeout(timeout time.Duration) *DeleteWorkerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete worker params
func (o *DeleteWorkerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete worker params
func (o *DeleteWorkerParams) WithContext(ctx context.Context) *DeleteWorkerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete worker params
func (o *DeleteWorkerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete worker params
func (o *DeleteWorkerParams) WithHTTPClient(client *http.Client) *DeleteWorkerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete worker params
func (o *DeleteWorkerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMachineID adds the machineID to the delete worker params
func (o *DeleteWorkerParams) WithMachineID(machineID string) *DeleteWorkerParams {
	o.SetMachineID(machineID)
	return o
}

// SetMachineID adds the machineId to the delete worker params
func (o *DeleteWorkerParams) SetMachineID(machineID string) {
	o.MachineID = machineID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param machineID
	if err := r.SetPathParam("machineID", o.MachineID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}