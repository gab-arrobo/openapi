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

// checks if the UeRegStatusUpdateRspData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeRegStatusUpdateRspData{}

// UeRegStatusUpdateRspData Data within a UE registration status update response to provide the status of UE context transfer status update at a source AMF
type UeRegStatusUpdateRspData struct {
	RegStatusTransferComplete bool `json:"regStatusTransferComplete"`
	AdditionalProperties      map[string]interface{}
}

type _UeRegStatusUpdateRspData UeRegStatusUpdateRspData

// NewUeRegStatusUpdateRspData instantiates a new UeRegStatusUpdateRspData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeRegStatusUpdateRspData(regStatusTransferComplete bool) *UeRegStatusUpdateRspData {
	this := UeRegStatusUpdateRspData{}
	this.RegStatusTransferComplete = regStatusTransferComplete
	return &this
}

// NewUeRegStatusUpdateRspDataWithDefaults instantiates a new UeRegStatusUpdateRspData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeRegStatusUpdateRspDataWithDefaults() *UeRegStatusUpdateRspData {
	this := UeRegStatusUpdateRspData{}
	return &this
}

// GetRegStatusTransferComplete returns the RegStatusTransferComplete field value
func (o *UeRegStatusUpdateRspData) GetRegStatusTransferComplete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RegStatusTransferComplete
}

// GetRegStatusTransferCompleteOk returns a tuple with the RegStatusTransferComplete field value
// and a boolean to check if the value has been set.
func (o *UeRegStatusUpdateRspData) GetRegStatusTransferCompleteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegStatusTransferComplete, true
}

// SetRegStatusTransferComplete sets field value
func (o *UeRegStatusUpdateRspData) SetRegStatusTransferComplete(v bool) {
	o.RegStatusTransferComplete = v
}

func (o UeRegStatusUpdateRspData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["regStatusTransferComplete"] = o.RegStatusTransferComplete

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableUeRegStatusUpdateRspData struct {
	value *UeRegStatusUpdateRspData
	isSet bool
}

func (v NullableUeRegStatusUpdateRspData) Get() *UeRegStatusUpdateRspData {
	return v.value
}

func (v *NullableUeRegStatusUpdateRspData) Set(val *UeRegStatusUpdateRspData) {
	v.value = val
	v.isSet = true
}

func (v NullableUeRegStatusUpdateRspData) IsSet() bool {
	return v.isSet
}

func (v *NullableUeRegStatusUpdateRspData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeRegStatusUpdateRspData(val *UeRegStatusUpdateRspData) *NullableUeRegStatusUpdateRspData {
	return &NullableUeRegStatusUpdateRspData{value: val, isSet: true}
}

func (v NullableUeRegStatusUpdateRspData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeRegStatusUpdateRspData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}