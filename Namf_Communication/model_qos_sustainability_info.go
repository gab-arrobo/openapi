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

// checks if the QosSustainabilityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosSustainabilityInfo{}

// QosSustainabilityInfo Represents the QoS Sustainability information.
type QosSustainabilityInfo struct {
	AreaInfo *NetworkAreaInfo
	// This attribute contains the geographical locations in a fine granularity.
	FineAreaInfos []GeographicalArea
	// string with format 'date-time' as defined in OpenAPI.
	StartTs *time.Time
	// string with format 'date-time' as defined in OpenAPI.
	EndTs         *time.Time
	QosFlowRetThd NullableRetainabilityThreshold
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	RanUeThrouThd *string
	Snssai        *Snssai
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence           *int32
	AdditionalProperties map[string]interface{}
}

type _QosSustainabilityInfo QosSustainabilityInfo

// NewQosSustainabilityInfo instantiates a new QosSustainabilityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosSustainabilityInfo() *QosSustainabilityInfo {
	this := QosSustainabilityInfo{}
	return &this
}

// NewQosSustainabilityInfoWithDefaults instantiates a new QosSustainabilityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosSustainabilityInfoWithDefaults() *QosSustainabilityInfo {
	this := QosSustainabilityInfo{}
	return &this
}

// GetAreaInfo returns the AreaInfo field value if set, zero value otherwise.
func (o *QosSustainabilityInfo) GetAreaInfo() NetworkAreaInfo {
	if o == nil || IsNil(o.AreaInfo) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.AreaInfo
}

// GetAreaInfoOk returns a tuple with the AreaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosSustainabilityInfo) GetAreaInfoOk() (*NetworkAreaInfo, bool) {
	if o == nil || IsNil(o.AreaInfo) {
		return nil, false
	}
	return o.AreaInfo, true
}

// HasAreaInfo returns a boolean if a field has been set.
func (o *QosSustainabilityInfo) HasAreaInfo() bool {
	if o != nil && !IsNil(o.AreaInfo) {
		return true
	}

	return false
}

// SetAreaInfo gets a reference to the given NetworkAreaInfo and assigns it to the AreaInfo field.
func (o *QosSustainabilityInfo) SetAreaInfo(v NetworkAreaInfo) {
	o.AreaInfo = &v
}

// GetFineAreaInfos returns the FineAreaInfos field value if set, zero value otherwise.
func (o *QosSustainabilityInfo) GetFineAreaInfos() []GeographicalArea {
	if o == nil || IsNil(o.FineAreaInfos) {
		var ret []GeographicalArea
		return ret
	}
	return o.FineAreaInfos
}

// GetFineAreaInfosOk returns a tuple with the FineAreaInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosSustainabilityInfo) GetFineAreaInfosOk() ([]GeographicalArea, bool) {
	if o == nil || IsNil(o.FineAreaInfos) {
		return nil, false
	}
	return o.FineAreaInfos, true
}

// HasFineAreaInfos returns a boolean if a field has been set.
func (o *QosSustainabilityInfo) HasFineAreaInfos() bool {
	if o != nil && !IsNil(o.FineAreaInfos) {
		return true
	}

	return false
}

// SetFineAreaInfos gets a reference to the given []GeographicalArea and assigns it to the FineAreaInfos field.
func (o *QosSustainabilityInfo) SetFineAreaInfos(v []GeographicalArea) {
	o.FineAreaInfos = v
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *QosSustainabilityInfo) GetStartTs() time.Time {
	if o == nil || IsNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosSustainabilityInfo) GetStartTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTs) {
		return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *QosSustainabilityInfo) HasStartTs() bool {
	if o != nil && !IsNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *QosSustainabilityInfo) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *QosSustainabilityInfo) GetEndTs() time.Time {
	if o == nil || IsNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosSustainabilityInfo) GetEndTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTs) {
		return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *QosSustainabilityInfo) HasEndTs() bool {
	if o != nil && !IsNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *QosSustainabilityInfo) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetQosFlowRetThd returns the QosFlowRetThd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosSustainabilityInfo) GetQosFlowRetThd() RetainabilityThreshold {
	if o == nil || IsNil(o.QosFlowRetThd.Get()) {
		var ret RetainabilityThreshold
		return ret
	}
	return *o.QosFlowRetThd.Get()
}

// GetQosFlowRetThdOk returns a tuple with the QosFlowRetThd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosSustainabilityInfo) GetQosFlowRetThdOk() (*RetainabilityThreshold, bool) {
	if o == nil {
		return nil, false
	}
	return o.QosFlowRetThd.Get(), o.QosFlowRetThd.IsSet()
}

// HasQosFlowRetThd returns a boolean if a field has been set.
func (o *QosSustainabilityInfo) HasQosFlowRetThd() bool {
	if o != nil && o.QosFlowRetThd.IsSet() {
		return true
	}

	return false
}

// SetQosFlowRetThd gets a reference to the given NullableRetainabilityThreshold and assigns it to the QosFlowRetThd field.
func (o *QosSustainabilityInfo) SetQosFlowRetThd(v RetainabilityThreshold) {
	o.QosFlowRetThd.Set(&v)
}

// SetQosFlowRetThdNil sets the value for QosFlowRetThd to be an explicit nil
func (o *QosSustainabilityInfo) SetQosFlowRetThdNil() {
	o.QosFlowRetThd.Set(nil)
}

// UnsetQosFlowRetThd ensures that no value is present for QosFlowRetThd, not even an explicit nil
func (o *QosSustainabilityInfo) UnsetQosFlowRetThd() {
	o.QosFlowRetThd.Unset()
}

// GetRanUeThrouThd returns the RanUeThrouThd field value if set, zero value otherwise.
func (o *QosSustainabilityInfo) GetRanUeThrouThd() string {
	if o == nil || IsNil(o.RanUeThrouThd) {
		var ret string
		return ret
	}
	return *o.RanUeThrouThd
}

// GetRanUeThrouThdOk returns a tuple with the RanUeThrouThd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosSustainabilityInfo) GetRanUeThrouThdOk() (*string, bool) {
	if o == nil || IsNil(o.RanUeThrouThd) {
		return nil, false
	}
	return o.RanUeThrouThd, true
}

// HasRanUeThrouThd returns a boolean if a field has been set.
func (o *QosSustainabilityInfo) HasRanUeThrouThd() bool {
	if o != nil && !IsNil(o.RanUeThrouThd) {
		return true
	}

	return false
}

// SetRanUeThrouThd gets a reference to the given string and assigns it to the RanUeThrouThd field.
func (o *QosSustainabilityInfo) SetRanUeThrouThd(v string) {
	o.RanUeThrouThd = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *QosSustainabilityInfo) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosSustainabilityInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *QosSustainabilityInfo) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *QosSustainabilityInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *QosSustainabilityInfo) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosSustainabilityInfo) GetConfidenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *QosSustainabilityInfo) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *QosSustainabilityInfo) SetConfidence(v int32) {
	o.Confidence = &v
}

func (o QosSustainabilityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AreaInfo) {
		toSerialize["areaInfo"] = o.AreaInfo
	}
	if !IsNil(o.FineAreaInfos) {
		toSerialize["fineAreaInfos"] = o.FineAreaInfos
	}
	if !IsNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !IsNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if o.QosFlowRetThd.IsSet() {
		toSerialize["qosFlowRetThd"] = o.QosFlowRetThd.Get()
	}
	if !IsNil(o.RanUeThrouThd) {
		toSerialize["ranUeThrouThd"] = o.RanUeThrouThd
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableQosSustainabilityInfo struct {
	value *QosSustainabilityInfo
	isSet bool
}

func (v NullableQosSustainabilityInfo) Get() *QosSustainabilityInfo {
	return v.value
}

func (v *NullableQosSustainabilityInfo) Set(val *QosSustainabilityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQosSustainabilityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQosSustainabilityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosSustainabilityInfo(val *QosSustainabilityInfo) *NullableQosSustainabilityInfo {
	return &NullableQosSustainabilityInfo{value: val, isSet: true}
}

func (v NullableQosSustainabilityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosSustainabilityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}