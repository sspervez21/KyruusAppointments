// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDoctorParams creates a new GetDoctorParams object
// no default values defined in spec.
func NewGetDoctorParams() GetDoctorParams {

	return GetDoctorParams{}
}

// GetDoctorParams contains all the bound params for the get doctor operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetDoctor
type GetDoctorParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	DoctorID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetDoctorParams() beforehand.
func (o *GetDoctorParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDoctorID, rhkDoctorID, _ := route.Params.GetOK("doctorId")
	if err := o.bindDoctorID(rDoctorID, rhkDoctorID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDoctorID binds and validates parameter DoctorID from path.
func (o *GetDoctorParams) bindDoctorID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("doctorId", "path", "int64", raw)
	}
	o.DoctorID = value

	return nil
}
