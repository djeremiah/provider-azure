/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta12 "github.com/upbound/provider-azure/apis/keyvault/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/network/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AccessConnector.
func (mg *AccessConnector) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Workspace.
func (mg *Workspace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CustomParameters); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomParameters[i3].PrivateSubnetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CustomParameters[i3].PrivateSubnetNameRef,
			Selector:     mg.Spec.ForProvider.CustomParameters[i3].PrivateSubnetNameSelector,
			To: reference.To{
				List:    &v1beta11.SubnetList{},
				Managed: &v1beta11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomParameters[i3].PrivateSubnetName")
		}
		mg.Spec.ForProvider.CustomParameters[i3].PrivateSubnetName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CustomParameters[i3].PrivateSubnetNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.CustomParameters); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomParameters[i3].PublicSubnetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CustomParameters[i3].PublicSubnetNameRef,
			Selector:     mg.Spec.ForProvider.CustomParameters[i3].PublicSubnetNameSelector,
			To: reference.To{
				List:    &v1beta11.SubnetList{},
				Managed: &v1beta11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomParameters[i3].PublicSubnetName")
		}
		mg.Spec.ForProvider.CustomParameters[i3].PublicSubnetName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CustomParameters[i3].PublicSubnetNameRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ManagedResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ManagedResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ManagedResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ManagedResourceGroupName")
	}
	mg.Spec.ForProvider.ManagedResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ManagedResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WorkspaceRootDbfsCustomerManagedKey.
func (mg *WorkspaceRootDbfsCustomerManagedKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyVaultKeyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.KeyVaultKeyIDRef,
		Selector:     mg.Spec.ForProvider.KeyVaultKeyIDSelector,
		To: reference.To{
			List:    &v1beta12.KeyList{},
			Managed: &v1beta12.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyVaultKeyID")
	}
	mg.Spec.ForProvider.KeyVaultKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyVaultKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference

	return nil
}
