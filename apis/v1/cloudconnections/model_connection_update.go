// Copyright 2023 Cisco Systems, Inc.
// 
// Licensed under the MPL License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     https://www.mozilla.org/en-US/MPL/2.0/
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/*
AppDynamics Cloud Connections API

Enables you to manage cloud connections

API version: 1.0.0
Contact: support@appdynamics.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudconnections

import (
	"encoding/json"
)

// ConnectionUpdate struct for ConnectionUpdate
type ConnectionUpdate struct {
	// Updated name for the connection
	DisplayName *string `json:"displayName,omitempty"`
	Description NullableString `json:"description,omitempty"`
	// Configuration ID assigned to the connection
	ConfigurationId *string `json:"configurationId,omitempty"`
	// Set the state of the connection
	State *string `json:"state,omitempty"`
	Details *ConnectionUpdateDetails `json:"details,omitempty"`
}

// NewConnectionUpdate instantiates a new ConnectionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionUpdate() *ConnectionUpdate {
	this := ConnectionUpdate{}
	return &this
}

// NewConnectionUpdateWithDefaults instantiates a new ConnectionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionUpdateWithDefaults() *ConnectionUpdate {
	this := ConnectionUpdate{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ConnectionUpdate) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ConnectionUpdate) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ConnectionUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionUpdate) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectionUpdate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ConnectionUpdate) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ConnectionUpdate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ConnectionUpdate) UnsetDescription() {
	o.Description.Unset()
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *ConnectionUpdate) GetConfigurationId() string {
	if o == nil || isNil(o.ConfigurationId) {
		var ret string
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetConfigurationIdOk() (*string, bool) {
	if o == nil || isNil(o.ConfigurationId) {
    return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *ConnectionUpdate) HasConfigurationId() bool {
	if o != nil && !isNil(o.ConfigurationId) {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given string and assigns it to the ConfigurationId field.
func (o *ConnectionUpdate) SetConfigurationId(v string) {
	o.ConfigurationId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ConnectionUpdate) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ConnectionUpdate) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ConnectionUpdate) SetState(v string) {
	o.State = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ConnectionUpdate) GetDetails() ConnectionUpdateDetails {
	if o == nil || isNil(o.Details) {
		var ret ConnectionUpdateDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetDetailsOk() (*ConnectionUpdateDetails, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ConnectionUpdate) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given ConnectionUpdateDetails and assigns it to the Details field.
func (o *ConnectionUpdate) SetDetails(v ConnectionUpdateDetails) {
	o.Details = &v
}

func (o ConnectionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.ConfigurationId) {
		toSerialize["configurationId"] = o.ConfigurationId
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionUpdate struct {
	value *ConnectionUpdate
	isSet bool
}

func (v NullableConnectionUpdate) Get() *ConnectionUpdate {
	return v.value
}

func (v *NullableConnectionUpdate) Set(val *ConnectionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionUpdate(val *ConnectionUpdate) *NullableConnectionUpdate {
	return &NullableConnectionUpdate{value: val, isSet: true}
}

func (v NullableConnectionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


