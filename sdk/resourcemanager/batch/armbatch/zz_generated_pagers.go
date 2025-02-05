//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ApplicationListPager provides operations for iterating over paged responses.
type ApplicationListPager struct {
	client    *ApplicationClient
	current   ApplicationListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ApplicationListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ApplicationListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ApplicationListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListApplicationsResult.NextLink == nil || len(*p.current.ListApplicationsResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ApplicationListResponse page.
func (p *ApplicationListPager) PageResponse() ApplicationListResponse {
	return p.current
}

// ApplicationPackageListPager provides operations for iterating over paged responses.
type ApplicationPackageListPager struct {
	client    *ApplicationPackageClient
	current   ApplicationPackageListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ApplicationPackageListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ApplicationPackageListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ApplicationPackageListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListApplicationPackagesResult.NextLink == nil || len(*p.current.ListApplicationPackagesResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ApplicationPackageListResponse page.
func (p *ApplicationPackageListPager) PageResponse() ApplicationPackageListResponse {
	return p.current
}

// BatchAccountListByResourceGroupPager provides operations for iterating over paged responses.
type BatchAccountListByResourceGroupPager struct {
	client    *BatchAccountClient
	current   BatchAccountListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BatchAccountListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BatchAccountListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BatchAccountListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.BatchAccountListResult.NextLink == nil || len(*p.current.BatchAccountListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BatchAccountListByResourceGroupResponse page.
func (p *BatchAccountListByResourceGroupPager) PageResponse() BatchAccountListByResourceGroupResponse {
	return p.current
}

// BatchAccountListOutboundNetworkDependenciesEndpointsPager provides operations for iterating over paged responses.
type BatchAccountListOutboundNetworkDependenciesEndpointsPager struct {
	client    *BatchAccountClient
	current   BatchAccountListOutboundNetworkDependenciesEndpointsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BatchAccountListOutboundNetworkDependenciesEndpointsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BatchAccountListOutboundNetworkDependenciesEndpointsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BatchAccountListOutboundNetworkDependenciesEndpointsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OutboundEnvironmentEndpointCollection.NextLink == nil || len(*p.current.OutboundEnvironmentEndpointCollection.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listOutboundNetworkDependenciesEndpointsHandleError(resp)
		return false
	}
	result, err := p.client.listOutboundNetworkDependenciesEndpointsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BatchAccountListOutboundNetworkDependenciesEndpointsResponse page.
func (p *BatchAccountListOutboundNetworkDependenciesEndpointsPager) PageResponse() BatchAccountListOutboundNetworkDependenciesEndpointsResponse {
	return p.current
}

// BatchAccountListPager provides operations for iterating over paged responses.
type BatchAccountListPager struct {
	client    *BatchAccountClient
	current   BatchAccountListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BatchAccountListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BatchAccountListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BatchAccountListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.BatchAccountListResult.NextLink == nil || len(*p.current.BatchAccountListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BatchAccountListResponse page.
func (p *BatchAccountListPager) PageResponse() BatchAccountListResponse {
	return p.current
}

// CertificateListByBatchAccountPager provides operations for iterating over paged responses.
type CertificateListByBatchAccountPager struct {
	client    *CertificateClient
	current   CertificateListByBatchAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CertificateListByBatchAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CertificateListByBatchAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CertificateListByBatchAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListCertificatesResult.NextLink == nil || len(*p.current.ListCertificatesResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByBatchAccountHandleError(resp)
		return false
	}
	result, err := p.client.listByBatchAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current CertificateListByBatchAccountResponse page.
func (p *CertificateListByBatchAccountPager) PageResponse() CertificateListByBatchAccountResponse {
	return p.current
}

// LocationListSupportedCloudServiceSKUsPager provides operations for iterating over paged responses.
type LocationListSupportedCloudServiceSKUsPager struct {
	client    *LocationClient
	current   LocationListSupportedCloudServiceSKUsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LocationListSupportedCloudServiceSKUsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LocationListSupportedCloudServiceSKUsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LocationListSupportedCloudServiceSKUsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SupportedSKUsResult.NextLink == nil || len(*p.current.SupportedSKUsResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listSupportedCloudServiceSKUsHandleError(resp)
		return false
	}
	result, err := p.client.listSupportedCloudServiceSKUsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current LocationListSupportedCloudServiceSKUsResponse page.
func (p *LocationListSupportedCloudServiceSKUsPager) PageResponse() LocationListSupportedCloudServiceSKUsResponse {
	return p.current
}

// LocationListSupportedVirtualMachineSKUsPager provides operations for iterating over paged responses.
type LocationListSupportedVirtualMachineSKUsPager struct {
	client    *LocationClient
	current   LocationListSupportedVirtualMachineSKUsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LocationListSupportedVirtualMachineSKUsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LocationListSupportedVirtualMachineSKUsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LocationListSupportedVirtualMachineSKUsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SupportedSKUsResult.NextLink == nil || len(*p.current.SupportedSKUsResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listSupportedVirtualMachineSKUsHandleError(resp)
		return false
	}
	result, err := p.client.listSupportedVirtualMachineSKUsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current LocationListSupportedVirtualMachineSKUsResponse page.
func (p *LocationListSupportedVirtualMachineSKUsPager) PageResponse() LocationListSupportedVirtualMachineSKUsResponse {
	return p.current
}

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// PoolListByBatchAccountPager provides operations for iterating over paged responses.
type PoolListByBatchAccountPager struct {
	client    *PoolClient
	current   PoolListByBatchAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PoolListByBatchAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PoolListByBatchAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PoolListByBatchAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListPoolsResult.NextLink == nil || len(*p.current.ListPoolsResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByBatchAccountHandleError(resp)
		return false
	}
	result, err := p.client.listByBatchAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PoolListByBatchAccountResponse page.
func (p *PoolListByBatchAccountPager) PageResponse() PoolListByBatchAccountResponse {
	return p.current
}

// PrivateEndpointConnectionListByBatchAccountPager provides operations for iterating over paged responses.
type PrivateEndpointConnectionListByBatchAccountPager struct {
	client    *PrivateEndpointConnectionClient
	current   PrivateEndpointConnectionListByBatchAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateEndpointConnectionListByBatchAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateEndpointConnectionListByBatchAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateEndpointConnectionListByBatchAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListPrivateEndpointConnectionsResult.NextLink == nil || len(*p.current.ListPrivateEndpointConnectionsResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByBatchAccountHandleError(resp)
		return false
	}
	result, err := p.client.listByBatchAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateEndpointConnectionListByBatchAccountResponse page.
func (p *PrivateEndpointConnectionListByBatchAccountPager) PageResponse() PrivateEndpointConnectionListByBatchAccountResponse {
	return p.current
}

// PrivateLinkResourceListByBatchAccountPager provides operations for iterating over paged responses.
type PrivateLinkResourceListByBatchAccountPager struct {
	client    *PrivateLinkResourceClient
	current   PrivateLinkResourceListByBatchAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateLinkResourceListByBatchAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateLinkResourceListByBatchAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateLinkResourceListByBatchAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListPrivateLinkResourcesResult.NextLink == nil || len(*p.current.ListPrivateLinkResourcesResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByBatchAccountHandleError(resp)
		return false
	}
	result, err := p.client.listByBatchAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateLinkResourceListByBatchAccountResponse page.
func (p *PrivateLinkResourceListByBatchAccountPager) PageResponse() PrivateLinkResourceListByBatchAccountResponse {
	return p.current
}
