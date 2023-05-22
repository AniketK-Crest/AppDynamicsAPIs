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
	"fmt"
)

// ExecuteQuery200Response struct for ExecuteQuery200Response
type ExecuteQuery200Response struct {
	DataResultChunk *DataResultChunk
	ErrorResultChunk *ErrorResultChunk
	HeartbeatResultChunk *HeartbeatResultChunk
	ModelResultChunk *ModelResultChunk
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ExecuteQuery200Response) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'DataResultChunk'
	if jsonDict["type"] == "DataResultChunk" {
		// try to unmarshal JSON data into DataResultChunk
		err = json.Unmarshal(data, &dst.DataResultChunk);
		if err == nil {
			jsonDataResultChunk, _ := json.Marshal(dst.DataResultChunk)
			if string(jsonDataResultChunk) == "{}" { // empty struct
				dst.DataResultChunk = nil
			} else {
				return nil // data stored in dst.DataResultChunk, return on the first match
			}
		} else {
			dst.DataResultChunk = nil
		}
	}

	// check if the discriminator value is 'ErrorResultChunk'
	if jsonDict["type"] == "ErrorResultChunk" {
		// try to unmarshal JSON data into ErrorResultChunk
		err = json.Unmarshal(data, &dst.ErrorResultChunk);
		if err == nil {
			jsonErrorResultChunk, _ := json.Marshal(dst.ErrorResultChunk)
			if string(jsonErrorResultChunk) == "{}" { // empty struct
				dst.ErrorResultChunk = nil
			} else {
				return nil // data stored in dst.ErrorResultChunk, return on the first match
			}
		} else {
			dst.ErrorResultChunk = nil
		}
	}

	// check if the discriminator value is 'HeartbeatResultChunk'
	if jsonDict["type"] == "HeartbeatResultChunk" {
		// try to unmarshal JSON data into HeartbeatResultChunk
		err = json.Unmarshal(data, &dst.HeartbeatResultChunk);
		if err == nil {
			jsonHeartbeatResultChunk, _ := json.Marshal(dst.HeartbeatResultChunk)
			if string(jsonHeartbeatResultChunk) == "{}" { // empty struct
				dst.HeartbeatResultChunk = nil
			} else {
				return nil // data stored in dst.HeartbeatResultChunk, return on the first match
			}
		} else {
			dst.HeartbeatResultChunk = nil
		}
	}

	// check if the discriminator value is 'ModelResultChunk'
	if jsonDict["type"] == "ModelResultChunk" {
		// try to unmarshal JSON data into ModelResultChunk
		err = json.Unmarshal(data, &dst.ModelResultChunk);
		if err == nil {
			jsonModelResultChunk, _ := json.Marshal(dst.ModelResultChunk)
			if string(jsonModelResultChunk) == "{}" { // empty struct
				dst.ModelResultChunk = nil
			} else {
				return nil // data stored in dst.ModelResultChunk, return on the first match
			}
		} else {
			dst.ModelResultChunk = nil
		}
	}

	// check if the discriminator value is 'data'
	if jsonDict["type"] == "data" {
		// try to unmarshal JSON data into DataResultChunk
		err = json.Unmarshal(data, &dst.DataResultChunk);
		if err == nil {
			jsonDataResultChunk, _ := json.Marshal(dst.DataResultChunk)
			if string(jsonDataResultChunk) == "{}" { // empty struct
				dst.DataResultChunk = nil
			} else {
				return nil // data stored in dst.DataResultChunk, return on the first match
			}
		} else {
			dst.DataResultChunk = nil
		}
	}

	// check if the discriminator value is 'error'
	if jsonDict["type"] == "error" {
		// try to unmarshal JSON data into ErrorResultChunk
		err = json.Unmarshal(data, &dst.ErrorResultChunk);
		if err == nil {
			jsonErrorResultChunk, _ := json.Marshal(dst.ErrorResultChunk)
			if string(jsonErrorResultChunk) == "{}" { // empty struct
				dst.ErrorResultChunk = nil
			} else {
				return nil // data stored in dst.ErrorResultChunk, return on the first match
			}
		} else {
			dst.ErrorResultChunk = nil
		}
	}

	// check if the discriminator value is 'heartbeat'
	if jsonDict["type"] == "heartbeat" {
		// try to unmarshal JSON data into HeartbeatResultChunk
		err = json.Unmarshal(data, &dst.HeartbeatResultChunk);
		if err == nil {
			jsonHeartbeatResultChunk, _ := json.Marshal(dst.HeartbeatResultChunk)
			if string(jsonHeartbeatResultChunk) == "{}" { // empty struct
				dst.HeartbeatResultChunk = nil
			} else {
				return nil // data stored in dst.HeartbeatResultChunk, return on the first match
			}
		} else {
			dst.HeartbeatResultChunk = nil
		}
	}

	// check if the discriminator value is 'model'
	if jsonDict["type"] == "model" {
		// try to unmarshal JSON data into ModelResultChunk
		err = json.Unmarshal(data, &dst.ModelResultChunk);
		if err == nil {
			jsonModelResultChunk, _ := json.Marshal(dst.ModelResultChunk)
			if string(jsonModelResultChunk) == "{}" { // empty struct
				dst.ModelResultChunk = nil
			} else {
				return nil // data stored in dst.ModelResultChunk, return on the first match
			}
		} else {
			dst.ModelResultChunk = nil
		}
	}

	// try to unmarshal JSON data into DataResultChunk
	err = json.Unmarshal(data, &dst.DataResultChunk);
	if err == nil {
		jsonDataResultChunk, _ := json.Marshal(dst.DataResultChunk)
		if string(jsonDataResultChunk) == "{}" { // empty struct
			dst.DataResultChunk = nil
		} else {
			return nil // data stored in dst.DataResultChunk, return on the first match
		}
	} else {
		dst.DataResultChunk = nil
	}

	// try to unmarshal JSON data into ErrorResultChunk
	err = json.Unmarshal(data, &dst.ErrorResultChunk);
	if err == nil {
		jsonErrorResultChunk, _ := json.Marshal(dst.ErrorResultChunk)
		if string(jsonErrorResultChunk) == "{}" { // empty struct
			dst.ErrorResultChunk = nil
		} else {
			return nil // data stored in dst.ErrorResultChunk, return on the first match
		}
	} else {
		dst.ErrorResultChunk = nil
	}

	// try to unmarshal JSON data into HeartbeatResultChunk
	err = json.Unmarshal(data, &dst.HeartbeatResultChunk);
	if err == nil {
		jsonHeartbeatResultChunk, _ := json.Marshal(dst.HeartbeatResultChunk)
		if string(jsonHeartbeatResultChunk) == "{}" { // empty struct
			dst.HeartbeatResultChunk = nil
		} else {
			return nil // data stored in dst.HeartbeatResultChunk, return on the first match
		}
	} else {
		dst.HeartbeatResultChunk = nil
	}

	// try to unmarshal JSON data into ModelResultChunk
	err = json.Unmarshal(data, &dst.ModelResultChunk);
	if err == nil {
		jsonModelResultChunk, _ := json.Marshal(dst.ModelResultChunk)
		if string(jsonModelResultChunk) == "{}" { // empty struct
			dst.ModelResultChunk = nil
		} else {
			return nil // data stored in dst.ModelResultChunk, return on the first match
		}
	} else {
		dst.ModelResultChunk = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ExecuteQuery200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ExecuteQuery200Response) MarshalJSON() ([]byte, error) {
	if src.DataResultChunk != nil {
		return json.Marshal(&src.DataResultChunk)
	}

	if src.ErrorResultChunk != nil {
		return json.Marshal(&src.ErrorResultChunk)
	}

	if src.HeartbeatResultChunk != nil {
		return json.Marshal(&src.HeartbeatResultChunk)
	}

	if src.ModelResultChunk != nil {
		return json.Marshal(&src.ModelResultChunk)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableExecuteQuery200Response struct {
	value *ExecuteQuery200Response
	isSet bool
}

func (v NullableExecuteQuery200Response) Get() *ExecuteQuery200Response {
	return v.value
}

func (v *NullableExecuteQuery200Response) Set(val *ExecuteQuery200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteQuery200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteQuery200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteQuery200Response(val *ExecuteQuery200Response) *NullableExecuteQuery200Response {
	return &NullableExecuteQuery200Response{value: val, isSet: true}
}

func (v NullableExecuteQuery200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteQuery200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


