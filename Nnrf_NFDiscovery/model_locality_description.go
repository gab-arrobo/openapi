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

// checks if the LocalityDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalityDescription{}

// LocalityDescription Locality description
type LocalityDescription struct {
	LocalityType         LocalityType              `json:"localityType"`
	LocalityValue        string                    `json:"localityValue"`
	AddlLocDescrItems    []LocalityDescriptionItem `json:"addlLocDescrItems,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LocalityDescription LocalityDescription

// NewLocalityDescription instantiates a new LocalityDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalityDescription(localityType LocalityType, localityValue string) *LocalityDescription {
	this := LocalityDescription{}
	this.LocalityType = localityType
	this.LocalityValue = localityValue
	return &this
}

// NewLocalityDescriptionWithDefaults instantiates a new LocalityDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalityDescriptionWithDefaults() *LocalityDescription {
	this := LocalityDescription{}
	return &this
}

// GetLocalityType returns the LocalityType field value
func (o *LocalityDescription) GetLocalityType() LocalityType {
	if o == nil {
		var ret LocalityType
		return ret
	}

	return o.LocalityType
}

// GetLocalityTypeOk returns a tuple with the LocalityType field value
// and a boolean to check if the value has been set.
func (o *LocalityDescription) GetLocalityTypeOk() (*LocalityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalityType, true
}

// SetLocalityType sets field value
func (o *LocalityDescription) SetLocalityType(v LocalityType) {
	o.LocalityType = v
}

// GetLocalityValue returns the LocalityValue field value
func (o *LocalityDescription) GetLocalityValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocalityValue
}

// GetLocalityValueOk returns a tuple with the LocalityValue field value
// and a boolean to check if the value has been set.
func (o *LocalityDescription) GetLocalityValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalityValue, true
}

// SetLocalityValue sets field value
func (o *LocalityDescription) SetLocalityValue(v string) {
	o.LocalityValue = v
}

// GetAddlLocDescrItems returns the AddlLocDescrItems field value if set, zero value otherwise.
func (o *LocalityDescription) GetAddlLocDescrItems() []LocalityDescriptionItem {
	if o == nil || IsNil(o.AddlLocDescrItems) {
		var ret []LocalityDescriptionItem
		return ret
	}
	return o.AddlLocDescrItems
}

// GetAddlLocDescrItemsOk returns a tuple with the AddlLocDescrItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalityDescription) GetAddlLocDescrItemsOk() ([]LocalityDescriptionItem, bool) {
	if o == nil || IsNil(o.AddlLocDescrItems) {
		return nil, false
	}
	return o.AddlLocDescrItems, true
}

// HasAddlLocDescrItems returns a boolean if a field has been set.
func (o *LocalityDescription) HasAddlLocDescrItems() bool {
	if o != nil && !IsNil(o.AddlLocDescrItems) {
		return true
	}

	return false
}

// SetAddlLocDescrItems gets a reference to the given []LocalityDescriptionItem and assigns it to the AddlLocDescrItems field.
func (o *LocalityDescription) SetAddlLocDescrItems(v []LocalityDescriptionItem) {
	o.AddlLocDescrItems = v
}

func (o LocalityDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["localityType"] = o.LocalityType
	toSerialize["localityValue"] = o.LocalityValue
	if !IsNil(o.AddlLocDescrItems) {
		toSerialize["addlLocDescrItems"] = o.AddlLocDescrItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableLocalityDescription struct {
	value *LocalityDescription
	isSet bool
}

func (v NullableLocalityDescription) Get() *LocalityDescription {
	return v.value
}

func (v *NullableLocalityDescription) Set(val *LocalityDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalityDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalityDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalityDescription(val *LocalityDescription) *NullableLocalityDescription {
	return &NullableLocalityDescription{value: val, isSet: true}
}

func (v NullableLocalityDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalityDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}