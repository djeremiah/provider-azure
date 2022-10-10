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

type CustomRouteObservation struct {
}

type CustomRouteParameters struct {

	// A list of address blocks reserved for this virtual network in CIDR notation.
	// +kubebuilder:validation:Optional
	AddressPrefixes []*string `json:"addressPrefixes,omitempty" tf:"address_prefixes,omitempty"`
}

type PeeringAddressesObservation struct {

	// A list of peering address assigned to the BGP peer of the Virtual Network Gateway.
	DefaultAddresses []*string `json:"defaultAddresses,omitempty" tf:"default_addresses,omitempty"`

	// A list of tunnel IP addresses assigned to the BGP peer of the Virtual Network Gateway.
	TunnelIPAddresses []*string `json:"tunnelIpAddresses,omitempty" tf:"tunnel_ip_addresses,omitempty"`
}

type PeeringAddressesParameters struct {

	// A list of Azure custom APIPA addresses assigned to the BGP peer of the Virtual Network Gateway.
	// +kubebuilder:validation:Optional
	ApipaAddresses []*string `json:"apipaAddresses,omitempty" tf:"apipa_addresses,omitempty"`

	// The name of the IP configuration of this Virtual Network Gateway. In case there are multiple ip_configuration blocks defined, this property is required to specify.
	// +kubebuilder:validation:Optional
	IPConfigurationName *string `json:"ipConfigurationName,omitempty" tf:"ip_configuration_name,omitempty"`
}

type RevokedCertificateObservation struct {
}

type RevokedCertificateParameters struct {

	// A user-defined name of the revoked certificate.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Thumbprint *string `json:"thumbprint" tf:"thumbprint,omitempty"`
}

type RootCertificateObservation struct {
}

type RootCertificateParameters struct {

	// A user-defined name of the revoked certificate.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The SHA1 thumbprint of the certificate to be
	// revoked.
	// +kubebuilder:validation:Required
	PublicCertData *string `json:"publicCertData" tf:"public_cert_data,omitempty"`
}

type VPNClientConfigurationObservation struct {
}

type VPNClientConfigurationParameters struct {

	// The client id of the Azure VPN application.
	// See Create an Active Directory (AD) tenant for P2S OpenVPN protocol connections for values
	// +kubebuilder:validation:Optional
	AADAudience *string `json:"aadAudience,omitempty" tf:"aad_audience,omitempty"`

	// The STS url for your tenant
	// +kubebuilder:validation:Optional
	AADIssuer *string `json:"aadIssuer,omitempty" tf:"aad_issuer,omitempty"`

	// AzureAD Tenant URL
	// +kubebuilder:validation:Optional
	AADTenant *string `json:"aadTenant,omitempty" tf:"aad_tenant,omitempty"`

	// The address space out of which IP addresses for
	// vpn clients will be taken. You can provide more than one address space, e.g.
	// in CIDR notation.
	// +kubebuilder:validation:Required
	AddressSpace []*string `json:"addressSpace" tf:"address_space,omitempty"`

	// The address of the Radius server.
	// +kubebuilder:validation:Optional
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty" tf:"radius_server_address,omitempty"`

	// The secret used by the Radius server.
	// +kubebuilder:validation:Optional
	RadiusServerSecret *string `json:"radiusServerSecret,omitempty" tf:"radius_server_secret,omitempty"`

	// One or more revoked_certificate blocks which
	// are defined below.
	// +kubebuilder:validation:Optional
	RevokedCertificate []RevokedCertificateParameters `json:"revokedCertificate,omitempty" tf:"revoked_certificate,omitempty"`

	// One or more root_certificate blocks which are
	// defined below. These root certificates are used to sign the client certificate
	// used by the VPN clients to connect to the gateway.
	// +kubebuilder:validation:Optional
	RootCertificate []RootCertificateParameters `json:"rootCertificate,omitempty" tf:"root_certificate,omitempty"`

	// List of the vpn authentication types for the virtual network gateway.
	// The supported values are AAD, Radius and Certificate.
	// +kubebuilder:validation:Optional
	VPNAuthTypes []*string `json:"vpnAuthTypes,omitempty" tf:"vpn_auth_types,omitempty"`

	// List of the protocols supported by the vpn client.
	// The supported values are SSTP, IkeV2 and OpenVPN.
	// Values SSTP and IkeV2 are incompatible with the use of
	// aad_tenant, aad_audience and aad_issuer.
	// +kubebuilder:validation:Optional
	VPNClientProtocols []*string `json:"vpnClientProtocols,omitempty" tf:"vpn_client_protocols,omitempty"`
}

type VirtualNetworkGatewayBGPSettingsObservation struct {

	// A list of peering_addresses as defined below. Only one peering_addresses block can be specified except when active_active of this Virtual Network Gateway is true.
	// +kubebuilder:validation:Optional
	PeeringAddresses []PeeringAddressesObservation `json:"peeringAddresses,omitempty" tf:"peering_addresses,omitempty"`
}

type VirtualNetworkGatewayBGPSettingsParameters struct {

	// The Autonomous System Number (ASN) to use as part of the BGP.
	// +kubebuilder:validation:Optional
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// The weight added to routes which have been learned
	// through BGP peering. Valid values can be between 0 and 100.
	// +kubebuilder:validation:Optional
	PeerWeight *float64 `json:"peerWeight,omitempty" tf:"peer_weight,omitempty"`

	// A list of peering_addresses as defined below. Only one peering_addresses block can be specified except when active_active of this Virtual Network Gateway is true.
	// +kubebuilder:validation:Optional
	PeeringAddresses []PeeringAddressesParameters `json:"peeringAddresses,omitempty" tf:"peering_addresses,omitempty"`
}

type VirtualNetworkGatewayIPConfigurationObservation struct {
}

type VirtualNetworkGatewayIPConfigurationParameters struct {

	// A user-defined name of the IP configuration. Defaults to
	// vnetGatewayConfig.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Defines how the private IP address
	// of the gateways virtual interface is assigned. Valid options are Static or
	// Dynamic. Defaults to Dynamic.
	// +kubebuilder:validation:Optional
	PrivateIPAddressAllocation *string `json:"privateIpAddressAllocation,omitempty" tf:"private_ip_address_allocation,omitempty"`

	// The ID of the public IP address to associate
	// with the Virtual Network Gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// The ID of the gateway subnet of a virtual network in
	// which the virtual network gateway will be created. It is mandatory that
	// the associated subnet is named GatewaySubnet. Therefore, each virtual
	// network can contain at most a single Virtual Network Gateway.
	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type VirtualNetworkGatewayObservation struct {

	// A bgp_settings block which is documented below. In this block the BGP specific settings can be defined.
	// +kubebuilder:validation:Optional
	BGPSettings []VirtualNetworkGatewayBGPSettingsObservation `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`

	// The ID of the Virtual Network Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VirtualNetworkGatewayParameters struct {

	// If true, an active-active Virtual Network Gateway
	// will be created. An active-active gateway requires a HighPerformance or an
	// UltraPerformance SKU. If false, an active-standby gateway will be created.
	// Defaults to false.
	// +kubebuilder:validation:Optional
	ActiveActive *bool `json:"activeActive,omitempty" tf:"active_active,omitempty"`

	// A bgp_settings block which is documented below. In this block the BGP specific settings can be defined.
	// +kubebuilder:validation:Optional
	BGPSettings []VirtualNetworkGatewayBGPSettingsParameters `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`

	// +kubebuilder:validation:Optional
	CustomRoute []CustomRouteParameters `json:"customRoute,omitempty" tf:"custom_route,omitempty"`

	// The ID of the local network gateway
	// through which outbound Internet traffic from the virtual network in which the
	// gateway is created will be routed (forced tunnelling). Refer to the
	// Azure documentation on forced tunnelling.
	// If not specified, forced tunnelling is disabled.
	// +kubebuilder:validation:Optional
	DefaultLocalNetworkGatewayID *string `json:"defaultLocalNetworkGatewayId,omitempty" tf:"default_local_network_gateway_id,omitempty"`

	// Specifies the Edge Zone within the Azure Region where this Virtual Network Gateway should exist. Changing this forces a new Virtual Network Gateway to be created.
	// +kubebuilder:validation:Optional
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// If true, BGP (Border Gateway Protocol) will be enabled
	// for this Virtual Network Gateway. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableBGP *bool `json:"enableBgp,omitempty" tf:"enable_bgp,omitempty"`

	// The Generation of the Virtual Network gateway. Possible values include Generation1, Generation2 or None.
	// +kubebuilder:validation:Optional
	Generation *string `json:"generation,omitempty" tf:"generation,omitempty"`

	// standby gateway requires exactly one ip_configuration block,
	// an active-active gateway requires exactly two ip_configuration blocks whereas
	// an active-active zone redundant gateway with P2S configuration requires exactly three ip_configuration blocks.
	// +kubebuilder:validation:Required
	IPConfiguration []VirtualNetworkGatewayIPConfigurationParameters `json:"ipConfiguration" tf:"ip_configuration,omitempty"`

	// The location/region where the Virtual Network Gateway is
	// located. Changing the location/region forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Should private IP be enabled on this gateway for connections? Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrivateIPAddressEnabled *bool `json:"privateIpAddressEnabled,omitempty" tf:"private_ip_address_enabled,omitempty"`

	// The name of the resource group in which to
	// create the Virtual Network Gateway. Changing the resource group name forces
	// a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Configuration of the size and capacity of the virtual network
	// gateway. Valid options are Basic, Standard, HighPerformance, UltraPerformance,
	// ErGw1AZ, ErGw2AZ, ErGw3AZ, VpnGw1, VpnGw2, VpnGw3, VpnGw4,VpnGw5, VpnGw1AZ,
	// VpnGw2AZ, VpnGw3AZ,VpnGw4AZ and VpnGw5AZ and depend on the type, vpn_type and
	// generation arguments.
	// A PolicyBased gateway only supports the Basic SKU. Further, the UltraPerformance
	// SKU is only supported by an ExpressRoute gateway.
	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the Virtual Network Gateway. Valid options are
	// Vpn or ExpressRoute. Changing the type forces a new resource to be created.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// to-site connections.
	// +kubebuilder:validation:Optional
	VPNClientConfiguration []VPNClientConfigurationParameters `json:"vpnClientConfiguration,omitempty" tf:"vpn_client_configuration,omitempty"`

	// The routing type of the Virtual Network Gateway. Valid
	// options are RouteBased or PolicyBased. Defaults to RouteBased.
	// +kubebuilder:validation:Optional
	VPNType *string `json:"vpnType,omitempty" tf:"vpn_type,omitempty"`
}

// VirtualNetworkGatewaySpec defines the desired state of VirtualNetworkGateway
type VirtualNetworkGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkGatewayParameters `json:"forProvider"`
}

// VirtualNetworkGatewayStatus defines the observed state of VirtualNetworkGateway.
type VirtualNetworkGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGateway is the Schema for the VirtualNetworkGateways API. Manages a virtual network gateway to establish secure, cross-premises connectivity.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetworkGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGatewaySpec   `json:"spec"`
	Status            VirtualNetworkGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGatewayList contains a list of VirtualNetworkGateways
type VirtualNetworkGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkGateway `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkGateway_Kind             = "VirtualNetworkGateway"
	VirtualNetworkGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualNetworkGateway_Kind}.String()
	VirtualNetworkGateway_KindAPIVersion   = VirtualNetworkGateway_Kind + "." + CRDGroupVersion.String()
	VirtualNetworkGateway_GroupVersionKind = CRDGroupVersion.WithKind(VirtualNetworkGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetworkGateway{}, &VirtualNetworkGatewayList{})
}