// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PartialSuccessReport struct {
	FailureCause FailureCause `json:"failureCause" yaml:"failureCause" bson:"failureCause" mapstructure:"FailureCause"`
	// Information about the PCC rules provisioned by the PCF not successfully installed/activated.
	RuleReports  []RuleReport  `json:"ruleReports" yaml:"ruleReports" bson:"ruleReports" mapstructure:"RuleReports"`
	UeCampingRep *UeCampingRep `json:"ueCampingRep,omitempty" yaml:"ueCampingRep" bson:"ueCampingRep" mapstructure:"UeCampingRep"`
}