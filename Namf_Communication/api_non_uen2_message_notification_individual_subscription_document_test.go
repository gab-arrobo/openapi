// SPDX-FileCopyrightText: 2024 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

/*
Namf_Communication

Testing NonUEN2MessageNotificationIndividualSubscriptionDocumentAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Namf_Communication

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Namf_Communication_NonUEN2MessageNotificationIndividualSubscriptionDocumentAPIService(t *testing.T) {

	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test NonUEN2MessageNotificationIndividualSubscriptionDocumentAPIService NonUeN2InfoUnSubscribe", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var n2NotifySubscriptionId string

		httpRes, err := apiClient.NonUEN2MessageNotificationIndividualSubscriptionDocumentAPI.NonUeN2InfoUnSubscribe(context.Background(), n2NotifySubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}