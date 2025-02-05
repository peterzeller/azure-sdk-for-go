//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_createorupdate.json
func ExampleRolloutsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewRolloutsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<rollout-name>",
		&armdeploymentmanager.RolloutsBeginCreateOrUpdateOptions{RolloutRequest: &armdeploymentmanager.RolloutRequest{
			TrackedResource: armdeploymentmanager.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags:     map[string]*string{},
			},
			Identity: &armdeploymentmanager.Identity{
				Type: to.StringPtr("<type>"),
				IdentityIDs: []*string{
					to.StringPtr("/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userassignedidentities/myuseridentity")},
			},
			Properties: &armdeploymentmanager.RolloutRequestProperties{
				ArtifactSourceID: to.StringPtr("<artifact-source-id>"),
				BuildVersion:     to.StringPtr("<build-version>"),
				StepGroups: []*armdeploymentmanager.StepGroup{
					{
						Name:               to.StringPtr("<name>"),
						DeploymentTargetID: to.StringPtr("<deployment-target-id>"),
						PostDeploymentSteps: []*armdeploymentmanager.PrePostStep{
							{
								StepID: to.StringPtr("<step-id>"),
							}},
						PreDeploymentSteps: []*armdeploymentmanager.PrePostStep{
							{
								StepID: to.StringPtr("<step-id>"),
							},
							{
								StepID: to.StringPtr("<step-id>"),
							}},
					},
					{
						Name: to.StringPtr("<name>"),
						DependsOnStepGroups: []*string{
							to.StringPtr("FirstRegion")},
						DeploymentTargetID: to.StringPtr("<deployment-target-id>"),
						PostDeploymentSteps: []*armdeploymentmanager.PrePostStep{
							{
								StepID: to.StringPtr("<step-id>"),
							}},
						PreDeploymentSteps: []*armdeploymentmanager.PrePostStep{
							{
								StepID: to.StringPtr("<step-id>"),
							},
							{
								StepID: to.StringPtr("<step-id>"),
							}},
					}},
				TargetServiceTopologyID: to.StringPtr("<target-service-topology-id>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RolloutRequest.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_get.json
func ExampleRolloutsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewRolloutsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<rollout-name>",
		&armdeploymentmanager.RolloutsGetOptions{RetryAttempt: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Rollout.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_delete.json
func ExampleRolloutsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewRolloutsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<rollout-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_post_cancel.json
func ExampleRolloutsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewRolloutsClient("<subscription-id>", cred, nil)
	res, err := client.Cancel(ctx,
		"<resource-group-name>",
		"<rollout-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Rollout.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_post_restart.json
func ExampleRolloutsClient_Restart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewRolloutsClient("<subscription-id>", cred, nil)
	res, err := client.Restart(ctx,
		"<resource-group-name>",
		"<rollout-name>",
		&armdeploymentmanager.RolloutsRestartOptions{SkipSucceeded: to.BoolPtr(true)})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Rollout.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollouts_list.json
func ExampleRolloutsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewRolloutsClient("<subscription-id>", cred, nil)
	_, err = client.List(ctx,
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
