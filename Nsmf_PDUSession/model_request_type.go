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

// RequestType Request Type in Create (SM context) service operation. Possible values are - INITIAL_REQUEST - EXISTING_PDU_SESSION - INITIAL_EMERGENCY_REQUEST - EXISTING_EMERGENCY_PDU_SESSION
type RequestType string

// List of RequestType
const (
	REQUESTTYPE_INITIAL_REQUEST                RequestType = "INITIAL_REQUEST"
	REQUESTTYPE_EXISTING_PDU_SESSION           RequestType = "EXISTING_PDU_SESSION"
	REQUESTTYPE_INITIAL_EMERGENCY_REQUEST      RequestType = "INITIAL_EMERGENCY_REQUEST"
	REQUESTTYPE_EXISTING_EMERGENCY_PDU_SESSION RequestType = "EXISTING_EMERGENCY_PDU_SESSION"
)

// All allowed values of RequestType enum
var AllowedRequestTypeEnumValues = []RequestType{
	"INITIAL_REQUEST",
	"EXISTING_PDU_SESSION",
	"INITIAL_EMERGENCY_REQUEST",
	"EXISTING_EMERGENCY_PDU_SESSION",
}

func (v *RequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestType(value)
	for _, existing := range AllowedRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestType", value)
}

// NewRequestTypeFromValue returns a pointer to a valid RequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestTypeFromValue(v string) (*RequestType, error) {
	ev := RequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestType: valid values are %v", v, AllowedRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestType) IsValid() bool {
	for _, existing := range AllowedRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestType value
func (v RequestType) Ptr() *RequestType {
	return &v
}

type NullableRequestType struct {
	value *RequestType
	isSet bool
}

func (v NullableRequestType) Get() *RequestType {
	return v.value
}

func (v *NullableRequestType) Set(val *RequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestType(val *RequestType) *NullableRequestType {
	return &NullableRequestType{value: val, isSet: true}
}

func (v NullableRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}