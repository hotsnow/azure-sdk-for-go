// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// QueueServicesOperations contains the methods for the QueueServices group.
type QueueServicesOperations interface {
	// GetServiceProperties - Gets the properties of a storage account’s Queue service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
	GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (*QueueServicePropertiesResponse, error)
	// List - List all queue services for the storage account
	List(ctx context.Context, resourceGroupName string, accountName string) (*ListQueueServicesResponse, error)
	// SetServiceProperties - Sets the properties of a storage account’s Queue service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
	SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters QueueServiceProperties) (*QueueServicePropertiesResponse, error)
}

// queueServicesOperations implements the QueueServicesOperations interface.
type queueServicesOperations struct {
	*Client
	subscriptionID string
}

// GetServiceProperties - Gets the properties of a storage account’s Queue service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
func (client *queueServicesOperations) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (*QueueServicePropertiesResponse, error) {
	req, err := client.getServicePropertiesCreateRequest(resourceGroupName, accountName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getServicePropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getServicePropertiesCreateRequest creates the GetServiceProperties request.
func (client *queueServicesOperations) getServicePropertiesCreateRequest(resourceGroupName string, accountName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/{queueServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueServiceName}", url.PathEscape("default"))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-06-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getServicePropertiesHandleResponse handles the GetServiceProperties response.
func (client *queueServicesOperations) getServicePropertiesHandleResponse(resp *azcore.Response) (*QueueServicePropertiesResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getServicePropertiesHandleError(resp)
	}
	result := QueueServicePropertiesResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.QueueServiceProperties)
}

// getServicePropertiesHandleError handles the GetServiceProperties error response.
func (client *queueServicesOperations) getServicePropertiesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - List all queue services for the storage account
func (client *queueServicesOperations) List(ctx context.Context, resourceGroupName string, accountName string) (*ListQueueServicesResponse, error) {
	req, err := client.listCreateRequest(resourceGroupName, accountName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// listCreateRequest creates the List request.
func (client *queueServicesOperations) listCreateRequest(resourceGroupName string, accountName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-06-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *queueServicesOperations) listHandleResponse(resp *azcore.Response) (*ListQueueServicesResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ListQueueServicesResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListQueueServices)
}

// listHandleError handles the List error response.
func (client *queueServicesOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// SetServiceProperties - Sets the properties of a storage account’s Queue service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
func (client *queueServicesOperations) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters QueueServiceProperties) (*QueueServicePropertiesResponse, error) {
	req, err := client.setServicePropertiesCreateRequest(resourceGroupName, accountName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.setServicePropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// setServicePropertiesCreateRequest creates the SetServiceProperties request.
func (client *queueServicesOperations) setServicePropertiesCreateRequest(resourceGroupName string, accountName string, parameters QueueServiceProperties) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/{queueServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueServiceName}", url.PathEscape("default"))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-06-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// setServicePropertiesHandleResponse handles the SetServiceProperties response.
func (client *queueServicesOperations) setServicePropertiesHandleResponse(resp *azcore.Response) (*QueueServicePropertiesResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.setServicePropertiesHandleError(resp)
	}
	result := QueueServicePropertiesResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.QueueServiceProperties)
}

// setServicePropertiesHandleError handles the SetServiceProperties error response.
func (client *queueServicesOperations) setServicePropertiesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}