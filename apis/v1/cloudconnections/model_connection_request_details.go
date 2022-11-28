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
	"fmt"
)

// ConnectionRequestDetails - struct for ConnectionRequestDetails
type ConnectionRequestDetails struct {
	AWSConnectionRequestDetails *AWSConnectionRequestDetails
	AzureDetails *AzureDetails
}

// AWSConnectionRequestDetailsAsConnectionRequestDetails is a convenience function that returns AWSConnectionRequestDetails wrapped in ConnectionRequestDetails
func AWSConnectionRequestDetailsAsConnectionRequestDetails(v *AWSConnectionRequestDetails) ConnectionRequestDetails {
	return ConnectionRequestDetails{
		AWSConnectionRequestDetails: v,
	}
}

// AzureDetailsAsConnectionRequestDetails is a convenience function that returns AzureDetails wrapped in ConnectionRequestDetails
func AzureDetailsAsConnectionRequestDetails(v *AzureDetails) ConnectionRequestDetails {
	return ConnectionRequestDetails{
		AzureDetails: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConnectionRequestDetails) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSConnectionRequestDetails
	err = newStrictDecoder(data).Decode(&dst.AWSConnectionRequestDetails)
	if err == nil {
		jsonAWSConnectionRequestDetails, _ := json.Marshal(dst.AWSConnectionRequestDetails)
		if string(jsonAWSConnectionRequestDetails) == "{}" { // empty struct
			dst.AWSConnectionRequestDetails = nil
		} else {
			match++
		}
	} else {
		dst.AWSConnectionRequestDetails = nil
	}

	// try to unmarshal data into AzureDetails
	err = newStrictDecoder(data).Decode(&dst.AzureDetails)
	if err == nil {
		jsonAzureDetails, _ := json.Marshal(dst.AzureDetails)
		if string(jsonAzureDetails) == "{}" { // empty struct
			dst.AzureDetails = nil
		} else {
			match++
		}
	} else {
		dst.AzureDetails = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AWSConnectionRequestDetails = nil
		dst.AzureDetails = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ConnectionRequestDetails)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ConnectionRequestDetails)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConnectionRequestDetails) MarshalJSON() ([]byte, error) {
	if src.AWSConnectionRequestDetails != nil {
		return json.Marshal(&src.AWSConnectionRequestDetails)
	}

	if src.AzureDetails != nil {
		return json.Marshal(&src.AzureDetails)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ConnectionRequestDetails) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AWSConnectionRequestDetails != nil {
		return obj.AWSConnectionRequestDetails
	}

	if obj.AzureDetails != nil {
		return obj.AzureDetails
	}

	// all schemas are nil
	return nil
}

type NullableConnectionRequestDetails struct {
	value *ConnectionRequestDetails
	isSet bool
}

func (v NullableConnectionRequestDetails) Get() *ConnectionRequestDetails {
	return v.value
}

func (v *NullableConnectionRequestDetails) Set(val *ConnectionRequestDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRequestDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRequestDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRequestDetails(val *ConnectionRequestDetails) *NullableConnectionRequestDetails {
	return &NullableConnectionRequestDetails{value: val, isSet: true}
}

func (v NullableConnectionRequestDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRequestDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

