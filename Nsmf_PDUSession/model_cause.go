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

// Cause Cause information. Possible values are - REL_DUE_TO_HO - EPS_FALLBACK - REL_DUE_TO_UP_SEC - DNN_CONGESTION - S_NSSAI_CONGESTION - REL_DUE_TO_REACTIVATION - 5G_AN_NOT_RESPONDING - REL_DUE_TO_SLICE_NOT_AVAILABLE - REL_DUE_TO_DUPLICATE_SESSION_ID - PDU_SESSION_STATUS_MISMATCH - HO_FAILURE - INSUFFICIENT_UP_RESOURCES - PDU_SESSION_HANDED_OVER - PDU_SESSION_RESUMED - CN_ASSISTED_RAN_PARAMETER_TUNING - ISMF_CONTEXT_TRANSFER - SMF_CONTEXT_TRANSFER - REL_DUE_TO_PS_TO_CS_HO - REL_DUE_TO_SUBSCRIPTION_CHANGE - HO_CANCEL - REL_DUE_TO_SLICE_NOT_AUTHORIZED - PDU_SESSION_HAND_OVER_FAILURE - DDN_FAILURE_STATUS - REL_DUE_TO_CP_ONLY_NOT_APPLICABLE - NOT_SUPPORTED_WITH_ISMF - CHANGED_ANCHOR_SMF - CHANGED_INTERMEDIATE_SMF - TARGET_DNAI_NOTIFICATION - REL_DUE_TO_VPLMN_QOS_FAILURE - REL_DUE_TO_SMF_NOT_SUPPORT_PSETR - REL_DUE_TO_SNPN_SNPN_MOBILITY - REL_DUE_TO_NO_HR_AGREEMENT - REL_DUE_TO_UNSPECIFIED_REASON - REL_DUE_TO_MOB_ACCESS_RESTRICTIONS - REL_DUE_TO_SLICE_INACTIVITY - REL_DUE_TO_NSI_NOT_AVAILABLE - REL_DUE_TO_DNN_DENIED - REL_DUE_TO_PDUTYPE_DENIED - REL_DUE_TO_SSC_DENIED - REL_DUE_TO_SUBSCRIPTION_DENIED - REL_DUE_TO_DNN_NOT_SUPPORTED - REL_DUE_TO_PDUTYPE_NOT_SUPPORTED - REL_DUE_TO_SSC_NOT_SUPPORTED - REL_DUE_TO_INSUFFICIENT_RESOURCES_SLICE - REL_DUE_TO_INSUFFICIENT_RESOURCES_SLICE_DNN - REL_DUE_TO_DNN_CONGESTION - REL_DUE_TO_S_NSSAI_CONGESTION - REL_DUE_TO_PEER_NOT_RESPONDING - REL_DUE_TO_NETWORK_FAILURE - REL_DUE_TO_UPF_NOT_RESPONDING - REL_DUE_TO_NO_EPS_5GS_CONTINUITY - REL_DUE_TO_NOT_SUPPORTED_WITH_ISMF - REL_DUE_TO_EXCEEDED_UE_SLICE_DATA_RATE - REL_DUE_TO_EXCEEDED_SLICE_DATA_RATE - REL_DUE_TO_CONTEXT_NOT_FOUND - REL_DUE_TO_MBSR_NOT_AUTHORIZED - DEACT_DUE_TO_UE_OUT_OF_SLICE_SUPPORT_AREA - REJECT_DUE_TO_N1_SM_ERROR
type Cause string

// List of Cause
const (
	CAUSE_REL_DUE_TO_HO                               Cause = "REL_DUE_TO_HO"
	CAUSE_EPS_FALLBACK                                Cause = "EPS_FALLBACK"
	CAUSE_REL_DUE_TO_UP_SEC                           Cause = "REL_DUE_TO_UP_SEC"
	CAUSE_DNN_CONGESTION                              Cause = "DNN_CONGESTION"
	CAUSE_S_NSSAI_CONGESTION                          Cause = "S_NSSAI_CONGESTION"
	CAUSE_REL_DUE_TO_REACTIVATION                     Cause = "REL_DUE_TO_REACTIVATION"
	CAUSE__5_G_AN_NOT_RESPONDING                      Cause = "5G_AN_NOT_RESPONDING"
	CAUSE_REL_DUE_TO_SLICE_NOT_AVAILABLE              Cause = "REL_DUE_TO_SLICE_NOT_AVAILABLE"
	CAUSE_REL_DUE_TO_DUPLICATE_SESSION_ID             Cause = "REL_DUE_TO_DUPLICATE_SESSION_ID"
	CAUSE_PDU_SESSION_STATUS_MISMATCH                 Cause = "PDU_SESSION_STATUS_MISMATCH"
	CAUSE_HO_FAILURE                                  Cause = "HO_FAILURE"
	CAUSE_INSUFFICIENT_UP_RESOURCES                   Cause = "INSUFFICIENT_UP_RESOURCES"
	CAUSE_PDU_SESSION_HANDED_OVER                     Cause = "PDU_SESSION_HANDED_OVER"
	CAUSE_PDU_SESSION_RESUMED                         Cause = "PDU_SESSION_RESUMED"
	CAUSE_CN_ASSISTED_RAN_PARAMETER_TUNING            Cause = "CN_ASSISTED_RAN_PARAMETER_TUNING"
	CAUSE_ISMF_CONTEXT_TRANSFER                       Cause = "ISMF_CONTEXT_TRANSFER"
	CAUSE_SMF_CONTEXT_TRANSFER                        Cause = "SMF_CONTEXT_TRANSFER"
	CAUSE_REL_DUE_TO_PS_TO_CS_HO                      Cause = "REL_DUE_TO_PS_TO_CS_HO"
	CAUSE_REL_DUE_TO_SUBSCRIPTION_CHANGE              Cause = "REL_DUE_TO_SUBSCRIPTION_CHANGE"
	CAUSE_HO_CANCEL                                   Cause = "HO_CANCEL"
	CAUSE_REL_DUE_TO_SLICE_NOT_AUTHORIZED             Cause = "REL_DUE_TO_SLICE_NOT_AUTHORIZED"
	CAUSE_PDU_SESSION_HAND_OVER_FAILURE               Cause = "PDU_SESSION_HAND_OVER_FAILURE"
	CAUSE_DDN_FAILURE_STATUS                          Cause = "DDN_FAILURE_STATUS"
	CAUSE_REL_DUE_TO_CP_ONLY_NOT_APPLICABLE           Cause = "REL_DUE_TO_CP_ONLY_NOT_APPLICABLE"
	CAUSE_NOT_SUPPORTED_WITH_ISMF                     Cause = "NOT_SUPPORTED_WITH_ISMF"
	CAUSE_CHANGED_ANCHOR_SMF                          Cause = "CHANGED_ANCHOR_SMF"
	CAUSE_CHANGED_INTERMEDIATE_SMF                    Cause = "CHANGED_INTERMEDIATE_SMF"
	CAUSE_TARGET_DNAI_NOTIFICATION                    Cause = "TARGET_DNAI_NOTIFICATION"
	CAUSE_REL_DUE_TO_VPLMN_QOS_FAILURE                Cause = "REL_DUE_TO_VPLMN_QOS_FAILURE"
	CAUSE_REL_DUE_TO_SMF_NOT_SUPPORT_PSETR            Cause = "REL_DUE_TO_SMF_NOT_SUPPORT_PSETR"
	CAUSE_REL_DUE_TO_SNPN_SNPN_MOBILITY               Cause = "REL_DUE_TO_SNPN_SNPN_MOBILITY"
	CAUSE_REL_DUE_TO_NO_HR_AGREEMENT                  Cause = "REL_DUE_TO_NO_HR_AGREEMENT"
	CAUSE_REL_DUE_TO_UNSPECIFIED_REASON               Cause = "REL_DUE_TO_UNSPECIFIED_REASON"
	CAUSE_REL_DUE_TO_MOB_ACCESS_RESTRICTIONS          Cause = "REL_DUE_TO_MOB_ACCESS_RESTRICTIONS"
	CAUSE_REL_DUE_TO_SLICE_INACTIVITY                 Cause = "REL_DUE_TO_SLICE_INACTIVITY"
	CAUSE_REL_DUE_TO_NSI_NOT_AVAILABLE                Cause = "REL_DUE_TO_NSI_NOT_AVAILABLE"
	CAUSE_REL_DUE_TO_DNN_DENIED                       Cause = "REL_DUE_TO_DNN_DENIED"
	CAUSE_REL_DUE_TO_PDUTYPE_DENIED                   Cause = "REL_DUE_TO_PDUTYPE_DENIED"
	CAUSE_REL_DUE_TO_SSC_DENIED                       Cause = "REL_DUE_TO_SSC_DENIED"
	CAUSE_REL_DUE_TO_SUBSCRIPTION_DENIED              Cause = "REL_DUE_TO_SUBSCRIPTION_DENIED"
	CAUSE_REL_DUE_TO_DNN_NOT_SUPPORTED                Cause = "REL_DUE_TO_DNN_NOT_SUPPORTED"
	CAUSE_REL_DUE_TO_PDUTYPE_NOT_SUPPORTED            Cause = "REL_DUE_TO_PDUTYPE_NOT_SUPPORTED"
	CAUSE_REL_DUE_TO_SSC_NOT_SUPPORTED                Cause = "REL_DUE_TO_SSC_NOT_SUPPORTED"
	CAUSE_REL_DUE_TO_INSUFFICIENT_RESOURCES_SLICE     Cause = "REL_DUE_TO_INSUFFICIENT_RESOURCES_SLICE"
	CAUSE_REL_DUE_TO_INSUFFICIENT_RESOURCES_SLICE_DNN Cause = "REL_DUE_TO_INSUFFICIENT_RESOURCES_SLICE_DNN"
	CAUSE_REL_DUE_TO_DNN_CONGESTION                   Cause = "REL_DUE_TO_DNN_CONGESTION"
	CAUSE_REL_DUE_TO_S_NSSAI_CONGESTION               Cause = "REL_DUE_TO_S_NSSAI_CONGESTION"
	CAUSE_REL_DUE_TO_PEER_NOT_RESPONDING              Cause = "REL_DUE_TO_PEER_NOT_RESPONDING"
	CAUSE_REL_DUE_TO_NETWORK_FAILURE                  Cause = "REL_DUE_TO_NETWORK_FAILURE"
	CAUSE_REL_DUE_TO_UPF_NOT_RESPONDING               Cause = "REL_DUE_TO_UPF_NOT_RESPONDING"
	CAUSE_REL_DUE_TO_NO_EPS_5_GS_CONTINUITY           Cause = "REL_DUE_TO_NO_EPS_5GS_CONTINUITY"
	CAUSE_REL_DUE_TO_NOT_SUPPORTED_WITH_ISMF          Cause = "REL_DUE_TO_NOT_SUPPORTED_WITH_ISMF"
	CAUSE_REL_DUE_TO_EXCEEDED_UE_SLICE_DATA_RATE      Cause = "REL_DUE_TO_EXCEEDED_UE_SLICE_DATA_RATE"
	CAUSE_REL_DUE_TO_EXCEEDED_SLICE_DATA_RATE         Cause = "REL_DUE_TO_EXCEEDED_SLICE_DATA_RATE"
	CAUSE_REL_DUE_TO_CONTEXT_NOT_FOUND                Cause = "REL_DUE_TO_CONTEXT_NOT_FOUND"
	CAUSE_REL_DUE_TO_MBSR_NOT_AUTHORIZED              Cause = "REL_DUE_TO_MBSR_NOT_AUTHORIZED"
	CAUSE_DEACT_DUE_TO_UE_OUT_OF_SLICE_SUPPORT_AREA   Cause = "DEACT_DUE_TO_UE_OUT_OF_SLICE_SUPPORT_AREA"
	CAUSE_REJECT_DUE_TO_N1_SM_ERROR                   Cause = "REJECT_DUE_TO_N1_SM_ERROR"
)

// All allowed values of Cause enum
var AllowedCauseEnumValues = []Cause{
	"REL_DUE_TO_HO",
	"EPS_FALLBACK",
	"REL_DUE_TO_UP_SEC",
	"DNN_CONGESTION",
	"S_NSSAI_CONGESTION",
	"REL_DUE_TO_REACTIVATION",
	"5G_AN_NOT_RESPONDING",
	"REL_DUE_TO_SLICE_NOT_AVAILABLE",
	"REL_DUE_TO_DUPLICATE_SESSION_ID",
	"PDU_SESSION_STATUS_MISMATCH",
	"HO_FAILURE",
	"INSUFFICIENT_UP_RESOURCES",
	"PDU_SESSION_HANDED_OVER",
	"PDU_SESSION_RESUMED",
	"CN_ASSISTED_RAN_PARAMETER_TUNING",
	"ISMF_CONTEXT_TRANSFER",
	"SMF_CONTEXT_TRANSFER",
	"REL_DUE_TO_PS_TO_CS_HO",
	"REL_DUE_TO_SUBSCRIPTION_CHANGE",
	"HO_CANCEL",
	"REL_DUE_TO_SLICE_NOT_AUTHORIZED",
	"PDU_SESSION_HAND_OVER_FAILURE",
	"DDN_FAILURE_STATUS",
	"REL_DUE_TO_CP_ONLY_NOT_APPLICABLE",
	"NOT_SUPPORTED_WITH_ISMF",
	"CHANGED_ANCHOR_SMF",
	"CHANGED_INTERMEDIATE_SMF",
	"TARGET_DNAI_NOTIFICATION",
	"REL_DUE_TO_VPLMN_QOS_FAILURE",
	"REL_DUE_TO_SMF_NOT_SUPPORT_PSETR",
	"REL_DUE_TO_SNPN_SNPN_MOBILITY",
	"REL_DUE_TO_NO_HR_AGREEMENT",
	"REL_DUE_TO_UNSPECIFIED_REASON",
	"REL_DUE_TO_MOB_ACCESS_RESTRICTIONS",
	"REL_DUE_TO_SLICE_INACTIVITY",
	"REL_DUE_TO_NSI_NOT_AVAILABLE",
	"REL_DUE_TO_DNN_DENIED",
	"REL_DUE_TO_PDUTYPE_DENIED",
	"REL_DUE_TO_SSC_DENIED",
	"REL_DUE_TO_SUBSCRIPTION_DENIED",
	"REL_DUE_TO_DNN_NOT_SUPPORTED",
	"REL_DUE_TO_PDUTYPE_NOT_SUPPORTED",
	"REL_DUE_TO_SSC_NOT_SUPPORTED",
	"REL_DUE_TO_INSUFFICIENT_RESOURCES_SLICE",
	"REL_DUE_TO_INSUFFICIENT_RESOURCES_SLICE_DNN",
	"REL_DUE_TO_DNN_CONGESTION",
	"REL_DUE_TO_S_NSSAI_CONGESTION",
	"REL_DUE_TO_PEER_NOT_RESPONDING",
	"REL_DUE_TO_NETWORK_FAILURE",
	"REL_DUE_TO_UPF_NOT_RESPONDING",
	"REL_DUE_TO_NO_EPS_5GS_CONTINUITY",
	"REL_DUE_TO_NOT_SUPPORTED_WITH_ISMF",
	"REL_DUE_TO_EXCEEDED_UE_SLICE_DATA_RATE",
	"REL_DUE_TO_EXCEEDED_SLICE_DATA_RATE",
	"REL_DUE_TO_CONTEXT_NOT_FOUND",
	"REL_DUE_TO_MBSR_NOT_AUTHORIZED",
	"DEACT_DUE_TO_UE_OUT_OF_SLICE_SUPPORT_AREA",
	"REJECT_DUE_TO_N1_SM_ERROR",
}

func (v *Cause) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Cause(value)
	for _, existing := range AllowedCauseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Cause", value)
}

// NewCauseFromValue returns a pointer to a valid Cause
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCauseFromValue(v string) (*Cause, error) {
	ev := Cause(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Cause: valid values are %v", v, AllowedCauseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Cause) IsValid() bool {
	for _, existing := range AllowedCauseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Cause value
func (v Cause) Ptr() *Cause {
	return &v
}

type NullableCause struct {
	value *Cause
	isSet bool
}

func (v NullableCause) Get() *Cause {
	return v.value
}

func (v *NullableCause) Set(val *Cause) {
	v.value = val
	v.isSet = true
}

func (v NullableCause) IsSet() bool {
	return v.isSet
}

func (v *NullableCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCause(val *Cause) *NullableCause {
	return &NullableCause{value: val, isSet: true}
}

func (v NullableCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}