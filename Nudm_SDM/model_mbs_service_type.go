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

// MbsServiceType Indicates the MBS service type of an MBS session
type MbsServiceType string

// List of MbsServiceType
const (
	MBSSERVICETYPE_MULTICAST MbsServiceType = "MULTICAST"
	MBSSERVICETYPE_BROADCAST MbsServiceType = "BROADCAST"
)

// All allowed values of MbsServiceType enum
var AllowedMbsServiceTypeEnumValues = []MbsServiceType{
	"MULTICAST",
	"BROADCAST",
}

func (v *MbsServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MbsServiceType(value)
	for _, existing := range AllowedMbsServiceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MbsServiceType", value)
}

// NewMbsServiceTypeFromValue returns a pointer to a valid MbsServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMbsServiceTypeFromValue(v string) (*MbsServiceType, error) {
	ev := MbsServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MbsServiceType: valid values are %v", v, AllowedMbsServiceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MbsServiceType) IsValid() bool {
	for _, existing := range AllowedMbsServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MbsServiceType value
func (v MbsServiceType) Ptr() *MbsServiceType {
	return &v
}

type NullableMbsServiceType struct {
	value *MbsServiceType
	isSet bool
}

func (v NullableMbsServiceType) Get() *MbsServiceType {
	return v.value
}

func (v *NullableMbsServiceType) Set(val *MbsServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsServiceType(val *MbsServiceType) *NullableMbsServiceType {
	return &NullableMbsServiceType{value: val, isSet: true}
}

func (v NullableMbsServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}