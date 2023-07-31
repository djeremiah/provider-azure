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

type PrivateDNSTXTRecordInitParameters struct {

	// One or more record blocks as defined below.
	Record []PrivateDNSTXTRecordRecordInitParameters `json:"record,omitempty" tf:"record,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSTXTRecordObservation struct {

	// The FQDN of the DNS TXT Record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The Private DNS TXT Record ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more record blocks as defined below.
	Record []PrivateDNSTXTRecordRecordObservation `json:"record,omitempty" tf:"record,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`
}

type PrivateDNSTXTRecordParameters struct {

	// One or more record blocks as defined below.
	// +kubebuilder:validation:Optional
	Record []PrivateDNSTXTRecordRecordParameters `json:"record,omitempty" tf:"record,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Time To Live (TTL) of the DNS record in seconds.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=PrivateDNSZone
	// +kubebuilder:validation:Optional
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`

	// Reference to a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameRef *v1.Reference `json:"zoneNameRef,omitempty" tf:"-"`

	// Selector for a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameSelector *v1.Selector `json:"zoneNameSelector,omitempty" tf:"-"`
}

type PrivateDNSTXTRecordRecordInitParameters struct {

	// The value of the TXT record. Max length: 1024 characters
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PrivateDNSTXTRecordRecordObservation struct {

	// The value of the TXT record. Max length: 1024 characters
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PrivateDNSTXTRecordRecordParameters struct {

	// The value of the TXT record. Max length: 1024 characters
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// PrivateDNSTXTRecordSpec defines the desired state of PrivateDNSTXTRecord
type PrivateDNSTXTRecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSTXTRecordParameters `json:"forProvider"`
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
	InitProvider PrivateDNSTXTRecordInitParameters `json:"initProvider,omitempty"`
}

// PrivateDNSTXTRecordStatus defines the observed state of PrivateDNSTXTRecord.
type PrivateDNSTXTRecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSTXTRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSTXTRecord is the Schema for the PrivateDNSTXTRecords API. Manages a Private DNS TXT Record.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDNSTXTRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.record) || has(self.initProvider.record)",message="record is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ttl) || has(self.initProvider.ttl)",message="ttl is a required parameter"
	Spec   PrivateDNSTXTRecordSpec   `json:"spec"`
	Status PrivateDNSTXTRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSTXTRecordList contains a list of PrivateDNSTXTRecords
type PrivateDNSTXTRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSTXTRecord `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSTXTRecord_Kind             = "PrivateDNSTXTRecord"
	PrivateDNSTXTRecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSTXTRecord_Kind}.String()
	PrivateDNSTXTRecord_KindAPIVersion   = PrivateDNSTXTRecord_Kind + "." + CRDGroupVersion.String()
	PrivateDNSTXTRecord_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSTXTRecord_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSTXTRecord{}, &PrivateDNSTXTRecordList{})
}
