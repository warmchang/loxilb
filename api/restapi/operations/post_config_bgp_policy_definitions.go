// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostConfigBgpPolicyDefinitionsHandlerFunc turns a function with the right signature into a post config bgp policy definitions handler
type PostConfigBgpPolicyDefinitionsHandlerFunc func(PostConfigBgpPolicyDefinitionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostConfigBgpPolicyDefinitionsHandlerFunc) Handle(params PostConfigBgpPolicyDefinitionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostConfigBgpPolicyDefinitionsHandler interface for that can handle valid post config bgp policy definitions params
type PostConfigBgpPolicyDefinitionsHandler interface {
	Handle(PostConfigBgpPolicyDefinitionsParams, interface{}) middleware.Responder
}

// NewPostConfigBgpPolicyDefinitions creates a new http.Handler for the post config bgp policy definitions operation
func NewPostConfigBgpPolicyDefinitions(ctx *middleware.Context, handler PostConfigBgpPolicyDefinitionsHandler) *PostConfigBgpPolicyDefinitions {
	return &PostConfigBgpPolicyDefinitions{Context: ctx, Handler: handler}
}

/*
	PostConfigBgpPolicyDefinitions swagger:route POST /config/bgp/policy/definitions postConfigBgpPolicyDefinitions

# Adds a BGP Policy

Adds a BGP Policy
*/
type PostConfigBgpPolicyDefinitions struct {
	Context *middleware.Context
	Handler PostConfigBgpPolicyDefinitionsHandler
}

func (o *PostConfigBgpPolicyDefinitions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostConfigBgpPolicyDefinitionsParams()
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
