// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
NRF NFManagement Service

NRF NFManagement Service.
© 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).
All rights reserved.

API version: 1.3.0-alpha.7

Authors: Aether SD-Core team
Contact: dev@lists.aetherproject.org

Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
*/

package Nnrf_NFManagement

import (
	"encoding/json"
	"fmt"
)

// N32Purpose Usage purpose of establishing N32 connectivity
type N32Purpose struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N32Purpose) UnmarshalJSON(data []byte) error {
	var err error
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

	return fmt.Errorf("data failed to match schemas in anyOf(N32Purpose)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N32Purpose) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableN32Purpose struct {
	value *N32Purpose
	isSet bool
}

func (v NullableN32Purpose) Get() *N32Purpose {
	return v.value
}

func (v *NullableN32Purpose) Set(val *N32Purpose) {
	v.value = val
	v.isSet = true
}

func (v NullableN32Purpose) IsSet() bool {
	return v.isSet
}

func (v *NullableN32Purpose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN32Purpose(val *N32Purpose) *NullableN32Purpose {
	return &NullableN32Purpose{value: val, isSet: true}
}

func (v NullableN32Purpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN32Purpose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}