// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostConfigMetricsHandlerFunc turns a function with the right signature into a post config metrics handler
type PostConfigMetricsHandlerFunc func(PostConfigMetricsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostConfigMetricsHandlerFunc) Handle(params PostConfigMetricsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostConfigMetricsHandler interface for that can handle valid post config metrics params
type PostConfigMetricsHandler interface {
	Handle(PostConfigMetricsParams, interface{}) middleware.Responder
}

// NewPostConfigMetrics creates a new http.Handler for the post config metrics operation
func NewPostConfigMetrics(ctx *middleware.Context, handler PostConfigMetricsHandler) *PostConfigMetrics {
	return &PostConfigMetrics{Context: ctx, Handler: handler}
}

/*
	PostConfigMetrics swagger:route POST /config/metrics postConfigMetrics

turn on prometheus option
*/
type PostConfigMetrics struct {
	Context *middleware.Context
	Handler PostConfigMetricsHandler
}

func (o *PostConfigMetrics) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostConfigMetricsParams()
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
