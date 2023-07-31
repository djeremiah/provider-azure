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

type PrivateEndpointApplicationSecurityGroupAssociationInitParameters struct {
}

type PrivateEndpointApplicationSecurityGroupAssociationObservation struct {

	// The id of application security group to associate. Changing this forces a new resource to be created.
	ApplicationSecurityGroupID *string `json:"applicationSecurityGroupId,omitempty" tf:"application_security_group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The id of private endpoint to associate. Changing this forces a new resource to be created.
	PrivateEndpointID *string `json:"privateEndpointId,omitempty" tf:"private_endpoint_id,omitempty"`
}

type PrivateEndpointApplicationSecurityGroupAssociationParameters struct {

	// The id of application security group to associate. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.ApplicationSecurityGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationSecurityGroupID *string `json:"applicationSecurityGroupId,omitempty" tf:"application_security_group_id,omitempty"`

	// Reference to a ApplicationSecurityGroup in network to populate applicationSecurityGroupId.
	// +kubebuilder:validation:Optional
	ApplicationSecurityGroupIDRef *v1.Reference `json:"applicationSecurityGroupIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationSecurityGroup in network to populate applicationSecurityGroupId.
	// +kubebuilder:validation:Optional
	ApplicationSecurityGroupIDSelector *v1.Selector `json:"applicationSecurityGroupIdSelector,omitempty" tf:"-"`

	// The id of private endpoint to associate. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PrivateEndpoint
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrivateEndpointID *string `json:"privateEndpointId,omitempty" tf:"private_endpoint_id,omitempty"`

	// Reference to a PrivateEndpoint in network to populate privateEndpointId.
	// +kubebuilder:validation:Optional
	PrivateEndpointIDRef *v1.Reference `json:"privateEndpointIdRef,omitempty" tf:"-"`

	// Selector for a PrivateEndpoint in network to populate privateEndpointId.
	// +kubebuilder:validation:Optional
	PrivateEndpointIDSelector *v1.Selector `json:"privateEndpointIdSelector,omitempty" tf:"-"`
}

// PrivateEndpointApplicationSecurityGroupAssociationSpec defines the desired state of PrivateEndpointApplicationSecurityGroupAssociation
type PrivateEndpointApplicationSecurityGroupAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateEndpointApplicationSecurityGroupAssociationParameters `json:"forProvider"`
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
	InitProvider PrivateEndpointApplicationSecurityGroupAssociationInitParameters `json:"initProvider,omitempty"`
}

// PrivateEndpointApplicationSecurityGroupAssociationStatus defines the observed state of PrivateEndpointApplicationSecurityGroupAssociation.
type PrivateEndpointApplicationSecurityGroupAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateEndpointApplicationSecurityGroupAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateEndpointApplicationSecurityGroupAssociation is the Schema for the PrivateEndpointApplicationSecurityGroupAssociations API. Manages an association between Private Endpoint and Application Security Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateEndpointApplicationSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateEndpointApplicationSecurityGroupAssociationSpec   `json:"spec"`
	Status            PrivateEndpointApplicationSecurityGroupAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateEndpointApplicationSecurityGroupAssociationList contains a list of PrivateEndpointApplicationSecurityGroupAssociations
type PrivateEndpointApplicationSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateEndpointApplicationSecurityGroupAssociation `json:"items"`
}

// Repository type metadata.
var (
	PrivateEndpointApplicationSecurityGroupAssociation_Kind             = "PrivateEndpointApplicationSecurityGroupAssociation"
	PrivateEndpointApplicationSecurityGroupAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateEndpointApplicationSecurityGroupAssociation_Kind}.String()
	PrivateEndpointApplicationSecurityGroupAssociation_KindAPIVersion   = PrivateEndpointApplicationSecurityGroupAssociation_Kind + "." + CRDGroupVersion.String()
	PrivateEndpointApplicationSecurityGroupAssociation_GroupVersionKind = CRDGroupVersion.WithKind(PrivateEndpointApplicationSecurityGroupAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateEndpointApplicationSecurityGroupAssociation{}, &PrivateEndpointApplicationSecurityGroupAssociationList{})
}
