// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostConfigEndpointHandlerFunc turns a function with the right signature into a post config endpoint handler
type PostConfigEndpointHandlerFunc func(PostConfigEndpointParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostConfigEndpointHandlerFunc) Handle(params PostConfigEndpointParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostConfigEndpointHandler interface for that can handle valid post config endpoint params
type PostConfigEndpointHandler interface {
	Handle(PostConfigEndpointParams, interface{}) middleware.Responder
}

// NewPostConfigEndpoint creates a new http.Handler for the post config endpoint operation
func NewPostConfigEndpoint(ctx *middleware.Context, handler PostConfigEndpointHandler) *PostConfigEndpoint {
	return &PostConfigEndpoint{Context: ctx, Handler: handler}
}

/*
	PostConfigEndpoint swagger:route POST /config/endpoint postConfigEndpoint

# Adds a LB endpoint for monitoring

Adds a LB endpoint for monitoring
*/
type PostConfigEndpoint struct {
	Context *middleware.Context
	Handler PostConfigEndpointHandler
}

func (o *PostConfigEndpoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostConfigEndpointParams()
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
