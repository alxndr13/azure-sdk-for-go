//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PoolList.json
func ExamplePoolClient_ListByBatchAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPoolClient("<subscription-id>", cred, nil)
	pager := client.ListByBatchAccount("<resource-group-name>",
		"<account-name>",
		&armbatch.PoolClientListByBatchAccountOptions{Maxresults: nil,
			Select: nil,
			Filter: nil,
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PoolCreate_SharedImageGallery.json
func ExamplePoolClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPoolClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pool-name>",
		armbatch.Pool{
			Properties: &armbatch.PoolProperties{
				DeploymentConfiguration: &armbatch.DeploymentConfiguration{
					VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
						ImageReference: &armbatch.ImageReference{
							ID: to.StringPtr("<id>"),
						},
						NodeAgentSKUID: to.StringPtr("<node-agent-skuid>"),
					},
				},
				VMSize: to.StringPtr("<vmsize>"),
			},
		},
		&armbatch.PoolClientCreateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PoolClientCreateResult)
}

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PoolUpdate_EnableAutoScale.json
func ExamplePoolClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPoolClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pool-name>",
		armbatch.Pool{
			Properties: &armbatch.PoolProperties{
				ScaleSettings: &armbatch.ScaleSettings{
					AutoScale: &armbatch.AutoScaleSettings{
						Formula: to.StringPtr("<formula>"),
					},
				},
			},
		},
		&armbatch.PoolClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PoolClientUpdateResult)
}

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PoolDelete.json
func ExamplePoolClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPoolClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pool-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PoolGet.json
func ExamplePoolClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPoolClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pool-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PoolClientGetResult)
}

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PoolDisableAutoScale.json
func ExamplePoolClient_DisableAutoScale() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPoolClient("<subscription-id>", cred, nil)
	res, err := client.DisableAutoScale(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pool-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PoolClientDisableAutoScaleResult)
}

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PoolStopResize.json
func ExamplePoolClient_StopResize() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPoolClient("<subscription-id>", cred, nil)
	res, err := client.StopResize(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pool-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PoolClientStopResizeResult)
}