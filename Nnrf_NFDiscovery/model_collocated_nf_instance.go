// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
NRF NFDiscovery Service

NRF NFDiscovery Service.
© 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).
All rights reserved.

API version: 1.3.0-alpha.7

Authors: Aether SD-Core team
Contact: dev@lists.aetherproject.org

Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
*/

package Nnrf_NFDiscovery

import (
	"encoding/json"
)

// checks if the CollocatedNfInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollocatedNfInstance{}

// CollocatedNfInstance Information of an collocated NF Instance registered in the NRF
type CollocatedNfInstance struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfInstanceId         string           `json:"nfInstanceId"`
	NfType               CollocatedNfType `json:"nfType"`
	AdditionalProperties map[string]interface{}
}

type _CollocatedNfInstance CollocatedNfInstance

// NewCollocatedNfInstance instantiates a new CollocatedNfInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollocatedNfInstance(nfInstanceId string, nfType CollocatedNfType) *CollocatedNfInstance {
	this := CollocatedNfInstance{}
	this.NfInstanceId = nfInstanceId
	this.NfType = nfType
	return &this
}

// NewCollocatedNfInstanceWithDefaults instantiates a new CollocatedNfInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollocatedNfInstanceWithDefaults() *CollocatedNfInstance {
	this := CollocatedNfInstance{}
	return &this
}

// GetNfInstanceId returns the NfInstanceId field value
func (o *CollocatedNfInstance) GetNfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfInstanceId
}

// GetNfInstanceIdOk returns a tuple with the NfInstanceId field value
// and a boolean to check if the value has been set.
func (o *CollocatedNfInstance) GetNfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfInstanceId, true
}

// SetNfInstanceId sets field value
func (o *CollocatedNfInstance) SetNfInstanceId(v string) {
	o.NfInstanceId = v
}

// GetNfType returns the NfType field value
func (o *CollocatedNfInstance) GetNfType() CollocatedNfType {
	if o == nil {
		var ret CollocatedNfType
		return ret
	}

	return o.NfType
}

// GetNfTypeOk returns a tuple with the NfType field value
// and a boolean to check if the value has been set.
func (o *CollocatedNfInstance) GetNfTypeOk() (*CollocatedNfType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfType, true
}

// SetNfType sets field value
func (o *CollocatedNfInstance) SetNfType(v CollocatedNfType) {
	o.NfType = v
}

func (o CollocatedNfInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nfInstanceId"] = o.NfInstanceId
	toSerialize["nfType"] = o.NfType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableCollocatedNfInstance struct {
	value *CollocatedNfInstance
	isSet bool
}

func (v NullableCollocatedNfInstance) Get() *CollocatedNfInstance {
	return v.value
}

func (v *NullableCollocatedNfInstance) Set(val *CollocatedNfInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableCollocatedNfInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableCollocatedNfInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollocatedNfInstance(val *CollocatedNfInstance) *NullableCollocatedNfInstance {
	return &NullableCollocatedNfInstance{value: val, isSet: true}
}

func (v NullableCollocatedNfInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollocatedNfInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}