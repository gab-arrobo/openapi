// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
Namf_Communication

AMF Communication Service.
© 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).
All rights reserved.

API version: 1.3.0-alpha.6

Authors: Aether SD-Core team
Contact: dev@lists.aetherproject.org

Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
*/

package Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// TransferReason Indicates UE Context Transfer Reason
type TransferReason string

// List of TransferReason
const (
	TRANSFERREASON_INIT_REG              TransferReason = "INIT_REG"
	TRANSFERREASON_MOBI_REG              TransferReason = "MOBI_REG"
	TRANSFERREASON_MOBI_REG_UE_VALIDATED TransferReason = "MOBI_REG_UE_VALIDATED"
)

// All allowed values of TransferReason enum
var AllowedTransferReasonEnumValues = []TransferReason{
	"INIT_REG",
	"MOBI_REG",
	"MOBI_REG_UE_VALIDATED",
}

func (v *TransferReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferReason(value)
	for _, existing := range AllowedTransferReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferReason", value)
}

// NewTransferReasonFromValue returns a pointer to a valid TransferReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferReasonFromValue(v string) (*TransferReason, error) {
	ev := TransferReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferReason: valid values are %v", v, AllowedTransferReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferReason) IsValid() bool {
	for _, existing := range AllowedTransferReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferReason value
func (v TransferReason) Ptr() *TransferReason {
	return &v
}

type NullableTransferReason struct {
	value *TransferReason
	isSet bool
}

func (v NullableTransferReason) Get() *TransferReason {
	return v.value
}

func (v *NullableTransferReason) Set(val *TransferReason) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferReason) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferReason(val *TransferReason) *NullableTransferReason {
	return &NullableTransferReason{value: val, isSet: true}
}

func (v NullableTransferReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}