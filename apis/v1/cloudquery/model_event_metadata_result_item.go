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

// EventMetadataResultItem Metadata for an event data chunk.
type EventMetadataResultItem struct {
	// Arbitrary map of keys and values associated with the event data.
	Schema map[string]interface{} `json:"schema,omitempty"`
	// Arbitrary map of statistics associated with the event data.
	Statistics map[string]interface{} `json:"statistics,omitempty"`
}

// NewEventMetadataResultItem instantiates a new EventMetadataResultItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventMetadataResultItem() *EventMetadataResultItem {
	this := EventMetadataResultItem{}
	return &this
}

// NewEventMetadataResultItemWithDefaults instantiates a new EventMetadataResultItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventMetadataResultItemWithDefaults() *EventMetadataResultItem {
	this := EventMetadataResultItem{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *EventMetadataResultItem) GetSchema() map[string]interface{} {
	if o == nil || isNil(o.Schema) {
		var ret map[string]interface{}
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadataResultItem) GetSchemaOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Schema) {
    return map[string]interface{}{}, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *EventMetadataResultItem) HasSchema() bool {
	if o != nil && !isNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given map[string]interface{} and assigns it to the Schema field.
func (o *EventMetadataResultItem) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *EventMetadataResultItem) GetStatistics() map[string]interface{} {
	if o == nil || isNil(o.Statistics) {
		var ret map[string]interface{}
		return ret
	}
	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadataResultItem) GetStatisticsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Statistics) {
    return map[string]interface{}{}, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *EventMetadataResultItem) HasStatistics() bool {
	if o != nil && !isNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given map[string]interface{} and assigns it to the Statistics field.
func (o *EventMetadataResultItem) SetStatistics(v map[string]interface{}) {
	o.Statistics = v
}

func (o EventMetadataResultItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !isNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	return json.Marshal(toSerialize)
}

type NullableEventMetadataResultItem struct {
	value *EventMetadataResultItem
	isSet bool
}

func (v NullableEventMetadataResultItem) Get() *EventMetadataResultItem {
	return v.value
}

func (v *NullableEventMetadataResultItem) Set(val *EventMetadataResultItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEventMetadataResultItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEventMetadataResultItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventMetadataResultItem(val *EventMetadataResultItem) *NullableEventMetadataResultItem {
	return &NullableEventMetadataResultItem{value: val, isSet: true}
}

func (v NullableEventMetadataResultItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventMetadataResultItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


