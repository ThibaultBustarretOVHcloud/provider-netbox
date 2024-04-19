/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Device.
func (mg *Device) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.DeviceTypeID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DeviceTypeIDRef,
		Selector:     mg.Spec.ForProvider.DeviceTypeIDSelector,
		To: reference.To{
			List:    &DeviceTypeList{},
			Managed: &DeviceType{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeviceTypeID")
	}
	mg.Spec.ForProvider.DeviceTypeID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeviceTypeIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.RoleID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RoleIDRef,
		Selector:     mg.Spec.ForProvider.RoleIDSelector,
		To: reference.To{
			List:    &DeviceRoleList{},
			Managed: &DeviceRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleID")
	}
	mg.Spec.ForProvider.RoleID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.SiteID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SiteIDRef,
		Selector:     mg.Spec.ForProvider.SiteIDSelector,
		To: reference.To{
			List:    &SiteList{},
			Managed: &Site{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SiteID")
	}
	mg.Spec.ForProvider.SiteID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SiteIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DeviceInterface.
func (mg *DeviceInterface) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.DeviceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DeviceIDRef,
		Selector:     mg.Spec.ForProvider.DeviceIDSelector,
		To: reference.To{
			List:    &DeviceList{},
			Managed: &Device{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeviceID")
	}
	mg.Spec.ForProvider.DeviceID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeviceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DeviceType.
func (mg *DeviceType) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.ManufacturerID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ManufacturerIDRef,
		Selector:     mg.Spec.ForProvider.ManufacturerIDSelector,
		To: reference.To{
			List:    &ManufacturerList{},
			Managed: &Manufacturer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ManufacturerID")
	}
	mg.Spec.ForProvider.ManufacturerID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ManufacturerIDRef = rsp.ResolvedReference

	return nil
}