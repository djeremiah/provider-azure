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

type IOTHubDPSInitParameters struct {

	// The allocation policy of the IoT Device Provisioning Service (Hashed, GeoLatency or Static). Defaults to Hashed.
	AllocationPolicy *string `json:"allocationPolicy,omitempty" tf:"allocation_policy,omitempty"`

	// Specifies if the IoT Device Provisioning Service has data residency and disaster recovery enabled. Defaults to false. Changing this forces a new resource to be created.
	DataResidencyEnabled *bool `json:"dataResidencyEnabled,omitempty" tf:"data_residency_enabled,omitempty"`

	// An ip_filter_rule block as defined below.
	IPFilterRule []IPFilterRuleInitParameters `json:"ipFilterRule,omitempty" tf:"ip_filter_rule,omitempty"`

	// A linked_hub block as defined below.
	LinkedHub []LinkedHubInitParameters `json:"linkedHub,omitempty" tf:"linked_hub,omitempty"`

	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether requests from Public Network are allowed. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A sku block as defined below.
	Sku []IOTHubDPSSkuInitParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IOTHubDPSObservation struct {

	// The allocation policy of the IoT Device Provisioning Service (Hashed, GeoLatency or Static). Defaults to Hashed.
	AllocationPolicy *string `json:"allocationPolicy,omitempty" tf:"allocation_policy,omitempty"`

	// Specifies if the IoT Device Provisioning Service has data residency and disaster recovery enabled. Defaults to false. Changing this forces a new resource to be created.
	DataResidencyEnabled *bool `json:"dataResidencyEnabled,omitempty" tf:"data_residency_enabled,omitempty"`

	// The device endpoint of the IoT Device Provisioning Service.
	DeviceProvisioningHostName *string `json:"deviceProvisioningHostName,omitempty" tf:"device_provisioning_host_name,omitempty"`

	// The ID of the IoT Device Provisioning Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier of the IoT Device Provisioning Service.
	IDScope *string `json:"idScope,omitempty" tf:"id_scope,omitempty"`

	// An ip_filter_rule block as defined below.
	IPFilterRule []IPFilterRuleObservation `json:"ipFilterRule,omitempty" tf:"ip_filter_rule,omitempty"`

	// A linked_hub block as defined below.
	LinkedHub []LinkedHubObservation `json:"linkedHub,omitempty" tf:"linked_hub,omitempty"`

	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether requests from Public Network are allowed. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The service endpoint of the IoT Device Provisioning Service.
	ServiceOperationsHostName *string `json:"serviceOperationsHostName,omitempty" tf:"service_operations_host_name,omitempty"`

	// A sku block as defined below.
	Sku []IOTHubDPSSkuObservation `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IOTHubDPSParameters struct {

	// The allocation policy of the IoT Device Provisioning Service (Hashed, GeoLatency or Static). Defaults to Hashed.
	// +kubebuilder:validation:Optional
	AllocationPolicy *string `json:"allocationPolicy,omitempty" tf:"allocation_policy,omitempty"`

	// Specifies if the IoT Device Provisioning Service has data residency and disaster recovery enabled. Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DataResidencyEnabled *bool `json:"dataResidencyEnabled,omitempty" tf:"data_residency_enabled,omitempty"`

	// An ip_filter_rule block as defined below.
	// +kubebuilder:validation:Optional
	IPFilterRule []IPFilterRuleParameters `json:"ipFilterRule,omitempty" tf:"ip_filter_rule,omitempty"`

	// A linked_hub block as defined below.
	// +kubebuilder:validation:Optional
	LinkedHub []LinkedHubParameters `json:"linkedHub,omitempty" tf:"linked_hub,omitempty"`

	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether requests from Public Network are allowed. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A sku block as defined below.
	// +kubebuilder:validation:Optional
	Sku []IOTHubDPSSkuParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IOTHubDPSSkuInitParameters struct {

	// The number of provisioned IoT Device Provisioning Service units.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The name of the sku. Currently can only be set to S1.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IOTHubDPSSkuObservation struct {

	// The number of provisioned IoT Device Provisioning Service units.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The name of the sku. Currently can only be set to S1.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IOTHubDPSSkuParameters struct {

	// The number of provisioned IoT Device Provisioning Service units.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The name of the sku. Currently can only be set to S1.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IPFilterRuleInitParameters struct {

	// The desired action for requests captured by this rule. Possible values are Accept, Reject
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The IP address range in CIDR notation for the rule.
	IPMask *string `json:"ipMask,omitempty" tf:"ip_mask,omitempty"`

	// The name of the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target for requests captured by this rule. Possible values are All, DeviceApi and ServiceApi.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type IPFilterRuleObservation struct {

	// The desired action for requests captured by this rule. Possible values are Accept, Reject
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The IP address range in CIDR notation for the rule.
	IPMask *string `json:"ipMask,omitempty" tf:"ip_mask,omitempty"`

	// The name of the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target for requests captured by this rule. Possible values are All, DeviceApi and ServiceApi.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type IPFilterRuleParameters struct {

	// The desired action for requests captured by this rule. Possible values are Accept, Reject
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The IP address range in CIDR notation for the rule.
	// +kubebuilder:validation:Optional
	IPMask *string `json:"ipMask,omitempty" tf:"ip_mask,omitempty"`

	// The name of the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target for requests captured by this rule. Possible values are All, DeviceApi and ServiceApi.
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type LinkedHubInitParameters struct {

	// The weight applied to the IoT Hub. Defaults to 1.
	AllocationWeight *float64 `json:"allocationWeight,omitempty" tf:"allocation_weight,omitempty"`

	// Determines whether to apply allocation policies to the IoT Hub. Defaults to true.
	ApplyAllocationPolicy *bool `json:"applyAllocationPolicy,omitempty" tf:"apply_allocation_policy,omitempty"`

	// The location of the IoT hub.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type LinkedHubObservation struct {

	// The weight applied to the IoT Hub. Defaults to 1.
	AllocationWeight *float64 `json:"allocationWeight,omitempty" tf:"allocation_weight,omitempty"`

	// Determines whether to apply allocation policies to the IoT Hub. Defaults to true.
	ApplyAllocationPolicy *bool `json:"applyAllocationPolicy,omitempty" tf:"apply_allocation_policy,omitempty"`

	// (Computed) The IoT Hub hostname.
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The location of the IoT hub.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type LinkedHubParameters struct {

	// The weight applied to the IoT Hub. Defaults to 1.
	// +kubebuilder:validation:Optional
	AllocationWeight *float64 `json:"allocationWeight,omitempty" tf:"allocation_weight,omitempty"`

	// Determines whether to apply allocation policies to the IoT Hub. Defaults to true.
	// +kubebuilder:validation:Optional
	ApplyAllocationPolicy *bool `json:"applyAllocationPolicy,omitempty" tf:"apply_allocation_policy,omitempty"`

	// The connection string to connect to the IoT Hub.
	// +kubebuilder:validation:Required
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// The location of the IoT hub.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

// IOTHubDPSSpec defines the desired state of IOTHubDPS
type IOTHubDPSSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubDPSParameters `json:"forProvider"`
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
	InitProvider IOTHubDPSInitParameters `json:"initProvider,omitempty"`
}

// IOTHubDPSStatus defines the observed state of IOTHubDPS.
type IOTHubDPSStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubDPSObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubDPS is the Schema for the IOTHubDPSs API. Manages an IoT Device Provisioning Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure},path=iothubdps
type IOTHubDPS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || has(self.initProvider.sku)",message="sku is a required parameter"
	Spec   IOTHubDPSSpec   `json:"spec"`
	Status IOTHubDPSStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubDPSList contains a list of IOTHubDPSs
type IOTHubDPSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubDPS `json:"items"`
}

// Repository type metadata.
var (
	IOTHubDPS_Kind             = "IOTHubDPS"
	IOTHubDPS_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubDPS_Kind}.String()
	IOTHubDPS_KindAPIVersion   = IOTHubDPS_Kind + "." + CRDGroupVersion.String()
	IOTHubDPS_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubDPS_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubDPS{}, &IOTHubDPSList{})
}
