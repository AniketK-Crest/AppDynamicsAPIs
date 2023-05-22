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
Application Principal Management Service

Handles all administrative requests to manage application identities, including both Agents and Service principals.

API version: 1.0
Contact: info@appdynamics.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package applicationprincipalmanagement

import (
	"encoding/json"
)

// CredentialsDto Data object for a client's credentials.
type CredentialsDto struct {
	// The unique client identifier.
	ClientId *string `json:"clientId,omitempty"`
	// The client's secret, used to authenticate during an oAuth token request.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The RFC3339 timestamp when the rotated client secret will expire.
	RotatedSecretExpiresAt *string `json:"rotatedSecretExpiresAt,omitempty"`
}

// NewCredentialsDto instantiates a new CredentialsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsDto() *CredentialsDto {
	this := CredentialsDto{}
	return &this
}

// NewCredentialsDtoWithDefaults instantiates a new CredentialsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsDtoWithDefaults() *CredentialsDto {
	this := CredentialsDto{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CredentialsDto) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsDto) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CredentialsDto) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CredentialsDto) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *CredentialsDto) GetClientSecret() string {
	if o == nil || isNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsDto) GetClientSecretOk() (*string, bool) {
	if o == nil || isNil(o.ClientSecret) {
    return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *CredentialsDto) HasClientSecret() bool {
	if o != nil && !isNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *CredentialsDto) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetRotatedSecretExpiresAt returns the RotatedSecretExpiresAt field value if set, zero value otherwise.
func (o *CredentialsDto) GetRotatedSecretExpiresAt() string {
	if o == nil || isNil(o.RotatedSecretExpiresAt) {
		var ret string
		return ret
	}
	return *o.RotatedSecretExpiresAt
}

// GetRotatedSecretExpiresAtOk returns a tuple with the RotatedSecretExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsDto) GetRotatedSecretExpiresAtOk() (*string, bool) {
	if o == nil || isNil(o.RotatedSecretExpiresAt) {
    return nil, false
	}
	return o.RotatedSecretExpiresAt, true
}

// HasRotatedSecretExpiresAt returns a boolean if a field has been set.
func (o *CredentialsDto) HasRotatedSecretExpiresAt() bool {
	if o != nil && !isNil(o.RotatedSecretExpiresAt) {
		return true
	}

	return false
}

// SetRotatedSecretExpiresAt gets a reference to the given string and assigns it to the RotatedSecretExpiresAt field.
func (o *CredentialsDto) SetRotatedSecretExpiresAt(v string) {
	o.RotatedSecretExpiresAt = &v
}

func (o CredentialsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !isNil(o.RotatedSecretExpiresAt) {
		toSerialize["rotatedSecretExpiresAt"] = o.RotatedSecretExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialsDto struct {
	value *CredentialsDto
	isSet bool
}

func (v NullableCredentialsDto) Get() *CredentialsDto {
	return v.value
}

func (v *NullableCredentialsDto) Set(val *CredentialsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsDto(val *CredentialsDto) *NullableCredentialsDto {
	return &NullableCredentialsDto{value: val, isSet: true}
}

func (v NullableCredentialsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


