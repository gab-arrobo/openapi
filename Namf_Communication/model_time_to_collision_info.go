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

// checks if the TimeToCollisionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeToCollisionInfo{}

// TimeToCollisionInfo Represents Time To Collision (TTC) information.
type TimeToCollisionInfo struct {
	// string with format 'date-time' as defined in OpenAPI.
	Ttc *time.Time `json:"ttc,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Accuracy *int32 `json:"accuracy,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence           *int32 `json:"confidence,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimeToCollisionInfo TimeToCollisionInfo

// NewTimeToCollisionInfo instantiates a new TimeToCollisionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeToCollisionInfo() *TimeToCollisionInfo {
	this := TimeToCollisionInfo{}
	return &this
}

// NewTimeToCollisionInfoWithDefaults instantiates a new TimeToCollisionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeToCollisionInfoWithDefaults() *TimeToCollisionInfo {
	this := TimeToCollisionInfo{}
	return &this
}

// GetTtc returns the Ttc field value if set, zero value otherwise.
func (o *TimeToCollisionInfo) GetTtc() time.Time {
	if o == nil || IsNil(o.Ttc) {
		var ret time.Time
		return ret
	}
	return *o.Ttc
}

// GetTtcOk returns a tuple with the Ttc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeToCollisionInfo) GetTtcOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ttc) {
		return nil, false
	}
	return o.Ttc, true
}

// HasTtc returns a boolean if a field has been set.
func (o *TimeToCollisionInfo) HasTtc() bool {
	if o != nil && !IsNil(o.Ttc) {
		return true
	}

	return false
}

// SetTtc gets a reference to the given time.Time and assigns it to the Ttc field.
func (o *TimeToCollisionInfo) SetTtc(v time.Time) {
	o.Ttc = &v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *TimeToCollisionInfo) GetAccuracy() int32 {
	if o == nil || IsNil(o.Accuracy) {
		var ret int32
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeToCollisionInfo) GetAccuracyOk() (*int32, bool) {
	if o == nil || IsNil(o.Accuracy) {
		return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *TimeToCollisionInfo) HasAccuracy() bool {
	if o != nil && !IsNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given int32 and assigns it to the Accuracy field.
func (o *TimeToCollisionInfo) SetAccuracy(v int32) {
	o.Accuracy = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *TimeToCollisionInfo) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeToCollisionInfo) GetConfidenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *TimeToCollisionInfo) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *TimeToCollisionInfo) SetConfidence(v int32) {
	o.Confidence = &v
}

func (o TimeToCollisionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ttc) {
		toSerialize["ttc"] = o.Ttc
	}
	if !IsNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableTimeToCollisionInfo struct {
	value *TimeToCollisionInfo
	isSet bool
}

func (v NullableTimeToCollisionInfo) Get() *TimeToCollisionInfo {
	return v.value
}

func (v *NullableTimeToCollisionInfo) Set(val *TimeToCollisionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeToCollisionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeToCollisionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeToCollisionInfo(val *TimeToCollisionInfo) *NullableTimeToCollisionInfo {
	return &NullableTimeToCollisionInfo{value: val, isSet: true}
}

func (v NullableTimeToCollisionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeToCollisionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}