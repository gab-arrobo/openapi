// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
Npcf_SMPolicyControl API

Session Management Policy Control Service
© 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).
All rights reserved.

API version: 1.3.0-alpha.6

Authors: Aether SD-Core team
Contact: dev@lists.aetherproject.org

Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
*/

package Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// SmPolicyAssociationReleaseCause Represents the cause due to which the PCF requests the termination of the SM policy association.
type SmPolicyAssociationReleaseCause string

// List of SmPolicyAssociationReleaseCause
const (
	SMPOLICYASSOCIATIONRELEASECAUSE_UNSPECIFIED                  SmPolicyAssociationReleaseCause = "UNSPECIFIED"
	SMPOLICYASSOCIATIONRELEASECAUSE_UE_SUBSCRIPTION              SmPolicyAssociationReleaseCause = "UE_SUBSCRIPTION"
	SMPOLICYASSOCIATIONRELEASECAUSE_INSUFFICIENT_RES             SmPolicyAssociationReleaseCause = "INSUFFICIENT_RES"
	SMPOLICYASSOCIATIONRELEASECAUSE_VALIDATION_CONDITION_NOT_MET SmPolicyAssociationReleaseCause = "VALIDATION_CONDITION_NOT_MET"
	SMPOLICYASSOCIATIONRELEASECAUSE_REACTIVATION_REQUESTED       SmPolicyAssociationReleaseCause = "REACTIVATION_REQUESTED"
)

// All allowed values of SmPolicyAssociationReleaseCause enum
var AllowedSmPolicyAssociationReleaseCauseEnumValues = []SmPolicyAssociationReleaseCause{
	"UNSPECIFIED",
	"UE_SUBSCRIPTION",
	"INSUFFICIENT_RES",
	"VALIDATION_CONDITION_NOT_MET",
	"REACTIVATION_REQUESTED",
}

func (v *SmPolicyAssociationReleaseCause) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SmPolicyAssociationReleaseCause(value)
	for _, existing := range AllowedSmPolicyAssociationReleaseCauseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SmPolicyAssociationReleaseCause", value)
}

// NewSmPolicyAssociationReleaseCauseFromValue returns a pointer to a valid SmPolicyAssociationReleaseCause
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSmPolicyAssociationReleaseCauseFromValue(v string) (*SmPolicyAssociationReleaseCause, error) {
	ev := SmPolicyAssociationReleaseCause(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SmPolicyAssociationReleaseCause: valid values are %v", v, AllowedSmPolicyAssociationReleaseCauseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SmPolicyAssociationReleaseCause) IsValid() bool {
	for _, existing := range AllowedSmPolicyAssociationReleaseCauseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SmPolicyAssociationReleaseCause value
func (v SmPolicyAssociationReleaseCause) Ptr() *SmPolicyAssociationReleaseCause {
	return &v
}

type NullableSmPolicyAssociationReleaseCause struct {
	value *SmPolicyAssociationReleaseCause
	isSet bool
}

func (v NullableSmPolicyAssociationReleaseCause) Get() *SmPolicyAssociationReleaseCause {
	return v.value
}

func (v *NullableSmPolicyAssociationReleaseCause) Set(val *SmPolicyAssociationReleaseCause) {
	v.value = val
	v.isSet = true
}

func (v NullableSmPolicyAssociationReleaseCause) IsSet() bool {
	return v.isSet
}

func (v *NullableSmPolicyAssociationReleaseCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmPolicyAssociationReleaseCause(val *SmPolicyAssociationReleaseCause) *NullableSmPolicyAssociationReleaseCause {
	return &NullableSmPolicyAssociationReleaseCause{value: val, isSet: true}
}

func (v NullableSmPolicyAssociationReleaseCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmPolicyAssociationReleaseCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}