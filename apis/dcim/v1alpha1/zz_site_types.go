// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SiteInitParameters struct {

	// (Set of Number)
	AsnIds []*float64 `json:"asnIds,omitempty" tf:"asn_ids,omitempty"`

	// (Map of String)
	CustomFields map[string]*string `json:"customFields,omitempty" tf:"custom_fields,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Facility *string `json:"facility,omitempty" tf:"facility,omitempty"`

	// (Number)
	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// (Number)
	Latitude *float64 `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// (Number)
	Longitude *float64 `json:"longitude,omitempty" tf:"longitude,omitempty"`

	// (String)
	PhysicalAddress *string `json:"physicalAddress,omitempty" tf:"physical_address,omitempty"`

	// (Number)
	RegionID *float64 `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// (String)
	ShippingAddress *string `json:"shippingAddress,omitempty" tf:"shipping_address,omitempty"`

	// (String)
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (String) Valid values are planned, staging, active, decommissioning and retired. Defaults to active.
	// Valid values are `planned`, `staging`, `active`, `decommissioning` and `retired`. Defaults to `active`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Set of String)
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Number)
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// (String)
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type SiteObservation struct {

	// (Set of Number)
	AsnIds []*float64 `json:"asnIds,omitempty" tf:"asn_ids,omitempty"`

	// (Map of String)
	CustomFields map[string]*string `json:"customFields,omitempty" tf:"custom_fields,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Facility *string `json:"facility,omitempty" tf:"facility,omitempty"`

	// (Number)
	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number)
	Latitude *float64 `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// (Number)
	Longitude *float64 `json:"longitude,omitempty" tf:"longitude,omitempty"`

	// (String)
	PhysicalAddress *string `json:"physicalAddress,omitempty" tf:"physical_address,omitempty"`

	// (Number)
	RegionID *float64 `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// (String)
	ShippingAddress *string `json:"shippingAddress,omitempty" tf:"shipping_address,omitempty"`

	// (String)
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (String) Valid values are planned, staging, active, decommissioning and retired. Defaults to active.
	// Valid values are `planned`, `staging`, `active`, `decommissioning` and `retired`. Defaults to `active`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Set of String)
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Number)
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// (String)
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type SiteParameters struct {

	// (Set of Number)
	// +kubebuilder:validation:Optional
	AsnIds []*float64 `json:"asnIds,omitempty" tf:"asn_ids,omitempty"`

	// (Map of String)
	// +kubebuilder:validation:Optional
	CustomFields map[string]*string `json:"customFields,omitempty" tf:"custom_fields,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Facility *string `json:"facility,omitempty" tf:"facility,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	Latitude *float64 `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	Longitude *float64 `json:"longitude,omitempty" tf:"longitude,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PhysicalAddress *string `json:"physicalAddress,omitempty" tf:"physical_address,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	RegionID *float64 `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ShippingAddress *string `json:"shippingAddress,omitempty" tf:"shipping_address,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (String) Valid values are planned, staging, active, decommissioning and retired. Defaults to active.
	// Valid values are `planned`, `staging`, `active`, `decommissioning` and `retired`. Defaults to `active`.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Set of String)
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

// SiteSpec defines the desired state of Site
type SiteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SiteInitParameters `json:"initProvider,omitempty"`
}

// SiteStatus defines the observed state of Site.
type SiteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Site is the Schema for the Sites API. From the official documentation https://docs.netbox.dev/en/stable/features/sites-and-racks/#sites: How you choose to employ sites when modeling your network may vary depending on the nature of your organization, but generally a site will equate to a building or campus. For example, a chain of banks might create a site to represent each of its branches, a site for its corporate headquarters, and two additional sites for its presence in two colocation facilities. Each site must be assigned a unique name and may optionally be assigned to a region and/or tenant.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type Site struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SiteSpec   `json:"spec"`
	Status            SiteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteList contains a list of Sites
type SiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Site `json:"items"`
}

// Repository type metadata.
var (
	Site_Kind             = "Site"
	Site_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Site_Kind}.String()
	Site_KindAPIVersion   = Site_Kind + "." + CRDGroupVersion.String()
	Site_GroupVersionKind = CRDGroupVersion.WithKind(Site_Kind)
)

func init() {
	SchemeBuilder.Register(&Site{}, &SiteList{})
}
