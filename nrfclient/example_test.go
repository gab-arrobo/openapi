// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nrfclient_test

import (
	"context"

	"github.com/omec-project/openapi/v2/models"
	"github.com/omec-project/openapi/v2/nrfclient"
	"github.com/omec-project/openapi/v2/nrfregistration"
)

// An NF keeps its own profile builder and its own record of the NF instance ID, and plugs the
// client into the registration service.
func Example_withNrfregistration() {
	var nfInstanceID string
	buildProfile := func(plmns []models.PlmnId) models.NFProfile {
		p := models.NewNFProfileWithDefaults()
		p.SetNfInstanceId(nfInstanceID)
		p.SetNfType(models.NFTYPE_UDM)
		p.SetPlmnList(plmns)
		return *p
	}
	client := nrfclient.New("http://nrf:8000")

	svc, err := nrfregistration.New(nrfregistration.Config[[]models.PlmnId]{
		NfName:  "UDM",
		IsEmpty: func(c []models.PlmnId) bool { return len(c) == 0 },
		Register: func(ctx context.Context, plmns []models.PlmnId) (*models.NFProfile, error) {
			res, err := client.Register(ctx, buildProfile(plmns))
			if err != nil {
				return nil, err
			}
			if res.NfInstanceID != "" {
				nfInstanceID = res.NfInstanceID
			}
			return res.Profile, nil
		},
		Update: func(ctx context.Context, patch []models.PatchItem) (*models.NFProfile, *models.ProblemDetails, error) {
			return client.Update(ctx, nfInstanceID, patch)
		},
		Deregister: func(ctx context.Context) error {
			return client.Deregister(ctx, nfInstanceID)
		},
	})
	if err != nil {
		panic(err)
	}
	_ = svc
}
