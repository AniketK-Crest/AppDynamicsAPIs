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
	"time"
)

// ConfigurationDetail struct for ConfigurationDetail
type ConfigurationDetail struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// Name of the connection or configuration
	DisplayName string `json:"displayName"`
	// Description for this connection or configuration
	Description *string                   `json:"description,omitempty"`
	Type        ProviderType              `json:"type"`
	Details     AzureConfigurationDetails `json:"details"`
}

// NewConfigurationDetail instantiates a new ConfigurationDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationDetail(id string, createdAt time.Time, updatedAt time.Time, displayName string, type_ ProviderType, details AzureConfigurationDetails) *ConfigurationDetail {
	this := ConfigurationDetail{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.DisplayName = displayName
	this.Type = type_
	this.Details = details
	return &this
}

// NewConfigurationDetailWithDefaults instantiates a new ConfigurationDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationDetailWithDefaults() *ConfigurationDetail {
	this := ConfigurationDetail{}
	return &this
}

// GetId returns the Id field value
func (o *ConfigurationDetail) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDetail) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConfigurationDetail) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ConfigurationDetail) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDetail) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ConfigurationDetail) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ConfigurationDetail) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDetail) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ConfigurationDetail) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDisplayName returns the DisplayName field value
func (o *ConfigurationDetail) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDetail) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ConfigurationDetail) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigurationDetail) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigurationDetail) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigurationDetail) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *ConfigurationDetail) GetType() ProviderType {
	if o == nil {
		var ret ProviderType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDetail) GetTypeOk() (*ProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConfigurationDetail) SetType(v ProviderType) {
	o.Type = v
}

// GetDetails returns the Details field value
func (o *ConfigurationDetail) GetDetails() AzureConfigurationDetails {
	if o == nil {
		var ret AzureConfigurationDetails
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDetail) GetDetailsOk() (*AzureConfigurationDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *ConfigurationDetail) SetDetails(v AzureConfigurationDetails) {
	o.Details = v
}

func (o ConfigurationDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationDetail struct {
	value *ConfigurationDetail
	isSet bool
}

func (v NullableConfigurationDetail) Get() *ConfigurationDetail {
	return v.value
}

func (v *NullableConfigurationDetail) Set(val *ConfigurationDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationDetail(val *ConfigurationDetail) *NullableConfigurationDetail {
	return &NullableConfigurationDetail{value: val, isSet: true}
}

func (v NullableConfigurationDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
