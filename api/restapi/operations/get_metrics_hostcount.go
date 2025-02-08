// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMetricsHostcountHandlerFunc turns a function with the right signature into a get metrics hostcount handler
type GetMetricsHostcountHandlerFunc func(GetMetricsHostcountParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMetricsHostcountHandlerFunc) Handle(params GetMetricsHostcountParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetMetricsHostcountHandler interface for that can handle valid get metrics hostcount params
type GetMetricsHostcountHandler interface {
	Handle(GetMetricsHostcountParams, interface{}) middleware.Responder
}

// NewGetMetricsHostcount creates a new http.Handler for the get metrics hostcount operation
func NewGetMetricsHostcount(ctx *middleware.Context, handler GetMetricsHostcountHandler) *GetMetricsHostcount {
	return &GetMetricsHostcount{Context: ctx, Handler: handler}
}

/*
	GetMetricsHostcount swagger:route GET /metrics/hostcount getMetricsHostcount

# Get host count metrics

Get metrics related to host counts.
*/
type GetMetricsHostcount struct {
	Context *middleware.Context
	Handler GetMetricsHostcountHandler
}

func (o *GetMetricsHostcount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMetricsHostcountParams()
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
