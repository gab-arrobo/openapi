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

// DlDataDeliveryStatus Possible values are: - BUFFERED: The first downlink data is buffered with extended buffering matching the   source of the downlink traffic. - TRANSMITTED: The first downlink data matching the source of the downlink traffic is   transmitted after previous buffering or discarding of corresponding packet(s) because   the UE of the PDU Session becomes ACTIVE, and buffered data can be delivered to UE. - DISCARDED: The first downlink data matching the source of the downlink traffic is   discarded because the Extended Buffering time, as determined by the SMF, expires or   the amount of downlink data to be buffered is exceeded.
type DlDataDeliveryStatus string

// List of DlDataDeliveryStatus
const (
	DLDATADELIVERYSTATUS_BUFFERED    DlDataDeliveryStatus = "BUFFERED"
	DLDATADELIVERYSTATUS_TRANSMITTED DlDataDeliveryStatus = "TRANSMITTED"
	DLDATADELIVERYSTATUS_DISCARDED   DlDataDeliveryStatus = "DISCARDED"
)

// All allowed values of DlDataDeliveryStatus enum
var AllowedDlDataDeliveryStatusEnumValues = []DlDataDeliveryStatus{
	"BUFFERED",
	"TRANSMITTED",
	"DISCARDED",
}

func (v *DlDataDeliveryStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DlDataDeliveryStatus(value)
	for _, existing := range AllowedDlDataDeliveryStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DlDataDeliveryStatus", value)
}

// NewDlDataDeliveryStatusFromValue returns a pointer to a valid DlDataDeliveryStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDlDataDeliveryStatusFromValue(v string) (*DlDataDeliveryStatus, error) {
	ev := DlDataDeliveryStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DlDataDeliveryStatus: valid values are %v", v, AllowedDlDataDeliveryStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DlDataDeliveryStatus) IsValid() bool {
	for _, existing := range AllowedDlDataDeliveryStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DlDataDeliveryStatus value
func (v DlDataDeliveryStatus) Ptr() *DlDataDeliveryStatus {
	return &v
}

type NullableDlDataDeliveryStatus struct {
	value *DlDataDeliveryStatus
	isSet bool
}

func (v NullableDlDataDeliveryStatus) Get() *DlDataDeliveryStatus {
	return v.value
}

func (v *NullableDlDataDeliveryStatus) Set(val *DlDataDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDlDataDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDlDataDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDlDataDeliveryStatus(val *DlDataDeliveryStatus) *NullableDlDataDeliveryStatus {
	return &NullableDlDataDeliveryStatus{value: val, isSet: true}
}

func (v NullableDlDataDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDlDataDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}