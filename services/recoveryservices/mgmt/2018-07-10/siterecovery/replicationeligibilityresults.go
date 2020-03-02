package siterecovery

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

// ReplicationEligibilityResultsClient is the client for the ReplicationEligibilityResults methods of the Siterecovery
// service.
type ReplicationEligibilityResultsClient struct {
	BaseClient
}

// NewReplicationEligibilityResultsClient creates an instance of the ReplicationEligibilityResultsClient client.
func NewReplicationEligibilityResultsClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationEligibilityResultsClient {
	return NewReplicationEligibilityResultsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationEligibilityResultsClientWithBaseURI creates an instance of the ReplicationEligibilityResultsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewReplicationEligibilityResultsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationEligibilityResultsClient {
	return ReplicationEligibilityResultsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Get validates whether a given VM can be protected or not in which case returns list of errors.
// Parameters:
// virtualMachineName - virtual Machine name.
func (client ReplicationEligibilityResultsClient) Get(ctx context.Context, virtualMachineName string) (result ReplicationEligibilityResultsGetFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationEligibilityResultsClient.Get")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, virtualMachineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEligibilityResultsClient", "Get", nil, "Failure preparing request")
		return
	}

	result, err = client.GetSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEligibilityResultsClient", "Get", result.Response(), "Failure sending request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationEligibilityResultsClient) GetPreparer(ctx context.Context, virtualMachineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2018-07-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/providers/Microsoft.RecoveryServices/replicationEligibilityResults/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationEligibilityResultsClient) GetSender(req *http.Request) (future ReplicationEligibilityResultsGetFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationEligibilityResultsClient) GetResponder(resp *http.Response) (result ReplicationEligibilityResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List validates whether a given VM can be protected or not in which case returns list of errors.
// Parameters:
// virtualMachineName - virtual Machine name.
func (client ReplicationEligibilityResultsClient) List(ctx context.Context, virtualMachineName string) (result ReplicationEligibilityResultsListFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationEligibilityResultsClient.List")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, virtualMachineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEligibilityResultsClient", "List", nil, "Failure preparing request")
		return
	}

	result, err = client.ListSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEligibilityResultsClient", "List", result.Response(), "Failure sending request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationEligibilityResultsClient) ListPreparer(ctx context.Context, virtualMachineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2018-07-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/providers/Microsoft.RecoveryServices/replicationEligibilityResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationEligibilityResultsClient) ListSender(req *http.Request) (future ReplicationEligibilityResultsListFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationEligibilityResultsClient) ListResponder(resp *http.Response) (result ReplicationEligibilityResultsCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
