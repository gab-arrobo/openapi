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
	"os"
)

// checks if the ReleaseSmContextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseSmContextRequest{}

// ReleaseSmContextRequest struct for ReleaseSmContextRequest
type ReleaseSmContextRequest struct {
	JsonData                  *SmContextReleaseData `json:"jsonData,omitempty"`
	BinaryDataN2SmInformation **os.File             `json:"binaryDataN2SmInformation,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _ReleaseSmContextRequest ReleaseSmContextRequest

// NewReleaseSmContextRequest instantiates a new ReleaseSmContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseSmContextRequest() *ReleaseSmContextRequest {
	this := ReleaseSmContextRequest{}
	return &this
}

// NewReleaseSmContextRequestWithDefaults instantiates a new ReleaseSmContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseSmContextRequestWithDefaults() *ReleaseSmContextRequest {
	this := ReleaseSmContextRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *ReleaseSmContextRequest) GetJsonData() SmContextReleaseData {
	if o == nil || IsNil(o.JsonData) {
		var ret SmContextReleaseData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseSmContextRequest) GetJsonDataOk() (*SmContextReleaseData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *ReleaseSmContextRequest) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given SmContextReleaseData and assigns it to the JsonData field.
func (o *ReleaseSmContextRequest) SetJsonData(v SmContextReleaseData) {
	o.JsonData = &v
}

// GetBinaryDataN2SmInformation returns the BinaryDataN2SmInformation field value if set, zero value otherwise.
func (o *ReleaseSmContextRequest) GetBinaryDataN2SmInformation() *os.File {
	if o == nil || IsNil(o.BinaryDataN2SmInformation) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2SmInformation
}

// GetBinaryDataN2SmInformationOk returns a tuple with the BinaryDataN2SmInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseSmContextRequest) GetBinaryDataN2SmInformationOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2SmInformation) {
		return nil, false
	}
	return o.BinaryDataN2SmInformation, true
}

// HasBinaryDataN2SmInformation returns a boolean if a field has been set.
func (o *ReleaseSmContextRequest) HasBinaryDataN2SmInformation() bool {
	if o != nil && !IsNil(o.BinaryDataN2SmInformation) {
		return true
	}

	return false
}

// SetBinaryDataN2SmInformation gets a reference to the given *os.File and assigns it to the BinaryDataN2SmInformation field.
func (o *ReleaseSmContextRequest) SetBinaryDataN2SmInformation(v *os.File) {
	o.BinaryDataN2SmInformation = &v
}

func (o ReleaseSmContextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryDataN2SmInformation) {
		toSerialize["binaryDataN2SmInformation"] = o.BinaryDataN2SmInformation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableReleaseSmContextRequest struct {
	value *ReleaseSmContextRequest
	isSet bool
}

func (v NullableReleaseSmContextRequest) Get() *ReleaseSmContextRequest {
	return v.value
}

func (v *NullableReleaseSmContextRequest) Set(val *ReleaseSmContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseSmContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseSmContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseSmContextRequest(val *ReleaseSmContextRequest) *NullableReleaseSmContextRequest {
	return &NullableReleaseSmContextRequest{value: val, isSet: true}
}

func (v NullableReleaseSmContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseSmContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}