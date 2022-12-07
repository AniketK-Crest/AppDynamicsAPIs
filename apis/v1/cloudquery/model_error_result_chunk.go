/*
AppDynamics Cloud Query Service API

An API providing access to observation data using an AppDynamics domain-specific language.  The language is declarative, read-only (it does not allow for data modification) and specific to the AppD MELT model. It presents MELT data (metrics, events, logs, traces) and State in the scope of monitored topology.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudquery

import (
	"encoding/json"
)

// ErrorResultChunk A signal that an error occurred during query execution.
type ErrorResultChunk struct {
	// Attribute identifying a response chunk as holding information of an error.
	Type *string `json:"type,omitempty"`
	Error *ErrorResultChunkError `json:"error,omitempty"`
}

// NewErrorResultChunk instantiates a new ErrorResultChunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResultChunk() *ErrorResultChunk {
	this := ErrorResultChunk{}
	return &this
}

// NewErrorResultChunkWithDefaults instantiates a new ErrorResultChunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResultChunkWithDefaults() *ErrorResultChunk {
	this := ErrorResultChunk{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErrorResultChunk) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResultChunk) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrorResultChunk) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ErrorResultChunk) SetType(v string) {
	o.Type = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ErrorResultChunk) GetError() ErrorResultChunkError {
	if o == nil || isNil(o.Error) {
		var ret ErrorResultChunkError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResultChunk) GetErrorOk() (*ErrorResultChunkError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ErrorResultChunk) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ErrorResultChunkError and assigns it to the Error field.
func (o *ErrorResultChunk) SetError(v ErrorResultChunkError) {
	o.Error = &v
}

func (o ErrorResultChunk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableErrorResultChunk struct {
	value *ErrorResultChunk
	isSet bool
}

func (v NullableErrorResultChunk) Get() *ErrorResultChunk {
	return v.value
}

func (v *NullableErrorResultChunk) Set(val *ErrorResultChunk) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResultChunk) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResultChunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResultChunk(val *ErrorResultChunk) *NullableErrorResultChunk {
	return &NullableErrorResultChunk{value: val, isSet: true}
}

func (v NullableErrorResultChunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResultChunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


