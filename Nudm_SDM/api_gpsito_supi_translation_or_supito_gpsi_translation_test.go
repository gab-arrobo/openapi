// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
Nudm_SDM

Testing GPSIToSUPITranslationOrSUPIToGPSITranslationAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudm_SDM

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Nudm_SDM_GPSIToSUPITranslationOrSUPIToGPSITranslationAPIService(t *testing.T) {

	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test GPSIToSUPITranslationOrSUPIToGPSITranslationAPIService GetSupiOrGpsi", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var ueId string

		resp, httpRes, err := apiClient.GPSIToSUPITranslationOrSUPIToGPSITranslationAPI.GetSupiOrGpsi(context.Background(), ueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}