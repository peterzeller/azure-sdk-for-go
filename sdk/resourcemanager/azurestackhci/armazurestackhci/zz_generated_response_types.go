//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// ArcSettingsCreateResponse contains the response from method ArcSettings.Create.
type ArcSettingsCreateResponse struct {
	ArcSettingsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArcSettingsCreateResult contains the result from method ArcSettings.Create.
type ArcSettingsCreateResult struct {
	ArcSetting
}

// ArcSettingsDeletePollerResponse contains the response from method ArcSettings.Delete.
type ArcSettingsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ArcSettingsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ArcSettingsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ArcSettingsDeleteResponse, error) {
	respType := ArcSettingsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ArcSettingsDeletePollerResponse from the provided client and resume token.
func (l *ArcSettingsDeletePollerResponse) Resume(ctx context.Context, client *ArcSettingsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ArcSettingsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &ArcSettingsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ArcSettingsDeleteResponse contains the response from method ArcSettings.Delete.
type ArcSettingsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArcSettingsGetResponse contains the response from method ArcSettings.Get.
type ArcSettingsGetResponse struct {
	ArcSettingsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArcSettingsGetResult contains the result from method ArcSettings.Get.
type ArcSettingsGetResult struct {
	ArcSetting
}

// ArcSettingsListByClusterResponse contains the response from method ArcSettings.ListByCluster.
type ArcSettingsListByClusterResponse struct {
	ArcSettingsListByClusterResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArcSettingsListByClusterResult contains the result from method ArcSettings.ListByCluster.
type ArcSettingsListByClusterResult struct {
	ArcSettingList
}

// ClustersCreateResponse contains the response from method Clusters.Create.
type ClustersCreateResponse struct {
	ClustersCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersCreateResult contains the result from method Clusters.Create.
type ClustersCreateResult struct {
	Cluster
}

// ClustersDeleteResponse contains the response from method Clusters.Delete.
type ClustersDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersGetResponse contains the response from method Clusters.Get.
type ClustersGetResponse struct {
	ClustersGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersGetResult contains the result from method Clusters.Get.
type ClustersGetResult struct {
	Cluster
}

// ClustersListByResourceGroupResponse contains the response from method Clusters.ListByResourceGroup.
type ClustersListByResourceGroupResponse struct {
	ClustersListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersListByResourceGroupResult contains the result from method Clusters.ListByResourceGroup.
type ClustersListByResourceGroupResult struct {
	ClusterList
}

// ClustersListBySubscriptionResponse contains the response from method Clusters.ListBySubscription.
type ClustersListBySubscriptionResponse struct {
	ClustersListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersListBySubscriptionResult contains the result from method Clusters.ListBySubscription.
type ClustersListBySubscriptionResult struct {
	ClusterList
}

// ClustersUpdateResponse contains the response from method Clusters.Update.
type ClustersUpdateResponse struct {
	ClustersUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersUpdateResult contains the result from method Clusters.Update.
type ClustersUpdateResult struct {
	Cluster
}

// ExtensionsCreatePollerResponse contains the response from method Extensions.Create.
type ExtensionsCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ExtensionsCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ExtensionsCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ExtensionsCreateResponse, error) {
	respType := ExtensionsCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Extension)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ExtensionsCreatePollerResponse from the provided client and resume token.
func (l *ExtensionsCreatePollerResponse) Resume(ctx context.Context, client *ExtensionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ExtensionsClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &ExtensionsCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ExtensionsCreateResponse contains the response from method Extensions.Create.
type ExtensionsCreateResponse struct {
	ExtensionsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsCreateResult contains the result from method Extensions.Create.
type ExtensionsCreateResult struct {
	Extension
}

// ExtensionsDeletePollerResponse contains the response from method Extensions.Delete.
type ExtensionsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ExtensionsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ExtensionsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ExtensionsDeleteResponse, error) {
	respType := ExtensionsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ExtensionsDeletePollerResponse from the provided client and resume token.
func (l *ExtensionsDeletePollerResponse) Resume(ctx context.Context, client *ExtensionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ExtensionsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &ExtensionsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ExtensionsDeleteResponse contains the response from method Extensions.Delete.
type ExtensionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsGetResponse contains the response from method Extensions.Get.
type ExtensionsGetResponse struct {
	ExtensionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsGetResult contains the result from method Extensions.Get.
type ExtensionsGetResult struct {
	Extension
}

// ExtensionsListByArcSettingResponse contains the response from method Extensions.ListByArcSetting.
type ExtensionsListByArcSettingResponse struct {
	ExtensionsListByArcSettingResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsListByArcSettingResult contains the result from method Extensions.ListByArcSetting.
type ExtensionsListByArcSettingResult struct {
	ExtensionList
}

// ExtensionsUpdatePollerResponse contains the response from method Extensions.Update.
type ExtensionsUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ExtensionsUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ExtensionsUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ExtensionsUpdateResponse, error) {
	respType := ExtensionsUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Extension)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ExtensionsUpdatePollerResponse from the provided client and resume token.
func (l *ExtensionsUpdatePollerResponse) Resume(ctx context.Context, client *ExtensionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ExtensionsClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &ExtensionsUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ExtensionsUpdateResponse contains the response from method Extensions.Update.
type ExtensionsUpdateResponse struct {
	ExtensionsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsUpdateResult contains the result from method Extensions.Update.
type ExtensionsUpdateResult struct {
	Extension
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}
