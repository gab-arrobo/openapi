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

// checks if the PduSesTrafficInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PduSesTrafficInfo{}

// PduSesTrafficInfo Represents the PDU Set traffic analytics information.
type PduSesTrafficInfo struct {
	Supis []string `json:"supis,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn                  *string     `json:"dnn,omitempty"`
	Snssai               *Snssai     `json:"snssai,omitempty"`
	TdMatchTrafs         []TdTraffic `json:"tdMatchTrafs,omitempty"`
	TdUnmatchTrafs       []TdTraffic `json:"tdUnmatchTrafs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PduSesTrafficInfo PduSesTrafficInfo

// NewPduSesTrafficInfo instantiates a new PduSesTrafficInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSesTrafficInfo() *PduSesTrafficInfo {
	this := PduSesTrafficInfo{}
	return &this
}

// NewPduSesTrafficInfoWithDefaults instantiates a new PduSesTrafficInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSesTrafficInfoWithDefaults() *PduSesTrafficInfo {
	this := PduSesTrafficInfo{}
	return &this
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *PduSesTrafficInfo) GetSupis() []string {
	if o == nil || IsNil(o.Supis) {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSesTrafficInfo) GetSupisOk() ([]string, bool) {
	if o == nil || IsNil(o.Supis) {
		return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *PduSesTrafficInfo) HasSupis() bool {
	if o != nil && !IsNil(o.Supis) {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *PduSesTrafficInfo) SetSupis(v []string) {
	o.Supis = v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *PduSesTrafficInfo) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSesTrafficInfo) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *PduSesTrafficInfo) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *PduSesTrafficInfo) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *PduSesTrafficInfo) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSesTrafficInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *PduSesTrafficInfo) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *PduSesTrafficInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetTdMatchTrafs returns the TdMatchTrafs field value if set, zero value otherwise.
func (o *PduSesTrafficInfo) GetTdMatchTrafs() []TdTraffic {
	if o == nil || IsNil(o.TdMatchTrafs) {
		var ret []TdTraffic
		return ret
	}
	return o.TdMatchTrafs
}

// GetTdMatchTrafsOk returns a tuple with the TdMatchTrafs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSesTrafficInfo) GetTdMatchTrafsOk() ([]TdTraffic, bool) {
	if o == nil || IsNil(o.TdMatchTrafs) {
		return nil, false
	}
	return o.TdMatchTrafs, true
}

// HasTdMatchTrafs returns a boolean if a field has been set.
func (o *PduSesTrafficInfo) HasTdMatchTrafs() bool {
	if o != nil && !IsNil(o.TdMatchTrafs) {
		return true
	}

	return false
}

// SetTdMatchTrafs gets a reference to the given []TdTraffic and assigns it to the TdMatchTrafs field.
func (o *PduSesTrafficInfo) SetTdMatchTrafs(v []TdTraffic) {
	o.TdMatchTrafs = v
}

// GetTdUnmatchTrafs returns the TdUnmatchTrafs field value if set, zero value otherwise.
func (o *PduSesTrafficInfo) GetTdUnmatchTrafs() []TdTraffic {
	if o == nil || IsNil(o.TdUnmatchTrafs) {
		var ret []TdTraffic
		return ret
	}
	return o.TdUnmatchTrafs
}

// GetTdUnmatchTrafsOk returns a tuple with the TdUnmatchTrafs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSesTrafficInfo) GetTdUnmatchTrafsOk() ([]TdTraffic, bool) {
	if o == nil || IsNil(o.TdUnmatchTrafs) {
		return nil, false
	}
	return o.TdUnmatchTrafs, true
}

// HasTdUnmatchTrafs returns a boolean if a field has been set.
func (o *PduSesTrafficInfo) HasTdUnmatchTrafs() bool {
	if o != nil && !IsNil(o.TdUnmatchTrafs) {
		return true
	}

	return false
}

// SetTdUnmatchTrafs gets a reference to the given []TdTraffic and assigns it to the TdUnmatchTrafs field.
func (o *PduSesTrafficInfo) SetTdUnmatchTrafs(v []TdTraffic) {
	o.TdUnmatchTrafs = v
}

func (o PduSesTrafficInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Supis) {
		toSerialize["supis"] = o.Supis
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !IsNil(o.TdMatchTrafs) {
		toSerialize["tdMatchTrafs"] = o.TdMatchTrafs
	}
	if !IsNil(o.TdUnmatchTrafs) {
		toSerialize["tdUnmatchTrafs"] = o.TdUnmatchTrafs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullablePduSesTrafficInfo struct {
	value *PduSesTrafficInfo
	isSet bool
}

func (v NullablePduSesTrafficInfo) Get() *PduSesTrafficInfo {
	return v.value
}

func (v *NullablePduSesTrafficInfo) Set(val *PduSesTrafficInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSesTrafficInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSesTrafficInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSesTrafficInfo(val *PduSesTrafficInfo) *NullablePduSesTrafficInfo {
	return &NullablePduSesTrafficInfo{value: val, isSet: true}
}

func (v NullablePduSesTrafficInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSesTrafficInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}