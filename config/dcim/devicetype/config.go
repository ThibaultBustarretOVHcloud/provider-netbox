package devicetype

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device_type", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "dcim"
		r.Kind = "DeviceType"
		r.References["manufacturer_id"] = config.Reference{
			Type:      "Manufacturer",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
