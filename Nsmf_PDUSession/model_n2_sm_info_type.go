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

// N2SmInfoType N2 SM Information Type. Possible values are - PDU_RES_SETUP_REQ - PDU_RES_SETUP_RSP - PDU_RES_SETUP_FAIL - PDU_RES_REL_CMD - PDU_RES_REL_RSP - PDU_RES_MOD_REQ - PDU_RES_MOD_RSP - PDU_RES_MOD_FAIL - PDU_RES_NTY - PDU_RES_NTY_REL - PDU_RES_MOD_IND - PDU_RES_MOD_CFM - PATH_SWITCH_REQ - PATH_SWITCH_SETUP_FAIL - PATH_SWITCH_REQ_ACK - PATH_SWITCH_REQ_FAIL - HANDOVER_REQUIRED - HANDOVER_CMD - HANDOVER_PREP_FAIL - HANDOVER_REQ_ACK - HANDOVER_RES_ALLOC_FAIL - SECONDARY_RAT_USAGE - PDU_RES_MOD_IND_FAIL - UE_CONTEXT_RESUME_REQ - UE_CONTEXT_RESUME_RSP - UE_CONTEXT_SUSPEND_REQ
type N2SmInfoType string

// List of N2SmInfoType
const (
	N2SMINFOTYPE_PDU_RES_SETUP_REQ       N2SmInfoType = "PDU_RES_SETUP_REQ"
	N2SMINFOTYPE_PDU_RES_SETUP_RSP       N2SmInfoType = "PDU_RES_SETUP_RSP"
	N2SMINFOTYPE_PDU_RES_SETUP_FAIL      N2SmInfoType = "PDU_RES_SETUP_FAIL"
	N2SMINFOTYPE_PDU_RES_REL_CMD         N2SmInfoType = "PDU_RES_REL_CMD"
	N2SMINFOTYPE_PDU_RES_REL_RSP         N2SmInfoType = "PDU_RES_REL_RSP"
	N2SMINFOTYPE_PDU_RES_MOD_REQ         N2SmInfoType = "PDU_RES_MOD_REQ"
	N2SMINFOTYPE_PDU_RES_MOD_RSP         N2SmInfoType = "PDU_RES_MOD_RSP"
	N2SMINFOTYPE_PDU_RES_MOD_FAIL        N2SmInfoType = "PDU_RES_MOD_FAIL"
	N2SMINFOTYPE_PDU_RES_NTY             N2SmInfoType = "PDU_RES_NTY"
	N2SMINFOTYPE_PDU_RES_NTY_REL         N2SmInfoType = "PDU_RES_NTY_REL"
	N2SMINFOTYPE_PDU_RES_MOD_IND         N2SmInfoType = "PDU_RES_MOD_IND"
	N2SMINFOTYPE_PDU_RES_MOD_CFM         N2SmInfoType = "PDU_RES_MOD_CFM"
	N2SMINFOTYPE_PATH_SWITCH_REQ         N2SmInfoType = "PATH_SWITCH_REQ"
	N2SMINFOTYPE_PATH_SWITCH_SETUP_FAIL  N2SmInfoType = "PATH_SWITCH_SETUP_FAIL"
	N2SMINFOTYPE_PATH_SWITCH_REQ_ACK     N2SmInfoType = "PATH_SWITCH_REQ_ACK"
	N2SMINFOTYPE_PATH_SWITCH_REQ_FAIL    N2SmInfoType = "PATH_SWITCH_REQ_FAIL"
	N2SMINFOTYPE_HANDOVER_REQUIRED       N2SmInfoType = "HANDOVER_REQUIRED"
	N2SMINFOTYPE_HANDOVER_CMD            N2SmInfoType = "HANDOVER_CMD"
	N2SMINFOTYPE_HANDOVER_PREP_FAIL      N2SmInfoType = "HANDOVER_PREP_FAIL"
	N2SMINFOTYPE_HANDOVER_REQ_ACK        N2SmInfoType = "HANDOVER_REQ_ACK"
	N2SMINFOTYPE_HANDOVER_RES_ALLOC_FAIL N2SmInfoType = "HANDOVER_RES_ALLOC_FAIL"
	N2SMINFOTYPE_SECONDARY_RAT_USAGE     N2SmInfoType = "SECONDARY_RAT_USAGE"
	N2SMINFOTYPE_PDU_RES_MOD_IND_FAIL    N2SmInfoType = "PDU_RES_MOD_IND_FAIL"
	N2SMINFOTYPE_UE_CONTEXT_RESUME_REQ   N2SmInfoType = "UE_CONTEXT_RESUME_REQ"
	N2SMINFOTYPE_UE_CONTEXT_RESUME_RSP   N2SmInfoType = "UE_CONTEXT_RESUME_RSP"
	N2SMINFOTYPE_UE_CONTEXT_SUSPEND_REQ  N2SmInfoType = "UE_CONTEXT_SUSPEND_REQ"
)

// All allowed values of N2SmInfoType enum
var AllowedN2SmInfoTypeEnumValues = []N2SmInfoType{
	"PDU_RES_SETUP_REQ",
	"PDU_RES_SETUP_RSP",
	"PDU_RES_SETUP_FAIL",
	"PDU_RES_REL_CMD",
	"PDU_RES_REL_RSP",
	"PDU_RES_MOD_REQ",
	"PDU_RES_MOD_RSP",
	"PDU_RES_MOD_FAIL",
	"PDU_RES_NTY",
	"PDU_RES_NTY_REL",
	"PDU_RES_MOD_IND",
	"PDU_RES_MOD_CFM",
	"PATH_SWITCH_REQ",
	"PATH_SWITCH_SETUP_FAIL",
	"PATH_SWITCH_REQ_ACK",
	"PATH_SWITCH_REQ_FAIL",
	"HANDOVER_REQUIRED",
	"HANDOVER_CMD",
	"HANDOVER_PREP_FAIL",
	"HANDOVER_REQ_ACK",
	"HANDOVER_RES_ALLOC_FAIL",
	"SECONDARY_RAT_USAGE",
	"PDU_RES_MOD_IND_FAIL",
	"UE_CONTEXT_RESUME_REQ",
	"UE_CONTEXT_RESUME_RSP",
	"UE_CONTEXT_SUSPEND_REQ",
}

func (v *N2SmInfoType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := N2SmInfoType(value)
	for _, existing := range AllowedN2SmInfoTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid N2SmInfoType", value)
}

// NewN2SmInfoTypeFromValue returns a pointer to a valid N2SmInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewN2SmInfoTypeFromValue(v string) (*N2SmInfoType, error) {
	ev := N2SmInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for N2SmInfoType: valid values are %v", v, AllowedN2SmInfoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v N2SmInfoType) IsValid() bool {
	for _, existing := range AllowedN2SmInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to N2SmInfoType value
func (v N2SmInfoType) Ptr() *N2SmInfoType {
	return &v
}

type NullableN2SmInfoType struct {
	value *N2SmInfoType
	isSet bool
}

func (v NullableN2SmInfoType) Get() *N2SmInfoType {
	return v.value
}

func (v *NullableN2SmInfoType) Set(val *N2SmInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableN2SmInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableN2SmInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2SmInfoType(val *N2SmInfoType) *NullableN2SmInfoType {
	return &NullableN2SmInfoType{value: val, isSet: true}
}

func (v NullableN2SmInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2SmInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}