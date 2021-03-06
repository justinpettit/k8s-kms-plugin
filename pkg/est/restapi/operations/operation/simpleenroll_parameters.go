// Code generated by go-swagger; DO NOT EDIT.

package operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewSimpleenrollParams creates a new SimpleenrollParams object
// no default values defined in spec.
func NewSimpleenrollParams() SimpleenrollParams {

	return SimpleenrollParams{}
}

// SimpleenrollParams contains all the bound params for the simpleenroll operation
// typically these are obtained from a http.Request
//
// swagger:parameters simpleenroll
type SimpleenrollParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*PKCS10 certificate request (DER format, base64 encoded)
	  Required: true
	  In: body
	*/
	Certrequest string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSimpleenrollParams() beforehand.
func (o *SimpleenrollParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body string
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("certrequest", "body", ""))
			} else {
				res = append(res, errors.NewParseError("certrequest", "body", "", err))
			}
		} else {
			// no validation required on inline body
			o.Certrequest = body
		}
	} else {
		res = append(res, errors.Required("certrequest", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
