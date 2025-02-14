// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetMetricsErrorcountOKCode is the HTTP code returned for type GetMetricsErrorcountOK
const GetMetricsErrorcountOKCode int = 200

/*
GetMetricsErrorcountOK OK

swagger:response getMetricsErrorcountOK
*/
type GetMetricsErrorcountOK struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorCountMetrics `json:"body,omitempty"`
}

// NewGetMetricsErrorcountOK creates GetMetricsErrorcountOK with default headers values
func NewGetMetricsErrorcountOK() *GetMetricsErrorcountOK {

	return &GetMetricsErrorcountOK{}
}

// WithPayload adds the payload to the get metrics errorcount o k response
func (o *GetMetricsErrorcountOK) WithPayload(payload *models.ErrorCountMetrics) *GetMetricsErrorcountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metrics errorcount o k response
func (o *GetMetricsErrorcountOK) SetPayload(payload *models.ErrorCountMetrics) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetricsErrorcountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMetricsErrorcountInternalServerErrorCode is the HTTP code returned for type GetMetricsErrorcountInternalServerError
const GetMetricsErrorcountInternalServerErrorCode int = 500

/*
GetMetricsErrorcountInternalServerError Internal service error

swagger:response getMetricsErrorcountInternalServerError
*/
type GetMetricsErrorcountInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMetricsErrorcountInternalServerError creates GetMetricsErrorcountInternalServerError with default headers values
func NewGetMetricsErrorcountInternalServerError() *GetMetricsErrorcountInternalServerError {

	return &GetMetricsErrorcountInternalServerError{}
}

// WithPayload adds the payload to the get metrics errorcount internal server error response
func (o *GetMetricsErrorcountInternalServerError) WithPayload(payload *models.Error) *GetMetricsErrorcountInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metrics errorcount internal server error response
func (o *GetMetricsErrorcountInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetricsErrorcountInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
