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

// ReportAmountMdt The enumeration ReportAmountMdt defines Report Amount for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.10-1.
type ReportAmountMdt string

// List of ReportAmountMdt
const (
	REPORTAMOUNTMDT__1       ReportAmountMdt = "1"
	REPORTAMOUNTMDT__2       ReportAmountMdt = "2"
	REPORTAMOUNTMDT__4       ReportAmountMdt = "4"
	REPORTAMOUNTMDT__8       ReportAmountMdt = "8"
	REPORTAMOUNTMDT__16      ReportAmountMdt = "16"
	REPORTAMOUNTMDT__32      ReportAmountMdt = "32"
	REPORTAMOUNTMDT__64      ReportAmountMdt = "64"
	REPORTAMOUNTMDT_INFINITY ReportAmountMdt = "infinity"
)

// All allowed values of ReportAmountMdt enum
var AllowedReportAmountMdtEnumValues = []ReportAmountMdt{
	"1",
	"2",
	"4",
	"8",
	"16",
	"32",
	"64",
	"infinity",
}

func (v *ReportAmountMdt) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportAmountMdt(value)
	for _, existing := range AllowedReportAmountMdtEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportAmountMdt", value)
}

// NewReportAmountMdtFromValue returns a pointer to a valid ReportAmountMdt
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAmountMdtFromValue(v string) (*ReportAmountMdt, error) {
	ev := ReportAmountMdt(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAmountMdt: valid values are %v", v, AllowedReportAmountMdtEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAmountMdt) IsValid() bool {
	for _, existing := range AllowedReportAmountMdtEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportAmountMdt value
func (v ReportAmountMdt) Ptr() *ReportAmountMdt {
	return &v
}

type NullableReportAmountMdt struct {
	value *ReportAmountMdt
	isSet bool
}

func (v NullableReportAmountMdt) Get() *ReportAmountMdt {
	return v.value
}

func (v *NullableReportAmountMdt) Set(val *ReportAmountMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableReportAmountMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableReportAmountMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportAmountMdt(val *ReportAmountMdt) *NullableReportAmountMdt {
	return &NullableReportAmountMdt{value: val, isSet: true}
}

func (v NullableReportAmountMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportAmountMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}