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
)

// checks if the L4sSupportInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &L4sSupportInfo{}

// L4sSupportInfo Contains the ECN marking for L4S support in 5GS information.
type L4sSupportInfo struct {
	// An array of PCC rule id references to the PCC rules associated with the ECN marking for L4S support info.
	RefPccRuleIds        []string     `json:"refPccRuleIds"`
	NotifType            L4sNotifType `json:"notifType"`
	AdditionalProperties map[string]interface{}
}

type _L4sSupportInfo L4sSupportInfo

// NewL4sSupportInfo instantiates a new L4sSupportInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewL4sSupportInfo(refPccRuleIds []string, notifType L4sNotifType) *L4sSupportInfo {
	this := L4sSupportInfo{}
	this.RefPccRuleIds = refPccRuleIds
	this.NotifType = notifType
	return &this
}

// NewL4sSupportInfoWithDefaults instantiates a new L4sSupportInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewL4sSupportInfoWithDefaults() *L4sSupportInfo {
	this := L4sSupportInfo{}
	return &this
}

// GetRefPccRuleIds returns the RefPccRuleIds field value
func (o *L4sSupportInfo) GetRefPccRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RefPccRuleIds
}

// GetRefPccRuleIdsOk returns a tuple with the RefPccRuleIds field value
// and a boolean to check if the value has been set.
func (o *L4sSupportInfo) GetRefPccRuleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefPccRuleIds, true
}

// SetRefPccRuleIds sets field value
func (o *L4sSupportInfo) SetRefPccRuleIds(v []string) {
	o.RefPccRuleIds = v
}

// GetNotifType returns the NotifType field value
func (o *L4sSupportInfo) GetNotifType() L4sNotifType {
	if o == nil {
		var ret L4sNotifType
		return ret
	}

	return o.NotifType
}

// GetNotifTypeOk returns a tuple with the NotifType field value
// and a boolean to check if the value has been set.
func (o *L4sSupportInfo) GetNotifTypeOk() (*L4sNotifType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifType, true
}

// SetNotifType sets field value
func (o *L4sSupportInfo) SetNotifType(v L4sNotifType) {
	o.NotifType = v
}

func (o L4sSupportInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["refPccRuleIds"] = o.RefPccRuleIds
	toSerialize["notifType"] = o.NotifType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableL4sSupportInfo struct {
	value *L4sSupportInfo
	isSet bool
}

func (v NullableL4sSupportInfo) Get() *L4sSupportInfo {
	return v.value
}

func (v *NullableL4sSupportInfo) Set(val *L4sSupportInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableL4sSupportInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableL4sSupportInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL4sSupportInfo(val *L4sSupportInfo) *NullableL4sSupportInfo {
	return &NullableL4sSupportInfo{value: val, isSet: true}
}

func (v NullableL4sSupportInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL4sSupportInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}