// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetConfigMetricsAllHandlerFunc turns a function with the right signature into a get config metrics all handler
type GetConfigMetricsAllHandlerFunc func(GetConfigMetricsAllParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigMetricsAllHandlerFunc) Handle(params GetConfigMetricsAllParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetConfigMetricsAllHandler interface for that can handle valid get config metrics all params
type GetConfigMetricsAllHandler interface {
	Handle(GetConfigMetricsAllParams, interface{}) middleware.Responder
}

// NewGetConfigMetricsAll creates a new http.Handler for the get config metrics all operation
func NewGetConfigMetricsAll(ctx *middleware.Context, handler GetConfigMetricsAllHandler) *GetConfigMetricsAll {
	return &GetConfigMetricsAll{Context: ctx, Handler: handler}
}

/*
	GetConfigMetricsAll swagger:route GET /config/metrics/all getConfigMetricsAll

Scrape shamred metrics from the API
*/
type GetConfigMetricsAll struct {
	Context *middleware.Context
	Handler GetConfigMetricsAllHandler
}

func (o *GetConfigMetricsAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetConfigMetricsAllParams()
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
