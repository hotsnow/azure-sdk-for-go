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

package desktopvirtualization

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/desktopvirtualization/mgmt/2019-12-10-preview/desktopvirtualization"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ApplicationGroupType = original.ApplicationGroupType

const (
	ApplicationGroupTypeDesktop   ApplicationGroupType = original.ApplicationGroupTypeDesktop
	ApplicationGroupTypeRemoteApp ApplicationGroupType = original.ApplicationGroupTypeRemoteApp
)

type ApplicationType = original.ApplicationType

const (
	ApplicationTypeDesktop   ApplicationType = original.ApplicationTypeDesktop
	ApplicationTypeRemoteApp ApplicationType = original.ApplicationTypeRemoteApp
)

type CommandLineSetting = original.CommandLineSetting

const (
	Allow      CommandLineSetting = original.Allow
	DoNotAllow CommandLineSetting = original.DoNotAllow
	Require    CommandLineSetting = original.Require
)

type HostPoolType = original.HostPoolType

const (
	Personal HostPoolType = original.Personal
	Pooled   HostPoolType = original.Pooled
)

type LoadBalancerType = original.LoadBalancerType

const (
	BreadthFirst LoadBalancerType = original.BreadthFirst
	DepthFirst   LoadBalancerType = original.DepthFirst
	Persistent   LoadBalancerType = original.Persistent
)

type PersonalDesktopAssignmentType = original.PersonalDesktopAssignmentType

const (
	Automatic PersonalDesktopAssignmentType = original.Automatic
	Direct    PersonalDesktopAssignmentType = original.Direct
)

type RegistrationTokenOperation = original.RegistrationTokenOperation

const (
	Delete RegistrationTokenOperation = original.Delete
	None   RegistrationTokenOperation = original.None
	Update RegistrationTokenOperation = original.Update
)

type SessionState = original.SessionState

const (
	Active                 SessionState = original.Active
	Disconnected           SessionState = original.Disconnected
	LogOff                 SessionState = original.LogOff
	Pending                SessionState = original.Pending
	Unknown                SessionState = original.Unknown
	UserProfileDiskMounted SessionState = original.UserProfileDiskMounted
)

type Status = original.Status

const (
	StatusAvailable     Status = original.StatusAvailable
	StatusDisconnected  Status = original.StatusDisconnected
	StatusShutdown      Status = original.StatusShutdown
	StatusUnavailable   Status = original.StatusUnavailable
	StatusUpgradeFailed Status = original.StatusUpgradeFailed
	StatusUpgrading     Status = original.StatusUpgrading
)

type UpdateState = original.UpdateState

const (
	UpdateStateFailed    UpdateState = original.UpdateStateFailed
	UpdateStateInitial   UpdateState = original.UpdateStateInitial
	UpdateStatePending   UpdateState = original.UpdateStatePending
	UpdateStateStarted   UpdateState = original.UpdateStateStarted
	UpdateStateSucceeded UpdateState = original.UpdateStateSucceeded
)

type ActiveApplicationsClient = original.ActiveApplicationsClient
type Application = original.Application
type ApplicationGroup = original.ApplicationGroup
type ApplicationGroupAssignmentsClient = original.ApplicationGroupAssignmentsClient
type ApplicationGroupList = original.ApplicationGroupList
type ApplicationGroupListIterator = original.ApplicationGroupListIterator
type ApplicationGroupListPage = original.ApplicationGroupListPage
type ApplicationGroupPatch = original.ApplicationGroupPatch
type ApplicationGroupPatchProperties = original.ApplicationGroupPatchProperties
type ApplicationGroupProperties = original.ApplicationGroupProperties
type ApplicationGroupsClient = original.ApplicationGroupsClient
type ApplicationList = original.ApplicationList
type ApplicationListIterator = original.ApplicationListIterator
type ApplicationListPage = original.ApplicationListPage
type ApplicationPatch = original.ApplicationPatch
type ApplicationPatchProperties = original.ApplicationPatchProperties
type ApplicationProperties = original.ApplicationProperties
type ApplicationsClient = original.ApplicationsClient
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CloudError = original.CloudError
type Desktop = original.Desktop
type DesktopList = original.DesktopList
type DesktopPatch = original.DesktopPatch
type DesktopPatchProperties = original.DesktopPatchProperties
type DesktopProperties = original.DesktopProperties
type DesktopsClient = original.DesktopsClient
type HostPool = original.HostPool
type HostPoolList = original.HostPoolList
type HostPoolListIterator = original.HostPoolListIterator
type HostPoolListPage = original.HostPoolListPage
type HostPoolPatch = original.HostPoolPatch
type HostPoolPatchProperties = original.HostPoolPatchProperties
type HostPoolProperties = original.HostPoolProperties
type HostPoolsClient = original.HostPoolsClient
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type RegistrationInfo = original.RegistrationInfo
type RegistrationInfoPatch = original.RegistrationInfoPatch
type Resource = original.Resource
type ResourceProviderOperation = original.ResourceProviderOperation
type ResourceProviderOperationDisplay = original.ResourceProviderOperationDisplay
type ResourceProviderOperationList = original.ResourceProviderOperationList
type SendMessage = original.SendMessage
type SessionHost = original.SessionHost
type SessionHostList = original.SessionHostList
type SessionHostListIterator = original.SessionHostListIterator
type SessionHostListPage = original.SessionHostListPage
type SessionHostPatch = original.SessionHostPatch
type SessionHostPatchProperties = original.SessionHostPatchProperties
type SessionHostProperties = original.SessionHostProperties
type SessionHostsClient = original.SessionHostsClient
type StartMenuItem = original.StartMenuItem
type StartMenuItemList = original.StartMenuItemList
type StartMenuItemListIterator = original.StartMenuItemListIterator
type StartMenuItemListPage = original.StartMenuItemListPage
type StartMenuItemProperties = original.StartMenuItemProperties
type StartMenuItemsClient = original.StartMenuItemsClient
type TrackedResource = original.TrackedResource
type UserSession = original.UserSession
type UserSessionList = original.UserSessionList
type UserSessionListIterator = original.UserSessionListIterator
type UserSessionListPage = original.UserSessionListPage
type UserSessionProperties = original.UserSessionProperties
type UserSessionsClient = original.UserSessionsClient
type Workspace = original.Workspace
type WorkspaceList = original.WorkspaceList
type WorkspaceListIterator = original.WorkspaceListIterator
type WorkspaceListPage = original.WorkspaceListPage
type WorkspacePatch = original.WorkspacePatch
type WorkspacePatchProperties = original.WorkspacePatchProperties
type WorkspaceProperties = original.WorkspaceProperties
type WorkspacesClient = original.WorkspacesClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewActiveApplicationsClient(subscriptionID string) ActiveApplicationsClient {
	return original.NewActiveApplicationsClient(subscriptionID)
}
func NewActiveApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ActiveApplicationsClient {
	return original.NewActiveApplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationGroupAssignmentsClient(subscriptionID string) ApplicationGroupAssignmentsClient {
	return original.NewApplicationGroupAssignmentsClient(subscriptionID)
}
func NewApplicationGroupAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGroupAssignmentsClient {
	return original.NewApplicationGroupAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationGroupListIterator(page ApplicationGroupListPage) ApplicationGroupListIterator {
	return original.NewApplicationGroupListIterator(page)
}
func NewApplicationGroupListPage(getNextPage func(context.Context, ApplicationGroupList) (ApplicationGroupList, error)) ApplicationGroupListPage {
	return original.NewApplicationGroupListPage(getNextPage)
}
func NewApplicationGroupsClient(subscriptionID string) ApplicationGroupsClient {
	return original.NewApplicationGroupsClient(subscriptionID)
}
func NewApplicationGroupsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGroupsClient {
	return original.NewApplicationGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationListIterator(page ApplicationListPage) ApplicationListIterator {
	return original.NewApplicationListIterator(page)
}
func NewApplicationListPage(getNextPage func(context.Context, ApplicationList) (ApplicationList, error)) ApplicationListPage {
	return original.NewApplicationListPage(getNextPage)
}
func NewApplicationsClient(subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClient(subscriptionID)
}
func NewApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDesktopsClient(subscriptionID string) DesktopsClient {
	return original.NewDesktopsClient(subscriptionID)
}
func NewDesktopsClientWithBaseURI(baseURI string, subscriptionID string) DesktopsClient {
	return original.NewDesktopsClientWithBaseURI(baseURI, subscriptionID)
}
func NewHostPoolListIterator(page HostPoolListPage) HostPoolListIterator {
	return original.NewHostPoolListIterator(page)
}
func NewHostPoolListPage(getNextPage func(context.Context, HostPoolList) (HostPoolList, error)) HostPoolListPage {
	return original.NewHostPoolListPage(getNextPage)
}
func NewHostPoolsClient(subscriptionID string) HostPoolsClient {
	return original.NewHostPoolsClient(subscriptionID)
}
func NewHostPoolsClientWithBaseURI(baseURI string, subscriptionID string) HostPoolsClient {
	return original.NewHostPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSessionHostListIterator(page SessionHostListPage) SessionHostListIterator {
	return original.NewSessionHostListIterator(page)
}
func NewSessionHostListPage(getNextPage func(context.Context, SessionHostList) (SessionHostList, error)) SessionHostListPage {
	return original.NewSessionHostListPage(getNextPage)
}
func NewSessionHostsClient(subscriptionID string) SessionHostsClient {
	return original.NewSessionHostsClient(subscriptionID)
}
func NewSessionHostsClientWithBaseURI(baseURI string, subscriptionID string) SessionHostsClient {
	return original.NewSessionHostsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStartMenuItemListIterator(page StartMenuItemListPage) StartMenuItemListIterator {
	return original.NewStartMenuItemListIterator(page)
}
func NewStartMenuItemListPage(getNextPage func(context.Context, StartMenuItemList) (StartMenuItemList, error)) StartMenuItemListPage {
	return original.NewStartMenuItemListPage(getNextPage)
}
func NewStartMenuItemsClient(subscriptionID string) StartMenuItemsClient {
	return original.NewStartMenuItemsClient(subscriptionID)
}
func NewStartMenuItemsClientWithBaseURI(baseURI string, subscriptionID string) StartMenuItemsClient {
	return original.NewStartMenuItemsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUserSessionListIterator(page UserSessionListPage) UserSessionListIterator {
	return original.NewUserSessionListIterator(page)
}
func NewUserSessionListPage(getNextPage func(context.Context, UserSessionList) (UserSessionList, error)) UserSessionListPage {
	return original.NewUserSessionListPage(getNextPage)
}
func NewUserSessionsClient(subscriptionID string) UserSessionsClient {
	return original.NewUserSessionsClient(subscriptionID)
}
func NewUserSessionsClientWithBaseURI(baseURI string, subscriptionID string) UserSessionsClient {
	return original.NewUserSessionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspaceListIterator(page WorkspaceListPage) WorkspaceListIterator {
	return original.NewWorkspaceListIterator(page)
}
func NewWorkspaceListPage(getNextPage func(context.Context, WorkspaceList) (WorkspaceList, error)) WorkspaceListPage {
	return original.NewWorkspaceListPage(getNextPage)
}
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleApplicationGroupTypeValues() []ApplicationGroupType {
	return original.PossibleApplicationGroupTypeValues()
}
func PossibleApplicationTypeValues() []ApplicationType {
	return original.PossibleApplicationTypeValues()
}
func PossibleCommandLineSettingValues() []CommandLineSetting {
	return original.PossibleCommandLineSettingValues()
}
func PossibleHostPoolTypeValues() []HostPoolType {
	return original.PossibleHostPoolTypeValues()
}
func PossibleLoadBalancerTypeValues() []LoadBalancerType {
	return original.PossibleLoadBalancerTypeValues()
}
func PossiblePersonalDesktopAssignmentTypeValues() []PersonalDesktopAssignmentType {
	return original.PossiblePersonalDesktopAssignmentTypeValues()
}
func PossibleRegistrationTokenOperationValues() []RegistrationTokenOperation {
	return original.PossibleRegistrationTokenOperationValues()
}
func PossibleSessionStateValues() []SessionState {
	return original.PossibleSessionStateValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleUpdateStateValues() []UpdateState {
	return original.PossibleUpdateStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}