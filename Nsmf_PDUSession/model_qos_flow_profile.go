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
)

// checks if the QosFlowProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosFlowProfile{}

// QosFlowProfile QoS flow profile
type QosFlowProfile struct {
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.
	Var5qi                int32                         `json:"5qi"`
	NonDynamic5Qi         *NonDynamic5Qi                `json:"nonDynamic5Qi,omitempty"`
	Dynamic5Qi            *Dynamic5Qi                   `json:"dynamic5Qi,omitempty"`
	Arp                   *Arp                          `json:"arp,omitempty"`
	GbrQosFlowInfo        *GbrQosFlowInformation        `json:"gbrQosFlowInfo,omitempty"`
	Rqa                   *ReflectiveQoSAttribute       `json:"rqa,omitempty"`
	AdditionalQosFlowInfo NullableAdditionalQosFlowInfo `json:"additionalQosFlowInfo,omitempty"`
	QosMonitoringReq      *QosMonitoringReq             `json:"qosMonitoringReq,omitempty"`
	// indicating a time in seconds.
	QosRepPeriod         *int32         `json:"qosRepPeriod,omitempty"`
	PduSetQosDl          *PduSetQosPara `json:"pduSetQosDl,omitempty"`
	PduSetQosUl          *PduSetQosPara `json:"pduSetQosUl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QosFlowProfile QosFlowProfile

// NewQosFlowProfile instantiates a new QosFlowProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosFlowProfile(var5qi int32) *QosFlowProfile {
	this := QosFlowProfile{}
	this.Var5qi = var5qi
	return &this
}

// NewQosFlowProfileWithDefaults instantiates a new QosFlowProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosFlowProfileWithDefaults() *QosFlowProfile {
	this := QosFlowProfile{}
	return &this
}

// GetVar5qi returns the Var5qi field value
func (o *QosFlowProfile) GetVar5qi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var5qi
}

// GetVar5qiOk returns a tuple with the Var5qi field value
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetVar5qiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var5qi, true
}

// SetVar5qi sets field value
func (o *QosFlowProfile) SetVar5qi(v int32) {
	o.Var5qi = v
}

// GetNonDynamic5Qi returns the NonDynamic5Qi field value if set, zero value otherwise.
func (o *QosFlowProfile) GetNonDynamic5Qi() NonDynamic5Qi {
	if o == nil || IsNil(o.NonDynamic5Qi) {
		var ret NonDynamic5Qi
		return ret
	}
	return *o.NonDynamic5Qi
}

// GetNonDynamic5QiOk returns a tuple with the NonDynamic5Qi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetNonDynamic5QiOk() (*NonDynamic5Qi, bool) {
	if o == nil || IsNil(o.NonDynamic5Qi) {
		return nil, false
	}
	return o.NonDynamic5Qi, true
}

// HasNonDynamic5Qi returns a boolean if a field has been set.
func (o *QosFlowProfile) HasNonDynamic5Qi() bool {
	if o != nil && !IsNil(o.NonDynamic5Qi) {
		return true
	}

	return false
}

// SetNonDynamic5Qi gets a reference to the given NonDynamic5Qi and assigns it to the NonDynamic5Qi field.
func (o *QosFlowProfile) SetNonDynamic5Qi(v NonDynamic5Qi) {
	o.NonDynamic5Qi = &v
}

// GetDynamic5Qi returns the Dynamic5Qi field value if set, zero value otherwise.
func (o *QosFlowProfile) GetDynamic5Qi() Dynamic5Qi {
	if o == nil || IsNil(o.Dynamic5Qi) {
		var ret Dynamic5Qi
		return ret
	}
	return *o.Dynamic5Qi
}

// GetDynamic5QiOk returns a tuple with the Dynamic5Qi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetDynamic5QiOk() (*Dynamic5Qi, bool) {
	if o == nil || IsNil(o.Dynamic5Qi) {
		return nil, false
	}
	return o.Dynamic5Qi, true
}

// HasDynamic5Qi returns a boolean if a field has been set.
func (o *QosFlowProfile) HasDynamic5Qi() bool {
	if o != nil && !IsNil(o.Dynamic5Qi) {
		return true
	}

	return false
}

// SetDynamic5Qi gets a reference to the given Dynamic5Qi and assigns it to the Dynamic5Qi field.
func (o *QosFlowProfile) SetDynamic5Qi(v Dynamic5Qi) {
	o.Dynamic5Qi = &v
}

// GetArp returns the Arp field value if set, zero value otherwise.
func (o *QosFlowProfile) GetArp() Arp {
	if o == nil || IsNil(o.Arp) {
		var ret Arp
		return ret
	}
	return *o.Arp
}

// GetArpOk returns a tuple with the Arp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetArpOk() (*Arp, bool) {
	if o == nil || IsNil(o.Arp) {
		return nil, false
	}
	return o.Arp, true
}

// HasArp returns a boolean if a field has been set.
func (o *QosFlowProfile) HasArp() bool {
	if o != nil && !IsNil(o.Arp) {
		return true
	}

	return false
}

// SetArp gets a reference to the given Arp and assigns it to the Arp field.
func (o *QosFlowProfile) SetArp(v Arp) {
	o.Arp = &v
}

// GetGbrQosFlowInfo returns the GbrQosFlowInfo field value if set, zero value otherwise.
func (o *QosFlowProfile) GetGbrQosFlowInfo() GbrQosFlowInformation {
	if o == nil || IsNil(o.GbrQosFlowInfo) {
		var ret GbrQosFlowInformation
		return ret
	}
	return *o.GbrQosFlowInfo
}

// GetGbrQosFlowInfoOk returns a tuple with the GbrQosFlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetGbrQosFlowInfoOk() (*GbrQosFlowInformation, bool) {
	if o == nil || IsNil(o.GbrQosFlowInfo) {
		return nil, false
	}
	return o.GbrQosFlowInfo, true
}

// HasGbrQosFlowInfo returns a boolean if a field has been set.
func (o *QosFlowProfile) HasGbrQosFlowInfo() bool {
	if o != nil && !IsNil(o.GbrQosFlowInfo) {
		return true
	}

	return false
}

// SetGbrQosFlowInfo gets a reference to the given GbrQosFlowInformation and assigns it to the GbrQosFlowInfo field.
func (o *QosFlowProfile) SetGbrQosFlowInfo(v GbrQosFlowInformation) {
	o.GbrQosFlowInfo = &v
}

// GetRqa returns the Rqa field value if set, zero value otherwise.
func (o *QosFlowProfile) GetRqa() ReflectiveQoSAttribute {
	if o == nil || IsNil(o.Rqa) {
		var ret ReflectiveQoSAttribute
		return ret
	}
	return *o.Rqa
}

// GetRqaOk returns a tuple with the Rqa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetRqaOk() (*ReflectiveQoSAttribute, bool) {
	if o == nil || IsNil(o.Rqa) {
		return nil, false
	}
	return o.Rqa, true
}

// HasRqa returns a boolean if a field has been set.
func (o *QosFlowProfile) HasRqa() bool {
	if o != nil && !IsNil(o.Rqa) {
		return true
	}

	return false
}

// SetRqa gets a reference to the given ReflectiveQoSAttribute and assigns it to the Rqa field.
func (o *QosFlowProfile) SetRqa(v ReflectiveQoSAttribute) {
	o.Rqa = &v
}

// GetAdditionalQosFlowInfo returns the AdditionalQosFlowInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosFlowProfile) GetAdditionalQosFlowInfo() AdditionalQosFlowInfo {
	if o == nil || IsNil(o.AdditionalQosFlowInfo.Get()) {
		var ret AdditionalQosFlowInfo
		return ret
	}
	return *o.AdditionalQosFlowInfo.Get()
}

// GetAdditionalQosFlowInfoOk returns a tuple with the AdditionalQosFlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosFlowProfile) GetAdditionalQosFlowInfoOk() (*AdditionalQosFlowInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalQosFlowInfo.Get(), o.AdditionalQosFlowInfo.IsSet()
}

// HasAdditionalQosFlowInfo returns a boolean if a field has been set.
func (o *QosFlowProfile) HasAdditionalQosFlowInfo() bool {
	if o != nil && o.AdditionalQosFlowInfo.IsSet() {
		return true
	}

	return false
}

// SetAdditionalQosFlowInfo gets a reference to the given NullableAdditionalQosFlowInfo and assigns it to the AdditionalQosFlowInfo field.
func (o *QosFlowProfile) SetAdditionalQosFlowInfo(v AdditionalQosFlowInfo) {
	o.AdditionalQosFlowInfo.Set(&v)
}

// SetAdditionalQosFlowInfoNil sets the value for AdditionalQosFlowInfo to be an explicit nil
func (o *QosFlowProfile) SetAdditionalQosFlowInfoNil() {
	o.AdditionalQosFlowInfo.Set(nil)
}

// UnsetAdditionalQosFlowInfo ensures that no value is present for AdditionalQosFlowInfo, not even an explicit nil
func (o *QosFlowProfile) UnsetAdditionalQosFlowInfo() {
	o.AdditionalQosFlowInfo.Unset()
}

// GetQosMonitoringReq returns the QosMonitoringReq field value if set, zero value otherwise.
func (o *QosFlowProfile) GetQosMonitoringReq() QosMonitoringReq {
	if o == nil || IsNil(o.QosMonitoringReq) {
		var ret QosMonitoringReq
		return ret
	}
	return *o.QosMonitoringReq
}

// GetQosMonitoringReqOk returns a tuple with the QosMonitoringReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetQosMonitoringReqOk() (*QosMonitoringReq, bool) {
	if o == nil || IsNil(o.QosMonitoringReq) {
		return nil, false
	}
	return o.QosMonitoringReq, true
}

// HasQosMonitoringReq returns a boolean if a field has been set.
func (o *QosFlowProfile) HasQosMonitoringReq() bool {
	if o != nil && !IsNil(o.QosMonitoringReq) {
		return true
	}

	return false
}

// SetQosMonitoringReq gets a reference to the given QosMonitoringReq and assigns it to the QosMonitoringReq field.
func (o *QosFlowProfile) SetQosMonitoringReq(v QosMonitoringReq) {
	o.QosMonitoringReq = &v
}

// GetQosRepPeriod returns the QosRepPeriod field value if set, zero value otherwise.
func (o *QosFlowProfile) GetQosRepPeriod() int32 {
	if o == nil || IsNil(o.QosRepPeriod) {
		var ret int32
		return ret
	}
	return *o.QosRepPeriod
}

// GetQosRepPeriodOk returns a tuple with the QosRepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetQosRepPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.QosRepPeriod) {
		return nil, false
	}
	return o.QosRepPeriod, true
}

// HasQosRepPeriod returns a boolean if a field has been set.
func (o *QosFlowProfile) HasQosRepPeriod() bool {
	if o != nil && !IsNil(o.QosRepPeriod) {
		return true
	}

	return false
}

// SetQosRepPeriod gets a reference to the given int32 and assigns it to the QosRepPeriod field.
func (o *QosFlowProfile) SetQosRepPeriod(v int32) {
	o.QosRepPeriod = &v
}

// GetPduSetQosDl returns the PduSetQosDl field value if set, zero value otherwise.
func (o *QosFlowProfile) GetPduSetQosDl() PduSetQosPara {
	if o == nil || IsNil(o.PduSetQosDl) {
		var ret PduSetQosPara
		return ret
	}
	return *o.PduSetQosDl
}

// GetPduSetQosDlOk returns a tuple with the PduSetQosDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetPduSetQosDlOk() (*PduSetQosPara, bool) {
	if o == nil || IsNil(o.PduSetQosDl) {
		return nil, false
	}
	return o.PduSetQosDl, true
}

// HasPduSetQosDl returns a boolean if a field has been set.
func (o *QosFlowProfile) HasPduSetQosDl() bool {
	if o != nil && !IsNil(o.PduSetQosDl) {
		return true
	}

	return false
}

// SetPduSetQosDl gets a reference to the given PduSetQosPara and assigns it to the PduSetQosDl field.
func (o *QosFlowProfile) SetPduSetQosDl(v PduSetQosPara) {
	o.PduSetQosDl = &v
}

// GetPduSetQosUl returns the PduSetQosUl field value if set, zero value otherwise.
func (o *QosFlowProfile) GetPduSetQosUl() PduSetQosPara {
	if o == nil || IsNil(o.PduSetQosUl) {
		var ret PduSetQosPara
		return ret
	}
	return *o.PduSetQosUl
}

// GetPduSetQosUlOk returns a tuple with the PduSetQosUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetPduSetQosUlOk() (*PduSetQosPara, bool) {
	if o == nil || IsNil(o.PduSetQosUl) {
		return nil, false
	}
	return o.PduSetQosUl, true
}

// HasPduSetQosUl returns a boolean if a field has been set.
func (o *QosFlowProfile) HasPduSetQosUl() bool {
	if o != nil && !IsNil(o.PduSetQosUl) {
		return true
	}

	return false
}

// SetPduSetQosUl gets a reference to the given PduSetQosPara and assigns it to the PduSetQosUl field.
func (o *QosFlowProfile) SetPduSetQosUl(v PduSetQosPara) {
	o.PduSetQosUl = &v
}

func (o QosFlowProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["5qi"] = o.Var5qi
	if !IsNil(o.NonDynamic5Qi) {
		toSerialize["nonDynamic5Qi"] = o.NonDynamic5Qi
	}
	if !IsNil(o.Dynamic5Qi) {
		toSerialize["dynamic5Qi"] = o.Dynamic5Qi
	}
	if !IsNil(o.Arp) {
		toSerialize["arp"] = o.Arp
	}
	if !IsNil(o.GbrQosFlowInfo) {
		toSerialize["gbrQosFlowInfo"] = o.GbrQosFlowInfo
	}
	if !IsNil(o.Rqa) {
		toSerialize["rqa"] = o.Rqa
	}
	if o.AdditionalQosFlowInfo.IsSet() {
		toSerialize["additionalQosFlowInfo"] = o.AdditionalQosFlowInfo.Get()
	}
	if !IsNil(o.QosMonitoringReq) {
		toSerialize["qosMonitoringReq"] = o.QosMonitoringReq
	}
	if !IsNil(o.QosRepPeriod) {
		toSerialize["qosRepPeriod"] = o.QosRepPeriod
	}
	if !IsNil(o.PduSetQosDl) {
		toSerialize["pduSetQosDl"] = o.PduSetQosDl
	}
	if !IsNil(o.PduSetQosUl) {
		toSerialize["pduSetQosUl"] = o.PduSetQosUl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableQosFlowProfile struct {
	value *QosFlowProfile
	isSet bool
}

func (v NullableQosFlowProfile) Get() *QosFlowProfile {
	return v.value
}

func (v *NullableQosFlowProfile) Set(val *QosFlowProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableQosFlowProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableQosFlowProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosFlowProfile(val *QosFlowProfile) *NullableQosFlowProfile {
	return &NullableQosFlowProfile{value: val, isSet: true}
}

func (v NullableQosFlowProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosFlowProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}