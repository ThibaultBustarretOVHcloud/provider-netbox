package ipaddress

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_ip_address", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "ipam"
		r.Kind = "IPAddress"

		r.References["device_interface_id"] = config.Reference{
			Type:      "DeviceInterface",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
