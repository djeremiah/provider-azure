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

type DashboardObservation struct {

	// The ID of the Dashboard.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DashboardParameters struct {

	// JSON data representing dashboard body. See above for details on how to obtain this from the Portal.
	// +kubebuilder:validation:Optional
	DashboardProperties *string `json:"dashboardProperties,omitempty" tf:"dashboard_properties,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the name of the Shared Dashboard. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The name of the resource group in which to create the dashboard. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DashboardSpec defines the desired state of Dashboard
type DashboardSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DashboardParameters `json:"forProvider"`
}

// DashboardStatus defines the observed state of Dashboard.
type DashboardStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DashboardObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Dashboard is the Schema for the Dashboards API. Manages a shared dashboard in the Azure Portal.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Dashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DashboardSpec   `json:"spec"`
	Status            DashboardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DashboardList contains a list of Dashboards
type DashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dashboard `json:"items"`
}

// Repository type metadata.
var (
	Dashboard_Kind             = "Dashboard"
	Dashboard_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dashboard_Kind}.String()
	Dashboard_KindAPIVersion   = Dashboard_Kind + "." + CRDGroupVersion.String()
	Dashboard_GroupVersionKind = CRDGroupVersion.WithKind(Dashboard_Kind)
)

func init() {
	SchemeBuilder.Register(&Dashboard{}, &DashboardList{})
}