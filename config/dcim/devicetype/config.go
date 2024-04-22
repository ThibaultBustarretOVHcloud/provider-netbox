package devicetype

import "github.com/crossplane/upjet/pkg/config"

// Configure configures how the provider interacts with the Netbox Terraform Provider for the DeviceType resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device_type", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "dcim"
		r.Kind = "DeviceType"
		r.References["manufacturer_id"] = config.Reference{
			Type:      "Manufacturer",
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
