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

type FlexibleServerFirewallRuleInitParameters struct {

	// Specifies the End IP Address associated with this Firewall Rule.
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// Specifies the Start IP Address associated with this Firewall Rule.
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`
}

type FlexibleServerFirewallRuleObservation struct {

	// Specifies the End IP Address associated with this Firewall Rule.
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// The ID of the MySQL Firewall Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Specifies the Start IP Address associated with this Firewall Rule.
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`
}

type FlexibleServerFirewallRuleParameters struct {

	// Specifies the End IP Address associated with this Firewall Rule.
	// +kubebuilder:validation:Optional
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=FlexibleServer
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Reference to a FlexibleServer to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// Selector for a FlexibleServer to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// Specifies the Start IP Address associated with this Firewall Rule.
	// +kubebuilder:validation:Optional
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`
}

// FlexibleServerFirewallRuleSpec defines the desired state of FlexibleServerFirewallRule
type FlexibleServerFirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleServerFirewallRuleParameters `json:"forProvider"`
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
	InitProvider FlexibleServerFirewallRuleInitParameters `json:"initProvider,omitempty"`
}

// FlexibleServerFirewallRuleStatus defines the observed state of FlexibleServerFirewallRule.
type FlexibleServerFirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleServerFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerFirewallRule is the Schema for the FlexibleServerFirewallRules API. Manages a Firewall Rule for a MySQL Flexible Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FlexibleServerFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endIpAddress) || has(self.initProvider.endIpAddress)",message="endIpAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.startIpAddress) || has(self.initProvider.startIpAddress)",message="startIpAddress is a required parameter"
	Spec   FlexibleServerFirewallRuleSpec   `json:"spec"`
	Status FlexibleServerFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerFirewallRuleList contains a list of FlexibleServerFirewallRules
type FlexibleServerFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServerFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	FlexibleServerFirewallRule_Kind             = "FlexibleServerFirewallRule"
	FlexibleServerFirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlexibleServerFirewallRule_Kind}.String()
	FlexibleServerFirewallRule_KindAPIVersion   = FlexibleServerFirewallRule_Kind + "." + CRDGroupVersion.String()
	FlexibleServerFirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(FlexibleServerFirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&FlexibleServerFirewallRule{}, &FlexibleServerFirewallRuleList{})
}
