//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ManagementGroupsAPIClient contains the methods for the ManagementGroupsAPI group.
// Don't use this type directly, use NewManagementGroupsAPIClient() instead.
type ManagementGroupsAPIClient struct {
	ep string
	pl runtime.Pipeline
}

// NewManagementGroupsAPIClient creates a new instance of ManagementGroupsAPIClient with the specified values.
func NewManagementGroupsAPIClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ManagementGroupsAPIClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ManagementGroupsAPIClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckNameAvailability - Checks if the specified management group name is valid and unique
// If the operation fails it returns the *ErrorResponse error type.
func (client *ManagementGroupsAPIClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityRequest CheckNameAvailabilityRequest, options *ManagementGroupsAPICheckNameAvailabilityOptions) (ManagementGroupsAPICheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, checkNameAvailabilityRequest, options)
	if err != nil {
		return ManagementGroupsAPICheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementGroupsAPICheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementGroupsAPICheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ManagementGroupsAPIClient) checkNameAvailabilityCreateRequest(ctx context.Context, checkNameAvailabilityRequest CheckNameAvailabilityRequest, options *ManagementGroupsAPICheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/checkNameAvailability"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, checkNameAvailabilityRequest)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ManagementGroupsAPIClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ManagementGroupsAPICheckNameAvailabilityResponse, error) {
	result := ManagementGroupsAPICheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return ManagementGroupsAPICheckNameAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *ManagementGroupsAPIClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// StartTenantBackfill - Starts backfilling subscriptions for the Tenant.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ManagementGroupsAPIClient) StartTenantBackfill(ctx context.Context, options *ManagementGroupsAPIStartTenantBackfillOptions) (ManagementGroupsAPIStartTenantBackfillResponse, error) {
	req, err := client.startTenantBackfillCreateRequest(ctx, options)
	if err != nil {
		return ManagementGroupsAPIStartTenantBackfillResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementGroupsAPIStartTenantBackfillResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementGroupsAPIStartTenantBackfillResponse{}, client.startTenantBackfillHandleError(resp)
	}
	return client.startTenantBackfillHandleResponse(resp)
}

// startTenantBackfillCreateRequest creates the StartTenantBackfill request.
func (client *ManagementGroupsAPIClient) startTenantBackfillCreateRequest(ctx context.Context, options *ManagementGroupsAPIStartTenantBackfillOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/startTenantBackfill"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// startTenantBackfillHandleResponse handles the StartTenantBackfill response.
func (client *ManagementGroupsAPIClient) startTenantBackfillHandleResponse(resp *http.Response) (ManagementGroupsAPIStartTenantBackfillResponse, error) {
	result := ManagementGroupsAPIStartTenantBackfillResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantBackfillStatusResult); err != nil {
		return ManagementGroupsAPIStartTenantBackfillResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// startTenantBackfillHandleError handles the StartTenantBackfill error response.
func (client *ManagementGroupsAPIClient) startTenantBackfillHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// TenantBackfillStatus - Gets tenant backfill status
// If the operation fails it returns the *ErrorResponse error type.
func (client *ManagementGroupsAPIClient) TenantBackfillStatus(ctx context.Context, options *ManagementGroupsAPITenantBackfillStatusOptions) (ManagementGroupsAPITenantBackfillStatusResponse, error) {
	req, err := client.tenantBackfillStatusCreateRequest(ctx, options)
	if err != nil {
		return ManagementGroupsAPITenantBackfillStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementGroupsAPITenantBackfillStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementGroupsAPITenantBackfillStatusResponse{}, client.tenantBackfillStatusHandleError(resp)
	}
	return client.tenantBackfillStatusHandleResponse(resp)
}

// tenantBackfillStatusCreateRequest creates the TenantBackfillStatus request.
func (client *ManagementGroupsAPIClient) tenantBackfillStatusCreateRequest(ctx context.Context, options *ManagementGroupsAPITenantBackfillStatusOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/tenantBackfillStatus"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// tenantBackfillStatusHandleResponse handles the TenantBackfillStatus response.
func (client *ManagementGroupsAPIClient) tenantBackfillStatusHandleResponse(resp *http.Response) (ManagementGroupsAPITenantBackfillStatusResponse, error) {
	result := ManagementGroupsAPITenantBackfillStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantBackfillStatusResult); err != nil {
		return ManagementGroupsAPITenantBackfillStatusResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// tenantBackfillStatusHandleError handles the TenantBackfillStatus error response.
func (client *ManagementGroupsAPIClient) tenantBackfillStatusHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
