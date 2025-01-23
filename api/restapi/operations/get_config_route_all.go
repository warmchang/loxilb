// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigRouteAllHandlerFunc turns a function with the right signature into a get config route all handler
type GetConfigRouteAllHandlerFunc func(GetConfigRouteAllParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigRouteAllHandlerFunc) Handle(params GetConfigRouteAllParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetConfigRouteAllHandler interface for that can handle valid get config route all params
type GetConfigRouteAllHandler interface {
	Handle(GetConfigRouteAllParams, interface{}) middleware.Responder
}

// NewGetConfigRouteAll creates a new http.Handler for the get config route all operation
func NewGetConfigRouteAll(ctx *middleware.Context, handler GetConfigRouteAllHandler) *GetConfigRouteAll {
	return &GetConfigRouteAll{Context: ctx, Handler: handler}
}

/*
	GetConfigRouteAll swagger:route GET /config/route/all getConfigRouteAll

# Get all route table

Get all route table
*/
type GetConfigRouteAll struct {
	Context *middleware.Context
	Handler GetConfigRouteAllHandler
}

func (o *GetConfigRouteAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetConfigRouteAllParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetConfigRouteAllOKBody get config route all o k body
//
// swagger:model GetConfigRouteAllOKBody
type GetConfigRouteAllOKBody struct {

	// route attr
	RouteAttr []*models.RouteGetEntry `json:"routeAttr"`
}

// Validate validates this get config route all o k body
func (o *GetConfigRouteAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRouteAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigRouteAllOKBody) validateRouteAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.RouteAttr) { // not required
		return nil
	}

	for i := 0; i < len(o.RouteAttr); i++ {
		if swag.IsZero(o.RouteAttr[i]) { // not required
			continue
		}

		if o.RouteAttr[i] != nil {
			if err := o.RouteAttr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigRouteAllOK" + "." + "routeAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigRouteAllOK" + "." + "routeAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get config route all o k body based on the context it is used
func (o *GetConfigRouteAllOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRouteAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigRouteAllOKBody) contextValidateRouteAttr(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RouteAttr); i++ {

		if o.RouteAttr[i] != nil {
			if err := o.RouteAttr[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigRouteAllOK" + "." + "routeAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigRouteAllOK" + "." + "routeAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigRouteAllOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigRouteAllOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigRouteAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
