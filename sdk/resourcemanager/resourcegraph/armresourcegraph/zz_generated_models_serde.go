//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcegraph

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type DateTimeInterval.
func (d DateTimeInterval) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "end", d.End)
	populateTimeRFC3339(objectMap, "start", d.Start)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DateTimeInterval.
func (d *DateTimeInterval) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "end":
			err = unpopulateTimeRFC3339(val, &d.End)
			delete(rawMsg, key)
		case "start":
			err = unpopulateTimeRFC3339(val, &d.Start)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Error.
func (e Error) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetails.
func (e ErrorDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	if e.AdditionalProperties != nil {
		for key, val := range e.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetails.
func (e *ErrorDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, &e.Code)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &e.Message)
			delete(rawMsg, key)
		default:
			if e.AdditionalProperties == nil {
				e.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				e.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetFacet implements the FacetClassification interface for type Facet.
func (f *Facet) GetFacet() *Facet { return f }

// GetFacet implements the FacetClassification interface for type FacetError.
func (f *FacetError) GetFacet() *Facet {
	return &Facet{
		Expression: f.Expression,
		ResultType: f.ResultType,
	}
}

// MarshalJSON implements the json.Marshaller interface for type FacetError.
func (f FacetError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "errors", f.Errors)
	populate(objectMap, "expression", f.Expression)
	objectMap["resultType"] = "FacetError"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetError.
func (f *FacetError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "errors":
			err = unpopulate(val, &f.Errors)
			delete(rawMsg, key)
		case "expression":
			err = unpopulate(val, &f.Expression)
			delete(rawMsg, key)
		case "resultType":
			err = unpopulate(val, &f.ResultType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetFacet implements the FacetClassification interface for type FacetResult.
func (f *FacetResult) GetFacet() *Facet {
	return &Facet{
		Expression: f.Expression,
		ResultType: f.ResultType,
	}
}

// MarshalJSON implements the json.Marshaller interface for type FacetResult.
func (f FacetResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "count", f.Count)
	populate(objectMap, "data", &f.Data)
	populate(objectMap, "expression", f.Expression)
	objectMap["resultType"] = "FacetResult"
	populate(objectMap, "totalRecords", f.TotalRecords)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetResult.
func (f *FacetResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, &f.Count)
			delete(rawMsg, key)
		case "data":
			err = unpopulate(val, &f.Data)
			delete(rawMsg, key)
		case "expression":
			err = unpopulate(val, &f.Expression)
			delete(rawMsg, key)
		case "resultType":
			err = unpopulate(val, &f.ResultType)
			delete(rawMsg, key)
		case "totalRecords":
			err = unpopulate(val, &f.TotalRecords)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type QueryRequest.
func (q QueryRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "facets", q.Facets)
	populate(objectMap, "managementGroups", q.ManagementGroups)
	populate(objectMap, "options", q.Options)
	populate(objectMap, "query", q.Query)
	populate(objectMap, "subscriptions", q.Subscriptions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type QueryResponse.
func (q QueryResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "count", q.Count)
	populate(objectMap, "data", &q.Data)
	populate(objectMap, "facets", q.Facets)
	populate(objectMap, "resultTruncated", q.ResultTruncated)
	populate(objectMap, "$skipToken", q.SkipToken)
	populate(objectMap, "totalRecords", q.TotalRecords)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type QueryResponse.
func (q *QueryResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, &q.Count)
			delete(rawMsg, key)
		case "data":
			err = unpopulate(val, &q.Data)
			delete(rawMsg, key)
		case "facets":
			q.Facets, err = unmarshalFacetClassificationArray(val)
			delete(rawMsg, key)
		case "resultTruncated":
			err = unpopulate(val, &q.ResultTruncated)
			delete(rawMsg, key)
		case "$skipToken":
			err = unpopulate(val, &q.SkipToken)
			delete(rawMsg, key)
		case "totalRecords":
			err = unpopulate(val, &q.TotalRecords)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourcesHistoryRequest.
func (r ResourcesHistoryRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "managementGroups", r.ManagementGroups)
	populate(objectMap, "options", r.Options)
	populate(objectMap, "query", r.Query)
	populate(objectMap, "subscriptions", r.Subscriptions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Table.
func (t Table) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "columns", t.Columns)
	populate(objectMap, "rows", t.Rows)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}