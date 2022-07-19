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

type BackupVaultObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`
}

type BackupVaultParameters struct {

	// +kubebuilder:validation:Required
	DatastoreType *string `json:"datastoreType" tf:"datastore_type,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Redundancy *string `json:"redundancy" tf:"redundancy,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// BackupVaultSpec defines the desired state of BackupVault
type BackupVaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupVaultParameters `json:"forProvider"`
}

// BackupVaultStatus defines the observed state of BackupVault.
type BackupVaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupVaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupVault is the Schema for the BackupVaults API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupVaultSpec   `json:"spec"`
	Status            BackupVaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupVaultList contains a list of BackupVaults
type BackupVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupVault `json:"items"`
}

// Repository type metadata.
var (
	BackupVault_Kind             = "BackupVault"
	BackupVault_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupVault_Kind}.String()
	BackupVault_KindAPIVersion   = BackupVault_Kind + "." + CRDGroupVersion.String()
	BackupVault_GroupVersionKind = CRDGroupVersion.WithKind(BackupVault_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupVault{}, &BackupVaultList{})
}
