package compute

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// GalleryImageVersionsClient is the compute Client
type GalleryImageVersionsClient struct {
	BaseClient
}

// NewGalleryImageVersionsClient creates an instance of the GalleryImageVersionsClient client.
func NewGalleryImageVersionsClient(subscriptionID string) GalleryImageVersionsClient {
	return NewGalleryImageVersionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGalleryImageVersionsClientWithBaseURI creates an instance of the GalleryImageVersionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewGalleryImageVersionsClientWithBaseURI(baseURI string, subscriptionID string) GalleryImageVersionsClient {
	return GalleryImageVersionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a gallery Image Version.
// Parameters:
// resourceGroupName - the name of the resource group.
// galleryName - the name of the Shared Image Gallery in which the Image Definition resides.
// galleryImageName - the name of the gallery Image Definition in which the Image Version is to be created.
// galleryImageVersionName - the name of the gallery Image Version to be created. Needs to follow semantic
// version name pattern: The allowed characters are digit and period. Digits must be within the range of a
// 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
// galleryImageVersion - parameters supplied to the create or update gallery Image Version operation.
func (client GalleryImageVersionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersion) (result GalleryImageVersionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImageVersionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: galleryImageVersion,
			Constraints: []validation.Constraint{{Target: "galleryImageVersion.GalleryImageVersionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "galleryImageVersion.GalleryImageVersionProperties.PublishingProfile", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "galleryImageVersion.GalleryImageVersionProperties.PublishingProfile.Source", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "galleryImageVersion.GalleryImageVersionProperties.PublishingProfile.Source.ManagedImage", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "galleryImageVersion.GalleryImageVersionProperties.PublishingProfile.Source.ManagedImage.ID", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("compute.GalleryImageVersionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, galleryName, galleryImageName, galleryImageVersionName, galleryImageVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GalleryImageVersionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersion) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":        autorest.Encode("path", galleryImageName),
		"galleryImageVersionName": autorest.Encode("path", galleryImageVersionName),
		"galleryName":             autorest.Encode("path", galleryName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}", pathParameters),
		autorest.WithJSON(galleryImageVersion),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryImageVersionsClient) CreateOrUpdateSender(req *http.Request) (future GalleryImageVersionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GalleryImageVersionsClient) CreateOrUpdateResponder(resp *http.Response) (result GalleryImageVersion, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a gallery Image Version.
// Parameters:
// resourceGroupName - the name of the resource group.
// galleryName - the name of the Shared Image Gallery in which the Image Definition resides.
// galleryImageName - the name of the gallery Image Definition in which the Image Version resides.
// galleryImageVersionName - the name of the gallery Image Version to be deleted.
func (client GalleryImageVersionsClient) Delete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string) (result GalleryImageVersionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImageVersionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, galleryName, galleryImageName, galleryImageVersionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GalleryImageVersionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":        autorest.Encode("path", galleryImageName),
		"galleryImageVersionName": autorest.Encode("path", galleryImageVersionName),
		"galleryName":             autorest.Encode("path", galleryName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryImageVersionsClient) DeleteSender(req *http.Request) (future GalleryImageVersionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GalleryImageVersionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves information about a gallery Image Version.
// Parameters:
// resourceGroupName - the name of the resource group.
// galleryName - the name of the Shared Image Gallery in which the Image Definition resides.
// galleryImageName - the name of the gallery Image Definition in which the Image Version resides.
// galleryImageVersionName - the name of the gallery Image Version to be retrieved.
// expand - the expand expression to apply on the operation.
func (client GalleryImageVersionsClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, expand ReplicationStatusTypes) (result GalleryImageVersion, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImageVersionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, galleryName, galleryImageName, galleryImageVersionName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GalleryImageVersionsClient) GetPreparer(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, expand ReplicationStatusTypes) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":        autorest.Encode("path", galleryImageName),
		"galleryImageVersionName": autorest.Encode("path", galleryImageVersionName),
		"galleryName":             autorest.Encode("path", galleryName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryImageVersionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GalleryImageVersionsClient) GetResponder(resp *http.Response) (result GalleryImageVersion, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByGalleryImage list gallery Image Versions in a gallery Image Definition.
// Parameters:
// resourceGroupName - the name of the resource group.
// galleryName - the name of the Shared Image Gallery in which the Image Definition resides.
// galleryImageName - the name of the Shared Image Gallery Image Definition from which the Image Versions are
// to be listed.
func (client GalleryImageVersionsClient) ListByGalleryImage(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string) (result GalleryImageVersionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImageVersionsClient.ListByGalleryImage")
		defer func() {
			sc := -1
			if result.givl.Response.Response != nil {
				sc = result.givl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByGalleryImageNextResults
	req, err := client.ListByGalleryImagePreparer(ctx, resourceGroupName, galleryName, galleryImageName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "ListByGalleryImage", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByGalleryImageSender(req)
	if err != nil {
		result.givl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "ListByGalleryImage", resp, "Failure sending request")
		return
	}

	result.givl, err = client.ListByGalleryImageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "ListByGalleryImage", resp, "Failure responding to request")
	}

	return
}

// ListByGalleryImagePreparer prepares the ListByGalleryImage request.
func (client GalleryImageVersionsClient) ListByGalleryImagePreparer(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":  autorest.Encode("path", galleryImageName),
		"galleryName":       autorest.Encode("path", galleryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByGalleryImageSender sends the ListByGalleryImage request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryImageVersionsClient) ListByGalleryImageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByGalleryImageResponder handles the response to the ListByGalleryImage request. The method always
// closes the http.Response Body.
func (client GalleryImageVersionsClient) ListByGalleryImageResponder(resp *http.Response) (result GalleryImageVersionList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByGalleryImageNextResults retrieves the next set of results, if any.
func (client GalleryImageVersionsClient) listByGalleryImageNextResults(ctx context.Context, lastResults GalleryImageVersionList) (result GalleryImageVersionList, err error) {
	req, err := lastResults.galleryImageVersionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "listByGalleryImageNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByGalleryImageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "listByGalleryImageNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByGalleryImageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryImageVersionsClient", "listByGalleryImageNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByGalleryImageComplete enumerates all values, automatically crossing page boundaries as required.
func (client GalleryImageVersionsClient) ListByGalleryImageComplete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string) (result GalleryImageVersionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImageVersionsClient.ListByGalleryImage")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByGalleryImage(ctx, resourceGroupName, galleryName, galleryImageName)
	return
}
