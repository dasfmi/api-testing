// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"axiom/models/components"
)

type QueryDatasetRequest struct {
	SaveAsKind              *string                            `queryParam:"style=form,explode=true,name=saveAsKind"`
	ID                      string                             `pathParam:"style=simple,explode=false,name=id"`
	StreamingDuration       *string                            `queryParam:"style=form,explode=true,name=streaming-duration"`
	Nocache                 *bool                              `queryParam:"style=form,explode=true,name=nocache"`
	QueryRequestWithOptions components.QueryRequestWithOptions `request:"mediaType=application/json"`
}

func (o *QueryDatasetRequest) GetSaveAsKind() *string {
	if o == nil {
		return nil
	}
	return o.SaveAsKind
}

func (o *QueryDatasetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *QueryDatasetRequest) GetStreamingDuration() *string {
	if o == nil {
		return nil
	}
	return o.StreamingDuration
}

func (o *QueryDatasetRequest) GetNocache() *bool {
	if o == nil {
		return nil
	}
	return o.Nocache
}

func (o *QueryDatasetRequest) GetQueryRequestWithOptions() components.QueryRequestWithOptions {
	if o == nil {
		return components.QueryRequestWithOptions{}
	}
	return o.QueryRequestWithOptions
}

type QueryDatasetResponse struct {
	HTTPMeta components.HTTPMetadata
	// Result
	Result  *components.Result
	Headers map[string][]string
}

func (o *QueryDatasetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *QueryDatasetResponse) GetResult() *components.Result {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *QueryDatasetResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
