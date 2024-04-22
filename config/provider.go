/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/thibaultbustarret-ovhcloud/provider-netbox/config/dcim/device"
	"github.com/thibaultbustarret-ovhcloud/provider-netbox/config/dcim/deviceinterface"
	"github.com/thibaultbustarret-ovhcloud/provider-netbox/config/dcim/devicerole"
	"github.com/thibaultbustarret-ovhcloud/provider-netbox/config/dcim/devicetype"
	"github.com/thibaultbustarret-ovhcloud/provider-netbox/config/dcim/manufacturer"
	"github.com/thibaultbustarret-ovhcloud/provider-netbox/config/dcim/site"
	"github.com/thibaultbustarret-ovhcloud/provider-netbox/config/ipam/ipaddress"
)

const (
	resourcePrefix = "netbox"
	modulePath     = "github.com/thibaultbustarret-ovhcloud/provider-netbox"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("thibaultbustarret-ovhcloud.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		manufacturer.Configure,
		site.Configure,
		devicetype.Configure,
		devicerole.Configure,
		device.Configure,
		deviceinterface.Configure,
		ipaddress.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
