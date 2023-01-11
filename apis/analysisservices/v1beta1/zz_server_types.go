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

type IPv4FirewallRuleObservation struct {
}

type IPv4FirewallRuleParameters struct {

	// Specifies the name of the firewall rule.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// End of the firewall rule range as IPv4 address.
	// +kubebuilder:validation:Required
	RangeEnd *string `json:"rangeEnd" tf:"range_end,omitempty"`

	// Start of the firewall rule range as IPv4 address.
	// +kubebuilder:validation:Required
	RangeStart *string `json:"rangeStart" tf:"range_start,omitempty"`
}

type ServerObservation struct {

	// The ID of the Analysis Services Server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The full name of the Analysis Services Server.
	ServerFullName *string `json:"serverFullName,omitempty" tf:"server_full_name,omitempty"`
}

type ServerParameters struct {

	// List of email addresses of admin users.
	// +kubebuilder:validation:Optional
	AdminUsers []*string `json:"adminUsers,omitempty" tf:"admin_users,omitempty"`

	// URI and SAS token for a blob container to store backups.
	// +kubebuilder:validation:Optional
	BackupBlobContainerURISecretRef *v1.SecretKeySelector `json:"backupBlobContainerUriSecretRef,omitempty" tf:"-"`

	// Indicates if the Power BI service is allowed to access or not.
	// +kubebuilder:validation:Optional
	EnablePowerBiService *bool `json:"enablePowerBiService,omitempty" tf:"enable_power_bi_service,omitempty"`

	// One or more ipv4_firewall_rule block(s) as defined below.
	// +kubebuilder:validation:Optional
	IPv4FirewallRule []IPv4FirewallRuleParameters `json:"ipv4FirewallRule,omitempty" tf:"ipv4_firewall_rule,omitempty"`

	// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the Analysis Services Server. Only lowercase Alphanumeric characters allowed, starting with a letter. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Controls how the read-write server is used in the query pool. If this value is set to All then read-write servers are also used for queries. Otherwise with ReadOnly these servers do not participate in query operations.
	// +kubebuilder:validation:Optional
	QuerypoolConnectionMode *string `json:"querypoolConnectionMode,omitempty" tf:"querypool_connection_mode,omitempty"`

	// The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// SKU for the Analysis Services Server. Possible values are: D1, B1, B2, S0, S1, S2, S4, S8, S9, S8v2 and S9v2.
	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerParameters `json:"forProvider"`
}

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Server is the Schema for the Servers API. Manages an Analysis Services Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec"`
	Status            ServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerList contains a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Repository type metadata.
var (
	Server_Kind             = "Server"
	Server_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Server_Kind}.String()
	Server_KindAPIVersion   = Server_Kind + "." + CRDGroupVersion.String()
	Server_GroupVersionKind = CRDGroupVersion.WithKind(Server_Kind)
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
