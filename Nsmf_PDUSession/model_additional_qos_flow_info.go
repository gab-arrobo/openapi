// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
Nsmf_PDUSession

SMF PDU Session Service.
© 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).
All rights reserved.

API version: 1.3.0-alpha.7

Authors: Aether SD-Core team
Contact: dev@lists.aetherproject.org

Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
*/

package Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// AdditionalQosFlowInfo struct for AdditionalQosFlowInfo
type AdditionalQosFlowInfo struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AdditionalQosFlowInfo) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AdditionalQosFlowInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AdditionalQosFlowInfo) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAdditionalQosFlowInfo struct {
	value *AdditionalQosFlowInfo
	isSet bool
}

func (v NullableAdditionalQosFlowInfo) Get() *AdditionalQosFlowInfo {
	return v.value
}

func (v *NullableAdditionalQosFlowInfo) Set(val *AdditionalQosFlowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalQosFlowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalQosFlowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalQosFlowInfo(val *AdditionalQosFlowInfo) *NullableAdditionalQosFlowInfo {
	return &NullableAdditionalQosFlowInfo{value: val, isSet: true}
}

func (v NullableAdditionalQosFlowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalQosFlowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}