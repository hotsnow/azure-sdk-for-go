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

package frontdoor

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-01-01/frontdoor"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionType = original.ActionType

const (
	Allow    ActionType = original.Allow
	Block    ActionType = original.Block
	Log      ActionType = original.Log
	Redirect ActionType = original.Redirect
)

type AggregationInterval = original.AggregationInterval

const (
	Daily  AggregationInterval = original.Daily
	Hourly AggregationInterval = original.Hourly
)

type Availability = original.Availability

const (
	Available   Availability = original.Available
	Unavailable Availability = original.Unavailable
)

type BackendEnabledState = original.BackendEnabledState

const (
	Disabled BackendEnabledState = original.Disabled
	Enabled  BackendEnabledState = original.Enabled
)

type CertificateSource = original.CertificateSource

const (
	CertificateSourceAzureKeyVault CertificateSource = original.CertificateSourceAzureKeyVault
	CertificateSourceFrontDoor     CertificateSource = original.CertificateSourceFrontDoor
)

type CertificateType = original.CertificateType

const (
	Dedicated CertificateType = original.Dedicated
)

type CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningState

const (
	CustomHTTPSProvisioningStateDisabled  CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningStateDisabled
	CustomHTTPSProvisioningStateDisabling CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningStateDisabling
	CustomHTTPSProvisioningStateEnabled   CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningStateEnabled
	CustomHTTPSProvisioningStateEnabling  CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningStateEnabling
	CustomHTTPSProvisioningStateFailed    CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningStateFailed
)

type CustomHTTPSProvisioningSubstate = original.CustomHTTPSProvisioningSubstate

const (
	CertificateDeleted                            CustomHTTPSProvisioningSubstate = original.CertificateDeleted
	CertificateDeployed                           CustomHTTPSProvisioningSubstate = original.CertificateDeployed
	DeletingCertificate                           CustomHTTPSProvisioningSubstate = original.DeletingCertificate
	DeployingCertificate                          CustomHTTPSProvisioningSubstate = original.DeployingCertificate
	DomainControlValidationRequestApproved        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestApproved
	DomainControlValidationRequestRejected        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestRejected
	DomainControlValidationRequestTimedOut        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestTimedOut
	IssuingCertificate                            CustomHTTPSProvisioningSubstate = original.IssuingCertificate
	PendingDomainControlValidationREquestApproval CustomHTTPSProvisioningSubstate = original.PendingDomainControlValidationREquestApproval
	SubmittingDomainControlValidationRequest      CustomHTTPSProvisioningSubstate = original.SubmittingDomainControlValidationRequest
)

type CustomRuleEnabledState = original.CustomRuleEnabledState

const (
	CustomRuleEnabledStateDisabled CustomRuleEnabledState = original.CustomRuleEnabledStateDisabled
	CustomRuleEnabledStateEnabled  CustomRuleEnabledState = original.CustomRuleEnabledStateEnabled
)

type DynamicCompressionEnabled = original.DynamicCompressionEnabled

const (
	DynamicCompressionEnabledDisabled DynamicCompressionEnabled = original.DynamicCompressionEnabledDisabled
	DynamicCompressionEnabledEnabled  DynamicCompressionEnabled = original.DynamicCompressionEnabledEnabled
)

type EnabledState = original.EnabledState

const (
	EnabledStateDisabled EnabledState = original.EnabledStateDisabled
	EnabledStateEnabled  EnabledState = original.EnabledStateEnabled
)

type EndpointType = original.EndpointType

const (
	AFD         EndpointType = original.AFD
	ATM         EndpointType = original.ATM
	AzureRegion EndpointType = original.AzureRegion
	CDN         EndpointType = original.CDN
)

type EnforceCertificateNameCheckEnabledState = original.EnforceCertificateNameCheckEnabledState

const (
	EnforceCertificateNameCheckEnabledStateDisabled EnforceCertificateNameCheckEnabledState = original.EnforceCertificateNameCheckEnabledStateDisabled
	EnforceCertificateNameCheckEnabledStateEnabled  EnforceCertificateNameCheckEnabledState = original.EnforceCertificateNameCheckEnabledStateEnabled
)

type ForwardingProtocol = original.ForwardingProtocol

const (
	HTTPOnly     ForwardingProtocol = original.HTTPOnly
	HTTPSOnly    ForwardingProtocol = original.HTTPSOnly
	MatchRequest ForwardingProtocol = original.MatchRequest
)

type HeaderActionType = original.HeaderActionType

const (
	Append    HeaderActionType = original.Append
	Delete    HeaderActionType = original.Delete
	Overwrite HeaderActionType = original.Overwrite
)

type HealthProbeEnabled = original.HealthProbeEnabled

const (
	HealthProbeEnabledDisabled HealthProbeEnabled = original.HealthProbeEnabledDisabled
	HealthProbeEnabledEnabled  HealthProbeEnabled = original.HealthProbeEnabledEnabled
)

type HealthProbeMethod = original.HealthProbeMethod

const (
	GET  HealthProbeMethod = original.GET
	HEAD HealthProbeMethod = original.HEAD
)

type LatencyScorecardAggregationInterval = original.LatencyScorecardAggregationInterval

const (
	LatencyScorecardAggregationIntervalDaily   LatencyScorecardAggregationInterval = original.LatencyScorecardAggregationIntervalDaily
	LatencyScorecardAggregationIntervalMonthly LatencyScorecardAggregationInterval = original.LatencyScorecardAggregationIntervalMonthly
	LatencyScorecardAggregationIntervalWeekly  LatencyScorecardAggregationInterval = original.LatencyScorecardAggregationIntervalWeekly
)

type ManagedRuleEnabledState = original.ManagedRuleEnabledState

const (
	ManagedRuleEnabledStateDisabled ManagedRuleEnabledState = original.ManagedRuleEnabledStateDisabled
	ManagedRuleEnabledStateEnabled  ManagedRuleEnabledState = original.ManagedRuleEnabledStateEnabled
)

type ManagedRuleExclusionMatchVariable = original.ManagedRuleExclusionMatchVariable

const (
	QueryStringArgNames     ManagedRuleExclusionMatchVariable = original.QueryStringArgNames
	RequestBodyPostArgNames ManagedRuleExclusionMatchVariable = original.RequestBodyPostArgNames
	RequestCookieNames      ManagedRuleExclusionMatchVariable = original.RequestCookieNames
	RequestHeaderNames      ManagedRuleExclusionMatchVariable = original.RequestHeaderNames
)

type ManagedRuleExclusionSelectorMatchOperator = original.ManagedRuleExclusionSelectorMatchOperator

const (
	Contains   ManagedRuleExclusionSelectorMatchOperator = original.Contains
	EndsWith   ManagedRuleExclusionSelectorMatchOperator = original.EndsWith
	Equals     ManagedRuleExclusionSelectorMatchOperator = original.Equals
	EqualsAny  ManagedRuleExclusionSelectorMatchOperator = original.EqualsAny
	StartsWith ManagedRuleExclusionSelectorMatchOperator = original.StartsWith
)

type MatchProcessingBehavior = original.MatchProcessingBehavior

const (
	Continue MatchProcessingBehavior = original.Continue
	Stop     MatchProcessingBehavior = original.Stop
)

type MatchVariable = original.MatchVariable

const (
	Cookies       MatchVariable = original.Cookies
	PostArgs      MatchVariable = original.PostArgs
	QueryString   MatchVariable = original.QueryString
	RemoteAddr    MatchVariable = original.RemoteAddr
	RequestBody   MatchVariable = original.RequestBody
	RequestHeader MatchVariable = original.RequestHeader
	RequestMethod MatchVariable = original.RequestMethod
	RequestURI    MatchVariable = original.RequestURI
	SocketAddr    MatchVariable = original.SocketAddr
)

type MinimumTLSVersion = original.MinimumTLSVersion

const (
	OneFullStopTwo  MinimumTLSVersion = original.OneFullStopTwo
	OneFullStopZero MinimumTLSVersion = original.OneFullStopZero
)

type NetworkExperimentResourceState = original.NetworkExperimentResourceState

const (
	NetworkExperimentResourceStateCreating  NetworkExperimentResourceState = original.NetworkExperimentResourceStateCreating
	NetworkExperimentResourceStateDeleting  NetworkExperimentResourceState = original.NetworkExperimentResourceStateDeleting
	NetworkExperimentResourceStateDisabled  NetworkExperimentResourceState = original.NetworkExperimentResourceStateDisabled
	NetworkExperimentResourceStateDisabling NetworkExperimentResourceState = original.NetworkExperimentResourceStateDisabling
	NetworkExperimentResourceStateEnabled   NetworkExperimentResourceState = original.NetworkExperimentResourceStateEnabled
	NetworkExperimentResourceStateEnabling  NetworkExperimentResourceState = original.NetworkExperimentResourceStateEnabling
)

type NetworkOperationStatus = original.NetworkOperationStatus

const (
	Failed     NetworkOperationStatus = original.Failed
	InProgress NetworkOperationStatus = original.InProgress
	Succeeded  NetworkOperationStatus = original.Succeeded
)

type OdataType = original.OdataType

const (
	OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorForwardingConfiguration OdataType = original.OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorForwardingConfiguration
	OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorRedirectConfiguration   OdataType = original.OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorRedirectConfiguration
	OdataTypeRouteConfiguration                                            OdataType = original.OdataTypeRouteConfiguration
)

type Operator = original.Operator

const (
	OperatorAny                Operator = original.OperatorAny
	OperatorBeginsWith         Operator = original.OperatorBeginsWith
	OperatorContains           Operator = original.OperatorContains
	OperatorEndsWith           Operator = original.OperatorEndsWith
	OperatorEqual              Operator = original.OperatorEqual
	OperatorGeoMatch           Operator = original.OperatorGeoMatch
	OperatorGreaterThan        Operator = original.OperatorGreaterThan
	OperatorGreaterThanOrEqual Operator = original.OperatorGreaterThanOrEqual
	OperatorIPMatch            Operator = original.OperatorIPMatch
	OperatorLessThan           Operator = original.OperatorLessThan
	OperatorLessThanOrEqual    Operator = original.OperatorLessThanOrEqual
	OperatorRegEx              Operator = original.OperatorRegEx
)

type PolicyEnabledState = original.PolicyEnabledState

const (
	PolicyEnabledStateDisabled PolicyEnabledState = original.PolicyEnabledStateDisabled
	PolicyEnabledStateEnabled  PolicyEnabledState = original.PolicyEnabledStateEnabled
)

type PolicyMode = original.PolicyMode

const (
	Detection  PolicyMode = original.Detection
	Prevention PolicyMode = original.Prevention
)

type PolicyResourceState = original.PolicyResourceState

const (
	PolicyResourceStateCreating  PolicyResourceState = original.PolicyResourceStateCreating
	PolicyResourceStateDeleting  PolicyResourceState = original.PolicyResourceStateDeleting
	PolicyResourceStateDisabled  PolicyResourceState = original.PolicyResourceStateDisabled
	PolicyResourceStateDisabling PolicyResourceState = original.PolicyResourceStateDisabling
	PolicyResourceStateEnabled   PolicyResourceState = original.PolicyResourceStateEnabled
	PolicyResourceStateEnabling  PolicyResourceState = original.PolicyResourceStateEnabling
)

type PrivateEndpointStatus = original.PrivateEndpointStatus

const (
	Approved     PrivateEndpointStatus = original.Approved
	Disconnected PrivateEndpointStatus = original.Disconnected
	Pending      PrivateEndpointStatus = original.Pending
	Rejected     PrivateEndpointStatus = original.Rejected
	Timeout      PrivateEndpointStatus = original.Timeout
)

type Protocol = original.Protocol

const (
	HTTP  Protocol = original.HTTP
	HTTPS Protocol = original.HTTPS
)

type Query = original.Query

const (
	StripAll  Query = original.StripAll
	StripNone Query = original.StripNone
)

type RedirectProtocol = original.RedirectProtocol

const (
	RedirectProtocolHTTPOnly     RedirectProtocol = original.RedirectProtocolHTTPOnly
	RedirectProtocolHTTPSOnly    RedirectProtocol = original.RedirectProtocolHTTPSOnly
	RedirectProtocolMatchRequest RedirectProtocol = original.RedirectProtocolMatchRequest
)

type RedirectType = original.RedirectType

const (
	Found             RedirectType = original.Found
	Moved             RedirectType = original.Moved
	PermanentRedirect RedirectType = original.PermanentRedirect
	TemporaryRedirect RedirectType = original.TemporaryRedirect
)

type ResourceState = original.ResourceState

const (
	ResourceStateCreating  ResourceState = original.ResourceStateCreating
	ResourceStateDeleting  ResourceState = original.ResourceStateDeleting
	ResourceStateDisabled  ResourceState = original.ResourceStateDisabled
	ResourceStateDisabling ResourceState = original.ResourceStateDisabling
	ResourceStateEnabled   ResourceState = original.ResourceStateEnabled
	ResourceStateEnabling  ResourceState = original.ResourceStateEnabling
)

type ResourceType = original.ResourceType

const (
	MicrosoftNetworkfrontDoors                  ResourceType = original.MicrosoftNetworkfrontDoors
	MicrosoftNetworkfrontDoorsfrontendEndpoints ResourceType = original.MicrosoftNetworkfrontDoorsfrontendEndpoints
)

type RoutingRuleEnabledState = original.RoutingRuleEnabledState

const (
	RoutingRuleEnabledStateDisabled RoutingRuleEnabledState = original.RoutingRuleEnabledStateDisabled
	RoutingRuleEnabledStateEnabled  RoutingRuleEnabledState = original.RoutingRuleEnabledStateEnabled
)

type RuleType = original.RuleType

const (
	MatchRule     RuleType = original.MatchRule
	RateLimitRule RuleType = original.RateLimitRule
)

type RulesEngineMatchVariable = original.RulesEngineMatchVariable

const (
	RulesEngineMatchVariableIsMobile                 RulesEngineMatchVariable = original.RulesEngineMatchVariableIsMobile
	RulesEngineMatchVariablePostArgs                 RulesEngineMatchVariable = original.RulesEngineMatchVariablePostArgs
	RulesEngineMatchVariableQueryString              RulesEngineMatchVariable = original.RulesEngineMatchVariableQueryString
	RulesEngineMatchVariableRemoteAddr               RulesEngineMatchVariable = original.RulesEngineMatchVariableRemoteAddr
	RulesEngineMatchVariableRequestBody              RulesEngineMatchVariable = original.RulesEngineMatchVariableRequestBody
	RulesEngineMatchVariableRequestFilename          RulesEngineMatchVariable = original.RulesEngineMatchVariableRequestFilename
	RulesEngineMatchVariableRequestFilenameExtension RulesEngineMatchVariable = original.RulesEngineMatchVariableRequestFilenameExtension
	RulesEngineMatchVariableRequestHeader            RulesEngineMatchVariable = original.RulesEngineMatchVariableRequestHeader
	RulesEngineMatchVariableRequestMethod            RulesEngineMatchVariable = original.RulesEngineMatchVariableRequestMethod
	RulesEngineMatchVariableRequestPath              RulesEngineMatchVariable = original.RulesEngineMatchVariableRequestPath
	RulesEngineMatchVariableRequestScheme            RulesEngineMatchVariable = original.RulesEngineMatchVariableRequestScheme
	RulesEngineMatchVariableRequestURI               RulesEngineMatchVariable = original.RulesEngineMatchVariableRequestURI
)

type RulesEngineOperator = original.RulesEngineOperator

const (
	RulesEngineOperatorAny                RulesEngineOperator = original.RulesEngineOperatorAny
	RulesEngineOperatorBeginsWith         RulesEngineOperator = original.RulesEngineOperatorBeginsWith
	RulesEngineOperatorContains           RulesEngineOperator = original.RulesEngineOperatorContains
	RulesEngineOperatorEndsWith           RulesEngineOperator = original.RulesEngineOperatorEndsWith
	RulesEngineOperatorEqual              RulesEngineOperator = original.RulesEngineOperatorEqual
	RulesEngineOperatorGeoMatch           RulesEngineOperator = original.RulesEngineOperatorGeoMatch
	RulesEngineOperatorGreaterThan        RulesEngineOperator = original.RulesEngineOperatorGreaterThan
	RulesEngineOperatorGreaterThanOrEqual RulesEngineOperator = original.RulesEngineOperatorGreaterThanOrEqual
	RulesEngineOperatorIPMatch            RulesEngineOperator = original.RulesEngineOperatorIPMatch
	RulesEngineOperatorLessThan           RulesEngineOperator = original.RulesEngineOperatorLessThan
	RulesEngineOperatorLessThanOrEqual    RulesEngineOperator = original.RulesEngineOperatorLessThanOrEqual
)

type SessionAffinityEnabledState = original.SessionAffinityEnabledState

const (
	SessionAffinityEnabledStateDisabled SessionAffinityEnabledState = original.SessionAffinityEnabledStateDisabled
	SessionAffinityEnabledStateEnabled  SessionAffinityEnabledState = original.SessionAffinityEnabledStateEnabled
)

type State = original.State

const (
	StateDisabled State = original.StateDisabled
	StateEnabled  State = original.StateEnabled
)

type TimeseriesAggregationInterval = original.TimeseriesAggregationInterval

const (
	TimeseriesAggregationIntervalDaily  TimeseriesAggregationInterval = original.TimeseriesAggregationIntervalDaily
	TimeseriesAggregationIntervalHourly TimeseriesAggregationInterval = original.TimeseriesAggregationIntervalHourly
)

type TimeseriesType = original.TimeseriesType

const (
	LatencyP50        TimeseriesType = original.LatencyP50
	LatencyP75        TimeseriesType = original.LatencyP75
	LatencyP95        TimeseriesType = original.LatencyP95
	MeasurementCounts TimeseriesType = original.MeasurementCounts
)

type Transform = original.Transform

const (
	Lowercase   Transform = original.Lowercase
	RemoveNulls Transform = original.RemoveNulls
	Trim        Transform = original.Trim
	Uppercase   Transform = original.Uppercase
	URLDecode   Transform = original.URLDecode
	URLEncode   Transform = original.URLEncode
)

type TransformType = original.TransformType

const (
	TransformTypeLowercase   TransformType = original.TransformTypeLowercase
	TransformTypeRemoveNulls TransformType = original.TransformTypeRemoveNulls
	TransformTypeTrim        TransformType = original.TransformTypeTrim
	TransformTypeUppercase   TransformType = original.TransformTypeUppercase
	TransformTypeURLDecode   TransformType = original.TransformTypeURLDecode
	TransformTypeURLEncode   TransformType = original.TransformTypeURLEncode
)

type AzureAsyncOperationResult = original.AzureAsyncOperationResult
type Backend = original.Backend
type BackendPool = original.BackendPool
type BackendPoolListResult = original.BackendPoolListResult
type BackendPoolProperties = original.BackendPoolProperties
type BackendPoolUpdateParameters = original.BackendPoolUpdateParameters
type BackendPoolsSettings = original.BackendPoolsSettings
type BaseClient = original.BaseClient
type BasicRouteConfiguration = original.BasicRouteConfiguration
type CacheConfiguration = original.CacheConfiguration
type CertificateSourceParameters = original.CertificateSourceParameters
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CustomHTTPSConfiguration = original.CustomHTTPSConfiguration
type CustomRule = original.CustomRule
type CustomRuleList = original.CustomRuleList
type Endpoint = original.Endpoint
type EndpointsClient = original.EndpointsClient
type EndpointsPurgeContentFuture = original.EndpointsPurgeContentFuture
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Experiment = original.Experiment
type ExperimentList = original.ExperimentList
type ExperimentListIterator = original.ExperimentListIterator
type ExperimentListPage = original.ExperimentListPage
type ExperimentProperties = original.ExperimentProperties
type ExperimentUpdateModel = original.ExperimentUpdateModel
type ExperimentUpdateProperties = original.ExperimentUpdateProperties
type ExperimentsClient = original.ExperimentsClient
type ExperimentsCreateOrUpdateFuture = original.ExperimentsCreateOrUpdateFuture
type ExperimentsDeleteFuture = original.ExperimentsDeleteFuture
type ExperimentsUpdateFuture = original.ExperimentsUpdateFuture
type ForwardingConfiguration = original.ForwardingConfiguration
type FrontDoor = original.FrontDoor
type FrontDoorsClient = original.FrontDoorsClient
type FrontDoorsCreateOrUpdateFutureType = original.FrontDoorsCreateOrUpdateFutureType
type FrontDoorsDeleteFutureType = original.FrontDoorsDeleteFutureType
type FrontendEndpoint = original.FrontendEndpoint
type FrontendEndpointLink = original.FrontendEndpointLink
type FrontendEndpointProperties = original.FrontendEndpointProperties
type FrontendEndpointUpdateParameters = original.FrontendEndpointUpdateParameters
type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink = original.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink
type FrontendEndpointsClient = original.FrontendEndpointsClient
type FrontendEndpointsDisableHTTPSFuture = original.FrontendEndpointsDisableHTTPSFuture
type FrontendEndpointsEnableHTTPSFuture = original.FrontendEndpointsEnableHTTPSFuture
type FrontendEndpointsListResult = original.FrontendEndpointsListResult
type FrontendEndpointsListResultIterator = original.FrontendEndpointsListResultIterator
type FrontendEndpointsListResultPage = original.FrontendEndpointsListResultPage
type HeaderAction = original.HeaderAction
type HealthProbeSettingsListResult = original.HealthProbeSettingsListResult
type HealthProbeSettingsModel = original.HealthProbeSettingsModel
type HealthProbeSettingsProperties = original.HealthProbeSettingsProperties
type HealthProbeSettingsUpdateParameters = original.HealthProbeSettingsUpdateParameters
type KeyVaultCertificateSourceParameters = original.KeyVaultCertificateSourceParameters
type KeyVaultCertificateSourceParametersVault = original.KeyVaultCertificateSourceParametersVault
type LatencyMetric = original.LatencyMetric
type LatencyScorecard = original.LatencyScorecard
type LatencyScorecardProperties = original.LatencyScorecardProperties
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type LoadBalancingSettingsListResult = original.LoadBalancingSettingsListResult
type LoadBalancingSettingsModel = original.LoadBalancingSettingsModel
type LoadBalancingSettingsProperties = original.LoadBalancingSettingsProperties
type LoadBalancingSettingsUpdateParameters = original.LoadBalancingSettingsUpdateParameters
type ManagedRuleDefinition = original.ManagedRuleDefinition
type ManagedRuleExclusion = original.ManagedRuleExclusion
type ManagedRuleGroupDefinition = original.ManagedRuleGroupDefinition
type ManagedRuleGroupOverride = original.ManagedRuleGroupOverride
type ManagedRuleOverride = original.ManagedRuleOverride
type ManagedRuleSet = original.ManagedRuleSet
type ManagedRuleSetDefinition = original.ManagedRuleSetDefinition
type ManagedRuleSetDefinitionList = original.ManagedRuleSetDefinitionList
type ManagedRuleSetDefinitionListIterator = original.ManagedRuleSetDefinitionListIterator
type ManagedRuleSetDefinitionListPage = original.ManagedRuleSetDefinitionListPage
type ManagedRuleSetDefinitionProperties = original.ManagedRuleSetDefinitionProperties
type ManagedRuleSetList = original.ManagedRuleSetList
type ManagedRuleSetsClient = original.ManagedRuleSetsClient
type MatchCondition = original.MatchCondition
type NetworkExperimentProfilesClient = original.NetworkExperimentProfilesClient
type NetworkExperimentProfilesCreateOrUpdateFuture = original.NetworkExperimentProfilesCreateOrUpdateFuture
type NetworkExperimentProfilesDeleteFuture = original.NetworkExperimentProfilesDeleteFuture
type NetworkExperimentProfilesUpdateFuture = original.NetworkExperimentProfilesUpdateFuture
type PoliciesClient = original.PoliciesClient
type PoliciesCreateOrUpdateFuture = original.PoliciesCreateOrUpdateFuture
type PoliciesDeleteFuture = original.PoliciesDeleteFuture
type PolicySettings = original.PolicySettings
type PreconfiguredEndpoint = original.PreconfiguredEndpoint
type PreconfiguredEndpointList = original.PreconfiguredEndpointList
type PreconfiguredEndpointListIterator = original.PreconfiguredEndpointListIterator
type PreconfiguredEndpointListPage = original.PreconfiguredEndpointListPage
type PreconfiguredEndpointProperties = original.PreconfiguredEndpointProperties
type PreconfiguredEndpointsClient = original.PreconfiguredEndpointsClient
type Profile = original.Profile
type ProfileList = original.ProfileList
type ProfileListIterator = original.ProfileListIterator
type ProfileListPage = original.ProfileListPage
type ProfileProperties = original.ProfileProperties
type ProfileUpdateModel = original.ProfileUpdateModel
type ProfileUpdateProperties = original.ProfileUpdateProperties
type Properties = original.Properties
type PurgeParameters = original.PurgeParameters
type RedirectConfiguration = original.RedirectConfiguration
type ReportsClient = original.ReportsClient
type Resource = original.Resource
type RouteConfiguration = original.RouteConfiguration
type RoutingRule = original.RoutingRule
type RoutingRuleListResult = original.RoutingRuleListResult
type RoutingRuleProperties = original.RoutingRuleProperties
type RoutingRuleUpdateParameters = original.RoutingRuleUpdateParameters
type RulesEngine = original.RulesEngine
type RulesEngineAction = original.RulesEngineAction
type RulesEngineListResult = original.RulesEngineListResult
type RulesEngineListResultIterator = original.RulesEngineListResultIterator
type RulesEngineListResultPage = original.RulesEngineListResultPage
type RulesEngineMatchCondition = original.RulesEngineMatchCondition
type RulesEngineProperties = original.RulesEngineProperties
type RulesEngineRule = original.RulesEngineRule
type RulesEngineUpdateParameters = original.RulesEngineUpdateParameters
type RulesEnginesClient = original.RulesEnginesClient
type RulesEnginesCreateOrUpdateFuture = original.RulesEnginesCreateOrUpdateFuture
type RulesEnginesDeleteFuture = original.RulesEnginesDeleteFuture
type SubResource = original.SubResource
type TagsObject = original.TagsObject
type Timeseries = original.Timeseries
type TimeseriesDataPoint = original.TimeseriesDataPoint
type TimeseriesProperties = original.TimeseriesProperties
type UpdateParameters = original.UpdateParameters
type ValidateCustomDomainInput = original.ValidateCustomDomainInput
type ValidateCustomDomainOutput = original.ValidateCustomDomainOutput
type WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicy
type WebApplicationFirewallPolicyList = original.WebApplicationFirewallPolicyList
type WebApplicationFirewallPolicyListIterator = original.WebApplicationFirewallPolicyListIterator
type WebApplicationFirewallPolicyListPage = original.WebApplicationFirewallPolicyListPage
type WebApplicationFirewallPolicyProperties = original.WebApplicationFirewallPolicyProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return original.NewEndpointsClient(subscriptionID)
}
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewExperimentListIterator(page ExperimentListPage) ExperimentListIterator {
	return original.NewExperimentListIterator(page)
}
func NewExperimentListPage(getNextPage func(context.Context, ExperimentList) (ExperimentList, error)) ExperimentListPage {
	return original.NewExperimentListPage(getNextPage)
}
func NewExperimentsClient(subscriptionID string) ExperimentsClient {
	return original.NewExperimentsClient(subscriptionID)
}
func NewExperimentsClientWithBaseURI(baseURI string, subscriptionID string) ExperimentsClient {
	return original.NewExperimentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontDoorsClient(subscriptionID string) FrontDoorsClient {
	return original.NewFrontDoorsClient(subscriptionID)
}
func NewFrontDoorsClientWithBaseURI(baseURI string, subscriptionID string) FrontDoorsClient {
	return original.NewFrontDoorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontendEndpointsClient(subscriptionID string) FrontendEndpointsClient {
	return original.NewFrontendEndpointsClient(subscriptionID)
}
func NewFrontendEndpointsClientWithBaseURI(baseURI string, subscriptionID string) FrontendEndpointsClient {
	return original.NewFrontendEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontendEndpointsListResultIterator(page FrontendEndpointsListResultPage) FrontendEndpointsListResultIterator {
	return original.NewFrontendEndpointsListResultIterator(page)
}
func NewFrontendEndpointsListResultPage(getNextPage func(context.Context, FrontendEndpointsListResult) (FrontendEndpointsListResult, error)) FrontendEndpointsListResultPage {
	return original.NewFrontendEndpointsListResultPage(getNextPage)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(getNextPage)
}
func NewManagedRuleSetDefinitionListIterator(page ManagedRuleSetDefinitionListPage) ManagedRuleSetDefinitionListIterator {
	return original.NewManagedRuleSetDefinitionListIterator(page)
}
func NewManagedRuleSetDefinitionListPage(getNextPage func(context.Context, ManagedRuleSetDefinitionList) (ManagedRuleSetDefinitionList, error)) ManagedRuleSetDefinitionListPage {
	return original.NewManagedRuleSetDefinitionListPage(getNextPage)
}
func NewManagedRuleSetsClient(subscriptionID string) ManagedRuleSetsClient {
	return original.NewManagedRuleSetsClient(subscriptionID)
}
func NewManagedRuleSetsClientWithBaseURI(baseURI string, subscriptionID string) ManagedRuleSetsClient {
	return original.NewManagedRuleSetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewNetworkExperimentProfilesClient(subscriptionID string) NetworkExperimentProfilesClient {
	return original.NewNetworkExperimentProfilesClient(subscriptionID)
}
func NewNetworkExperimentProfilesClientWithBaseURI(baseURI string, subscriptionID string) NetworkExperimentProfilesClient {
	return original.NewNetworkExperimentProfilesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPoliciesClient(subscriptionID string) PoliciesClient {
	return original.NewPoliciesClient(subscriptionID)
}
func NewPoliciesClientWithBaseURI(baseURI string, subscriptionID string) PoliciesClient {
	return original.NewPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPreconfiguredEndpointListIterator(page PreconfiguredEndpointListPage) PreconfiguredEndpointListIterator {
	return original.NewPreconfiguredEndpointListIterator(page)
}
func NewPreconfiguredEndpointListPage(getNextPage func(context.Context, PreconfiguredEndpointList) (PreconfiguredEndpointList, error)) PreconfiguredEndpointListPage {
	return original.NewPreconfiguredEndpointListPage(getNextPage)
}
func NewPreconfiguredEndpointsClient(subscriptionID string) PreconfiguredEndpointsClient {
	return original.NewPreconfiguredEndpointsClient(subscriptionID)
}
func NewPreconfiguredEndpointsClientWithBaseURI(baseURI string, subscriptionID string) PreconfiguredEndpointsClient {
	return original.NewPreconfiguredEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProfileListIterator(page ProfileListPage) ProfileListIterator {
	return original.NewProfileListIterator(page)
}
func NewProfileListPage(getNextPage func(context.Context, ProfileList) (ProfileList, error)) ProfileListPage {
	return original.NewProfileListPage(getNextPage)
}
func NewReportsClient(subscriptionID string) ReportsClient {
	return original.NewReportsClient(subscriptionID)
}
func NewReportsClientWithBaseURI(baseURI string, subscriptionID string) ReportsClient {
	return original.NewReportsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRulesEngineListResultIterator(page RulesEngineListResultPage) RulesEngineListResultIterator {
	return original.NewRulesEngineListResultIterator(page)
}
func NewRulesEngineListResultPage(getNextPage func(context.Context, RulesEngineListResult) (RulesEngineListResult, error)) RulesEngineListResultPage {
	return original.NewRulesEngineListResultPage(getNextPage)
}
func NewRulesEnginesClient(subscriptionID string) RulesEnginesClient {
	return original.NewRulesEnginesClient(subscriptionID)
}
func NewRulesEnginesClientWithBaseURI(baseURI string, subscriptionID string) RulesEnginesClient {
	return original.NewRulesEnginesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWebApplicationFirewallPolicyListIterator(page WebApplicationFirewallPolicyListPage) WebApplicationFirewallPolicyListIterator {
	return original.NewWebApplicationFirewallPolicyListIterator(page)
}
func NewWebApplicationFirewallPolicyListPage(getNextPage func(context.Context, WebApplicationFirewallPolicyList) (WebApplicationFirewallPolicyList, error)) WebApplicationFirewallPolicyListPage {
	return original.NewWebApplicationFirewallPolicyListPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleAggregationIntervalValues() []AggregationInterval {
	return original.PossibleAggregationIntervalValues()
}
func PossibleAvailabilityValues() []Availability {
	return original.PossibleAvailabilityValues()
}
func PossibleBackendEnabledStateValues() []BackendEnabledState {
	return original.PossibleBackendEnabledStateValues()
}
func PossibleCertificateSourceValues() []CertificateSource {
	return original.PossibleCertificateSourceValues()
}
func PossibleCertificateTypeValues() []CertificateType {
	return original.PossibleCertificateTypeValues()
}
func PossibleCustomHTTPSProvisioningStateValues() []CustomHTTPSProvisioningState {
	return original.PossibleCustomHTTPSProvisioningStateValues()
}
func PossibleCustomHTTPSProvisioningSubstateValues() []CustomHTTPSProvisioningSubstate {
	return original.PossibleCustomHTTPSProvisioningSubstateValues()
}
func PossibleCustomRuleEnabledStateValues() []CustomRuleEnabledState {
	return original.PossibleCustomRuleEnabledStateValues()
}
func PossibleDynamicCompressionEnabledValues() []DynamicCompressionEnabled {
	return original.PossibleDynamicCompressionEnabledValues()
}
func PossibleEnabledStateValues() []EnabledState {
	return original.PossibleEnabledStateValues()
}
func PossibleEndpointTypeValues() []EndpointType {
	return original.PossibleEndpointTypeValues()
}
func PossibleEnforceCertificateNameCheckEnabledStateValues() []EnforceCertificateNameCheckEnabledState {
	return original.PossibleEnforceCertificateNameCheckEnabledStateValues()
}
func PossibleForwardingProtocolValues() []ForwardingProtocol {
	return original.PossibleForwardingProtocolValues()
}
func PossibleHeaderActionTypeValues() []HeaderActionType {
	return original.PossibleHeaderActionTypeValues()
}
func PossibleHealthProbeEnabledValues() []HealthProbeEnabled {
	return original.PossibleHealthProbeEnabledValues()
}
func PossibleHealthProbeMethodValues() []HealthProbeMethod {
	return original.PossibleHealthProbeMethodValues()
}
func PossibleLatencyScorecardAggregationIntervalValues() []LatencyScorecardAggregationInterval {
	return original.PossibleLatencyScorecardAggregationIntervalValues()
}
func PossibleManagedRuleEnabledStateValues() []ManagedRuleEnabledState {
	return original.PossibleManagedRuleEnabledStateValues()
}
func PossibleManagedRuleExclusionMatchVariableValues() []ManagedRuleExclusionMatchVariable {
	return original.PossibleManagedRuleExclusionMatchVariableValues()
}
func PossibleManagedRuleExclusionSelectorMatchOperatorValues() []ManagedRuleExclusionSelectorMatchOperator {
	return original.PossibleManagedRuleExclusionSelectorMatchOperatorValues()
}
func PossibleMatchProcessingBehaviorValues() []MatchProcessingBehavior {
	return original.PossibleMatchProcessingBehaviorValues()
}
func PossibleMatchVariableValues() []MatchVariable {
	return original.PossibleMatchVariableValues()
}
func PossibleMinimumTLSVersionValues() []MinimumTLSVersion {
	return original.PossibleMinimumTLSVersionValues()
}
func PossibleNetworkExperimentResourceStateValues() []NetworkExperimentResourceState {
	return original.PossibleNetworkExperimentResourceStateValues()
}
func PossibleNetworkOperationStatusValues() []NetworkOperationStatus {
	return original.PossibleNetworkOperationStatusValues()
}
func PossibleOdataTypeValues() []OdataType {
	return original.PossibleOdataTypeValues()
}
func PossibleOperatorValues() []Operator {
	return original.PossibleOperatorValues()
}
func PossiblePolicyEnabledStateValues() []PolicyEnabledState {
	return original.PossiblePolicyEnabledStateValues()
}
func PossiblePolicyModeValues() []PolicyMode {
	return original.PossiblePolicyModeValues()
}
func PossiblePolicyResourceStateValues() []PolicyResourceState {
	return original.PossiblePolicyResourceStateValues()
}
func PossiblePrivateEndpointStatusValues() []PrivateEndpointStatus {
	return original.PossiblePrivateEndpointStatusValues()
}
func PossibleProtocolValues() []Protocol {
	return original.PossibleProtocolValues()
}
func PossibleQueryValues() []Query {
	return original.PossibleQueryValues()
}
func PossibleRedirectProtocolValues() []RedirectProtocol {
	return original.PossibleRedirectProtocolValues()
}
func PossibleRedirectTypeValues() []RedirectType {
	return original.PossibleRedirectTypeValues()
}
func PossibleResourceStateValues() []ResourceState {
	return original.PossibleResourceStateValues()
}
func PossibleResourceTypeValues() []ResourceType {
	return original.PossibleResourceTypeValues()
}
func PossibleRoutingRuleEnabledStateValues() []RoutingRuleEnabledState {
	return original.PossibleRoutingRuleEnabledStateValues()
}
func PossibleRuleTypeValues() []RuleType {
	return original.PossibleRuleTypeValues()
}
func PossibleRulesEngineMatchVariableValues() []RulesEngineMatchVariable {
	return original.PossibleRulesEngineMatchVariableValues()
}
func PossibleRulesEngineOperatorValues() []RulesEngineOperator {
	return original.PossibleRulesEngineOperatorValues()
}
func PossibleSessionAffinityEnabledStateValues() []SessionAffinityEnabledState {
	return original.PossibleSessionAffinityEnabledStateValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleTimeseriesAggregationIntervalValues() []TimeseriesAggregationInterval {
	return original.PossibleTimeseriesAggregationIntervalValues()
}
func PossibleTimeseriesTypeValues() []TimeseriesType {
	return original.PossibleTimeseriesTypeValues()
}
func PossibleTransformTypeValues() []TransformType {
	return original.PossibleTransformTypeValues()
}
func PossibleTransformValues() []Transform {
	return original.PossibleTransformValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
