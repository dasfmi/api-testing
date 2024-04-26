// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"axiom/models/components"
)

type IngestIntoDatasetRequestBody struct {
}

type IngestIntoDatasetRawRequest struct {
	// Unique ID of the dataset where you want to send data.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Data you want to send to Axiom in a supported format.
	RequestBody []byte `request:"mediaType=text/csv"`
	// The name of the field to use as the timestamp. If not specified, the timestamp will be the time the event was received by Axiom.
	TimestampField *string `queryParam:"style=form,explode=true,name=timestamp-field"`
	// The date-time format of the timestamp field. The reference time is `Mon Jan 2 15:04:05 -0700 MST 2006`, as specified in https://pkg.go.dev/time/?tab=doc#Parse
	TimestampFormat *string `queryParam:"style=form,explode=true,name=timestamp-format"`
	// The delimiter to use when parsing CSV data. If not specified, the default delimiter is `,`.
	CsvDelimiter *string `queryParam:"style=form,explode=true,name=csv-delimiter"`
	// A list of optional comma separated fields to use for CSV ingestion. If not specified, the first line of the CSV will be used as the field names.
	XAxiomCSVFields []string `header:"style=simple,explode=false,name=X-Axiom-CSV-Fields"`
	// An optional JSON encoded object with additional labels to add to all events in the request
	XAxiomEventLabels *string `header:"style=simple,explode=false,name=X-Axiom-Event-Labels"`
}

func (o *IngestIntoDatasetRawRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *IngestIntoDatasetRawRequest) GetRequestBody() []byte {
	if o == nil {
		return []byte{}
	}
	return o.RequestBody
}

func (o *IngestIntoDatasetRawRequest) GetTimestampField() *string {
	if o == nil {
		return nil
	}
	return o.TimestampField
}

func (o *IngestIntoDatasetRawRequest) GetTimestampFormat() *string {
	if o == nil {
		return nil
	}
	return o.TimestampFormat
}

func (o *IngestIntoDatasetRawRequest) GetCsvDelimiter() *string {
	if o == nil {
		return nil
	}
	return o.CsvDelimiter
}

func (o *IngestIntoDatasetRawRequest) GetXAxiomCSVFields() []string {
	if o == nil {
		return nil
	}
	return o.XAxiomCSVFields
}

func (o *IngestIntoDatasetRawRequest) GetXAxiomEventLabels() *string {
	if o == nil {
		return nil
	}
	return o.XAxiomEventLabels
}

type IngestIntoDatasetRawResponse struct {
	HTTPMeta components.HTTPMetadata
	// ingest status
	IngestStatus *components.IngestStatus
}

func (o *IngestIntoDatasetRawResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *IngestIntoDatasetRawResponse) GetIngestStatus() *components.IngestStatus {
	if o == nil {
		return nil
	}
	return o.IngestStatus
}
