// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaims struct {
	Iss   string      `json:"iss" yaml:"iss" bson:"iss" mapstructure:"Iss"`
	Sub   string      `json:"sub" yaml:"sub" bson:"sub" mapstructure:"Sub"`
	Aud   interface{} `json:"aud" yaml:"aud" bson:"aud" mapstructure:"Aud"`
	Scope string      `json:"scope" yaml:"scope" bson:"scope" mapstructure:"Scope"`
	Exp   int32       `json:"exp" yaml:"exp" bson:"exp" mapstructure:"Exp"`
	jwt.RegisteredClaims
}
