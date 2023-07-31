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

type DatabaseMigrationProjectInitParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specify the name of the database migration project. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The platform type of the migration source. Currently only support: SQL(on-premises SQL Server). Changing this forces a new resource to be created.
	SourcePlatform *string `json:"sourcePlatform,omitempty" tf:"source_platform,omitempty"`

	// A mapping of tags to assigned to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The platform type of the migration target. Currently only support: SQLDB(Azure SQL Database). Changing this forces a new resource to be created.
	TargetPlatform *string `json:"targetPlatform,omitempty" tf:"target_platform,omitempty"`
}

type DatabaseMigrationProjectObservation struct {

	// The ID of Database Migration Project.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specify the name of the database migration project. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The platform type of the migration source. Currently only support: SQL(on-premises SQL Server). Changing this forces a new resource to be created.
	SourcePlatform *string `json:"sourcePlatform,omitempty" tf:"source_platform,omitempty"`

	// A mapping of tags to assigned to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The platform type of the migration target. Currently only support: SQLDB(Azure SQL Database). Changing this forces a new resource to be created.
	TargetPlatform *string `json:"targetPlatform,omitempty" tf:"target_platform,omitempty"`
}

type DatabaseMigrationProjectParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specify the name of the database migration project. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datamigration/v1beta1.DatabaseMigrationService
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Reference to a DatabaseMigrationService in datamigration to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameRef *v1.Reference `json:"serviceNameRef,omitempty" tf:"-"`

	// Selector for a DatabaseMigrationService in datamigration to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameSelector *v1.Selector `json:"serviceNameSelector,omitempty" tf:"-"`

	// The platform type of the migration source. Currently only support: SQL(on-premises SQL Server). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SourcePlatform *string `json:"sourcePlatform,omitempty" tf:"source_platform,omitempty"`

	// A mapping of tags to assigned to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The platform type of the migration target. Currently only support: SQLDB(Azure SQL Database). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TargetPlatform *string `json:"targetPlatform,omitempty" tf:"target_platform,omitempty"`
}

// DatabaseMigrationProjectSpec defines the desired state of DatabaseMigrationProject
type DatabaseMigrationProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseMigrationProjectParameters `json:"forProvider"`
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
	InitProvider DatabaseMigrationProjectInitParameters `json:"initProvider,omitempty"`
}

// DatabaseMigrationProjectStatus defines the observed state of DatabaseMigrationProject.
type DatabaseMigrationProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseMigrationProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseMigrationProject is the Schema for the DatabaseMigrationProjects API. Manage Azure Database Migration Project instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DatabaseMigrationProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourcePlatform) || has(self.initProvider.sourcePlatform)",message="sourcePlatform is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetPlatform) || has(self.initProvider.targetPlatform)",message="targetPlatform is a required parameter"
	Spec   DatabaseMigrationProjectSpec   `json:"spec"`
	Status DatabaseMigrationProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseMigrationProjectList contains a list of DatabaseMigrationProjects
type DatabaseMigrationProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseMigrationProject `json:"items"`
}

// Repository type metadata.
var (
	DatabaseMigrationProject_Kind             = "DatabaseMigrationProject"
	DatabaseMigrationProject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseMigrationProject_Kind}.String()
	DatabaseMigrationProject_KindAPIVersion   = DatabaseMigrationProject_Kind + "." + CRDGroupVersion.String()
	DatabaseMigrationProject_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseMigrationProject_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseMigrationProject{}, &DatabaseMigrationProjectList{})
}
