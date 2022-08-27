// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigConntrackAllHandlerFunc turns a function with the right signature into a get config conntrack all handler
type GetConfigConntrackAllHandlerFunc func(GetConfigConntrackAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigConntrackAllHandlerFunc) Handle(params GetConfigConntrackAllParams) middleware.Responder {
	return fn(params)
}

// GetConfigConntrackAllHandler interface for that can handle valid get config conntrack all params
type GetConfigConntrackAllHandler interface {
	Handle(GetConfigConntrackAllParams) middleware.Responder
}

// NewGetConfigConntrackAll creates a new http.Handler for the get config conntrack all operation
func NewGetConfigConntrackAll(ctx *middleware.Context, handler GetConfigConntrackAllHandler) *GetConfigConntrackAll {
	return &GetConfigConntrackAll{Context: ctx, Handler: handler}
}

/* GetConfigConntrackAll swagger:route GET /config/conntrack/all getConfigConntrackAll

Get all of the conntrack entries.

Get all of the conntrack infomation for all of the service.

*/
type GetConfigConntrackAll struct {
	Context *middleware.Context
	Handler GetConfigConntrackAllHandler
}

func (o *GetConfigConntrackAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetConfigConntrackAllParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetConfigConntrackAllOKBody get config conntrack all o k body
//
// swagger:model GetConfigConntrackAllOKBody
type GetConfigConntrackAllOKBody struct {

	// ct attr
	CtAttr models.ConntrackEntry `json:"ctAttr"`
}

// Validate validates this get config conntrack all o k body
func (o *GetConfigConntrackAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCtAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigConntrackAllOKBody) validateCtAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.CtAttr) { // not required
		return nil
	}

	if err := o.CtAttr.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getConfigConntrackAllOK" + "." + "ctAttr")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getConfigConntrackAllOK" + "." + "ctAttr")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get config conntrack all o k body based on the context it is used
func (o *GetConfigConntrackAllOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCtAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigConntrackAllOKBody) contextValidateCtAttr(ctx context.Context, formats strfmt.Registry) error {

	if err := o.CtAttr.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getConfigConntrackAllOK" + "." + "ctAttr")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getConfigConntrackAllOK" + "." + "ctAttr")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigConntrackAllOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigConntrackAllOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigConntrackAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
