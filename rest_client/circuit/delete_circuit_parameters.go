// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package circuit

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

	"github.com/openziti/fabric/rest_model"
)

// NewDeleteCircuitParams creates a new DeleteCircuitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCircuitParams() *DeleteCircuitParams {
	return &DeleteCircuitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCircuitParamsWithTimeout creates a new DeleteCircuitParams object
// with the ability to set a timeout on a request.
func NewDeleteCircuitParamsWithTimeout(timeout time.Duration) *DeleteCircuitParams {
	return &DeleteCircuitParams{
		timeout: timeout,
	}
}

// NewDeleteCircuitParamsWithContext creates a new DeleteCircuitParams object
// with the ability to set a context for a request.
func NewDeleteCircuitParamsWithContext(ctx context.Context) *DeleteCircuitParams {
	return &DeleteCircuitParams{
		Context: ctx,
	}
}

// NewDeleteCircuitParamsWithHTTPClient creates a new DeleteCircuitParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCircuitParamsWithHTTPClient(client *http.Client) *DeleteCircuitParams {
	return &DeleteCircuitParams{
		HTTPClient: client,
	}
}

/* DeleteCircuitParams contains all the parameters to send to the API endpoint
   for the delete circuit operation.

   Typically these are written to a http.Request.
*/
type DeleteCircuitParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	/* Options.

	   A circuit delete object
	*/
	Options *rest_model.CircuitDelete

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete circuit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCircuitParams) WithDefaults() *DeleteCircuitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete circuit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCircuitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete circuit params
func (o *DeleteCircuitParams) WithTimeout(timeout time.Duration) *DeleteCircuitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete circuit params
func (o *DeleteCircuitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete circuit params
func (o *DeleteCircuitParams) WithContext(ctx context.Context) *DeleteCircuitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete circuit params
func (o *DeleteCircuitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete circuit params
func (o *DeleteCircuitParams) WithHTTPClient(client *http.Client) *DeleteCircuitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete circuit params
func (o *DeleteCircuitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete circuit params
func (o *DeleteCircuitParams) WithID(id string) *DeleteCircuitParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete circuit params
func (o *DeleteCircuitParams) SetID(id string) {
	o.ID = id
}

// WithOptions adds the options to the delete circuit params
func (o *DeleteCircuitParams) WithOptions(options *rest_model.CircuitDelete) *DeleteCircuitParams {
	o.SetOptions(options)
	return o
}

// SetOptions adds the options to the delete circuit params
func (o *DeleteCircuitParams) SetOptions(options *rest_model.CircuitDelete) {
	o.Options = options
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCircuitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Options != nil {
		if err := r.SetBodyParam(o.Options); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
