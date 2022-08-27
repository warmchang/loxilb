// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/loxilb-io/loxilb/api/models"
)

// NewPostConfigSessionParams creates a new PostConfigSessionParams object
//
// There are no default values defined in the spec.
func NewPostConfigSessionParams() PostConfigSessionParams {

	return PostConfigSessionParams{}
}

// PostConfigSessionParams contains all the bound params for the post config session operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostConfigSession
type PostConfigSessionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Attributes for 5G service session
	  Required: true
	  In: body
	*/
	Attr *models.SessionEntry
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostConfigSessionParams() beforehand.
func (o *PostConfigSessionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.SessionEntry
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("attr", "body", ""))
			} else {
				res = append(res, errors.NewParseError("attr", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Attr = &body
			}
		}
	} else {
		res = append(res, errors.Required("attr", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
