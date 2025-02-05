//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// EntitiesListPager provides operations for iterating over paged responses.
type EntitiesListPager struct {
	client    *EntitiesClient
	current   EntitiesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, EntitiesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *EntitiesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *EntitiesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EntityListResult.NextLink == nil || len(*p.current.EntityListResult.NextLink) == 0 {
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

// PageResponse returns the current EntitiesListResponse page.
func (p *EntitiesListPager) PageResponse() EntitiesListResponse {
	return p.current
}

// ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroupPager provides operations for iterating over paged responses.
type ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroupPager struct {
	client    *ManagementGroupSubscriptionsClient
	current   ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListSubscriptionUnderManagementGroup.NextLink == nil || len(*p.current.ListSubscriptionUnderManagementGroup.NextLink) == 0 {
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
		p.err = p.client.getSubscriptionsUnderManagementGroupHandleError(resp)
		return false
	}
	result, err := p.client.getSubscriptionsUnderManagementGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroupResponse page.
func (p *ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroupPager) PageResponse() ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroupResponse {
	return p.current
}

// ManagementGroupsGetDescendantsPager provides operations for iterating over paged responses.
type ManagementGroupsGetDescendantsPager struct {
	client    *ManagementGroupsClient
	current   ManagementGroupsGetDescendantsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementGroupsGetDescendantsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementGroupsGetDescendantsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementGroupsGetDescendantsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DescendantListResult.NextLink == nil || len(*p.current.DescendantListResult.NextLink) == 0 {
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
		p.err = p.client.getDescendantsHandleError(resp)
		return false
	}
	result, err := p.client.getDescendantsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementGroupsGetDescendantsResponse page.
func (p *ManagementGroupsGetDescendantsPager) PageResponse() ManagementGroupsGetDescendantsResponse {
	return p.current
}

// ManagementGroupsListPager provides operations for iterating over paged responses.
type ManagementGroupsListPager struct {
	client    *ManagementGroupsClient
	current   ManagementGroupsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementGroupsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementGroupsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementGroupsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagementGroupListResult.NextLink == nil || len(*p.current.ManagementGroupListResult.NextLink) == 0 {
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

// PageResponse returns the current ManagementGroupsListResponse page.
func (p *ManagementGroupsListPager) PageResponse() ManagementGroupsListResponse {
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
