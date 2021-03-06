// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateAppointmentHandlerFunc turns a function with the right signature into a create appointment handler
type CreateAppointmentHandlerFunc func(CreateAppointmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAppointmentHandlerFunc) Handle(params CreateAppointmentParams) middleware.Responder {
	return fn(params)
}

// CreateAppointmentHandler interface for that can handle valid create appointment params
type CreateAppointmentHandler interface {
	Handle(CreateAppointmentParams) middleware.Responder
}

// NewCreateAppointment creates a new http.Handler for the create appointment operation
func NewCreateAppointment(ctx *middleware.Context, handler CreateAppointmentHandler) *CreateAppointment {
	return &CreateAppointment{Context: ctx, Handler: handler}
}

/*CreateAppointment swagger:route POST /appointments createAppointment

CreateAppointment create appointment API

*/
type CreateAppointment struct {
	Context *middleware.Context
	Handler CreateAppointmentHandler
}

func (o *CreateAppointment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAppointmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
