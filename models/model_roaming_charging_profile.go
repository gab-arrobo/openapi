// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RoamingChargingProfile struct {
	Triggers            []Trigger           `json:"triggers,omitempty"`
	PartialRecordMethod PartialRecordMethod `json:"partialRecordMethod,omitempty"`
}