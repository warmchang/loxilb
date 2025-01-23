// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteConfigBgpPolicyApplyHandlerFunc turns a function with the right signature into a delete config bgp policy apply handler
type DeleteConfigBgpPolicyApplyHandlerFunc func(DeleteConfigBgpPolicyApplyParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteConfigBgpPolicyApplyHandlerFunc) Handle(params DeleteConfigBgpPolicyApplyParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteConfigBgpPolicyApplyHandler interface for that can handle valid delete config bgp policy apply params
type DeleteConfigBgpPolicyApplyHandler interface {
	Handle(DeleteConfigBgpPolicyApplyParams, interface{}) middleware.Responder
}

// NewDeleteConfigBgpPolicyApply creates a new http.Handler for the delete config bgp policy apply operation
func NewDeleteConfigBgpPolicyApply(ctx *middleware.Context, handler DeleteConfigBgpPolicyApplyHandler) *DeleteConfigBgpPolicyApply {
	return &DeleteConfigBgpPolicyApply{Context: ctx, Handler: handler}
}

/*
	DeleteConfigBgpPolicyApply swagger:route DELETE /config/bgp/policy/apply deleteConfigBgpPolicyApply

# Delete BGP Policy in neighbor

Delete BGP Policy in neighbor. It don't need "routeAction" in the attr body
*/
type DeleteConfigBgpPolicyApply struct {
	Context *middleware.Context
	Handler DeleteConfigBgpPolicyApplyHandler
}

func (o *DeleteConfigBgpPolicyApply) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteConfigBgpPolicyApplyParams()
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
