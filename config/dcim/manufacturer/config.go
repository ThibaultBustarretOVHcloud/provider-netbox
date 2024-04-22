package manufacturer

import "github.com/crossplane/upjet/pkg/config"

// Configure configures how the provider interacts with the Netbox Terraform Provider for the Manufacturer resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_manufacturer", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "dcim"
		r.Kind = "Manufacturer"
	})
}
