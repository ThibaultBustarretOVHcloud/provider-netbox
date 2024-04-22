package site

import "github.com/crossplane/upjet/pkg/config"

// Configure configures how the provider interacts with the Netbox Terraform Provider for the Site resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_site", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "dcim"
		r.Kind = "Site"
	})
}
