// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
Nudm_SDM

Nudm Subscriber Data Management Service.
© 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).
All rights reserved.

API version: 2.3.0-alpha.6

Authors: Aether SD-Core team
Contact: dev@lists.aetherproject.org

Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
*/

package Nudm_SDM

import (
	"encoding/json"
)

// checks if the Polygon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Polygon{}

// Polygon Polygon.
type Polygon struct {
	GADShape
	// List of points.
	PointList            []GeographicalCoordinates `json:"pointList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Polygon Polygon

// NewPolygon instantiates a new Polygon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolygon(shape SupportedGADShapes) *Polygon {
	this := Polygon{}
	this.Shape = shape
	return &this
}

// NewPolygonWithDefaults instantiates a new Polygon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolygonWithDefaults() *Polygon {
	this := Polygon{}
	return &this
}

// GetPointList returns the PointList field value if set, zero value otherwise.
func (o *Polygon) GetPointList() []GeographicalCoordinates {
	if o == nil || IsNil(o.PointList) {
		var ret []GeographicalCoordinates
		return ret
	}
	return o.PointList
}

// GetPointListOk returns a tuple with the PointList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Polygon) GetPointListOk() ([]GeographicalCoordinates, bool) {
	if o == nil || IsNil(o.PointList) {
		return nil, false
	}
	return o.PointList, true
}

// HasPointList returns a boolean if a field has been set.
func (o *Polygon) HasPointList() bool {
	if o != nil && !IsNil(o.PointList) {
		return true
	}

	return false
}

// SetPointList gets a reference to the given []GeographicalCoordinates and assigns it to the PointList field.
func (o *Polygon) SetPointList(v []GeographicalCoordinates) {
	o.PointList = v
}

func (o Polygon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedGADShape, errGADShape := json.Marshal(o.GADShape)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	errGADShape = json.Unmarshal([]byte(serializedGADShape), &toSerialize)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	if !IsNil(o.PointList) {
		toSerialize["pointList"] = o.PointList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullablePolygon struct {
	value *Polygon
	isSet bool
}

func (v NullablePolygon) Get() *Polygon {
	return v.value
}

func (v *NullablePolygon) Set(val *Polygon) {
	v.value = val
	v.isSet = true
}

func (v NullablePolygon) IsSet() bool {
	return v.isSet
}

func (v *NullablePolygon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolygon(val *Polygon) *NullablePolygon {
	return &NullablePolygon{value: val, isSet: true}
}

func (v NullablePolygon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolygon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}