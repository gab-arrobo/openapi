// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
Namf_Communication

AMF Communication Service.
© 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).
All rights reserved.

API version: 1.3.0-alpha.6

Authors: Aether SD-Core team
Contact: dev@lists.aetherproject.org

Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
*/

package Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// RedTransExpOrderingCriterion Represents the order criterion for the list of Redundant Transmission Experience.   Possible values are: - TIME_SLOT_START: Indicates the order of time slot start. - RED_TRANS_EXP: Indicates the order of Redundant Transmission Experience.
type RedTransExpOrderingCriterion struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RedTransExpOrderingCriterion) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(RedTransExpOrderingCriterion)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RedTransExpOrderingCriterion) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRedTransExpOrderingCriterion struct {
	value *RedTransExpOrderingCriterion
	isSet bool
}

func (v NullableRedTransExpOrderingCriterion) Get() *RedTransExpOrderingCriterion {
	return v.value
}

func (v *NullableRedTransExpOrderingCriterion) Set(val *RedTransExpOrderingCriterion) {
	v.value = val
	v.isSet = true
}

func (v NullableRedTransExpOrderingCriterion) IsSet() bool {
	return v.isSet
}

func (v *NullableRedTransExpOrderingCriterion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedTransExpOrderingCriterion(val *RedTransExpOrderingCriterion) *NullableRedTransExpOrderingCriterion {
	return &NullableRedTransExpOrderingCriterion{value: val, isSet: true}
}

func (v NullableRedTransExpOrderingCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedTransExpOrderingCriterion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}