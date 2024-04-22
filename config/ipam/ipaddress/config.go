package ipaddress

import "github.com/crossplane/upjet/pkg/config"

// Configure configures how the provider interacts with the Netbox Terraform Provider for the IPAddress resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_ip_address", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "ipam"
		r.Kind = "IPAddress"

		r.References["device_interface_id"] = config.Reference{
			Type:      "github.com/thibaultbustarret-ovhcloud/provider-netbox/apis/dcim/v1alpha1.DeviceInterface",
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
