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

// checks if the ReleasePduSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleasePduSessionRequest{}

// ReleasePduSessionRequest struct for ReleasePduSessionRequest
type ReleasePduSessionRequest struct {
	JsonData                    *ReleaseData `json:"jsonData,omitempty"`
	BinaryDataN4Information     **os.File    `json:"binaryDataN4Information,omitempty"`
	BinaryDataN4InformationExt1 **os.File    `json:"binaryDataN4InformationExt1,omitempty"`
	BinaryDataN4InformationExt2 **os.File    `json:"binaryDataN4InformationExt2,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _ReleasePduSessionRequest ReleasePduSessionRequest

// NewReleasePduSessionRequest instantiates a new ReleasePduSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleasePduSessionRequest() *ReleasePduSessionRequest {
	this := ReleasePduSessionRequest{}
	return &this
}

// NewReleasePduSessionRequestWithDefaults instantiates a new ReleasePduSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleasePduSessionRequestWithDefaults() *ReleasePduSessionRequest {
	this := ReleasePduSessionRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *ReleasePduSessionRequest) GetJsonData() ReleaseData {
	if o == nil || IsNil(o.JsonData) {
		var ret ReleaseData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasePduSessionRequest) GetJsonDataOk() (*ReleaseData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *ReleasePduSessionRequest) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given ReleaseData and assigns it to the JsonData field.
func (o *ReleasePduSessionRequest) SetJsonData(v ReleaseData) {
	o.JsonData = &v
}

// GetBinaryDataN4Information returns the BinaryDataN4Information field value if set, zero value otherwise.
func (o *ReleasePduSessionRequest) GetBinaryDataN4Information() *os.File {
	if o == nil || IsNil(o.BinaryDataN4Information) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN4Information
}

// GetBinaryDataN4InformationOk returns a tuple with the BinaryDataN4Information field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasePduSessionRequest) GetBinaryDataN4InformationOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN4Information) {
		return nil, false
	}
	return o.BinaryDataN4Information, true
}

// HasBinaryDataN4Information returns a boolean if a field has been set.
func (o *ReleasePduSessionRequest) HasBinaryDataN4Information() bool {
	if o != nil && !IsNil(o.BinaryDataN4Information) {
		return true
	}

	return false
}

// SetBinaryDataN4Information gets a reference to the given *os.File and assigns it to the BinaryDataN4Information field.
func (o *ReleasePduSessionRequest) SetBinaryDataN4Information(v *os.File) {
	o.BinaryDataN4Information = &v
}

// GetBinaryDataN4InformationExt1 returns the BinaryDataN4InformationExt1 field value if set, zero value otherwise.
func (o *ReleasePduSessionRequest) GetBinaryDataN4InformationExt1() *os.File {
	if o == nil || IsNil(o.BinaryDataN4InformationExt1) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN4InformationExt1
}

// GetBinaryDataN4InformationExt1Ok returns a tuple with the BinaryDataN4InformationExt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasePduSessionRequest) GetBinaryDataN4InformationExt1Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN4InformationExt1) {
		return nil, false
	}
	return o.BinaryDataN4InformationExt1, true
}

// HasBinaryDataN4InformationExt1 returns a boolean if a field has been set.
func (o *ReleasePduSessionRequest) HasBinaryDataN4InformationExt1() bool {
	if o != nil && !IsNil(o.BinaryDataN4InformationExt1) {
		return true
	}

	return false
}

// SetBinaryDataN4InformationExt1 gets a reference to the given *os.File and assigns it to the BinaryDataN4InformationExt1 field.
func (o *ReleasePduSessionRequest) SetBinaryDataN4InformationExt1(v *os.File) {
	o.BinaryDataN4InformationExt1 = &v
}

// GetBinaryDataN4InformationExt2 returns the BinaryDataN4InformationExt2 field value if set, zero value otherwise.
func (o *ReleasePduSessionRequest) GetBinaryDataN4InformationExt2() *os.File {
	if o == nil || IsNil(o.BinaryDataN4InformationExt2) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN4InformationExt2
}

// GetBinaryDataN4InformationExt2Ok returns a tuple with the BinaryDataN4InformationExt2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasePduSessionRequest) GetBinaryDataN4InformationExt2Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN4InformationExt2) {
		return nil, false
	}
	return o.BinaryDataN4InformationExt2, true
}

// HasBinaryDataN4InformationExt2 returns a boolean if a field has been set.
func (o *ReleasePduSessionRequest) HasBinaryDataN4InformationExt2() bool {
	if o != nil && !IsNil(o.BinaryDataN4InformationExt2) {
		return true
	}

	return false
}

// SetBinaryDataN4InformationExt2 gets a reference to the given *os.File and assigns it to the BinaryDataN4InformationExt2 field.
func (o *ReleasePduSessionRequest) SetBinaryDataN4InformationExt2(v *os.File) {
	o.BinaryDataN4InformationExt2 = &v
}

func (o ReleasePduSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryDataN4Information) {
		toSerialize["binaryDataN4Information"] = o.BinaryDataN4Information
	}
	if !IsNil(o.BinaryDataN4InformationExt1) {
		toSerialize["binaryDataN4InformationExt1"] = o.BinaryDataN4InformationExt1
	}
	if !IsNil(o.BinaryDataN4InformationExt2) {
		toSerialize["binaryDataN4InformationExt2"] = o.BinaryDataN4InformationExt2
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableReleasePduSessionRequest struct {
	value *ReleasePduSessionRequest
	isSet bool
}

func (v NullableReleasePduSessionRequest) Get() *ReleasePduSessionRequest {
	return v.value
}

func (v *NullableReleasePduSessionRequest) Set(val *ReleasePduSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReleasePduSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReleasePduSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleasePduSessionRequest(val *ReleasePduSessionRequest) *NullableReleasePduSessionRequest {
	return &NullableReleasePduSessionRequest{value: val, isSet: true}
}

func (v NullableReleasePduSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleasePduSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}