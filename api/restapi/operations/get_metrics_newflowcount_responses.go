// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetMetricsNewflowcountOKCode is the HTTP code returned for type GetMetricsNewflowcountOK
const GetMetricsNewflowcountOKCode int = 200

/*
GetMetricsNewflowcountOK OK

swagger:response getMetricsNewflowcountOK
*/
type GetMetricsNewflowcountOK struct {

	/*
	  In: Body
	*/
	Payload *models.NewFlowCountMetrics `json:"body,omitempty"`
}

// NewGetMetricsNewflowcountOK creates GetMetricsNewflowcountOK with default headers values
func NewGetMetricsNewflowcountOK() *GetMetricsNewflowcountOK {

	return &GetMetricsNewflowcountOK{}
}

// WithPayload adds the payload to the get metrics newflowcount o k response
func (o *GetMetricsNewflowcountOK) WithPayload(payload *models.NewFlowCountMetrics) *GetMetricsNewflowcountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metrics newflowcount o k response
func (o *GetMetricsNewflowcountOK) SetPayload(payload *models.NewFlowCountMetrics) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetricsNewflowcountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMetricsNewflowcountInternalServerErrorCode is the HTTP code returned for type GetMetricsNewflowcountInternalServerError
const GetMetricsNewflowcountInternalServerErrorCode int = 500

/*
GetMetricsNewflowcountInternalServerError Internal service error

swagger:response getMetricsNewflowcountInternalServerError
*/
type GetMetricsNewflowcountInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMetricsNewflowcountInternalServerError creates GetMetricsNewflowcountInternalServerError with default headers values
func NewGetMetricsNewflowcountInternalServerError() *GetMetricsNewflowcountInternalServerError {

	return &GetMetricsNewflowcountInternalServerError{}
}

// WithPayload adds the payload to the get metrics newflowcount internal server error response
func (o *GetMetricsNewflowcountInternalServerError) WithPayload(payload *models.Error) *GetMetricsNewflowcountInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metrics newflowcount internal server error response
func (o *GetMetricsNewflowcountInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetricsNewflowcountInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
