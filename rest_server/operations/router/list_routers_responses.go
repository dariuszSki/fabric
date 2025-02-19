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

package router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/fabric/rest_model"
)

// ListRoutersOKCode is the HTTP code returned for type ListRoutersOK
const ListRoutersOKCode int = 200

/*ListRoutersOK A list of routers

swagger:response listRoutersOK
*/
type ListRoutersOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListRoutersEnvelope `json:"body,omitempty"`
}

// NewListRoutersOK creates ListRoutersOK with default headers values
func NewListRoutersOK() *ListRoutersOK {

	return &ListRoutersOK{}
}

// WithPayload adds the payload to the list routers o k response
func (o *ListRoutersOK) WithPayload(payload *rest_model.ListRoutersEnvelope) *ListRoutersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list routers o k response
func (o *ListRoutersOK) SetPayload(payload *rest_model.ListRoutersEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRoutersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListRoutersUnauthorizedCode is the HTTP code returned for type ListRoutersUnauthorized
const ListRoutersUnauthorizedCode int = 401

/*ListRoutersUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listRoutersUnauthorized
*/
type ListRoutersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListRoutersUnauthorized creates ListRoutersUnauthorized with default headers values
func NewListRoutersUnauthorized() *ListRoutersUnauthorized {

	return &ListRoutersUnauthorized{}
}

// WithPayload adds the payload to the list routers unauthorized response
func (o *ListRoutersUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListRoutersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list routers unauthorized response
func (o *ListRoutersUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRoutersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
