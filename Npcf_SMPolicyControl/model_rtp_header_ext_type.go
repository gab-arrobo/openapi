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

// RtpHeaderExtType The enumeration indicates the type of Rtp Header Extension type
type RtpHeaderExtType string

// List of RtpHeaderExtType
const (
	RTPHEADEREXTTYPE_PDU_SET_MARKING RtpHeaderExtType = "PDU_SET_MARKING"
)

// All allowed values of RtpHeaderExtType enum
var AllowedRtpHeaderExtTypeEnumValues = []RtpHeaderExtType{
	"PDU_SET_MARKING",
}

func (v *RtpHeaderExtType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RtpHeaderExtType(value)
	for _, existing := range AllowedRtpHeaderExtTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RtpHeaderExtType", value)
}

// NewRtpHeaderExtTypeFromValue returns a pointer to a valid RtpHeaderExtType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRtpHeaderExtTypeFromValue(v string) (*RtpHeaderExtType, error) {
	ev := RtpHeaderExtType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RtpHeaderExtType: valid values are %v", v, AllowedRtpHeaderExtTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RtpHeaderExtType) IsValid() bool {
	for _, existing := range AllowedRtpHeaderExtTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RtpHeaderExtType value
func (v RtpHeaderExtType) Ptr() *RtpHeaderExtType {
	return &v
}

type NullableRtpHeaderExtType struct {
	value *RtpHeaderExtType
	isSet bool
}

func (v NullableRtpHeaderExtType) Get() *RtpHeaderExtType {
	return v.value
}

func (v *NullableRtpHeaderExtType) Set(val *RtpHeaderExtType) {
	v.value = val
	v.isSet = true
}

func (v NullableRtpHeaderExtType) IsSet() bool {
	return v.isSet
}

func (v *NullableRtpHeaderExtType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRtpHeaderExtType(val *RtpHeaderExtType) *NullableRtpHeaderExtType {
	return &NullableRtpHeaderExtType{value: val, isSet: true}
}

func (v NullableRtpHeaderExtType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRtpHeaderExtType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}