// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
Nudm_SDM

Nudm Subscriber Data Management Service.
© 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).
All rights reserved.

API version: 2.3.0-alpha.6

Authors: Aether SD-Core team
Contact: dev@lists.aetherproject.org

Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
*/

package Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// OdbPacketServices struct for OdbPacketServices
type OdbPacketServices struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *OdbPacketServices) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(OdbPacketServices)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *OdbPacketServices) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableOdbPacketServices struct {
	value *OdbPacketServices
	isSet bool
}

func (v NullableOdbPacketServices) Get() *OdbPacketServices {
	return v.value
}

func (v *NullableOdbPacketServices) Set(val *OdbPacketServices) {
	v.value = val
	v.isSet = true
}

func (v NullableOdbPacketServices) IsSet() bool {
	return v.isSet
}

func (v *NullableOdbPacketServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOdbPacketServices(val *OdbPacketServices) *NullableOdbPacketServices {
	return &NullableOdbPacketServices{value: val, isSet: true}
}

func (v NullableOdbPacketServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOdbPacketServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}