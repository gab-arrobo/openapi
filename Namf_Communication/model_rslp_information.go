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

// checks if the RslpInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RslpInformation{}

// RslpInformation Represents Ranging/SL positioning related N2 information.
type RslpInformation struct {
	N2Pc5RslpPol         *N2InfoContent `json:"n2Pc5RslpPol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RslpInformation RslpInformation

// NewRslpInformation instantiates a new RslpInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRslpInformation() *RslpInformation {
	this := RslpInformation{}
	return &this
}

// NewRslpInformationWithDefaults instantiates a new RslpInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRslpInformationWithDefaults() *RslpInformation {
	this := RslpInformation{}
	return &this
}

// GetN2Pc5RslpPol returns the N2Pc5RslpPol field value if set, zero value otherwise.
func (o *RslpInformation) GetN2Pc5RslpPol() N2InfoContent {
	if o == nil || IsNil(o.N2Pc5RslpPol) {
		var ret N2InfoContent
		return ret
	}
	return *o.N2Pc5RslpPol
}

// GetN2Pc5RslpPolOk returns a tuple with the N2Pc5RslpPol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RslpInformation) GetN2Pc5RslpPolOk() (*N2InfoContent, bool) {
	if o == nil || IsNil(o.N2Pc5RslpPol) {
		return nil, false
	}
	return o.N2Pc5RslpPol, true
}

// HasN2Pc5RslpPol returns a boolean if a field has been set.
func (o *RslpInformation) HasN2Pc5RslpPol() bool {
	if o != nil && !IsNil(o.N2Pc5RslpPol) {
		return true
	}

	return false
}

// SetN2Pc5RslpPol gets a reference to the given N2InfoContent and assigns it to the N2Pc5RslpPol field.
func (o *RslpInformation) SetN2Pc5RslpPol(v N2InfoContent) {
	o.N2Pc5RslpPol = &v
}

func (o RslpInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.N2Pc5RslpPol) {
		toSerialize["n2Pc5RslpPol"] = o.N2Pc5RslpPol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableRslpInformation struct {
	value *RslpInformation
	isSet bool
}

func (v NullableRslpInformation) Get() *RslpInformation {
	return v.value
}

func (v *NullableRslpInformation) Set(val *RslpInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableRslpInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableRslpInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRslpInformation(val *RslpInformation) *NullableRslpInformation {
	return &NullableRslpInformation{value: val, isSet: true}
}

func (v NullableRslpInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRslpInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}