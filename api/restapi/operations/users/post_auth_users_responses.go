// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// PostAuthUsersCreatedCode is the HTTP code returned for type PostAuthUsersCreated
const PostAuthUsersCreatedCode int = 201

/*
PostAuthUsersCreated Created

swagger:response postAuthUsersCreated
*/
type PostAuthUsersCreated struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPostAuthUsersCreated creates PostAuthUsersCreated with default headers values
func NewPostAuthUsersCreated() *PostAuthUsersCreated {

	return &PostAuthUsersCreated{}
}

// WithPayload adds the payload to the post auth users created response
func (o *PostAuthUsersCreated) WithPayload(payload *models.User) *PostAuthUsersCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth users created response
func (o *PostAuthUsersCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthUsersCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthUsersBadRequestCode is the HTTP code returned for type PostAuthUsersBadRequest
const PostAuthUsersBadRequestCode int = 400

/*
PostAuthUsersBadRequest Bad Request

swagger:response postAuthUsersBadRequest
*/
type PostAuthUsersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPostAuthUsersBadRequest creates PostAuthUsersBadRequest with default headers values
func NewPostAuthUsersBadRequest() *PostAuthUsersBadRequest {

	return &PostAuthUsersBadRequest{}
}

// WithPayload adds the payload to the post auth users bad request response
func (o *PostAuthUsersBadRequest) WithPayload(payload *models.ErrorResponse) *PostAuthUsersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth users bad request response
func (o *PostAuthUsersBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthUsersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthUsersInternalServerErrorCode is the HTTP code returned for type PostAuthUsersInternalServerError
const PostAuthUsersInternalServerErrorCode int = 500

/*
PostAuthUsersInternalServerError Internal Server Error

swagger:response postAuthUsersInternalServerError
*/
type PostAuthUsersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPostAuthUsersInternalServerError creates PostAuthUsersInternalServerError with default headers values
func NewPostAuthUsersInternalServerError() *PostAuthUsersInternalServerError {

	return &PostAuthUsersInternalServerError{}
}

// WithPayload adds the payload to the post auth users internal server error response
func (o *PostAuthUsersInternalServerError) WithPayload(payload *models.ErrorResponse) *PostAuthUsersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth users internal server error response
func (o *PostAuthUsersInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
