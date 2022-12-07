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

// ModelResultItem A description of a simple data type of a single fetched expression in the result.
type ModelResultItem struct {
	Alias *string `json:"alias,omitempty"`
	// Specifies a type of the value in the JSON response.
	Type *string `json:"type,omitempty"`
	Hints *Hints `json:"hints,omitempty"`
}

// NewModelResultItem instantiates a new ModelResultItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelResultItem() *ModelResultItem {
	this := ModelResultItem{}
	return &this
}

// NewModelResultItemWithDefaults instantiates a new ModelResultItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelResultItemWithDefaults() *ModelResultItem {
	this := ModelResultItem{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ModelResultItem) GetAlias() string {
	if o == nil || isNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResultItem) GetAliasOk() (*string, bool) {
	if o == nil || isNil(o.Alias) {
    return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ModelResultItem) HasAlias() bool {
	if o != nil && !isNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ModelResultItem) SetAlias(v string) {
	o.Alias = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelResultItem) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResultItem) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelResultItem) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ModelResultItem) SetType(v string) {
	o.Type = &v
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *ModelResultItem) GetHints() Hints {
	if o == nil || isNil(o.Hints) {
		var ret Hints
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResultItem) GetHintsOk() (*Hints, bool) {
	if o == nil || isNil(o.Hints) {
    return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *ModelResultItem) HasHints() bool {
	if o != nil && !isNil(o.Hints) {
		return true
	}

	return false
}

// SetHints gets a reference to the given Hints and assigns it to the Hints field.
func (o *ModelResultItem) SetHints(v Hints) {
	o.Hints = &v
}

func (o ModelResultItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Hints) {
		toSerialize["hints"] = o.Hints
	}
	return json.Marshal(toSerialize)
}

type NullableModelResultItem struct {
	value *ModelResultItem
	isSet bool
}

func (v NullableModelResultItem) Get() *ModelResultItem {
	return v.value
}

func (v *NullableModelResultItem) Set(val *ModelResultItem) {
	v.value = val
	v.isSet = true
}

func (v NullableModelResultItem) IsSet() bool {
	return v.isSet
}

func (v *NullableModelResultItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelResultItem(val *ModelResultItem) *NullableModelResultItem {
	return &NullableModelResultItem{value: val, isSet: true}
}

func (v NullableModelResultItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelResultItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


