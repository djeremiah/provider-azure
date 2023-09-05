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

type WorkspaceCustomerManagedKeyInitParameters struct {
}

type WorkspaceCustomerManagedKeyObservation struct {

	// The ID of the Databricks Workspace.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Key Vault.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// The ID of the Databricks Workspace..
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type WorkspaceCustomerManagedKeyParameters struct {

	// The ID of the Key Vault.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Reference to a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDRef *v1.Reference `json:"keyVaultKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDSelector *v1.Selector `json:"keyVaultKeyIdSelector,omitempty" tf:"-"`

	// The ID of the Databricks Workspace..
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/databricks/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in databricks to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in databricks to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

// WorkspaceCustomerManagedKeySpec defines the desired state of WorkspaceCustomerManagedKey
type WorkspaceCustomerManagedKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceCustomerManagedKeyParameters `json:"forProvider"`
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
	InitProvider WorkspaceCustomerManagedKeyInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceCustomerManagedKeyStatus defines the observed state of WorkspaceCustomerManagedKey.
type WorkspaceCustomerManagedKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceCustomerManagedKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceCustomerManagedKey is the Schema for the WorkspaceCustomerManagedKeys API. Manages a Customer Managed Key for a Databricks Workspace root DBFS
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WorkspaceCustomerManagedKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceCustomerManagedKeySpec   `json:"spec"`
	Status            WorkspaceCustomerManagedKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceCustomerManagedKeyList contains a list of WorkspaceCustomerManagedKeys
type WorkspaceCustomerManagedKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceCustomerManagedKey `json:"items"`
}

// Repository type metadata.
var (
	WorkspaceCustomerManagedKey_Kind             = "WorkspaceCustomerManagedKey"
	WorkspaceCustomerManagedKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkspaceCustomerManagedKey_Kind}.String()
	WorkspaceCustomerManagedKey_KindAPIVersion   = WorkspaceCustomerManagedKey_Kind + "." + CRDGroupVersion.String()
	WorkspaceCustomerManagedKey_GroupVersionKind = CRDGroupVersion.WithKind(WorkspaceCustomerManagedKey_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkspaceCustomerManagedKey{}, &WorkspaceCustomerManagedKeyList{})
}
