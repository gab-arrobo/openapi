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

// checks if the AnchorSmfFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnchorSmfFeatures{}

// AnchorSmfFeatures Anchor SMF supported features
type AnchorSmfFeatures struct {
	PsetrSupportInd      *bool `json:"psetrSupportInd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AnchorSmfFeatures AnchorSmfFeatures

// NewAnchorSmfFeatures instantiates a new AnchorSmfFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnchorSmfFeatures() *AnchorSmfFeatures {
	this := AnchorSmfFeatures{}
	return &this
}

// NewAnchorSmfFeaturesWithDefaults instantiates a new AnchorSmfFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnchorSmfFeaturesWithDefaults() *AnchorSmfFeatures {
	this := AnchorSmfFeatures{}
	return &this
}

// GetPsetrSupportInd returns the PsetrSupportInd field value if set, zero value otherwise.
func (o *AnchorSmfFeatures) GetPsetrSupportInd() bool {
	if o == nil || IsNil(o.PsetrSupportInd) {
		var ret bool
		return ret
	}
	return *o.PsetrSupportInd
}

// GetPsetrSupportIndOk returns a tuple with the PsetrSupportInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchorSmfFeatures) GetPsetrSupportIndOk() (*bool, bool) {
	if o == nil || IsNil(o.PsetrSupportInd) {
		return nil, false
	}
	return o.PsetrSupportInd, true
}

// HasPsetrSupportInd returns a boolean if a field has been set.
func (o *AnchorSmfFeatures) HasPsetrSupportInd() bool {
	if o != nil && !IsNil(o.PsetrSupportInd) {
		return true
	}

	return false
}

// SetPsetrSupportInd gets a reference to the given bool and assigns it to the PsetrSupportInd field.
func (o *AnchorSmfFeatures) SetPsetrSupportInd(v bool) {
	o.PsetrSupportInd = &v
}

func (o AnchorSmfFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PsetrSupportInd) {
		toSerialize["psetrSupportInd"] = o.PsetrSupportInd
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableAnchorSmfFeatures struct {
	value *AnchorSmfFeatures
	isSet bool
}

func (v NullableAnchorSmfFeatures) Get() *AnchorSmfFeatures {
	return v.value
}

func (v *NullableAnchorSmfFeatures) Set(val *AnchorSmfFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableAnchorSmfFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableAnchorSmfFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnchorSmfFeatures(val *AnchorSmfFeatures) *NullableAnchorSmfFeatures {
	return &NullableAnchorSmfFeatures{value: val, isSet: true}
}

func (v NullableAnchorSmfFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnchorSmfFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}