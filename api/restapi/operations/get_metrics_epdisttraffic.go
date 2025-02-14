// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMetricsEpdisttrafficHandlerFunc turns a function with the right signature into a get metrics epdisttraffic handler
type GetMetricsEpdisttrafficHandlerFunc func(GetMetricsEpdisttrafficParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMetricsEpdisttrafficHandlerFunc) Handle(params GetMetricsEpdisttrafficParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetMetricsEpdisttrafficHandler interface for that can handle valid get metrics epdisttraffic params
type GetMetricsEpdisttrafficHandler interface {
	Handle(GetMetricsEpdisttrafficParams, interface{}) middleware.Responder
}

// NewGetMetricsEpdisttraffic creates a new http.Handler for the get metrics epdisttraffic operation
func NewGetMetricsEpdisttraffic(ctx *middleware.Context, handler GetMetricsEpdisttrafficHandler) *GetMetricsEpdisttraffic {
	return &GetMetricsEpdisttraffic{Context: ctx, Handler: handler}
}

/*
	GetMetricsEpdisttraffic swagger:route GET /metrics/epdisttraffic getMetricsEpdisttraffic

# Get endpoint distribution traffic metrics

Get metrics related to endpoint distribution traffic per service. The additionalProp is service name.
*/
type GetMetricsEpdisttraffic struct {
	Context *middleware.Context
	Handler GetMetricsEpdisttrafficHandler
}

func (o *GetMetricsEpdisttraffic) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMetricsEpdisttrafficParams()
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
