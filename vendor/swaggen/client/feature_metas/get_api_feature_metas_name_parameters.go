package feature_metas

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

// NewGetAPIFeatureMetasNameParams creates a new GetAPIFeatureMetasNameParams object
// with the default values initialized.
func NewGetAPIFeatureMetasNameParams() *GetAPIFeatureMetasNameParams {
	var ()
	return &GetAPIFeatureMetasNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIFeatureMetasNameParamsWithTimeout creates a new GetAPIFeatureMetasNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIFeatureMetasNameParamsWithTimeout(timeout time.Duration) *GetAPIFeatureMetasNameParams {
	var ()
	return &GetAPIFeatureMetasNameParams{

		timeout: timeout,
	}
}

// NewGetAPIFeatureMetasNameParamsWithContext creates a new GetAPIFeatureMetasNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIFeatureMetasNameParamsWithContext(ctx context.Context) *GetAPIFeatureMetasNameParams {
	var ()
	return &GetAPIFeatureMetasNameParams{

		Context: ctx,
	}
}

/*GetAPIFeatureMetasNameParams contains all the parameters to send to the API endpoint
for the get API feature metas name operation typically these are written to a http.Request
*/
type GetAPIFeatureMetasNameParams struct {

	/*Name
	  the feature name.

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API feature metas name params
func (o *GetAPIFeatureMetasNameParams) WithTimeout(timeout time.Duration) *GetAPIFeatureMetasNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API feature metas name params
func (o *GetAPIFeatureMetasNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API feature metas name params
func (o *GetAPIFeatureMetasNameParams) WithContext(ctx context.Context) *GetAPIFeatureMetasNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API feature metas name params
func (o *GetAPIFeatureMetasNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithName adds the name to the get API feature metas name params
func (o *GetAPIFeatureMetasNameParams) WithName(name string) *GetAPIFeatureMetasNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API feature metas name params
func (o *GetAPIFeatureMetasNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIFeatureMetasNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
