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

// ConnectionUpdateDetailsOneOf1 struct for ConnectionUpdateDetailsOneOf1
type ConnectionUpdateDetailsOneOf1 struct {
	AccessKeyId string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
}

// NewConnectionUpdateDetailsOneOf1 instantiates a new ConnectionUpdateDetailsOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionUpdateDetailsOneOf1(accessKeyId string, secretAccessKey string) *ConnectionUpdateDetailsOneOf1 {
	this := ConnectionUpdateDetailsOneOf1{}
	this.AccessKeyId = accessKeyId
	this.SecretAccessKey = secretAccessKey
	return &this
}

// NewConnectionUpdateDetailsOneOf1WithDefaults instantiates a new ConnectionUpdateDetailsOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionUpdateDetailsOneOf1WithDefaults() *ConnectionUpdateDetailsOneOf1 {
	this := ConnectionUpdateDetailsOneOf1{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value
func (o *ConnectionUpdateDetailsOneOf1) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *ConnectionUpdateDetailsOneOf1) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value
func (o *ConnectionUpdateDetailsOneOf1) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetSecretAccessKey returns the SecretAccessKey field value
func (o *ConnectionUpdateDetailsOneOf1) GetSecretAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *ConnectionUpdateDetailsOneOf1) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SecretAccessKey, true
}

// SetSecretAccessKey sets field value
func (o *ConnectionUpdateDetailsOneOf1) SetSecretAccessKey(v string) {
	o.SecretAccessKey = v
}

func (o ConnectionUpdateDetailsOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if true {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionUpdateDetailsOneOf1 struct {
	value *ConnectionUpdateDetailsOneOf1
	isSet bool
}

func (v NullableConnectionUpdateDetailsOneOf1) Get() *ConnectionUpdateDetailsOneOf1 {
	return v.value
}

func (v *NullableConnectionUpdateDetailsOneOf1) Set(val *ConnectionUpdateDetailsOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionUpdateDetailsOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionUpdateDetailsOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionUpdateDetailsOneOf1(val *ConnectionUpdateDetailsOneOf1) *NullableConnectionUpdateDetailsOneOf1 {
	return &NullableConnectionUpdateDetailsOneOf1{value: val, isSet: true}
}

func (v NullableConnectionUpdateDetailsOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionUpdateDetailsOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

