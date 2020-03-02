package backup

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RecoveryPointsCrrClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type RecoveryPointsCrrClient struct {
	BaseClient
}

// NewRecoveryPointsCrrClient creates an instance of the RecoveryPointsCrrClient client.
func NewRecoveryPointsCrrClient(subscriptionID string) RecoveryPointsCrrClient {
	return NewRecoveryPointsCrrClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecoveryPointsCrrClientWithBaseURI creates an instance of the RecoveryPointsCrrClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewRecoveryPointsCrrClientWithBaseURI(baseURI string, subscriptionID string) RecoveryPointsCrrClient {
	return RecoveryPointsCrrClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the backup copies for the backed up item.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// fabricName - fabric name associated with the backed up item.
// containerName - container name associated with the backed up item.
// protectedItemName - backed up item whose backup copies are to be fetched.
// filter - oData filter options.
func (client RecoveryPointsCrrClient) List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, filter string) (result RecoveryPointResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoveryPointsCrrClient.List")
		defer func() {
			sc := -1
			if result.rprl.Response.Response != nil {
				sc = result.rprl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.RecoveryPointsCrrClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rprl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.RecoveryPointsCrrClient", "List", resp, "Failure sending request")
		return
	}

	result.rprl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.RecoveryPointsCrrClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RecoveryPointsCrrClient) ListPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2018-12-20"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RecoveryPointsCrrClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RecoveryPointsCrrClient) ListResponder(resp *http.Response) (result RecoveryPointResourceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RecoveryPointsCrrClient) listNextResults(ctx context.Context, lastResults RecoveryPointResourceList) (result RecoveryPointResourceList, err error) {
	req, err := lastResults.recoveryPointResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "backup.RecoveryPointsCrrClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "backup.RecoveryPointsCrrClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.RecoveryPointsCrrClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecoveryPointsCrrClient) ListComplete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, filter string) (result RecoveryPointResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoveryPointsCrrClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, filter)
	return
}
