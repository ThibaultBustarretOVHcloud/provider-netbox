// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	device "github.com/ThibaultBustarretOVHcloud/provider-netbox/internal/controller/dcim/device"
	deviceinterface "github.com/ThibaultBustarretOVHcloud/provider-netbox/internal/controller/dcim/deviceinterface"
	devicerole "github.com/ThibaultBustarretOVHcloud/provider-netbox/internal/controller/dcim/devicerole"
	devicetype "github.com/ThibaultBustarretOVHcloud/provider-netbox/internal/controller/dcim/devicetype"
	manufacturer "github.com/ThibaultBustarretOVHcloud/provider-netbox/internal/controller/dcim/manufacturer"
	site "github.com/ThibaultBustarretOVHcloud/provider-netbox/internal/controller/dcim/site"
	ipaddress "github.com/ThibaultBustarretOVHcloud/provider-netbox/internal/controller/ipam/ipaddress"
	providerconfig "github.com/ThibaultBustarretOVHcloud/provider-netbox/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		device.Setup,
		deviceinterface.Setup,
		devicerole.Setup,
		devicetype.Setup,
		manufacturer.Setup,
		site.Setup,
		ipaddress.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
