// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
Npcf_SMPolicyControl API

Session Management Policy Control Service
© 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).
All rights reserved.

API version: 1.3.0-alpha.6

Authors: Aether SD-Core team
Contact: dev@lists.aetherproject.org

Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
*/

package Npcf_SMPolicyControl

import (
	"encoding/json"
)

// checks if the AnGwAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnGwAddress{}

// AnGwAddress Describes the address of the access network gateway control node.
type AnGwAddress struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	AnGwIpv4Addr         *string
	AnGwIpv6Addr         *Ipv6Addr
	AdditionalProperties map[string]interface{}
}

type _AnGwAddress AnGwAddress

// NewAnGwAddress instantiates a new AnGwAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnGwAddress() *AnGwAddress {
	this := AnGwAddress{}
	return &this
}

// NewAnGwAddressWithDefaults instantiates a new AnGwAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnGwAddressWithDefaults() *AnGwAddress {
	this := AnGwAddress{}
	return &this
}

// GetAnGwIpv4Addr returns the AnGwIpv4Addr field value if set, zero value otherwise.
func (o *AnGwAddress) GetAnGwIpv4Addr() string {
	if o == nil || IsNil(o.AnGwIpv4Addr) {
		var ret string
		return ret
	}
	return *o.AnGwIpv4Addr
}

// GetAnGwIpv4AddrOk returns a tuple with the AnGwIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnGwAddress) GetAnGwIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.AnGwIpv4Addr) {
		return nil, false
	}
	return o.AnGwIpv4Addr, true
}

// HasAnGwIpv4Addr returns a boolean if a field has been set.
func (o *AnGwAddress) HasAnGwIpv4Addr() bool {
	if o != nil && !IsNil(o.AnGwIpv4Addr) {
		return true
	}

	return false
}

// SetAnGwIpv4Addr gets a reference to the given string and assigns it to the AnGwIpv4Addr field.
func (o *AnGwAddress) SetAnGwIpv4Addr(v string) {
	o.AnGwIpv4Addr = &v
}

// GetAnGwIpv6Addr returns the AnGwIpv6Addr field value if set, zero value otherwise.
func (o *AnGwAddress) GetAnGwIpv6Addr() Ipv6Addr {
	if o == nil || IsNil(o.AnGwIpv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.AnGwIpv6Addr
}

// GetAnGwIpv6AddrOk returns a tuple with the AnGwIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnGwAddress) GetAnGwIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.AnGwIpv6Addr) {
		return nil, false
	}
	return o.AnGwIpv6Addr, true
}

// HasAnGwIpv6Addr returns a boolean if a field has been set.
func (o *AnGwAddress) HasAnGwIpv6Addr() bool {
	if o != nil && !IsNil(o.AnGwIpv6Addr) {
		return true
	}

	return false
}

// SetAnGwIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the AnGwIpv6Addr field.
func (o *AnGwAddress) SetAnGwIpv6Addr(v Ipv6Addr) {
	o.AnGwIpv6Addr = &v
}

func (o AnGwAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnGwIpv4Addr) {
		toSerialize["anGwIpv4Addr"] = o.AnGwIpv4Addr
	}
	if !IsNil(o.AnGwIpv6Addr) {
		toSerialize["anGwIpv6Addr"] = o.AnGwIpv6Addr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableAnGwAddress struct {
	value *AnGwAddress
	isSet bool
}

func (v NullableAnGwAddress) Get() *AnGwAddress {
	return v.value
}

func (v *NullableAnGwAddress) Set(val *AnGwAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAnGwAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAnGwAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnGwAddress(val *AnGwAddress) *NullableAnGwAddress {
	return &NullableAnGwAddress{value: val, isSet: true}
}

func (v NullableAnGwAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnGwAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}