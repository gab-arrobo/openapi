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

// MeteringMethod Indicates the metering method.   Possible values are: - DURATION: Indicates that the duration of the service data flow traffic shall be metered. - VOLUME: Indicates that volume of the service data flow traffic shall be metered. - DURATION_VOLUME: Indicates that the duration and the volume of the service data flow traffic shall be metered. - EVENT: Indicates that events of the service data flow traffic shall be metered.
type MeteringMethod struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MeteringMethod) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(MeteringMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MeteringMethod) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMeteringMethod struct {
	value *MeteringMethod
	isSet bool
}

func (v NullableMeteringMethod) Get() *MeteringMethod {
	return v.value
}

func (v *NullableMeteringMethod) Set(val *MeteringMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableMeteringMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableMeteringMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeteringMethod(val *MeteringMethod) *NullableMeteringMethod {
	return &NullableMeteringMethod{value: val, isSet: true}
}

func (v NullableMeteringMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeteringMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}