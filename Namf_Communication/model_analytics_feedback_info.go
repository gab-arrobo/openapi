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
	"time"
)

// checks if the AnalyticsFeedbackInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsFeedbackInfo{}

// AnalyticsFeedbackInfo Analytics feedback information.
type AnalyticsFeedbackInfo struct {
	// The times at which an action was taken.
	ActionTimes []time.Time `json:"actionTimes"`
	// The analytics types that were used to take the action.
	UsedAnaTypes []NwdafEvent `json:"usedAnaTypes,omitempty"`
	// Indication about the impact of an action on the ground truth data.
	ImpactInd            *bool `json:"impactInd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AnalyticsFeedbackInfo AnalyticsFeedbackInfo

// NewAnalyticsFeedbackInfo instantiates a new AnalyticsFeedbackInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsFeedbackInfo(actionTimes []time.Time) *AnalyticsFeedbackInfo {
	this := AnalyticsFeedbackInfo{}
	this.ActionTimes = actionTimes
	return &this
}

// NewAnalyticsFeedbackInfoWithDefaults instantiates a new AnalyticsFeedbackInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsFeedbackInfoWithDefaults() *AnalyticsFeedbackInfo {
	this := AnalyticsFeedbackInfo{}
	return &this
}

// GetActionTimes returns the ActionTimes field value
func (o *AnalyticsFeedbackInfo) GetActionTimes() []time.Time {
	if o == nil {
		var ret []time.Time
		return ret
	}

	return o.ActionTimes
}

// GetActionTimesOk returns a tuple with the ActionTimes field value
// and a boolean to check if the value has been set.
func (o *AnalyticsFeedbackInfo) GetActionTimesOk() ([]time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionTimes, true
}

// SetActionTimes sets field value
func (o *AnalyticsFeedbackInfo) SetActionTimes(v []time.Time) {
	o.ActionTimes = v
}

// GetUsedAnaTypes returns the UsedAnaTypes field value if set, zero value otherwise.
func (o *AnalyticsFeedbackInfo) GetUsedAnaTypes() []NwdafEvent {
	if o == nil || IsNil(o.UsedAnaTypes) {
		var ret []NwdafEvent
		return ret
	}
	return o.UsedAnaTypes
}

// GetUsedAnaTypesOk returns a tuple with the UsedAnaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsFeedbackInfo) GetUsedAnaTypesOk() ([]NwdafEvent, bool) {
	if o == nil || IsNil(o.UsedAnaTypes) {
		return nil, false
	}
	return o.UsedAnaTypes, true
}

// HasUsedAnaTypes returns a boolean if a field has been set.
func (o *AnalyticsFeedbackInfo) HasUsedAnaTypes() bool {
	if o != nil && !IsNil(o.UsedAnaTypes) {
		return true
	}

	return false
}

// SetUsedAnaTypes gets a reference to the given []NwdafEvent and assigns it to the UsedAnaTypes field.
func (o *AnalyticsFeedbackInfo) SetUsedAnaTypes(v []NwdafEvent) {
	o.UsedAnaTypes = v
}

// GetImpactInd returns the ImpactInd field value if set, zero value otherwise.
func (o *AnalyticsFeedbackInfo) GetImpactInd() bool {
	if o == nil || IsNil(o.ImpactInd) {
		var ret bool
		return ret
	}
	return *o.ImpactInd
}

// GetImpactIndOk returns a tuple with the ImpactInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsFeedbackInfo) GetImpactIndOk() (*bool, bool) {
	if o == nil || IsNil(o.ImpactInd) {
		return nil, false
	}
	return o.ImpactInd, true
}

// HasImpactInd returns a boolean if a field has been set.
func (o *AnalyticsFeedbackInfo) HasImpactInd() bool {
	if o != nil && !IsNil(o.ImpactInd) {
		return true
	}

	return false
}

// SetImpactInd gets a reference to the given bool and assigns it to the ImpactInd field.
func (o *AnalyticsFeedbackInfo) SetImpactInd(v bool) {
	o.ImpactInd = &v
}

func (o AnalyticsFeedbackInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actionTimes"] = o.ActionTimes
	if !IsNil(o.UsedAnaTypes) {
		toSerialize["usedAnaTypes"] = o.UsedAnaTypes
	}
	if !IsNil(o.ImpactInd) {
		toSerialize["impactInd"] = o.ImpactInd
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableAnalyticsFeedbackInfo struct {
	value *AnalyticsFeedbackInfo
	isSet bool
}

func (v NullableAnalyticsFeedbackInfo) Get() *AnalyticsFeedbackInfo {
	return v.value
}

func (v *NullableAnalyticsFeedbackInfo) Set(val *AnalyticsFeedbackInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsFeedbackInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsFeedbackInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsFeedbackInfo(val *AnalyticsFeedbackInfo) *NullableAnalyticsFeedbackInfo {
	return &NullableAnalyticsFeedbackInfo{value: val, isSet: true}
}

func (v NullableAnalyticsFeedbackInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsFeedbackInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}