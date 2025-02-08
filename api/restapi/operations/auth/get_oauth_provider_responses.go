// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetOauthProviderFoundCode is the HTTP code returned for type GetOauthProviderFound
const GetOauthProviderFoundCode int = 302

/*
GetOauthProviderFound Found

swagger:response getOauthProviderFound
*/
type GetOauthProviderFound struct {

	/*
	  In: Body
	*/
	Payload *models.OauthMessageResponse `json:"body,omitempty"`
}

// NewGetOauthProviderFound creates GetOauthProviderFound with default headers values
func NewGetOauthProviderFound() *GetOauthProviderFound {

	return &GetOauthProviderFound{}
}

// WithPayload adds the payload to the get oauth provider found response
func (o *GetOauthProviderFound) WithPayload(payload *models.OauthMessageResponse) *GetOauthProviderFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get oauth provider found response
func (o *GetOauthProviderFound) SetPayload(payload *models.OauthMessageResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOauthProviderFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(302)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOauthProviderBadRequestCode is the HTTP code returned for type GetOauthProviderBadRequest
const GetOauthProviderBadRequestCode int = 400

/*
GetOauthProviderBadRequest Bad Request

swagger:response getOauthProviderBadRequest
*/
type GetOauthProviderBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.OauthErrorResponse `json:"body,omitempty"`
}

// NewGetOauthProviderBadRequest creates GetOauthProviderBadRequest with default headers values
func NewGetOauthProviderBadRequest() *GetOauthProviderBadRequest {

	return &GetOauthProviderBadRequest{}
}

// WithPayload adds the payload to the get oauth provider bad request response
func (o *GetOauthProviderBadRequest) WithPayload(payload *models.OauthErrorResponse) *GetOauthProviderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get oauth provider bad request response
func (o *GetOauthProviderBadRequest) SetPayload(payload *models.OauthErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOauthProviderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
