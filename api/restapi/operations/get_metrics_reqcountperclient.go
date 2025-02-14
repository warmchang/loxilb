// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMetricsReqcountperclientHandlerFunc turns a function with the right signature into a get metrics reqcountperclient handler
type GetMetricsReqcountperclientHandlerFunc func(GetMetricsReqcountperclientParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMetricsReqcountperclientHandlerFunc) Handle(params GetMetricsReqcountperclientParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetMetricsReqcountperclientHandler interface for that can handle valid get metrics reqcountperclient params
type GetMetricsReqcountperclientHandler interface {
	Handle(GetMetricsReqcountperclientParams, interface{}) middleware.Responder
}

// NewGetMetricsReqcountperclient creates a new http.Handler for the get metrics reqcountperclient operation
func NewGetMetricsReqcountperclient(ctx *middleware.Context, handler GetMetricsReqcountperclientHandler) *GetMetricsReqcountperclient {
	return &GetMetricsReqcountperclient{Context: ctx, Handler: handler}
}

/*
	GetMetricsReqcountperclient swagger:route GET /metrics/reqcountperclient getMetricsReqcountperclient

# Get request count per client metrics

Get metrics related to request counts per client. The additionalProp is client IP address.
*/
type GetMetricsReqcountperclient struct {
	Context *middleware.Context
	Handler GetMetricsReqcountperclientHandler
}

func (o *GetMetricsReqcountperclient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMetricsReqcountperclientParams()
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
