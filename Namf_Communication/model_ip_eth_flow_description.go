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

// checks if the IpEthFlowDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpEthFlowDescription{}

// IpEthFlowDescription Contains the description of an Uplink and/or Downlink Ethernet flow.
type IpEthFlowDescription struct {
	// Defines a packet filter of an IP flow.
	IpTrafficFilter      *string
	EthTrafficFilter     *EthFlowDescription
	AdditionalProperties map[string]interface{}
}

type _IpEthFlowDescription IpEthFlowDescription

// NewIpEthFlowDescription instantiates a new IpEthFlowDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpEthFlowDescription() *IpEthFlowDescription {
	this := IpEthFlowDescription{}
	return &this
}

// NewIpEthFlowDescriptionWithDefaults instantiates a new IpEthFlowDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpEthFlowDescriptionWithDefaults() *IpEthFlowDescription {
	this := IpEthFlowDescription{}
	return &this
}

// GetIpTrafficFilter returns the IpTrafficFilter field value if set, zero value otherwise.
func (o *IpEthFlowDescription) GetIpTrafficFilter() string {
	if o == nil || IsNil(o.IpTrafficFilter) {
		var ret string
		return ret
	}
	return *o.IpTrafficFilter
}

// GetIpTrafficFilterOk returns a tuple with the IpTrafficFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEthFlowDescription) GetIpTrafficFilterOk() (*string, bool) {
	if o == nil || IsNil(o.IpTrafficFilter) {
		return nil, false
	}
	return o.IpTrafficFilter, true
}

// HasIpTrafficFilter returns a boolean if a field has been set.
func (o *IpEthFlowDescription) HasIpTrafficFilter() bool {
	if o != nil && !IsNil(o.IpTrafficFilter) {
		return true
	}

	return false
}

// SetIpTrafficFilter gets a reference to the given string and assigns it to the IpTrafficFilter field.
func (o *IpEthFlowDescription) SetIpTrafficFilter(v string) {
	o.IpTrafficFilter = &v
}

// GetEthTrafficFilter returns the EthTrafficFilter field value if set, zero value otherwise.
func (o *IpEthFlowDescription) GetEthTrafficFilter() EthFlowDescription {
	if o == nil || IsNil(o.EthTrafficFilter) {
		var ret EthFlowDescription
		return ret
	}
	return *o.EthTrafficFilter
}

// GetEthTrafficFilterOk returns a tuple with the EthTrafficFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpEthFlowDescription) GetEthTrafficFilterOk() (*EthFlowDescription, bool) {
	if o == nil || IsNil(o.EthTrafficFilter) {
		return nil, false
	}
	return o.EthTrafficFilter, true
}

// HasEthTrafficFilter returns a boolean if a field has been set.
func (o *IpEthFlowDescription) HasEthTrafficFilter() bool {
	if o != nil && !IsNil(o.EthTrafficFilter) {
		return true
	}

	return false
}

// SetEthTrafficFilter gets a reference to the given EthFlowDescription and assigns it to the EthTrafficFilter field.
func (o *IpEthFlowDescription) SetEthTrafficFilter(v EthFlowDescription) {
	o.EthTrafficFilter = &v
}

func (o IpEthFlowDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpTrafficFilter) {
		toSerialize["ipTrafficFilter"] = o.IpTrafficFilter
	}
	if !IsNil(o.EthTrafficFilter) {
		toSerialize["ethTrafficFilter"] = o.EthTrafficFilter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableIpEthFlowDescription struct {
	value *IpEthFlowDescription
	isSet bool
}

func (v NullableIpEthFlowDescription) Get() *IpEthFlowDescription {
	return v.value
}

func (v *NullableIpEthFlowDescription) Set(val *IpEthFlowDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableIpEthFlowDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableIpEthFlowDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpEthFlowDescription(val *IpEthFlowDescription) *NullableIpEthFlowDescription {
	return &NullableIpEthFlowDescription{value: val, isSet: true}
}

func (v NullableIpEthFlowDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpEthFlowDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}