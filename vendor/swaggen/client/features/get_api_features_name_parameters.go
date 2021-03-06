package features

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAPIFeaturesNameParams creates a new GetAPIFeaturesNameParams object
// with the default values initialized.
func NewGetAPIFeaturesNameParams() *GetAPIFeaturesNameParams {
	var ()
	return &GetAPIFeaturesNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIFeaturesNameParamsWithTimeout creates a new GetAPIFeaturesNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIFeaturesNameParamsWithTimeout(timeout time.Duration) *GetAPIFeaturesNameParams {
	var ()
	return &GetAPIFeaturesNameParams{

		timeout: timeout,
	}
}

// NewGetAPIFeaturesNameParamsWithContext creates a new GetAPIFeaturesNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIFeaturesNameParamsWithContext(ctx context.Context) *GetAPIFeaturesNameParams {
	var ()
	return &GetAPIFeaturesNameParams{

		Context: ctx,
	}
}

/*GetAPIFeaturesNameParams contains all the parameters to send to the API endpoint
for the get API features name operation typically these are written to a http.Request
*/
type GetAPIFeaturesNameParams struct {

	/*Name
	  the feature name.

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API features name params
func (o *GetAPIFeaturesNameParams) WithTimeout(timeout time.Duration) *GetAPIFeaturesNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API features name params
func (o *GetAPIFeaturesNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API features name params
func (o *GetAPIFeaturesNameParams) WithContext(ctx context.Context) *GetAPIFeaturesNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API features name params
func (o *GetAPIFeaturesNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithName adds the name to the get API features name params
func (o *GetAPIFeaturesNameParams) WithName(name string) *GetAPIFeaturesNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API features name params
func (o *GetAPIFeaturesNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIFeaturesNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
