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

// SessionRuleFailureCode Indicates the reason of the session rule failure.   Possible values are - NF_MAL: Indicates that the PCC rule could not be successfully installed (for those provisioned from the PCF) or activated (for those pre-defined in SMF) or enforced (for those already successfully installed) due to SMF/UPF malfunction. - RES_LIM: Indicates that the PCC rule could not be successfully installed (for those provisioned from PCF) or activated (for those pre-defined in SMF) or enforced (for those already successfully installed) due to a limitation of resources at the SMF/UPF. - SESSION_RESOURCE_ALLOCATION_FAILURE: Indicates the session rule could not be successfully enforced due to failure during the allocation of resources for the PDU session in the UE, RAN or AMF. - UNSUCC_QOS_VAL: indicates that the QoS validation has failed. - INCORRECT_UM: The usage monitoring data of the enforced session rule is not the same for all the provisioned session rule(s). - UE_STA_SUSP: Indicates that the UE is in suspend state. - UNKNOWN_REF_ID: Indicates that the session rule could not be successfully  installed/modified because the referenced identifier to a Policy Decision Data or to a Condition Data is unknown to the SMF. - INCORRECT_COND_DATA: Indicates that the session rule could not be successfully installed/modified because the referenced Condition data are incorrect. - REF_ID_COLLISION: Indicates that the session rule could not be successfully installed/modified because the same Policy Decision is referenced by a PCC rule (e.g. the session rule and the PCC rule refer to the same Usage Monitoring decision data). - AN_GW_FAILED: Indicates that the AN-Gateway has failed and that the PCF should refrain from sending policy decisions to the SMF until it is informed that the S-GW has been recovered. This value shall not be used if the SM Policy association modification procedure is initiated for session rule removal only. - DEFAULT_QOS_MODIFICATION_FAILURE: Indicates that the enforcement of the default QoS modification failed. The SMF shall use this value to indicate to the PCF that the default QoS modification has failed. - SESSION_AMBR_MODIFICATION_FAILURE: Indicates that the enforcement of the session-AMBR modification failed. The SMF shall use this value to indicate to the PCF that the session-AMBR modification has failed.
type SessionRuleFailureCode string

// List of SessionRuleFailureCode
const (
	SESSIONRULEFAILURECODE_NF_MAL                              SessionRuleFailureCode = "NF_MAL"
	SESSIONRULEFAILURECODE_RES_LIM                             SessionRuleFailureCode = "RES_LIM"
	SESSIONRULEFAILURECODE_SESSION_RESOURCE_ALLOCATION_FAILURE SessionRuleFailureCode = "SESSION_RESOURCE_ALLOCATION_FAILURE"
	SESSIONRULEFAILURECODE_UNSUCC_QOS_VAL                      SessionRuleFailureCode = "UNSUCC_QOS_VAL"
	SESSIONRULEFAILURECODE_INCORRECT_UM                        SessionRuleFailureCode = "INCORRECT_UM"
	SESSIONRULEFAILURECODE_UE_STA_SUSP                         SessionRuleFailureCode = "UE_STA_SUSP"
	SESSIONRULEFAILURECODE_UNKNOWN_REF_ID                      SessionRuleFailureCode = "UNKNOWN_REF_ID"
	SESSIONRULEFAILURECODE_INCORRECT_COND_DATA                 SessionRuleFailureCode = "INCORRECT_COND_DATA"
	SESSIONRULEFAILURECODE_REF_ID_COLLISION                    SessionRuleFailureCode = "REF_ID_COLLISION"
	SESSIONRULEFAILURECODE_AN_GW_FAILED                        SessionRuleFailureCode = "AN_GW_FAILED"
	SESSIONRULEFAILURECODE_DEFAULT_QOS_MODIFICATION_FAILURE    SessionRuleFailureCode = "DEFAULT_QOS_MODIFICATION_FAILURE"
	SESSIONRULEFAILURECODE_SESSION_AMBR_MODIFICATION_FAILURE   SessionRuleFailureCode = "SESSION_AMBR_MODIFICATION_FAILURE"
)

// All allowed values of SessionRuleFailureCode enum
var AllowedSessionRuleFailureCodeEnumValues = []SessionRuleFailureCode{
	"NF_MAL",
	"RES_LIM",
	"SESSION_RESOURCE_ALLOCATION_FAILURE",
	"UNSUCC_QOS_VAL",
	"INCORRECT_UM",
	"UE_STA_SUSP",
	"UNKNOWN_REF_ID",
	"INCORRECT_COND_DATA",
	"REF_ID_COLLISION",
	"AN_GW_FAILED",
	"DEFAULT_QOS_MODIFICATION_FAILURE",
	"SESSION_AMBR_MODIFICATION_FAILURE",
}

func (v *SessionRuleFailureCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SessionRuleFailureCode(value)
	for _, existing := range AllowedSessionRuleFailureCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SessionRuleFailureCode", value)
}

// NewSessionRuleFailureCodeFromValue returns a pointer to a valid SessionRuleFailureCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSessionRuleFailureCodeFromValue(v string) (*SessionRuleFailureCode, error) {
	ev := SessionRuleFailureCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SessionRuleFailureCode: valid values are %v", v, AllowedSessionRuleFailureCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SessionRuleFailureCode) IsValid() bool {
	for _, existing := range AllowedSessionRuleFailureCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SessionRuleFailureCode value
func (v SessionRuleFailureCode) Ptr() *SessionRuleFailureCode {
	return &v
}

type NullableSessionRuleFailureCode struct {
	value *SessionRuleFailureCode
	isSet bool
}

func (v NullableSessionRuleFailureCode) Get() *SessionRuleFailureCode {
	return v.value
}

func (v *NullableSessionRuleFailureCode) Set(val *SessionRuleFailureCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionRuleFailureCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionRuleFailureCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionRuleFailureCode(val *SessionRuleFailureCode) *NullableSessionRuleFailureCode {
	return &NullableSessionRuleFailureCode{value: val, isSet: true}
}

func (v NullableSessionRuleFailureCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionRuleFailureCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}