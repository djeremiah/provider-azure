/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RetentionPolicyInitParameters struct {

	// The number of days to retain flow log records.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Boolean flag to enable/disable retention.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type RetentionPolicyObservation struct {

	// The number of days to retain flow log records.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Boolean flag to enable/disable retention.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type RetentionPolicyParameters struct {

	// The number of days to retain flow log records.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Boolean flag to enable/disable retention.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type TrafficAnalyticsInitParameters struct {

	// Boolean flag to enable/disable traffic analytics.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// How frequently service should do flow analytics in minutes. Defaults to 60.
	IntervalInMinutes *float64 `json:"intervalInMinutes,omitempty" tf:"interval_in_minutes,omitempty"`

	// The location of the attached workspace.
	WorkspaceRegion *string `json:"workspaceRegion,omitempty" tf:"workspace_region,omitempty"`
}

type TrafficAnalyticsObservation struct {

	// Boolean flag to enable/disable traffic analytics.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// How frequently service should do flow analytics in minutes. Defaults to 60.
	IntervalInMinutes *float64 `json:"intervalInMinutes,omitempty" tf:"interval_in_minutes,omitempty"`

	// The resource GUID of the attached workspace.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// The location of the attached workspace.
	WorkspaceRegion *string `json:"workspaceRegion,omitempty" tf:"workspace_region,omitempty"`

	// The resource ID of the attached workspace.
	WorkspaceResourceID *string `json:"workspaceResourceId,omitempty" tf:"workspace_resource_id,omitempty"`
}

type TrafficAnalyticsParameters struct {

	// Boolean flag to enable/disable traffic analytics.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// How frequently service should do flow analytics in minutes. Defaults to 60.
	// +kubebuilder:validation:Optional
	IntervalInMinutes *float64 `json:"intervalInMinutes,omitempty" tf:"interval_in_minutes,omitempty"`

	// The resource GUID of the attached workspace.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("workspace_id",true)
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`

	// The location of the attached workspace.
	// +kubebuilder:validation:Optional
	WorkspaceRegion *string `json:"workspaceRegion,omitempty" tf:"workspace_region,omitempty"`

	// The resource ID of the attached workspace.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceResourceID *string `json:"workspaceResourceId,omitempty" tf:"workspace_resource_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceResourceId.
	// +kubebuilder:validation:Optional
	WorkspaceResourceIDRef *v1.Reference `json:"workspaceResourceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceResourceId.
	// +kubebuilder:validation:Optional
	WorkspaceResourceIDSelector *v1.Selector `json:"workspaceResourceIdSelector,omitempty" tf:"-"`
}

type WatcherFlowLogInitParameters struct {

	// Should Network Flow Logging be Enabled?
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The location where the Network Watcher Flow Log resides. Changing this forces a new resource to be created. Defaults to the location of the Network Watcher.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A retention_policy block as documented below.
	RetentionPolicy []RetentionPolicyInitParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// A mapping of tags which should be assigned to the Network Watcher Flow Log.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A traffic_analytics block as documented below.
	TrafficAnalytics []TrafficAnalyticsInitParameters `json:"trafficAnalytics,omitempty" tf:"traffic_analytics,omitempty"`

	// The version (revision) of the flow log. Possible values are 1 and 2.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type WatcherFlowLogObservation struct {

	// Should Network Flow Logging be Enabled?
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Network Watcher.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location where the Network Watcher Flow Log resides. Changing this forces a new resource to be created. Defaults to the location of the Network Watcher.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupID *string `json:"networkSecurityGroupId,omitempty" tf:"network_security_group_id,omitempty"`

	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName *string `json:"networkWatcherName,omitempty" tf:"network_watcher_name,omitempty"`

	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A retention_policy block as documented below.
	RetentionPolicy []RetentionPolicyObservation `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// The ID of the Storage Account where flow logs are stored.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// A mapping of tags which should be assigned to the Network Watcher Flow Log.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A traffic_analytics block as documented below.
	TrafficAnalytics []TrafficAnalyticsObservation `json:"trafficAnalytics,omitempty" tf:"traffic_analytics,omitempty"`

	// The version (revision) of the flow log. Possible values are 1 and 2.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type WatcherFlowLogParameters struct {

	// Should Network Flow Logging be Enabled?
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The location where the Network Watcher Flow Log resides. Changing this forces a new resource to be created. Defaults to the location of the Network Watcher.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupID *string `json:"networkSecurityGroupId,omitempty" tf:"network_security_group_id,omitempty"`

	// Reference to a SecurityGroup to populate networkSecurityGroupId.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIDRef *v1.Reference `json:"networkSecurityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup to populate networkSecurityGroupId.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIDSelector *v1.Selector `json:"networkSecurityGroupIdSelector,omitempty" tf:"-"`

	// The name of the Network Watcher. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Watcher
	// +kubebuilder:validation:Optional
	NetworkWatcherName *string `json:"networkWatcherName,omitempty" tf:"network_watcher_name,omitempty"`

	// Reference to a Watcher to populate networkWatcherName.
	// +kubebuilder:validation:Optional
	NetworkWatcherNameRef *v1.Reference `json:"networkWatcherNameRef,omitempty" tf:"-"`

	// Selector for a Watcher to populate networkWatcherName.
	// +kubebuilder:validation:Optional
	NetworkWatcherNameSelector *v1.Selector `json:"networkWatcherNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A retention_policy block as documented below.
	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// The ID of the Storage Account where flow logs are stored.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Network Watcher Flow Log.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A traffic_analytics block as documented below.
	// +kubebuilder:validation:Optional
	TrafficAnalytics []TrafficAnalyticsParameters `json:"trafficAnalytics,omitempty" tf:"traffic_analytics,omitempty"`

	// The version (revision) of the flow log. Possible values are 1 and 2.
	// +kubebuilder:validation:Optional
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

// WatcherFlowLogSpec defines the desired state of WatcherFlowLog
type WatcherFlowLogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WatcherFlowLogParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider WatcherFlowLogInitParameters `json:"initProvider,omitempty"`
}

// WatcherFlowLogStatus defines the observed state of WatcherFlowLog.
type WatcherFlowLogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WatcherFlowLogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WatcherFlowLog is the Schema for the WatcherFlowLogs API. Manages a Network Watcher Flow Log.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WatcherFlowLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || has(self.initProvider.enabled)",message="enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.retentionPolicy) || has(self.initProvider.retentionPolicy)",message="retentionPolicy is a required parameter"
	Spec   WatcherFlowLogSpec   `json:"spec"`
	Status WatcherFlowLogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WatcherFlowLogList contains a list of WatcherFlowLogs
type WatcherFlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WatcherFlowLog `json:"items"`
}

// Repository type metadata.
var (
	WatcherFlowLog_Kind             = "WatcherFlowLog"
	WatcherFlowLog_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WatcherFlowLog_Kind}.String()
	WatcherFlowLog_KindAPIVersion   = WatcherFlowLog_Kind + "." + CRDGroupVersion.String()
	WatcherFlowLog_GroupVersionKind = CRDGroupVersion.WithKind(WatcherFlowLog_Kind)
)

func init() {
	SchemeBuilder.Register(&WatcherFlowLog{}, &WatcherFlowLogList{})
}
