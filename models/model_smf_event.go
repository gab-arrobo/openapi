// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SmfEvent string

// List of SmfEvent
const (
	SmfEvent_AC_TY_CH    SmfEvent = "AC_TY_CH"
	SmfEvent_UP_PATH_CH  SmfEvent = "UP_PATH_CH"
	SmfEvent_PDU_SES_REL SmfEvent = "PDU_SES_REL"
	SmfEvent_PLMN_CH     SmfEvent = "PLMN_CH"
	SmfEvent_UE_IP_CH    SmfEvent = "UE_IP_CH"
)