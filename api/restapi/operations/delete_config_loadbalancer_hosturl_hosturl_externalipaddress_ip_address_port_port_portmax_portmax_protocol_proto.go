// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProtoHandlerFunc turns a function with the right signature into a delete config loadbalancer hosturl hosturl externalipaddress IP address port port portmax portmax protocol proto handler
type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProtoHandlerFunc func(DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProtoParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProtoHandlerFunc) Handle(params DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProtoParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProtoHandler interface for that can handle valid delete config loadbalancer hosturl hosturl externalipaddress IP address port port portmax portmax protocol proto params
type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProtoHandler interface {
	Handle(DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProtoParams, interface{}) middleware.Responder
}

// NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProto creates a new http.Handler for the delete config loadbalancer hosturl hosturl externalipaddress IP address port port portmax portmax protocol proto operation
func NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProto(ctx *middleware.Context, handler DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProtoHandler) *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProto {
	return &DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProto{Context: ctx, Handler: handler}
}

/*
	DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProto swagger:route DELETE /config/loadbalancer/hosturl/{hosturl}/externalipaddress/{ip_address}/port/{port}/portmax/{portmax}/protocol/{proto} deleteConfigLoadbalancerHosturlHosturlExternalipaddressIpAddressPortPortPortmaxPortmaxProtocolProto

# Delete an existing Load balancer service

Delete an existing load balancer service with .
*/
type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProto struct {
	Context *middleware.Context
	Handler DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProtoHandler
}

func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProto) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortPortmaxPortmaxProtocolProtoParams()
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
