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

// checks if the AddrFqdn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddrFqdn{}

// AddrFqdn IP address and/or FQDN.
type AddrFqdn struct {
	IpAddr NullableIpAddr `json:"ipAddr,omitempty"`
	// Indicates an FQDN.
	Fqdn                 *string `json:"fqdn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AddrFqdn AddrFqdn

// NewAddrFqdn instantiates a new AddrFqdn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddrFqdn() *AddrFqdn {
	this := AddrFqdn{}
	return &this
}

// NewAddrFqdnWithDefaults instantiates a new AddrFqdn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddrFqdnWithDefaults() *AddrFqdn {
	this := AddrFqdn{}
	return &this
}

// GetIpAddr returns the IpAddr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddrFqdn) GetIpAddr() IpAddr {
	if o == nil || IsNil(o.IpAddr.Get()) {
		var ret IpAddr
		return ret
	}
	return *o.IpAddr.Get()
}

// GetIpAddrOk returns a tuple with the IpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddrFqdn) GetIpAddrOk() (*IpAddr, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddr.Get(), o.IpAddr.IsSet()
}

// HasIpAddr returns a boolean if a field has been set.
func (o *AddrFqdn) HasIpAddr() bool {
	if o != nil && o.IpAddr.IsSet() {
		return true
	}

	return false
}

// SetIpAddr gets a reference to the given NullableIpAddr and assigns it to the IpAddr field.
func (o *AddrFqdn) SetIpAddr(v IpAddr) {
	o.IpAddr.Set(&v)
}

// SetIpAddrNil sets the value for IpAddr to be an explicit nil
func (o *AddrFqdn) SetIpAddrNil() {
	o.IpAddr.Set(nil)
}

// UnsetIpAddr ensures that no value is present for IpAddr, not even an explicit nil
func (o *AddrFqdn) UnsetIpAddr() {
	o.IpAddr.Unset()
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *AddrFqdn) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddrFqdn) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *AddrFqdn) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *AddrFqdn) SetFqdn(v string) {
	o.Fqdn = &v
}

func (o AddrFqdn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAddr.IsSet() {
		toSerialize["ipAddr"] = o.IpAddr.Get()
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableAddrFqdn struct {
	value *AddrFqdn
	isSet bool
}

func (v NullableAddrFqdn) Get() *AddrFqdn {
	return v.value
}

func (v *NullableAddrFqdn) Set(val *AddrFqdn) {
	v.value = val
	v.isSet = true
}

func (v NullableAddrFqdn) IsSet() bool {
	return v.isSet
}

func (v *NullableAddrFqdn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddrFqdn(val *AddrFqdn) *NullableAddrFqdn {
	return &NullableAddrFqdn{value: val, isSet: true}
}

func (v NullableAddrFqdn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddrFqdn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}