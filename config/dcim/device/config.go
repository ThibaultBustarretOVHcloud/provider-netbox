package device

import "github.com/crossplane/upjet/pkg/config"

// Configure configures how the provider interacts with the Netbox Terraform Provider for the Device resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "dcim"
		r.Kind = "Device"

		r.References["role_id"] = config.Reference{
			Type:      "DeviceRole",
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}

		r.References["device_type_id"] = config.Reference{
			Type:      "DeviceType",
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}

		r.References["site_id"] = config.Reference{
			Type:      "Site",
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}

		//TODO: Check if it's necessary
		if s, ok := r.TerraformResource.Schema["role_id"]; ok {
			s.Optional = true
			s.Computed = true
		}

		if s, ok := r.TerraformResource.Schema["device_type_id"]; ok {
			s.Optional = true
			s.Computed = true
		}

		if s, ok := r.TerraformResource.Schema["site_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
	})
}
