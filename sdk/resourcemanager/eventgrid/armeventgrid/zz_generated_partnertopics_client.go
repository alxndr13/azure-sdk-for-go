//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PartnerTopicsClient contains the methods for the PartnerTopics group.
// Don't use this type directly, use NewPartnerTopicsClient() instead.
type PartnerTopicsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPartnerTopicsClient creates a new instance of PartnerTopicsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPartnerTopicsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PartnerTopicsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &PartnerTopicsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Activate - Activate a newly created partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// partnerTopicName - Name of the partner topic.
// options - PartnerTopicsClientActivateOptions contains the optional parameters for the PartnerTopicsClient.Activate method.
func (client *PartnerTopicsClient) Activate(ctx context.Context, resourceGroupName string, partnerTopicName string, options *PartnerTopicsClientActivateOptions) (PartnerTopicsClientActivateResponse, error) {
	req, err := client.activateCreateRequest(ctx, resourceGroupName, partnerTopicName, options)
	if err != nil {
		return PartnerTopicsClientActivateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerTopicsClientActivateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PartnerTopicsClientActivateResponse{}, runtime.NewResponseError(resp)
	}
	return client.activateHandleResponse(resp)
}

// activateCreateRequest creates the Activate request.
func (client *PartnerTopicsClient) activateCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, options *PartnerTopicsClientActivateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/activate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// activateHandleResponse handles the Activate response.
func (client *PartnerTopicsClient) activateHandleResponse(resp *http.Response) (PartnerTopicsClientActivateResponse, error) {
	result := PartnerTopicsClientActivateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerTopic); err != nil {
		return PartnerTopicsClientActivateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - Asynchronously creates a new partner topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// partnerTopicName - Name of the partner topic.
// partnerTopicInfo - Partner Topic information.
// options - PartnerTopicsClientCreateOrUpdateOptions contains the optional parameters for the PartnerTopicsClient.CreateOrUpdate
// method.
func (client *PartnerTopicsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, partnerTopicName string, partnerTopicInfo PartnerTopic, options *PartnerTopicsClientCreateOrUpdateOptions) (PartnerTopicsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, partnerTopicName, partnerTopicInfo, options)
	if err != nil {
		return PartnerTopicsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerTopicsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PartnerTopicsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PartnerTopicsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, partnerTopicInfo PartnerTopic, options *PartnerTopicsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, partnerTopicInfo)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PartnerTopicsClient) createOrUpdateHandleResponse(resp *http.Response) (PartnerTopicsClientCreateOrUpdateResponse, error) {
	result := PartnerTopicsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerTopic); err != nil {
		return PartnerTopicsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Deactivate - Deactivate specific partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// partnerTopicName - Name of the partner topic.
// options - PartnerTopicsClientDeactivateOptions contains the optional parameters for the PartnerTopicsClient.Deactivate
// method.
func (client *PartnerTopicsClient) Deactivate(ctx context.Context, resourceGroupName string, partnerTopicName string, options *PartnerTopicsClientDeactivateOptions) (PartnerTopicsClientDeactivateResponse, error) {
	req, err := client.deactivateCreateRequest(ctx, resourceGroupName, partnerTopicName, options)
	if err != nil {
		return PartnerTopicsClientDeactivateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerTopicsClientDeactivateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PartnerTopicsClientDeactivateResponse{}, runtime.NewResponseError(resp)
	}
	return client.deactivateHandleResponse(resp)
}

// deactivateCreateRequest creates the Deactivate request.
func (client *PartnerTopicsClient) deactivateCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, options *PartnerTopicsClientDeactivateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/deactivate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deactivateHandleResponse handles the Deactivate response.
func (client *PartnerTopicsClient) deactivateHandleResponse(resp *http.Response) (PartnerTopicsClientDeactivateResponse, error) {
	result := PartnerTopicsClientDeactivateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerTopic); err != nil {
		return PartnerTopicsClientDeactivateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete existing partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// partnerTopicName - Name of the partner topic.
// options - PartnerTopicsClientBeginDeleteOptions contains the optional parameters for the PartnerTopicsClient.BeginDelete
// method.
func (client *PartnerTopicsClient) BeginDelete(ctx context.Context, resourceGroupName string, partnerTopicName string, options *PartnerTopicsClientBeginDeleteOptions) (*armruntime.Poller[PartnerTopicsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, partnerTopicName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[PartnerTopicsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[PartnerTopicsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete existing partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PartnerTopicsClient) deleteOperation(ctx context.Context, resourceGroupName string, partnerTopicName string, options *PartnerTopicsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, partnerTopicName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PartnerTopicsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, options *PartnerTopicsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get properties of a partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// partnerTopicName - Name of the partner topic.
// options - PartnerTopicsClientGetOptions contains the optional parameters for the PartnerTopicsClient.Get method.
func (client *PartnerTopicsClient) Get(ctx context.Context, resourceGroupName string, partnerTopicName string, options *PartnerTopicsClientGetOptions) (PartnerTopicsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, partnerTopicName, options)
	if err != nil {
		return PartnerTopicsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerTopicsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PartnerTopicsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PartnerTopicsClient) getCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, options *PartnerTopicsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PartnerTopicsClient) getHandleResponse(resp *http.Response) (PartnerTopicsClientGetResponse, error) {
	result := PartnerTopicsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerTopic); err != nil {
		return PartnerTopicsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all the partner topics under a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// options - PartnerTopicsClientListByResourceGroupOptions contains the optional parameters for the PartnerTopicsClient.ListByResourceGroup
// method.
func (client *PartnerTopicsClient) NewListByResourceGroupPager(resourceGroupName string, options *PartnerTopicsClientListByResourceGroupOptions) *runtime.Pager[PartnerTopicsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[PartnerTopicsClientListByResourceGroupResponse]{
		More: func(page PartnerTopicsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PartnerTopicsClientListByResourceGroupResponse) (PartnerTopicsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PartnerTopicsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PartnerTopicsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PartnerTopicsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PartnerTopicsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PartnerTopicsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PartnerTopicsClient) listByResourceGroupHandleResponse(resp *http.Response) (PartnerTopicsClientListByResourceGroupResponse, error) {
	result := PartnerTopicsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerTopicsListResult); err != nil {
		return PartnerTopicsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all the partner topics under an Azure subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PartnerTopicsClientListBySubscriptionOptions contains the optional parameters for the PartnerTopicsClient.ListBySubscription
// method.
func (client *PartnerTopicsClient) NewListBySubscriptionPager(options *PartnerTopicsClientListBySubscriptionOptions) *runtime.Pager[PartnerTopicsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[PartnerTopicsClientListBySubscriptionResponse]{
		More: func(page PartnerTopicsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PartnerTopicsClientListBySubscriptionResponse) (PartnerTopicsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PartnerTopicsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PartnerTopicsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PartnerTopicsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PartnerTopicsClient) listBySubscriptionCreateRequest(ctx context.Context, options *PartnerTopicsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/partnerTopics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PartnerTopicsClient) listBySubscriptionHandleResponse(resp *http.Response) (PartnerTopicsClientListBySubscriptionResponse, error) {
	result := PartnerTopicsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerTopicsListResult); err != nil {
		return PartnerTopicsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Asynchronously updates a partner topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// partnerTopicName - Name of the partner topic.
// partnerTopicUpdateParameters - PartnerTopic update information.
// options - PartnerTopicsClientUpdateOptions contains the optional parameters for the PartnerTopicsClient.Update method.
func (client *PartnerTopicsClient) Update(ctx context.Context, resourceGroupName string, partnerTopicName string, partnerTopicUpdateParameters PartnerTopicUpdateParameters, options *PartnerTopicsClientUpdateOptions) (PartnerTopicsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, partnerTopicName, partnerTopicUpdateParameters, options)
	if err != nil {
		return PartnerTopicsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerTopicsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PartnerTopicsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PartnerTopicsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, partnerTopicUpdateParameters PartnerTopicUpdateParameters, options *PartnerTopicsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, partnerTopicUpdateParameters)
}

// updateHandleResponse handles the Update response.
func (client *PartnerTopicsClient) updateHandleResponse(resp *http.Response) (PartnerTopicsClientUpdateResponse, error) {
	result := PartnerTopicsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerTopic); err != nil {
		return PartnerTopicsClientUpdateResponse{}, err
	}
	return result, nil
}