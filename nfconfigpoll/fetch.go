// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nfconfigpoll

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"
)

const contentTypeJSON = "application/json"

// FetchJSON makes an HTTP GET request to url and decodes the JSON response body into a T. Anything
// other than a 200 response with a JSON Content-Type is an error.
func FetchJSON[T any](ctx context.Context, client *http.Client, url string) (T, error) {
	var zero T

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return zero, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Accept", contentTypeJSON)

	resp, err := client.Do(req)
	if err != nil {
		return zero, fmt.Errorf("HTTP GET %v failed: %w", url, err)
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil || !strings.EqualFold(mediaType, contentTypeJSON) {
		return zero, fmt.Errorf("unexpected Content-Type: got %q, want %s", contentType, contentTypeJSON)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return zero, fmt.Errorf("failed to read response body: %w", err)
		}
		var config T
		if err := json.Unmarshal(body, &config); err != nil {
			return zero, fmt.Errorf("failed to parse JSON response: %w", err)
		}
		return config, nil
	case http.StatusBadRequest, http.StatusInternalServerError:
		return zero, fmt.Errorf("server returned %d error code", resp.StatusCode)
	default:
		return zero, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}
