// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
NRF NFManagement Service

NRF NFManagement Service.
© 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).
All rights reserved.

API version: 1.3.0-alpha.7

Authors: Aether SD-Core team
Contact: dev@lists.aetherproject.org

Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
*/

package Nnrf_NFManagement

import (
	"encoding/json"
)

// checks if the SubscriptionContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionContext{}

// SubscriptionContext Context data related to a created subscription, to be included in notifications sent by NRF
type SubscriptionContext struct {
	SubscriptionId       string      `json:"subscriptionId"`
	SubscrCond           *SubscrCond `json:"subscrCond,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionContext SubscriptionContext

// NewSubscriptionContext instantiates a new SubscriptionContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionContext(subscriptionId string) *SubscriptionContext {
	this := SubscriptionContext{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewSubscriptionContextWithDefaults instantiates a new SubscriptionContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionContextWithDefaults() *SubscriptionContext {
	this := SubscriptionContext{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *SubscriptionContext) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionContext) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *SubscriptionContext) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetSubscrCond returns the SubscrCond field value if set, zero value otherwise.
func (o *SubscriptionContext) GetSubscrCond() SubscrCond {
	if o == nil || IsNil(o.SubscrCond) {
		var ret SubscrCond
		return ret
	}
	return *o.SubscrCond
}

// GetSubscrCondOk returns a tuple with the SubscrCond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionContext) GetSubscrCondOk() (*SubscrCond, bool) {
	if o == nil || IsNil(o.SubscrCond) {
		return nil, false
	}
	return o.SubscrCond, true
}

// HasSubscrCond returns a boolean if a field has been set.
func (o *SubscriptionContext) HasSubscrCond() bool {
	if o != nil && !IsNil(o.SubscrCond) {
		return true
	}

	return false
}

// SetSubscrCond gets a reference to the given SubscrCond and assigns it to the SubscrCond field.
func (o *SubscriptionContext) SetSubscrCond(v SubscrCond) {
	o.SubscrCond = &v
}

func (o SubscriptionContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionId"] = o.SubscriptionId
	if !IsNil(o.SubscrCond) {
		toSerialize["subscrCond"] = o.SubscrCond
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableSubscriptionContext struct {
	value *SubscriptionContext
	isSet bool
}

func (v NullableSubscriptionContext) Get() *SubscriptionContext {
	return v.value
}

func (v *NullableSubscriptionContext) Set(val *SubscriptionContext) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionContext) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionContext(val *SubscriptionContext) *NullableSubscriptionContext {
	return &NullableSubscriptionContext{value: val, isSet: true}
}

func (v NullableSubscriptionContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}