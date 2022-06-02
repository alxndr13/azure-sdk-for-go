//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armworkloads

// MonitorsClientCreateResponse contains the response from method MonitorsClient.Create.
type MonitorsClientCreateResponse struct {
	Monitor
}

// MonitorsClientDeleteResponse contains the response from method MonitorsClient.Delete.
type MonitorsClientDeleteResponse struct {
	OperationStatusResult
}

// MonitorsClientGetResponse contains the response from method MonitorsClient.Get.
type MonitorsClientGetResponse struct {
	Monitor
}

// MonitorsClientListByResourceGroupResponse contains the response from method MonitorsClient.ListByResourceGroup.
type MonitorsClientListByResourceGroupResponse struct {
	MonitorListResult
}

// MonitorsClientListResponse contains the response from method MonitorsClient.List.
type MonitorsClientListResponse struct {
	MonitorListResult
}

// MonitorsClientUpdateResponse contains the response from method MonitorsClient.Update.
type MonitorsClientUpdateResponse struct {
	Monitor
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// ProviderInstancesClientCreateResponse contains the response from method ProviderInstancesClient.Create.
type ProviderInstancesClientCreateResponse struct {
	ProviderInstance
}

// ProviderInstancesClientDeleteResponse contains the response from method ProviderInstancesClient.Delete.
type ProviderInstancesClientDeleteResponse struct {
	OperationStatusResult
}

// ProviderInstancesClientGetResponse contains the response from method ProviderInstancesClient.Get.
type ProviderInstancesClientGetResponse struct {
	ProviderInstance
}

// ProviderInstancesClientListResponse contains the response from method ProviderInstancesClient.List.
type ProviderInstancesClientListResponse struct {
	ProviderInstanceListResult
}