package devicerole

import "github.com/crossplane/upjet/pkg/config"

// Configure configures how the provider interacts with the Netbox Terraform Provider for the DeviceRole resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device_role", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "dcim"
		r.Kind = "DeviceRole"
	})
}
