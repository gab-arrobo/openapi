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

// QosMonitoringReq QoS monitoring request. Possible values are   - UL   - DL   - BOTH   - NONE
type QosMonitoringReq string

// List of QosMonitoringReq
const (
	QOSMONITORINGREQ_UL   QosMonitoringReq = "UL"
	QOSMONITORINGREQ_DL   QosMonitoringReq = "DL"
	QOSMONITORINGREQ_BOTH QosMonitoringReq = "BOTH"
	QOSMONITORINGREQ_NONE QosMonitoringReq = "NONE"
)

// All allowed values of QosMonitoringReq enum
var AllowedQosMonitoringReqEnumValues = []QosMonitoringReq{
	"UL",
	"DL",
	"BOTH",
	"NONE",
}

func (v *QosMonitoringReq) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QosMonitoringReq(value)
	for _, existing := range AllowedQosMonitoringReqEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QosMonitoringReq", value)
}

// NewQosMonitoringReqFromValue returns a pointer to a valid QosMonitoringReq
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQosMonitoringReqFromValue(v string) (*QosMonitoringReq, error) {
	ev := QosMonitoringReq(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QosMonitoringReq: valid values are %v", v, AllowedQosMonitoringReqEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QosMonitoringReq) IsValid() bool {
	for _, existing := range AllowedQosMonitoringReqEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QosMonitoringReq value
func (v QosMonitoringReq) Ptr() *QosMonitoringReq {
	return &v
}

type NullableQosMonitoringReq struct {
	value *QosMonitoringReq
	isSet bool
}

func (v NullableQosMonitoringReq) Get() *QosMonitoringReq {
	return v.value
}

func (v *NullableQosMonitoringReq) Set(val *QosMonitoringReq) {
	v.value = val
	v.isSet = true
}

func (v NullableQosMonitoringReq) IsSet() bool {
	return v.isSet
}

func (v *NullableQosMonitoringReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosMonitoringReq(val *QosMonitoringReq) *NullableQosMonitoringReq {
	return &NullableQosMonitoringReq{value: val, isSet: true}
}

func (v NullableQosMonitoringReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosMonitoringReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}