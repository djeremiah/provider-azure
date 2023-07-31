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

type PrivateDNSAAAARecordInitParameters struct {

	// A list of IPv6 Addresses.
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSAAAARecordObservation struct {

	// The FQDN of the DNS AAAA Record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The Private DNS AAAA Record ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of IPv6 Addresses.
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`
}

type PrivateDNSAAAARecordParameters struct {

	// A list of IPv6 Addresses.
	// +kubebuilder:validation:Optional
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

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

// PrivateDNSAAAARecordSpec defines the desired state of PrivateDNSAAAARecord
type PrivateDNSAAAARecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSAAAARecordParameters `json:"forProvider"`
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
	InitProvider PrivateDNSAAAARecordInitParameters `json:"initProvider,omitempty"`
}

// PrivateDNSAAAARecordStatus defines the observed state of PrivateDNSAAAARecord.
type PrivateDNSAAAARecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSAAAARecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSAAAARecord is the Schema for the PrivateDNSAAAARecords API. Manages a Private DNS AAAA Record.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDNSAAAARecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.records) || has(self.initProvider.records)",message="records is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ttl) || has(self.initProvider.ttl)",message="ttl is a required parameter"
	Spec   PrivateDNSAAAARecordSpec   `json:"spec"`
	Status PrivateDNSAAAARecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSAAAARecordList contains a list of PrivateDNSAAAARecords
type PrivateDNSAAAARecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSAAAARecord `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSAAAARecord_Kind             = "PrivateDNSAAAARecord"
	PrivateDNSAAAARecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSAAAARecord_Kind}.String()
	PrivateDNSAAAARecord_KindAPIVersion   = PrivateDNSAAAARecord_Kind + "." + CRDGroupVersion.String()
	PrivateDNSAAAARecord_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSAAAARecord_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSAAAARecord{}, &PrivateDNSAAAARecordList{})
}
