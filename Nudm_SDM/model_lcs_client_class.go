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

// LcsClientClass Indicates LCS Client class.
type LcsClientClass string

// List of LcsClientClass
const (
	LCSCLIENTCLASS_BROADCAST_SERVICE          LcsClientClass = "BROADCAST_SERVICE"
	LCSCLIENTCLASS_OM_IN_HPLMN                LcsClientClass = "OM_IN_HPLMN"
	LCSCLIENTCLASS_OM_IN_VPLMN                LcsClientClass = "OM_IN_VPLMN"
	LCSCLIENTCLASS_ANONYMOUS_LOCATION_SERVICE LcsClientClass = "ANONYMOUS_LOCATION_SERVICE"
	LCSCLIENTCLASS_SPECIFIC_SERVICE           LcsClientClass = "SPECIFIC_SERVICE"
	LCSCLIENTCLASS_NWDAF_IN_HPLMN             LcsClientClass = "NWDAF_IN_HPLMN"
	LCSCLIENTCLASS_NWDAF_IN_VPLMN             LcsClientClass = "NWDAF_IN_VPLMN"
)

// All allowed values of LcsClientClass enum
var AllowedLcsClientClassEnumValues = []LcsClientClass{
	"BROADCAST_SERVICE",
	"OM_IN_HPLMN",
	"OM_IN_VPLMN",
	"ANONYMOUS_LOCATION_SERVICE",
	"SPECIFIC_SERVICE",
	"NWDAF_IN_HPLMN",
	"NWDAF_IN_VPLMN",
}

func (v *LcsClientClass) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LcsClientClass(value)
	for _, existing := range AllowedLcsClientClassEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LcsClientClass", value)
}

// NewLcsClientClassFromValue returns a pointer to a valid LcsClientClass
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLcsClientClassFromValue(v string) (*LcsClientClass, error) {
	ev := LcsClientClass(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LcsClientClass: valid values are %v", v, AllowedLcsClientClassEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LcsClientClass) IsValid() bool {
	for _, existing := range AllowedLcsClientClassEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LcsClientClass value
func (v LcsClientClass) Ptr() *LcsClientClass {
	return &v
}

type NullableLcsClientClass struct {
	value *LcsClientClass
	isSet bool
}

func (v NullableLcsClientClass) Get() *LcsClientClass {
	return v.value
}

func (v *NullableLcsClientClass) Set(val *LcsClientClass) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsClientClass) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsClientClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsClientClass(val *LcsClientClass) *NullableLcsClientClass {
	return &NullableLcsClientClass{value: val, isSet: true}
}

func (v NullableLcsClientClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsClientClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}