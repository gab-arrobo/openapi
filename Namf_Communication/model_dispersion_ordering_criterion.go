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

// DispersionOrderingCriterion Represents the order criterion for the list of dispersion.   Possible values are: - TIME_SLOT_START: Indicates the order of time slot start. - DISPERSION: Indicates the order of data/transaction dispersion. - CLASSIFICATION: Indicates the order of data/transaction classification. - RANKING: Indicates the order of data/transaction ranking. - PERCENTILE_RANKING: Indicates the order of data/transaction percentile ranking.
type DispersionOrderingCriterion struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DispersionOrderingCriterion) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DispersionOrderingCriterion)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DispersionOrderingCriterion) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDispersionOrderingCriterion struct {
	value *DispersionOrderingCriterion
	isSet bool
}

func (v NullableDispersionOrderingCriterion) Get() *DispersionOrderingCriterion {
	return v.value
}

func (v *NullableDispersionOrderingCriterion) Set(val *DispersionOrderingCriterion) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionOrderingCriterion) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionOrderingCriterion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionOrderingCriterion(val *DispersionOrderingCriterion) *NullableDispersionOrderingCriterion {
	return &NullableDispersionOrderingCriterion{value: val, isSet: true}
}

func (v NullableDispersionOrderingCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionOrderingCriterion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}