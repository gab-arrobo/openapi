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

// SharedDataTreatmentInstruction Indicates the presence of this attribute in the individual data. Otherwise, the individual data takes precedence, by default.
type SharedDataTreatmentInstruction string

// List of SharedDataTreatmentInstruction
const (
	SHAREDDATATREATMENTINSTRUCTION_USE_IF_NO_CLASH SharedDataTreatmentInstruction = "USE_IF_NO_CLASH"
	SHAREDDATATREATMENTINSTRUCTION_OVERWRITE       SharedDataTreatmentInstruction = "OVERWRITE"
	SHAREDDATATREATMENTINSTRUCTION_MAX             SharedDataTreatmentInstruction = "MAX"
	SHAREDDATATREATMENTINSTRUCTION_MIN             SharedDataTreatmentInstruction = "MIN"
)

// All allowed values of SharedDataTreatmentInstruction enum
var AllowedSharedDataTreatmentInstructionEnumValues = []SharedDataTreatmentInstruction{
	"USE_IF_NO_CLASH",
	"OVERWRITE",
	"MAX",
	"MIN",
}

func (v *SharedDataTreatmentInstruction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SharedDataTreatmentInstruction(value)
	for _, existing := range AllowedSharedDataTreatmentInstructionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SharedDataTreatmentInstruction", value)
}

// NewSharedDataTreatmentInstructionFromValue returns a pointer to a valid SharedDataTreatmentInstruction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedDataTreatmentInstructionFromValue(v string) (*SharedDataTreatmentInstruction, error) {
	ev := SharedDataTreatmentInstruction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedDataTreatmentInstruction: valid values are %v", v, AllowedSharedDataTreatmentInstructionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedDataTreatmentInstruction) IsValid() bool {
	for _, existing := range AllowedSharedDataTreatmentInstructionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SharedDataTreatmentInstruction value
func (v SharedDataTreatmentInstruction) Ptr() *SharedDataTreatmentInstruction {
	return &v
}

type NullableSharedDataTreatmentInstruction struct {
	value *SharedDataTreatmentInstruction
	isSet bool
}

func (v NullableSharedDataTreatmentInstruction) Get() *SharedDataTreatmentInstruction {
	return v.value
}

func (v *NullableSharedDataTreatmentInstruction) Set(val *SharedDataTreatmentInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedDataTreatmentInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedDataTreatmentInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedDataTreatmentInstruction(val *SharedDataTreatmentInstruction) *NullableSharedDataTreatmentInstruction {
	return &NullableSharedDataTreatmentInstruction{value: val, isSet: true}
}

func (v NullableSharedDataTreatmentInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedDataTreatmentInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}