// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAuthUsersIDHandlerFunc turns a function with the right signature into a delete auth users ID handler
type DeleteAuthUsersIDHandlerFunc func(DeleteAuthUsersIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAuthUsersIDHandlerFunc) Handle(params DeleteAuthUsersIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteAuthUsersIDHandler interface for that can handle valid delete auth users ID params
type DeleteAuthUsersIDHandler interface {
	Handle(DeleteAuthUsersIDParams, interface{}) middleware.Responder
}

// NewDeleteAuthUsersID creates a new http.Handler for the delete auth users ID operation
func NewDeleteAuthUsersID(ctx *middleware.Context, handler DeleteAuthUsersIDHandler) *DeleteAuthUsersID {
	return &DeleteAuthUsersID{Context: ctx, Handler: handler}
}

/*
	DeleteAuthUsersID swagger:route DELETE /auth/users/{id} users deleteAuthUsersId

# Delete user

Deletes a user by its ID
*/
type DeleteAuthUsersID struct {
	Context *middleware.Context
	Handler DeleteAuthUsersIDHandler
}

func (o *DeleteAuthUsersID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteAuthUsersIDParams()
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
