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

// checks if the AccNetChId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccNetChId{}

// AccNetChId Contains the access network charging identifier for the PCC rule(s) or for the whole PDU session.
type AccNetChId struct {
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	// Deprecated
	AccNetChaIdValue *int32
	// A character string containing the access network charging id.
	AccNetChargId *string
	// Contains the identifier of the PCC rule(s) associated to the provided Access Network Charging Identifier.
	RefPccRuleIds []string
	// When it is included and set to true, indicates the Access Network Charging Identifier applies to the whole PDU Session
	SessionChScope       *bool
	AdditionalProperties map[string]interface{}
}

type _AccNetChId AccNetChId

// NewAccNetChId instantiates a new AccNetChId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccNetChId() *AccNetChId {
	this := AccNetChId{}
	return &this
}

// NewAccNetChIdWithDefaults instantiates a new AccNetChId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccNetChIdWithDefaults() *AccNetChId {
	this := AccNetChId{}
	return &this
}

// GetAccNetChaIdValue returns the AccNetChaIdValue field value if set, zero value otherwise.
// Deprecated
func (o *AccNetChId) GetAccNetChaIdValue() int32 {
	if o == nil || IsNil(o.AccNetChaIdValue) {
		var ret int32
		return ret
	}
	return *o.AccNetChaIdValue
}

// GetAccNetChaIdValueOk returns a tuple with the AccNetChaIdValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AccNetChId) GetAccNetChaIdValueOk() (*int32, bool) {
	if o == nil || IsNil(o.AccNetChaIdValue) {
		return nil, false
	}
	return o.AccNetChaIdValue, true
}

// HasAccNetChaIdValue returns a boolean if a field has been set.
func (o *AccNetChId) HasAccNetChaIdValue() bool {
	if o != nil && !IsNil(o.AccNetChaIdValue) {
		return true
	}

	return false
}

// SetAccNetChaIdValue gets a reference to the given int32 and assigns it to the AccNetChaIdValue field.
// Deprecated
func (o *AccNetChId) SetAccNetChaIdValue(v int32) {
	o.AccNetChaIdValue = &v
}

// GetAccNetChargId returns the AccNetChargId field value if set, zero value otherwise.
func (o *AccNetChId) GetAccNetChargId() string {
	if o == nil || IsNil(o.AccNetChargId) {
		var ret string
		return ret
	}
	return *o.AccNetChargId
}

// GetAccNetChargIdOk returns a tuple with the AccNetChargId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccNetChId) GetAccNetChargIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccNetChargId) {
		return nil, false
	}
	return o.AccNetChargId, true
}

// HasAccNetChargId returns a boolean if a field has been set.
func (o *AccNetChId) HasAccNetChargId() bool {
	if o != nil && !IsNil(o.AccNetChargId) {
		return true
	}

	return false
}

// SetAccNetChargId gets a reference to the given string and assigns it to the AccNetChargId field.
func (o *AccNetChId) SetAccNetChargId(v string) {
	o.AccNetChargId = &v
}

// GetRefPccRuleIds returns the RefPccRuleIds field value if set, zero value otherwise.
func (o *AccNetChId) GetRefPccRuleIds() []string {
	if o == nil || IsNil(o.RefPccRuleIds) {
		var ret []string
		return ret
	}
	return o.RefPccRuleIds
}

// GetRefPccRuleIdsOk returns a tuple with the RefPccRuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccNetChId) GetRefPccRuleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RefPccRuleIds) {
		return nil, false
	}
	return o.RefPccRuleIds, true
}

// HasRefPccRuleIds returns a boolean if a field has been set.
func (o *AccNetChId) HasRefPccRuleIds() bool {
	if o != nil && !IsNil(o.RefPccRuleIds) {
		return true
	}

	return false
}

// SetRefPccRuleIds gets a reference to the given []string and assigns it to the RefPccRuleIds field.
func (o *AccNetChId) SetRefPccRuleIds(v []string) {
	o.RefPccRuleIds = v
}

// GetSessionChScope returns the SessionChScope field value if set, zero value otherwise.
func (o *AccNetChId) GetSessionChScope() bool {
	if o == nil || IsNil(o.SessionChScope) {
		var ret bool
		return ret
	}
	return *o.SessionChScope
}

// GetSessionChScopeOk returns a tuple with the SessionChScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccNetChId) GetSessionChScopeOk() (*bool, bool) {
	if o == nil || IsNil(o.SessionChScope) {
		return nil, false
	}
	return o.SessionChScope, true
}

// HasSessionChScope returns a boolean if a field has been set.
func (o *AccNetChId) HasSessionChScope() bool {
	if o != nil && !IsNil(o.SessionChScope) {
		return true
	}

	return false
}

// SetSessionChScope gets a reference to the given bool and assigns it to the SessionChScope field.
func (o *AccNetChId) SetSessionChScope(v bool) {
	o.SessionChScope = &v
}

func (o AccNetChId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccNetChaIdValue) {
		toSerialize["accNetChaIdValue"] = o.AccNetChaIdValue
	}
	if !IsNil(o.AccNetChargId) {
		toSerialize["accNetChargId"] = o.AccNetChargId
	}
	if !IsNil(o.RefPccRuleIds) {
		toSerialize["refPccRuleIds"] = o.RefPccRuleIds
	}
	if !IsNil(o.SessionChScope) {
		toSerialize["sessionChScope"] = o.SessionChScope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableAccNetChId struct {
	value *AccNetChId
	isSet bool
}

func (v NullableAccNetChId) Get() *AccNetChId {
	return v.value
}

func (v *NullableAccNetChId) Set(val *AccNetChId) {
	v.value = val
	v.isSet = true
}

func (v NullableAccNetChId) IsSet() bool {
	return v.isSet
}

func (v *NullableAccNetChId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccNetChId(val *AccNetChId) *NullableAccNetChId {
	return &NullableAccNetChId{value: val, isSet: true}
}

func (v NullableAccNetChId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccNetChId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}