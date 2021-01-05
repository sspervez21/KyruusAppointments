// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateDoctorHandlerFunc turns a function with the right signature into a create doctor handler
type CreateDoctorHandlerFunc func(CreateDoctorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateDoctorHandlerFunc) Handle(params CreateDoctorParams) middleware.Responder {
	return fn(params)
}

// CreateDoctorHandler interface for that can handle valid create doctor params
type CreateDoctorHandler interface {
	Handle(CreateDoctorParams) middleware.Responder
}

// NewCreateDoctor creates a new http.Handler for the create doctor operation
func NewCreateDoctor(ctx *middleware.Context, handler CreateDoctorHandler) *CreateDoctor {
	return &CreateDoctor{Context: ctx, Handler: handler}
}

/*CreateDoctor swagger:route POST /doctors createDoctor

CreateDoctor create doctor API

*/
type CreateDoctor struct {
	Context *middleware.Context
	Handler CreateDoctorHandler
}

func (o *CreateDoctor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateDoctorParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
