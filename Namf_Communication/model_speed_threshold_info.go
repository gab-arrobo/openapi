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
)

// checks if the SpeedThresholdInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpeedThresholdInfo{}

// SpeedThresholdInfo UEs information whose speed is faster than the speed threshold.
type SpeedThresholdInfo struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NumOfUe *int32 `json:"numOfUe,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio                *int32 `json:"ratio,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SpeedThresholdInfo SpeedThresholdInfo

// NewSpeedThresholdInfo instantiates a new SpeedThresholdInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpeedThresholdInfo() *SpeedThresholdInfo {
	this := SpeedThresholdInfo{}
	return &this
}

// NewSpeedThresholdInfoWithDefaults instantiates a new SpeedThresholdInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpeedThresholdInfoWithDefaults() *SpeedThresholdInfo {
	this := SpeedThresholdInfo{}
	return &this
}

// GetNumOfUe returns the NumOfUe field value if set, zero value otherwise.
func (o *SpeedThresholdInfo) GetNumOfUe() int32 {
	if o == nil || IsNil(o.NumOfUe) {
		var ret int32
		return ret
	}
	return *o.NumOfUe
}

// GetNumOfUeOk returns a tuple with the NumOfUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpeedThresholdInfo) GetNumOfUeOk() (*int32, bool) {
	if o == nil || IsNil(o.NumOfUe) {
		return nil, false
	}
	return o.NumOfUe, true
}

// HasNumOfUe returns a boolean if a field has been set.
func (o *SpeedThresholdInfo) HasNumOfUe() bool {
	if o != nil && !IsNil(o.NumOfUe) {
		return true
	}

	return false
}

// SetNumOfUe gets a reference to the given int32 and assigns it to the NumOfUe field.
func (o *SpeedThresholdInfo) SetNumOfUe(v int32) {
	o.NumOfUe = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *SpeedThresholdInfo) GetRatio() int32 {
	if o == nil || IsNil(o.Ratio) {
		var ret int32
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpeedThresholdInfo) GetRatioOk() (*int32, bool) {
	if o == nil || IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *SpeedThresholdInfo) HasRatio() bool {
	if o != nil && !IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given int32 and assigns it to the Ratio field.
func (o *SpeedThresholdInfo) SetRatio(v int32) {
	o.Ratio = &v
}

func (o SpeedThresholdInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumOfUe) {
		toSerialize["numOfUe"] = o.NumOfUe
	}
	if !IsNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableSpeedThresholdInfo struct {
	value *SpeedThresholdInfo
	isSet bool
}

func (v NullableSpeedThresholdInfo) Get() *SpeedThresholdInfo {
	return v.value
}

func (v *NullableSpeedThresholdInfo) Set(val *SpeedThresholdInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSpeedThresholdInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSpeedThresholdInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpeedThresholdInfo(val *SpeedThresholdInfo) *NullableSpeedThresholdInfo {
	return &NullableSpeedThresholdInfo{value: val, isSet: true}
}

func (v NullableSpeedThresholdInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpeedThresholdInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}