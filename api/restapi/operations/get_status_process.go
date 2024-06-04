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

// GetStatusProcessHandlerFunc turns a function with the right signature into a get status process handler
type GetStatusProcessHandlerFunc func(GetStatusProcessParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStatusProcessHandlerFunc) Handle(params GetStatusProcessParams) middleware.Responder {
	return fn(params)
}

// GetStatusProcessHandler interface for that can handle valid get status process params
type GetStatusProcessHandler interface {
	Handle(GetStatusProcessParams) middleware.Responder
}

// NewGetStatusProcess creates a new http.Handler for the get status process operation
func NewGetStatusProcess(ctx *middleware.Context, handler GetStatusProcessHandler) *GetStatusProcess {
	return &GetStatusProcess{Context: ctx, Handler: handler}
}

/*
	GetStatusProcess swagger:route GET /status/process getStatusProcess

# Get a process based on CPU usage info in the device

Get a process based on high usage CPU(linux command "top") in the device or system.
*/
type GetStatusProcess struct {
	Context *middleware.Context
	Handler GetStatusProcessHandler
}

func (o *GetStatusProcess) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetStatusProcessParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetStatusProcessOKBody get status process o k body
//
// swagger:model GetStatusProcessOKBody
type GetStatusProcessOKBody struct {

	// process attr
	ProcessAttr []*models.ProcessInfoEntry `json:"processAttr"`
}

// Validate validates this get status process o k body
func (o *GetStatusProcessOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProcessAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStatusProcessOKBody) validateProcessAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.ProcessAttr) { // not required
		return nil
	}

	for i := 0; i < len(o.ProcessAttr); i++ {
		if swag.IsZero(o.ProcessAttr[i]) { // not required
			continue
		}

		if o.ProcessAttr[i] != nil {
			if err := o.ProcessAttr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getStatusProcessOK" + "." + "processAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getStatusProcessOK" + "." + "processAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get status process o k body based on the context it is used
func (o *GetStatusProcessOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateProcessAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStatusProcessOKBody) contextValidateProcessAttr(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ProcessAttr); i++ {

		if o.ProcessAttr[i] != nil {

			if swag.IsZero(o.ProcessAttr[i]) { // not required
				return nil
			}

			if err := o.ProcessAttr[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getStatusProcessOK" + "." + "processAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getStatusProcessOK" + "." + "processAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStatusProcessOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStatusProcessOKBody) UnmarshalBinary(b []byte) error {
	var res GetStatusProcessOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
