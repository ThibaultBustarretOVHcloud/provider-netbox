package deviceinterface

import "github.com/crossplane/upjet/pkg/config"

// Configure configures how the provider interacts with the Netbox Terraform Provider for the DeviceInterface resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device_interface", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "dcim"
		r.Kind = "DeviceInterface"
		r.References["device_id"] = config.Reference{
			Type:      "Device",
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()", // Needed ?
		}
	})
}
