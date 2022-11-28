/*
AppDynamics Cloud Query Service API

An API providing access to observation data using an AppDynamics domain-specific language.  The language is declarative, read-only (it does not allow for data modification) and specific to the AppD MELT model. It presents MELT data (metrics, events, logs, traces) and State in the scope of monitored topology.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudquery

import (
	"encoding/json"
	"fmt"
)

// DataResultItemInner struct for DataResultItemInner
type DataResultItemInner struct {
	DatasetReference *DatasetReference
	bool             *bool
	float32          *float32
	int32            *int32
	mapStrInterface  *map[string]interface{}
	string           *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DataResultItemInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DatasetReference
	err = json.Unmarshal(data, &dst.DatasetReference)
	if err == nil {
		jsonDatasetReference, _ := json.Marshal(dst.DatasetReference)
		if string(jsonDatasetReference) == "{}" { // empty struct
			dst.DatasetReference = nil
		} else {
			return nil // data stored in dst.DatasetReference, return on the first match
		}
	} else {
		dst.DatasetReference = nil
	}

	// try to unmarshal JSON data into bool
	err = json.Unmarshal(data, &dst.bool)
	if err == nil {
		jsonbool, _ := json.Marshal(dst.bool)
		if string(jsonbool) == "{}" { // empty struct
			dst.bool = nil
		} else {
			return nil // data stored in dst.bool, return on the first match
		}
	} else {
		dst.bool = nil
	}

	// try to unmarshal JSON data into float32
	err = json.Unmarshal(data, &dst.float32)
	if err == nil {
		jsonfloat32, _ := json.Marshal(dst.float32)
		if string(jsonfloat32) == "{}" { // empty struct
			dst.float32 = nil
		} else {
			return nil // data stored in dst.float32, return on the first match
		}
	} else {
		dst.float32 = nil
	}

	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.int32)
	if err == nil {
		jsonint32, _ := json.Marshal(dst.int32)
		if string(jsonint32) == "{}" { // empty struct
			dst.int32 = nil
		} else {
			return nil // data stored in dst.int32, return on the first match
		}
	} else {
		dst.int32 = nil
	}

	// try to unmarshal JSON data into mapStrInterface
	err = json.Unmarshal(data, &dst.mapStrInterface)
	if err == nil {
		jsonmapStrInterface, _ := json.Marshal(dst.mapStrInterface)
		if string(jsonmapStrInterface) == "{}" { // empty struct
			dst.mapStrInterface = nil
		} else {
			return nil // data stored in dst.mapStrInterface, return on the first match
		}
	} else {
		dst.mapStrInterface = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DataResultItemInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DataResultItemInner) MarshalJSON() ([]byte, error) {
	if src.DatasetReference != nil {
		return json.Marshal(&src.DatasetReference)
	}

	if src.bool != nil {
		return json.Marshal(&src.bool)
	}

	if src.float32 != nil {
		return json.Marshal(&src.float32)
	}

	if src.int32 != nil {
		return json.Marshal(&src.int32)
	}

	if src.mapStrInterface != nil {
		return json.Marshal(&src.mapStrInterface)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDataResultItemInner struct {
	value *DataResultItemInner
	isSet bool
}

func (v NullableDataResultItemInner) Get() *DataResultItemInner {
	return v.value
}

func (v *NullableDataResultItemInner) Set(val *DataResultItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDataResultItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDataResultItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataResultItemInner(val *DataResultItemInner) *NullableDataResultItemInner {
	return &NullableDataResultItemInner{value: val, isSet: true}
}

func (v NullableDataResultItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataResultItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
