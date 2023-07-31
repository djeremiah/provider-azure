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

type QuotaInitParameters struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 1Gi if not specified.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type QuotaObservation struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 1Gi if not specified.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type QuotaParameters struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	// +kubebuilder:validation:Optional
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 1Gi if not specified.
	// +kubebuilder:validation:Optional
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type SpringCloudBuildDeploymentInitParameters struct {

	// A JSON object that contains the addon configurations of the Spring Cloud Build Deployment.
	AddonJSON *string `json:"addonJson,omitempty" tf:"addon_json,omitempty"`

	// The ID of the Spring Cloud Build Result.
	BuildResultID *string `json:"buildResultId,omitempty" tf:"build_result_id,omitempty"`

	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// A quota block as defined below.
	Quota []QuotaInitParameters `json:"quota,omitempty" tf:"quota,omitempty"`
}

type SpringCloudBuildDeploymentObservation struct {

	// A JSON object that contains the addon configurations of the Spring Cloud Build Deployment.
	AddonJSON *string `json:"addonJson,omitempty" tf:"addon_json,omitempty"`

	// The ID of the Spring Cloud Build Result.
	BuildResultID *string `json:"buildResultId,omitempty" tf:"build_result_id,omitempty"`

	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// The ID of the Spring Cloud Build Deployment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// A quota block as defined below.
	Quota []QuotaObservation `json:"quota,omitempty" tf:"quota,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Build Deployment to be created.
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`
}

type SpringCloudBuildDeploymentParameters struct {

	// A JSON object that contains the addon configurations of the Spring Cloud Build Deployment.
	// +kubebuilder:validation:Optional
	AddonJSON *string `json:"addonJson,omitempty" tf:"addon_json,omitempty"`

	// The ID of the Spring Cloud Build Result.
	// +kubebuilder:validation:Optional
	BuildResultID *string `json:"buildResultId,omitempty" tf:"build_result_id,omitempty"`

	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	// +kubebuilder:validation:Optional
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// A quota block as defined below.
	// +kubebuilder:validation:Optional
	Quota []QuotaParameters `json:"quota,omitempty" tf:"quota,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Build Deployment to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudApp
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`

	// Reference to a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDRef *v1.Reference `json:"springCloudAppIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDSelector *v1.Selector `json:"springCloudAppIdSelector,omitempty" tf:"-"`
}

// SpringCloudBuildDeploymentSpec defines the desired state of SpringCloudBuildDeployment
type SpringCloudBuildDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudBuildDeploymentParameters `json:"forProvider"`
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
	InitProvider SpringCloudBuildDeploymentInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudBuildDeploymentStatus defines the observed state of SpringCloudBuildDeployment.
type SpringCloudBuildDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudBuildDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudBuildDeployment is the Schema for the SpringCloudBuildDeployments API. Manages a Spring Cloud Build Deployment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudBuildDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.buildResultId) || has(self.initProvider.buildResultId)",message="buildResultId is a required parameter"
	Spec   SpringCloudBuildDeploymentSpec   `json:"spec"`
	Status SpringCloudBuildDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudBuildDeploymentList contains a list of SpringCloudBuildDeployments
type SpringCloudBuildDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudBuildDeployment `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudBuildDeployment_Kind             = "SpringCloudBuildDeployment"
	SpringCloudBuildDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudBuildDeployment_Kind}.String()
	SpringCloudBuildDeployment_KindAPIVersion   = SpringCloudBuildDeployment_Kind + "." + CRDGroupVersion.String()
	SpringCloudBuildDeployment_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudBuildDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudBuildDeployment{}, &SpringCloudBuildDeploymentList{})
}
