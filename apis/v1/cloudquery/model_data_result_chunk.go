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
AppDynamics Cloud Query Service API

An API providing access to observation data using an AppDynamics domain-specific language.  The language is declarative, read-only (it does not allow for data modification) and specific to the AppD MELT model. It presents MELT data (metrics, events, logs, traces) and State in the scope of monitored topology.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudquery

import (
	"encoding/json"
)

// DataResultChunk A part of the result data. There may be multiple chunks of this type in the response.
type DataResultChunk struct {
	// Attribute identifying a response chunk holding data.
	Type *string `json:"type,omitempty"`
	Links *PaginationReference `json:"_links,omitempty"`
	// Name of the dataset. May be used as a reference in other datasets when returning multi dimensional data.
	Dataset *string `json:"dataset,omitempty"`
	Model *ModelReference `json:"model,omitempty"`
	Metadata *MetadataResultItem `json:"metadata,omitempty"`
	Data [][]DataResultItemInner `json:"data,omitempty"`
}

// NewDataResultChunk instantiates a new DataResultChunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataResultChunk() *DataResultChunk {
	this := DataResultChunk{}
	return &this
}

// NewDataResultChunkWithDefaults instantiates a new DataResultChunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataResultChunkWithDefaults() *DataResultChunk {
	this := DataResultChunk{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DataResultChunk) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataResultChunk) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DataResultChunk) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DataResultChunk) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DataResultChunk) GetLinks() PaginationReference {
	if o == nil || isNil(o.Links) {
		var ret PaginationReference
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataResultChunk) GetLinksOk() (*PaginationReference, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DataResultChunk) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationReference and assigns it to the Links field.
func (o *DataResultChunk) SetLinks(v PaginationReference) {
	o.Links = &v
}

// GetDataset returns the Dataset field value if set, zero value otherwise.
func (o *DataResultChunk) GetDataset() string {
	if o == nil || isNil(o.Dataset) {
		var ret string
		return ret
	}
	return *o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataResultChunk) GetDatasetOk() (*string, bool) {
	if o == nil || isNil(o.Dataset) {
    return nil, false
	}
	return o.Dataset, true
}

// HasDataset returns a boolean if a field has been set.
func (o *DataResultChunk) HasDataset() bool {
	if o != nil && !isNil(o.Dataset) {
		return true
	}

	return false
}

// SetDataset gets a reference to the given string and assigns it to the Dataset field.
func (o *DataResultChunk) SetDataset(v string) {
	o.Dataset = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *DataResultChunk) GetModel() ModelReference {
	if o == nil || isNil(o.Model) {
		var ret ModelReference
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataResultChunk) GetModelOk() (*ModelReference, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *DataResultChunk) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given ModelReference and assigns it to the Model field.
func (o *DataResultChunk) SetModel(v ModelReference) {
	o.Model = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DataResultChunk) GetMetadata() MetadataResultItem {
	if o == nil || isNil(o.Metadata) {
		var ret MetadataResultItem
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataResultChunk) GetMetadataOk() (*MetadataResultItem, bool) {
	if o == nil || isNil(o.Metadata) {
    return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DataResultChunk) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given MetadataResultItem and assigns it to the Metadata field.
func (o *DataResultChunk) SetMetadata(v MetadataResultItem) {
	o.Metadata = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DataResultChunk) GetData() [][]DataResultItemInner {
	if o == nil || isNil(o.Data) {
		var ret [][]DataResultItemInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataResultChunk) GetDataOk() ([][]DataResultItemInner, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DataResultChunk) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given [][]DataResultItemInner and assigns it to the Data field.
func (o *DataResultChunk) SetData(v [][]DataResultItemInner) {
	o.Data = v
}

func (o DataResultChunk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !isNil(o.Dataset) {
		toSerialize["dataset"] = o.Dataset
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDataResultChunk struct {
	value *DataResultChunk
	isSet bool
}

func (v NullableDataResultChunk) Get() *DataResultChunk {
	return v.value
}

func (v *NullableDataResultChunk) Set(val *DataResultChunk) {
	v.value = val
	v.isSet = true
}

func (v NullableDataResultChunk) IsSet() bool {
	return v.isSet
}

func (v *NullableDataResultChunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataResultChunk(val *DataResultChunk) *NullableDataResultChunk {
	return &NullableDataResultChunk{value: val, isSet: true}
}

func (v NullableDataResultChunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataResultChunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


