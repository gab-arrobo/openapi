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
	"time"
)

// checks if the UeCommunication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeCommunication{}

// UeCommunication Represents UE communication information.
type UeCommunication struct {
	// indicating a time in seconds.
	CommDur int32
	// string with format 'float' as defined in OpenAPI.
	CommDurVariance *float32
	// indicating a time in seconds.
	PerioTime *int32
	// string with format 'float' as defined in OpenAPI.
	PerioTimeVariance *float32
	// string with format 'date-time' as defined in OpenAPI.
	Ts *time.Time
	// string with format 'float' as defined in OpenAPI.
	TsVariance    *float32
	RecurringTime *ScheduledCommunicationTime1
	TrafChar      NullableTrafficCharacterization
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio *int32
	// This attribute indicates whether the UE communicates periodically or not. Set to \"true\" to indicate the UE communicates periodically, otherwise set to \"false\" or omitted.
	PerioCommInd *bool
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence           *int32
	AnaOfAppList         *AppListForUeComm
	SessInactTimer       *SessInactTimerForUeComm
	AdditionalProperties map[string]interface{}
}

type _UeCommunication UeCommunication

// NewUeCommunication instantiates a new UeCommunication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeCommunication(commDur int32, trafChar NullableTrafficCharacterization) *UeCommunication {
	this := UeCommunication{}
	return &this
}

// NewUeCommunicationWithDefaults instantiates a new UeCommunication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeCommunicationWithDefaults() *UeCommunication {
	this := UeCommunication{}
	return &this
}

// GetCommDur returns the CommDur field value
func (o *UeCommunication) GetCommDur() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CommDur
}

// GetCommDurOk returns a tuple with the CommDur field value
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetCommDurOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommDur, true
}

// SetCommDur sets field value
func (o *UeCommunication) SetCommDur(v int32) {
	o.CommDur = v
}

// GetCommDurVariance returns the CommDurVariance field value if set, zero value otherwise.
func (o *UeCommunication) GetCommDurVariance() float32 {
	if o == nil || IsNil(o.CommDurVariance) {
		var ret float32
		return ret
	}
	return *o.CommDurVariance
}

// GetCommDurVarianceOk returns a tuple with the CommDurVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetCommDurVarianceOk() (*float32, bool) {
	if o == nil || IsNil(o.CommDurVariance) {
		return nil, false
	}
	return o.CommDurVariance, true
}

// HasCommDurVariance returns a boolean if a field has been set.
func (o *UeCommunication) HasCommDurVariance() bool {
	if o != nil && !IsNil(o.CommDurVariance) {
		return true
	}

	return false
}

// SetCommDurVariance gets a reference to the given float32 and assigns it to the CommDurVariance field.
func (o *UeCommunication) SetCommDurVariance(v float32) {
	o.CommDurVariance = &v
}

// GetPerioTime returns the PerioTime field value if set, zero value otherwise.
func (o *UeCommunication) GetPerioTime() int32 {
	if o == nil || IsNil(o.PerioTime) {
		var ret int32
		return ret
	}
	return *o.PerioTime
}

// GetPerioTimeOk returns a tuple with the PerioTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetPerioTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.PerioTime) {
		return nil, false
	}
	return o.PerioTime, true
}

// HasPerioTime returns a boolean if a field has been set.
func (o *UeCommunication) HasPerioTime() bool {
	if o != nil && !IsNil(o.PerioTime) {
		return true
	}

	return false
}

// SetPerioTime gets a reference to the given int32 and assigns it to the PerioTime field.
func (o *UeCommunication) SetPerioTime(v int32) {
	o.PerioTime = &v
}

// GetPerioTimeVariance returns the PerioTimeVariance field value if set, zero value otherwise.
func (o *UeCommunication) GetPerioTimeVariance() float32 {
	if o == nil || IsNil(o.PerioTimeVariance) {
		var ret float32
		return ret
	}
	return *o.PerioTimeVariance
}

// GetPerioTimeVarianceOk returns a tuple with the PerioTimeVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetPerioTimeVarianceOk() (*float32, bool) {
	if o == nil || IsNil(o.PerioTimeVariance) {
		return nil, false
	}
	return o.PerioTimeVariance, true
}

// HasPerioTimeVariance returns a boolean if a field has been set.
func (o *UeCommunication) HasPerioTimeVariance() bool {
	if o != nil && !IsNil(o.PerioTimeVariance) {
		return true
	}

	return false
}

// SetPerioTimeVariance gets a reference to the given float32 and assigns it to the PerioTimeVariance field.
func (o *UeCommunication) SetPerioTimeVariance(v float32) {
	o.PerioTimeVariance = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *UeCommunication) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *UeCommunication) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *UeCommunication) SetTs(v time.Time) {
	o.Ts = &v
}

// GetTsVariance returns the TsVariance field value if set, zero value otherwise.
func (o *UeCommunication) GetTsVariance() float32 {
	if o == nil || IsNil(o.TsVariance) {
		var ret float32
		return ret
	}
	return *o.TsVariance
}

// GetTsVarianceOk returns a tuple with the TsVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetTsVarianceOk() (*float32, bool) {
	if o == nil || IsNil(o.TsVariance) {
		return nil, false
	}
	return o.TsVariance, true
}

// HasTsVariance returns a boolean if a field has been set.
func (o *UeCommunication) HasTsVariance() bool {
	if o != nil && !IsNil(o.TsVariance) {
		return true
	}

	return false
}

// SetTsVariance gets a reference to the given float32 and assigns it to the TsVariance field.
func (o *UeCommunication) SetTsVariance(v float32) {
	o.TsVariance = &v
}

// GetRecurringTime returns the RecurringTime field value if set, zero value otherwise.
func (o *UeCommunication) GetRecurringTime() ScheduledCommunicationTime1 {
	if o == nil || IsNil(o.RecurringTime) {
		var ret ScheduledCommunicationTime1
		return ret
	}
	return *o.RecurringTime
}

// GetRecurringTimeOk returns a tuple with the RecurringTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetRecurringTimeOk() (*ScheduledCommunicationTime1, bool) {
	if o == nil || IsNil(o.RecurringTime) {
		return nil, false
	}
	return o.RecurringTime, true
}

// HasRecurringTime returns a boolean if a field has been set.
func (o *UeCommunication) HasRecurringTime() bool {
	if o != nil && !IsNil(o.RecurringTime) {
		return true
	}

	return false
}

// SetRecurringTime gets a reference to the given ScheduledCommunicationTime1 and assigns it to the RecurringTime field.
func (o *UeCommunication) SetRecurringTime(v ScheduledCommunicationTime1) {
	o.RecurringTime = &v
}

// GetTrafChar returns the TrafChar field value
// If the value is explicit nil, the zero value for TrafficCharacterization will be returned
func (o *UeCommunication) GetTrafChar() TrafficCharacterization {
	if o == nil || o.TrafChar.Get() == nil {
		var ret TrafficCharacterization
		return ret
	}

	return *o.TrafChar.Get()
}

// GetTrafCharOk returns a tuple with the TrafChar field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UeCommunication) GetTrafCharOk() (*TrafficCharacterization, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrafChar.Get(), o.TrafChar.IsSet()
}

// SetTrafChar sets field value
func (o *UeCommunication) SetTrafChar(v TrafficCharacterization) {
	o.TrafChar.Set(&v)
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *UeCommunication) GetRatio() int32 {
	if o == nil || IsNil(o.Ratio) {
		var ret int32
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetRatioOk() (*int32, bool) {
	if o == nil || IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *UeCommunication) HasRatio() bool {
	if o != nil && !IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given int32 and assigns it to the Ratio field.
func (o *UeCommunication) SetRatio(v int32) {
	o.Ratio = &v
}

// GetPerioCommInd returns the PerioCommInd field value if set, zero value otherwise.
func (o *UeCommunication) GetPerioCommInd() bool {
	if o == nil || IsNil(o.PerioCommInd) {
		var ret bool
		return ret
	}
	return *o.PerioCommInd
}

// GetPerioCommIndOk returns a tuple with the PerioCommInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetPerioCommIndOk() (*bool, bool) {
	if o == nil || IsNil(o.PerioCommInd) {
		return nil, false
	}
	return o.PerioCommInd, true
}

// HasPerioCommInd returns a boolean if a field has been set.
func (o *UeCommunication) HasPerioCommInd() bool {
	if o != nil && !IsNil(o.PerioCommInd) {
		return true
	}

	return false
}

// SetPerioCommInd gets a reference to the given bool and assigns it to the PerioCommInd field.
func (o *UeCommunication) SetPerioCommInd(v bool) {
	o.PerioCommInd = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *UeCommunication) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetConfidenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *UeCommunication) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *UeCommunication) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetAnaOfAppList returns the AnaOfAppList field value if set, zero value otherwise.
func (o *UeCommunication) GetAnaOfAppList() AppListForUeComm {
	if o == nil || IsNil(o.AnaOfAppList) {
		var ret AppListForUeComm
		return ret
	}
	return *o.AnaOfAppList
}

// GetAnaOfAppListOk returns a tuple with the AnaOfAppList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetAnaOfAppListOk() (*AppListForUeComm, bool) {
	if o == nil || IsNil(o.AnaOfAppList) {
		return nil, false
	}
	return o.AnaOfAppList, true
}

// HasAnaOfAppList returns a boolean if a field has been set.
func (o *UeCommunication) HasAnaOfAppList() bool {
	if o != nil && !IsNil(o.AnaOfAppList) {
		return true
	}

	return false
}

// SetAnaOfAppList gets a reference to the given AppListForUeComm and assigns it to the AnaOfAppList field.
func (o *UeCommunication) SetAnaOfAppList(v AppListForUeComm) {
	o.AnaOfAppList = &v
}

// GetSessInactTimer returns the SessInactTimer field value if set, zero value otherwise.
func (o *UeCommunication) GetSessInactTimer() SessInactTimerForUeComm {
	if o == nil || IsNil(o.SessInactTimer) {
		var ret SessInactTimerForUeComm
		return ret
	}
	return *o.SessInactTimer
}

// GetSessInactTimerOk returns a tuple with the SessInactTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetSessInactTimerOk() (*SessInactTimerForUeComm, bool) {
	if o == nil || IsNil(o.SessInactTimer) {
		return nil, false
	}
	return o.SessInactTimer, true
}

// HasSessInactTimer returns a boolean if a field has been set.
func (o *UeCommunication) HasSessInactTimer() bool {
	if o != nil && !IsNil(o.SessInactTimer) {
		return true
	}

	return false
}

// SetSessInactTimer gets a reference to the given SessInactTimerForUeComm and assigns it to the SessInactTimer field.
func (o *UeCommunication) SetSessInactTimer(v SessInactTimerForUeComm) {
	o.SessInactTimer = &v
}

func (o UeCommunication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commDur"] = o.CommDur
	if !IsNil(o.CommDurVariance) {
		toSerialize["commDurVariance"] = o.CommDurVariance
	}
	if !IsNil(o.PerioTime) {
		toSerialize["perioTime"] = o.PerioTime
	}
	if !IsNil(o.PerioTimeVariance) {
		toSerialize["perioTimeVariance"] = o.PerioTimeVariance
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.TsVariance) {
		toSerialize["tsVariance"] = o.TsVariance
	}
	if !IsNil(o.RecurringTime) {
		toSerialize["recurringTime"] = o.RecurringTime
	}
	toSerialize["trafChar"] = o.TrafChar.Get()
	if !IsNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}
	if !IsNil(o.PerioCommInd) {
		toSerialize["perioCommInd"] = o.PerioCommInd
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !IsNil(o.AnaOfAppList) {
		toSerialize["anaOfAppList"] = o.AnaOfAppList
	}
	if !IsNil(o.SessInactTimer) {
		toSerialize["sessInactTimer"] = o.SessInactTimer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableUeCommunication struct {
	value *UeCommunication
	isSet bool
}

func (v NullableUeCommunication) Get() *UeCommunication {
	return v.value
}

func (v *NullableUeCommunication) Set(val *UeCommunication) {
	v.value = val
	v.isSet = true
}

func (v NullableUeCommunication) IsSet() bool {
	return v.isSet
}

func (v *NullableUeCommunication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeCommunication(val *UeCommunication) *NullableUeCommunication {
	return &NullableUeCommunication{value: val, isSet: true}
}

func (v NullableUeCommunication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeCommunication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}