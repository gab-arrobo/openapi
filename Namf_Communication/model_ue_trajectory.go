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

// checks if the UeTrajectory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeTrajectory{}

// UeTrajectory Represents timestamped UE positions.
type UeTrajectory struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi *string
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi                 *string
	TimestampedLocs      []TimestampedLocation
	AdditionalProperties map[string]interface{}
}

type _UeTrajectory UeTrajectory

// NewUeTrajectory instantiates a new UeTrajectory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeTrajectory(timestampedLocs []TimestampedLocation) *UeTrajectory {
	this := UeTrajectory{}
	return &this
}

// NewUeTrajectoryWithDefaults instantiates a new UeTrajectory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeTrajectoryWithDefaults() *UeTrajectory {
	this := UeTrajectory{}
	return &this
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *UeTrajectory) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeTrajectory) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *UeTrajectory) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *UeTrajectory) SetSupi(v string) {
	o.Supi = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *UeTrajectory) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeTrajectory) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *UeTrajectory) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *UeTrajectory) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetTimestampedLocs returns the TimestampedLocs field value
func (o *UeTrajectory) GetTimestampedLocs() []TimestampedLocation {
	if o == nil {
		var ret []TimestampedLocation
		return ret
	}

	return o.TimestampedLocs
}

// GetTimestampedLocsOk returns a tuple with the TimestampedLocs field value
// and a boolean to check if the value has been set.
func (o *UeTrajectory) GetTimestampedLocsOk() ([]TimestampedLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimestampedLocs, true
}

// SetTimestampedLocs sets field value
func (o *UeTrajectory) SetTimestampedLocs(v []TimestampedLocation) {
	o.TimestampedLocs = v
}

func (o UeTrajectory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	toSerialize["timestampedLocs"] = o.TimestampedLocs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableUeTrajectory struct {
	value *UeTrajectory
	isSet bool
}

func (v NullableUeTrajectory) Get() *UeTrajectory {
	return v.value
}

func (v *NullableUeTrajectory) Set(val *UeTrajectory) {
	v.value = val
	v.isSet = true
}

func (v NullableUeTrajectory) IsSet() bool {
	return v.isSet
}

func (v *NullableUeTrajectory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeTrajectory(val *UeTrajectory) *NullableUeTrajectory {
	return &NullableUeTrajectory{value: val, isSet: true}
}

func (v NullableUeTrajectory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeTrajectory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}