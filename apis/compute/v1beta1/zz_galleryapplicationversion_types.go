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

type GalleryApplicationVersionInitParameters struct {

	// Should the Gallery Application reports health. Defaults to false.
	EnableHealthCheck *bool `json:"enableHealthCheck,omitempty" tf:"enable_health_check,omitempty"`

	// The end of life date in RFC3339 format of the Gallery Application Version.
	EndOfLifeDate *string `json:"endOfLifeDate,omitempty" tf:"end_of_life_date,omitempty"`

	// Should the Gallery Application Version be excluded from the latest filter? If set to true this Gallery Application Version won't be returned for the latest version. Defaults to false.
	ExcludeFromLatest *bool `json:"excludeFromLatest,omitempty" tf:"exclude_from_latest,omitempty"`

	// The Azure Region where the Gallery Application Version exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A manage_action block as defined below.
	ManageAction []ManageActionInitParameters `json:"manageAction,omitempty" tf:"manage_action,omitempty"`

	// The version name of the Gallery Application Version, such as 1.0.0. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A source block as defined below.
	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`

	// A mapping of tags to assign to the Gallery Application Version.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more target_region blocks as defined below.
	TargetRegion []TargetRegionInitParameters `json:"targetRegion,omitempty" tf:"target_region,omitempty"`
}

type GalleryApplicationVersionObservation struct {

	// Should the Gallery Application reports health. Defaults to false.
	EnableHealthCheck *bool `json:"enableHealthCheck,omitempty" tf:"enable_health_check,omitempty"`

	// The end of life date in RFC3339 format of the Gallery Application Version.
	EndOfLifeDate *string `json:"endOfLifeDate,omitempty" tf:"end_of_life_date,omitempty"`

	// Should the Gallery Application Version be excluded from the latest filter? If set to true this Gallery Application Version won't be returned for the latest version. Defaults to false.
	ExcludeFromLatest *bool `json:"excludeFromLatest,omitempty" tf:"exclude_from_latest,omitempty"`

	// The ID of the Gallery Application. Changing this forces a new resource to be created.
	GalleryApplicationID *string `json:"galleryApplicationId,omitempty" tf:"gallery_application_id,omitempty"`

	// The ID of the Gallery Application Version.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Gallery Application Version exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A manage_action block as defined below.
	ManageAction []ManageActionObservation `json:"manageAction,omitempty" tf:"manage_action,omitempty"`

	// The version name of the Gallery Application Version, such as 1.0.0. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A source block as defined below.
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`

	// A mapping of tags to assign to the Gallery Application Version.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more target_region blocks as defined below.
	TargetRegion []TargetRegionObservation `json:"targetRegion,omitempty" tf:"target_region,omitempty"`
}

type GalleryApplicationVersionParameters struct {

	// Should the Gallery Application reports health. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableHealthCheck *bool `json:"enableHealthCheck,omitempty" tf:"enable_health_check,omitempty"`

	// The end of life date in RFC3339 format of the Gallery Application Version.
	// +kubebuilder:validation:Optional
	EndOfLifeDate *string `json:"endOfLifeDate,omitempty" tf:"end_of_life_date,omitempty"`

	// Should the Gallery Application Version be excluded from the latest filter? If set to true this Gallery Application Version won't be returned for the latest version. Defaults to false.
	// +kubebuilder:validation:Optional
	ExcludeFromLatest *bool `json:"excludeFromLatest,omitempty" tf:"exclude_from_latest,omitempty"`

	// The ID of the Gallery Application. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.GalleryApplication
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GalleryApplicationID *string `json:"galleryApplicationId,omitempty" tf:"gallery_application_id,omitempty"`

	// Reference to a GalleryApplication in compute to populate galleryApplicationId.
	// +kubebuilder:validation:Optional
	GalleryApplicationIDRef *v1.Reference `json:"galleryApplicationIdRef,omitempty" tf:"-"`

	// Selector for a GalleryApplication in compute to populate galleryApplicationId.
	// +kubebuilder:validation:Optional
	GalleryApplicationIDSelector *v1.Selector `json:"galleryApplicationIdSelector,omitempty" tf:"-"`

	// The Azure Region where the Gallery Application Version exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A manage_action block as defined below.
	// +kubebuilder:validation:Optional
	ManageAction []ManageActionParameters `json:"manageAction,omitempty" tf:"manage_action,omitempty"`

	// The version name of the Gallery Application Version, such as 1.0.0. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A source block as defined below.
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// A mapping of tags to assign to the Gallery Application Version.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more target_region blocks as defined below.
	// +kubebuilder:validation:Optional
	TargetRegion []TargetRegionParameters `json:"targetRegion,omitempty" tf:"target_region,omitempty"`
}

type ManageActionInitParameters struct {

	// The command to install the Gallery Application. Changing this forces a new resource to be created.
	Install *string `json:"install,omitempty" tf:"install,omitempty"`

	// The command to remove the Gallery Application. Changing this forces a new resource to be created.
	Remove *string `json:"remove,omitempty" tf:"remove,omitempty"`

	// The command to update the Gallery Application. Changing this forces a new resource to be created.
	Update *string `json:"update,omitempty" tf:"update,omitempty"`
}

type ManageActionObservation struct {

	// The command to install the Gallery Application. Changing this forces a new resource to be created.
	Install *string `json:"install,omitempty" tf:"install,omitempty"`

	// The command to remove the Gallery Application. Changing this forces a new resource to be created.
	Remove *string `json:"remove,omitempty" tf:"remove,omitempty"`

	// The command to update the Gallery Application. Changing this forces a new resource to be created.
	Update *string `json:"update,omitempty" tf:"update,omitempty"`
}

type ManageActionParameters struct {

	// The command to install the Gallery Application. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Install *string `json:"install,omitempty" tf:"install,omitempty"`

	// The command to remove the Gallery Application. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Remove *string `json:"remove,omitempty" tf:"remove,omitempty"`

	// The command to update the Gallery Application. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Update *string `json:"update,omitempty" tf:"update,omitempty"`
}

type SourceInitParameters struct {

	// The Storage Blob URI of the default configuration. Changing this forces a new resource to be created.
	DefaultConfigurationLink *string `json:"defaultConfigurationLink,omitempty" tf:"default_configuration_link,omitempty"`
}

type SourceObservation struct {

	// The Storage Blob URI of the default configuration. Changing this forces a new resource to be created.
	DefaultConfigurationLink *string `json:"defaultConfigurationLink,omitempty" tf:"default_configuration_link,omitempty"`

	// The Storage Blob URI of the source application package. Changing this forces a new resource to be created.
	MediaLink *string `json:"mediaLink,omitempty" tf:"media_link,omitempty"`
}

type SourceParameters struct {

	// The Storage Blob URI of the default configuration. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DefaultConfigurationLink *string `json:"defaultConfigurationLink,omitempty" tf:"default_configuration_link,omitempty"`

	// The Storage Blob URI of the source application package. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Blob
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MediaLink *string `json:"mediaLink,omitempty" tf:"media_link,omitempty"`

	// Reference to a Blob in storage to populate mediaLink.
	// +kubebuilder:validation:Optional
	MediaLinkRef *v1.Reference `json:"mediaLinkRef,omitempty" tf:"-"`

	// Selector for a Blob in storage to populate mediaLink.
	// +kubebuilder:validation:Optional
	MediaLinkSelector *v1.Selector `json:"mediaLinkSelector,omitempty" tf:"-"`
}

type TargetRegionInitParameters struct {

	// The number of replicas of the Gallery Application Version to be created per region. Possible values are between 1 and 10.
	RegionalReplicaCount *float64 `json:"regionalReplicaCount,omitempty" tf:"regional_replica_count,omitempty"`

	// The storage account type for the Gallery Application Version. Possible values are Standard_LRS, Premium_LRS and Standard_ZRS. Defaults to Standard_LRS.
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type,omitempty"`
}

type TargetRegionObservation struct {

	// The Azure Region in which the Gallery Application Version exists.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of replicas of the Gallery Application Version to be created per region. Possible values are between 1 and 10.
	RegionalReplicaCount *float64 `json:"regionalReplicaCount,omitempty" tf:"regional_replica_count,omitempty"`

	// The storage account type for the Gallery Application Version. Possible values are Standard_LRS, Premium_LRS and Standard_ZRS. Defaults to Standard_LRS.
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type,omitempty"`
}

type TargetRegionParameters struct {

	// The Azure Region in which the Gallery Application Version exists.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.GalleryApplication
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("location",false)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a GalleryApplication in compute to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a GalleryApplication in compute to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// The number of replicas of the Gallery Application Version to be created per region. Possible values are between 1 and 10.
	// +kubebuilder:validation:Optional
	RegionalReplicaCount *float64 `json:"regionalReplicaCount,omitempty" tf:"regional_replica_count,omitempty"`

	// The storage account type for the Gallery Application Version. Possible values are Standard_LRS, Premium_LRS and Standard_ZRS. Defaults to Standard_LRS.
	// +kubebuilder:validation:Optional
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type,omitempty"`
}

// GalleryApplicationVersionSpec defines the desired state of GalleryApplicationVersion
type GalleryApplicationVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GalleryApplicationVersionParameters `json:"forProvider"`
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
	InitProvider GalleryApplicationVersionInitParameters `json:"initProvider,omitempty"`
}

// GalleryApplicationVersionStatus defines the observed state of GalleryApplicationVersion.
type GalleryApplicationVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GalleryApplicationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GalleryApplicationVersion is the Schema for the GalleryApplicationVersions API. Manages a Gallery Application Version.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type GalleryApplicationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.manageAction) || has(self.initProvider.manageAction)",message="manageAction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.source) || has(self.initProvider.source)",message="source is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetRegion) || has(self.initProvider.targetRegion)",message="targetRegion is a required parameter"
	Spec   GalleryApplicationVersionSpec   `json:"spec"`
	Status GalleryApplicationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GalleryApplicationVersionList contains a list of GalleryApplicationVersions
type GalleryApplicationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GalleryApplicationVersion `json:"items"`
}

// Repository type metadata.
var (
	GalleryApplicationVersion_Kind             = "GalleryApplicationVersion"
	GalleryApplicationVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GalleryApplicationVersion_Kind}.String()
	GalleryApplicationVersion_KindAPIVersion   = GalleryApplicationVersion_Kind + "." + CRDGroupVersion.String()
	GalleryApplicationVersion_GroupVersionKind = CRDGroupVersion.WithKind(GalleryApplicationVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&GalleryApplicationVersion{}, &GalleryApplicationVersionList{})
}
