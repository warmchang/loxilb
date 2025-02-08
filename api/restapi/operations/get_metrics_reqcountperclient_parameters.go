// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetMetricsReqcountperclientParams creates a new GetMetricsReqcountperclientParams object
//
// There are no default values defined in the spec.
func NewGetMetricsReqcountperclientParams() GetMetricsReqcountperclientParams {

	return GetMetricsReqcountperclientParams{}
}

// GetMetricsReqcountperclientParams contains all the bound params for the get metrics reqcountperclient operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetMetricsReqcountperclient
type GetMetricsReqcountperclientParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMetricsReqcountperclientParams() beforehand.
func (o *GetMetricsReqcountperclientParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
