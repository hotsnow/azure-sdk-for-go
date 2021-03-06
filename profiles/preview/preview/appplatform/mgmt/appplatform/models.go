// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package appplatform

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/appplatform/mgmt/2019-05-01-preview/appplatform"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AppResourceProvisioningState = original.AppResourceProvisioningState

const (
	Creating  AppResourceProvisioningState = original.Creating
	Failed    AppResourceProvisioningState = original.Failed
	Succeeded AppResourceProvisioningState = original.Succeeded
	Updating  AppResourceProvisioningState = original.Updating
)

type ConfigServerState = original.ConfigServerState

const (
	ConfigServerStateDeleted      ConfigServerState = original.ConfigServerStateDeleted
	ConfigServerStateFailed       ConfigServerState = original.ConfigServerStateFailed
	ConfigServerStateNotAvailable ConfigServerState = original.ConfigServerStateNotAvailable
	ConfigServerStateSucceeded    ConfigServerState = original.ConfigServerStateSucceeded
	ConfigServerStateUpdating     ConfigServerState = original.ConfigServerStateUpdating
)

type DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningState

const (
	DeploymentResourceProvisioningStateCreating  DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateCreating
	DeploymentResourceProvisioningStateFailed    DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateFailed
	DeploymentResourceProvisioningStateSucceeded DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateSucceeded
	DeploymentResourceProvisioningStateUpdating  DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateUpdating
)

type DeploymentResourceStatus = original.DeploymentResourceStatus

const (
	DeploymentResourceStatusAllocating DeploymentResourceStatus = original.DeploymentResourceStatusAllocating
	DeploymentResourceStatusCompiling  DeploymentResourceStatus = original.DeploymentResourceStatusCompiling
	DeploymentResourceStatusFailed     DeploymentResourceStatus = original.DeploymentResourceStatusFailed
	DeploymentResourceStatusRunning    DeploymentResourceStatus = original.DeploymentResourceStatusRunning
	DeploymentResourceStatusStopped    DeploymentResourceStatus = original.DeploymentResourceStatusStopped
	DeploymentResourceStatusUnknown    DeploymentResourceStatus = original.DeploymentResourceStatusUnknown
	DeploymentResourceStatusUpgrading  DeploymentResourceStatus = original.DeploymentResourceStatusUpgrading
)

type ManagedIdentityType = original.ManagedIdentityType

const (
	None                       ManagedIdentityType = original.None
	SystemAssigned             ManagedIdentityType = original.SystemAssigned
	SystemAssignedUserAssigned ManagedIdentityType = original.SystemAssignedUserAssigned
	UserAssigned               ManagedIdentityType = original.UserAssigned
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCreating   ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleted    ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting   ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed     ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateMoved      ProvisioningState = original.ProvisioningStateMoved
	ProvisioningStateMoveFailed ProvisioningState = original.ProvisioningStateMoveFailed
	ProvisioningStateMoving     ProvisioningState = original.ProvisioningStateMoving
	ProvisioningStateSucceeded  ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating   ProvisioningState = original.ProvisioningStateUpdating
)

type ResourceSkuRestrictionsReasonCode = original.ResourceSkuRestrictionsReasonCode

const (
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = original.NotAvailableForSubscription
	QuotaID                     ResourceSkuRestrictionsReasonCode = original.QuotaID
)

type ResourceSkuRestrictionsType = original.ResourceSkuRestrictionsType

const (
	Location ResourceSkuRestrictionsType = original.Location
	Zone     ResourceSkuRestrictionsType = original.Zone
)

type RuntimeVersion = original.RuntimeVersion

const (
	Java11 RuntimeVersion = original.Java11
	Java8  RuntimeVersion = original.Java8
)

type SkuScaleType = original.SkuScaleType

const (
	SkuScaleTypeAutomatic SkuScaleType = original.SkuScaleTypeAutomatic
	SkuScaleTypeManual    SkuScaleType = original.SkuScaleTypeManual
	SkuScaleTypeNone      SkuScaleType = original.SkuScaleTypeNone
)

type TestKeyType = original.TestKeyType

const (
	Primary   TestKeyType = original.Primary
	Secondary TestKeyType = original.Secondary
)

type TraceProxyState = original.TraceProxyState

const (
	TraceProxyStateFailed       TraceProxyState = original.TraceProxyStateFailed
	TraceProxyStateNotAvailable TraceProxyState = original.TraceProxyStateNotAvailable
	TraceProxyStateSucceeded    TraceProxyState = original.TraceProxyStateSucceeded
	TraceProxyStateUpdating     TraceProxyState = original.TraceProxyStateUpdating
)

type UserSourceType = original.UserSourceType

const (
	Jar    UserSourceType = original.Jar
	Source UserSourceType = original.Source
)

type AppResource = original.AppResource
type AppResourceCollection = original.AppResourceCollection
type AppResourceCollectionIterator = original.AppResourceCollectionIterator
type AppResourceCollectionPage = original.AppResourceCollectionPage
type AppResourceProperties = original.AppResourceProperties
type AppsClient = original.AppsClient
type AppsCreateOrUpdateFuture = original.AppsCreateOrUpdateFuture
type AppsUpdateFuture = original.AppsUpdateFuture
type AvailableOperations = original.AvailableOperations
type AvailableOperationsIterator = original.AvailableOperationsIterator
type AvailableOperationsPage = original.AvailableOperationsPage
type BaseClient = original.BaseClient
type BindingResource = original.BindingResource
type BindingResourceCollection = original.BindingResourceCollection
type BindingResourceCollectionIterator = original.BindingResourceCollectionIterator
type BindingResourceCollectionPage = original.BindingResourceCollectionPage
type BindingResourceProperties = original.BindingResourceProperties
type BindingsClient = original.BindingsClient
type CertificateProperties = original.CertificateProperties
type CertificateResource = original.CertificateResource
type CertificateResourceCollection = original.CertificateResourceCollection
type CertificateResourceCollectionIterator = original.CertificateResourceCollectionIterator
type CertificateResourceCollectionPage = original.CertificateResourceCollectionPage
type CertificatesClient = original.CertificatesClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ClusterResourceProperties = original.ClusterResourceProperties
type ConfigServerGitProperty = original.ConfigServerGitProperty
type ConfigServerProperties = original.ConfigServerProperties
type ConfigServerSettings = original.ConfigServerSettings
type CustomDomainProperties = original.CustomDomainProperties
type CustomDomainResource = original.CustomDomainResource
type CustomDomainResourceCollection = original.CustomDomainResourceCollection
type CustomDomainResourceCollectionIterator = original.CustomDomainResourceCollectionIterator
type CustomDomainResourceCollectionPage = original.CustomDomainResourceCollectionPage
type CustomDomainValidatePayload = original.CustomDomainValidatePayload
type CustomDomainValidateResult = original.CustomDomainValidateResult
type CustomDomainsClient = original.CustomDomainsClient
type DeploymentInstance = original.DeploymentInstance
type DeploymentResource = original.DeploymentResource
type DeploymentResourceCollection = original.DeploymentResourceCollection
type DeploymentResourceCollectionIterator = original.DeploymentResourceCollectionIterator
type DeploymentResourceCollectionPage = original.DeploymentResourceCollectionPage
type DeploymentResourceProperties = original.DeploymentResourceProperties
type DeploymentSettings = original.DeploymentSettings
type DeploymentsClient = original.DeploymentsClient
type DeploymentsCreateOrUpdateFuture = original.DeploymentsCreateOrUpdateFuture
type DeploymentsRestartFuture = original.DeploymentsRestartFuture
type DeploymentsStartFuture = original.DeploymentsStartFuture
type DeploymentsStopFuture = original.DeploymentsStopFuture
type DeploymentsUpdateFuture = original.DeploymentsUpdateFuture
type Error = original.Error
type GitPatternRepository = original.GitPatternRepository
type LogFileURLResponse = original.LogFileURLResponse
type LogSpecification = original.LogSpecification
type ManagedIdentityProperties = original.ManagedIdentityProperties
type MetricDimension = original.MetricDimension
type MetricSpecification = original.MetricSpecification
type NameAvailability = original.NameAvailability
type NameAvailabilityParameters = original.NameAvailabilityParameters
type NetworkProfile = original.NetworkProfile
type OperationDetail = original.OperationDetail
type OperationDisplay = original.OperationDisplay
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PersistentDisk = original.PersistentDisk
type ProxyResource = original.ProxyResource
type RegenerateTestKeyRequestPayload = original.RegenerateTestKeyRequestPayload
type Resource = original.Resource
type ResourceSku = original.ResourceSku
type ResourceSkuCapabilities = original.ResourceSkuCapabilities
type ResourceSkuCollection = original.ResourceSkuCollection
type ResourceSkuCollectionIterator = original.ResourceSkuCollectionIterator
type ResourceSkuCollectionPage = original.ResourceSkuCollectionPage
type ResourceSkuLocationInfo = original.ResourceSkuLocationInfo
type ResourceSkuRestrictionInfo = original.ResourceSkuRestrictionInfo
type ResourceSkuRestrictions = original.ResourceSkuRestrictions
type ResourceSkuZoneDetails = original.ResourceSkuZoneDetails
type ResourceUploadDefinition = original.ResourceUploadDefinition
type ServiceResource = original.ServiceResource
type ServiceResourceList = original.ServiceResourceList
type ServiceResourceListIterator = original.ServiceResourceListIterator
type ServiceResourceListPage = original.ServiceResourceListPage
type ServiceSpecification = original.ServiceSpecification
type ServicesClient = original.ServicesClient
type ServicesCreateOrUpdateFuture = original.ServicesCreateOrUpdateFuture
type ServicesDeleteFuture = original.ServicesDeleteFuture
type ServicesUpdateFuture = original.ServicesUpdateFuture
type Sku = original.Sku
type SkuCapacity = original.SkuCapacity
type SkuClient = original.SkuClient
type TemporaryDisk = original.TemporaryDisk
type TestKeys = original.TestKeys
type TraceProperties = original.TraceProperties
type TrackedResource = original.TrackedResource
type UserSourceInfo = original.UserSourceInfo

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAppResourceCollectionIterator(page AppResourceCollectionPage) AppResourceCollectionIterator {
	return original.NewAppResourceCollectionIterator(page)
}
func NewAppResourceCollectionPage(getNextPage func(context.Context, AppResourceCollection) (AppResourceCollection, error)) AppResourceCollectionPage {
	return original.NewAppResourceCollectionPage(getNextPage)
}
func NewAppsClient(subscriptionID string) AppsClient {
	return original.NewAppsClient(subscriptionID)
}
func NewAppsClientWithBaseURI(baseURI string, subscriptionID string) AppsClient {
	return original.NewAppsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAvailableOperationsIterator(page AvailableOperationsPage) AvailableOperationsIterator {
	return original.NewAvailableOperationsIterator(page)
}
func NewAvailableOperationsPage(getNextPage func(context.Context, AvailableOperations) (AvailableOperations, error)) AvailableOperationsPage {
	return original.NewAvailableOperationsPage(getNextPage)
}
func NewBindingResourceCollectionIterator(page BindingResourceCollectionPage) BindingResourceCollectionIterator {
	return original.NewBindingResourceCollectionIterator(page)
}
func NewBindingResourceCollectionPage(getNextPage func(context.Context, BindingResourceCollection) (BindingResourceCollection, error)) BindingResourceCollectionPage {
	return original.NewBindingResourceCollectionPage(getNextPage)
}
func NewBindingsClient(subscriptionID string) BindingsClient {
	return original.NewBindingsClient(subscriptionID)
}
func NewBindingsClientWithBaseURI(baseURI string, subscriptionID string) BindingsClient {
	return original.NewBindingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCertificateResourceCollectionIterator(page CertificateResourceCollectionPage) CertificateResourceCollectionIterator {
	return original.NewCertificateResourceCollectionIterator(page)
}
func NewCertificateResourceCollectionPage(getNextPage func(context.Context, CertificateResourceCollection) (CertificateResourceCollection, error)) CertificateResourceCollectionPage {
	return original.NewCertificateResourceCollectionPage(getNextPage)
}
func NewCertificatesClient(subscriptionID string) CertificatesClient {
	return original.NewCertificatesClient(subscriptionID)
}
func NewCertificatesClientWithBaseURI(baseURI string, subscriptionID string) CertificatesClient {
	return original.NewCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewCustomDomainResourceCollectionIterator(page CustomDomainResourceCollectionPage) CustomDomainResourceCollectionIterator {
	return original.NewCustomDomainResourceCollectionIterator(page)
}
func NewCustomDomainResourceCollectionPage(getNextPage func(context.Context, CustomDomainResourceCollection) (CustomDomainResourceCollection, error)) CustomDomainResourceCollectionPage {
	return original.NewCustomDomainResourceCollectionPage(getNextPage)
}
func NewCustomDomainsClient(subscriptionID string) CustomDomainsClient {
	return original.NewCustomDomainsClient(subscriptionID)
}
func NewCustomDomainsClientWithBaseURI(baseURI string, subscriptionID string) CustomDomainsClient {
	return original.NewCustomDomainsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeploymentResourceCollectionIterator(page DeploymentResourceCollectionPage) DeploymentResourceCollectionIterator {
	return original.NewDeploymentResourceCollectionIterator(page)
}
func NewDeploymentResourceCollectionPage(getNextPage func(context.Context, DeploymentResourceCollection) (DeploymentResourceCollection, error)) DeploymentResourceCollectionPage {
	return original.NewDeploymentResourceCollectionPage(getNextPage)
}
func NewDeploymentsClient(subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClient(subscriptionID)
}
func NewDeploymentsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceSkuCollectionIterator(page ResourceSkuCollectionPage) ResourceSkuCollectionIterator {
	return original.NewResourceSkuCollectionIterator(page)
}
func NewResourceSkuCollectionPage(getNextPage func(context.Context, ResourceSkuCollection) (ResourceSkuCollection, error)) ResourceSkuCollectionPage {
	return original.NewResourceSkuCollectionPage(getNextPage)
}
func NewServiceResourceListIterator(page ServiceResourceListPage) ServiceResourceListIterator {
	return original.NewServiceResourceListIterator(page)
}
func NewServiceResourceListPage(getNextPage func(context.Context, ServiceResourceList) (ServiceResourceList, error)) ServiceResourceListPage {
	return original.NewServiceResourceListPage(getNextPage)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSkuClient(subscriptionID string) SkuClient {
	return original.NewSkuClient(subscriptionID)
}
func NewSkuClientWithBaseURI(baseURI string, subscriptionID string) SkuClient {
	return original.NewSkuClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAppResourceProvisioningStateValues() []AppResourceProvisioningState {
	return original.PossibleAppResourceProvisioningStateValues()
}
func PossibleConfigServerStateValues() []ConfigServerState {
	return original.PossibleConfigServerStateValues()
}
func PossibleDeploymentResourceProvisioningStateValues() []DeploymentResourceProvisioningState {
	return original.PossibleDeploymentResourceProvisioningStateValues()
}
func PossibleDeploymentResourceStatusValues() []DeploymentResourceStatus {
	return original.PossibleDeploymentResourceStatusValues()
}
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return original.PossibleManagedIdentityTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return original.PossibleResourceSkuRestrictionsReasonCodeValues()
}
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return original.PossibleResourceSkuRestrictionsTypeValues()
}
func PossibleRuntimeVersionValues() []RuntimeVersion {
	return original.PossibleRuntimeVersionValues()
}
func PossibleSkuScaleTypeValues() []SkuScaleType {
	return original.PossibleSkuScaleTypeValues()
}
func PossibleTestKeyTypeValues() []TestKeyType {
	return original.PossibleTestKeyTypeValues()
}
func PossibleTraceProxyStateValues() []TraceProxyState {
	return original.PossibleTraceProxyStateValues()
}
func PossibleUserSourceTypeValues() []UserSourceType {
	return original.PossibleUserSourceTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
