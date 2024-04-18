/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Subset of resources used to test the provider.
	"netbox_ip_address":       config.IdentifierFromProvider,
	"netbox_device_interface": config.NameAsIdentifier,
	"netbox_device":           config.NameAsIdentifier,
	"netbox_device_role":      config.NameAsIdentifier,
	"netbox_site":             config.NameAsIdentifier,
	"netbox_device_type":      config.IdentifierFromProvider,
	"netbox_manufacturer":     config.NameAsIdentifier,

	/* TODO: Implement all resources.
	This is the complete list of resources for provider v3.8.5 (w/o datasources)
	"netbox_aggregate":                  config.IdentifierFromProvider,
	"netbox_asn":                        config.IdentifierFromProvider,
	"netbox_available_ip_address":       config.IdentifierFromProvider,
	"netbox_available_prefix":           config.IdentifierFromProvider,
	"netbox_cable":                      config.IdentifierFromProvider,
	"netbox_circuit_provider":           config.NameAsIdentifier,
	"netbox_circuit_termination":        config.IdentifierFromProvider,
	"netbox_circuit_type":               config.NameAsIdentifier,
	"netbox_circuit":                    config.IdentifierFromProvider,
	"netbox_cluster_group":              config.NameAsIdentifier,
	"netbox_cluster_type":               config.NameAsIdentifier,
	"netbox_cluster":                    config.NameAsIdentifier,
	"netbox_contact_assignment":         config.NameAsIdentifier,
	"netbox_contact_group":              config.NameAsIdentifier,
	"netbox_contact_role":               config.NameAsIdentifier,
	"netbox_contact":                    config.NameAsIdentifier,
	"netbox_custom_field_choice_set":    config.NameAsIdentifier,
	"netbox_custom_field":               config.NameAsIdentifier,
	"netbox_device_console_port":        config.NameAsIdentifier,
	"netbox_device_console_server_port": config.NameAsIdentifier,
	"netbox_device_front_port":          config.NameAsIdentifier,
	"netbox_device_interface":           config.NameAsIdentifier,
	"netbox_device_module_bay":          config.NameAsIdentifier,
	"netbox_device_power_outlet":        config.NameAsIdentifier,
	"netbox_device_power_port":          config.NameAsIdentifier,
	"netbox_device_primary_ip":          config.IdentifierFromProvider,
	"netbox_device_rear_port":           config.NameAsIdentifier,
	"netbox_device_role":                config.NameAsIdentifier,
	"netbox_device_type":                config.IdentifierFromProvider,
	"netbox_device":                     config.NameAsIdentifier,
	"netbox_event_rule":                 config.NameAsIdentifier,
	"netbox_interface":                  config.NameAsIdentifier,
	"netbox_inventory_item":             config.NameAsIdentifier,
	"netbox_inventory_item_role":        config.NameAsIdentifier,
	"netbox_ip_address":                 config.IdentifierFromProvider,
	"netbox_ip_range":                   config.IdentifierFromProvider,
	"netbox_ipam_role":                  config.NameAsIdentifier,
	"netbox_location":                   config.NameAsIdentifier,
	"netbox_manufacturer":               config.NameAsIdentifier,
	"netbox_module":                     config.IdentifierFromProvider,
	"netbox_module_type":                config.IdentifierFromProvider,
	"netbox_permission":                 config.NameAsIdentifier,
	"netbox_platform":                   config.NameAsIdentifier,
	"netbox_power_feed":                 config.NameAsIdentifier,
	"netbox_power_panel":                config.NameAsIdentifier,
	"netbox_prefix":                     config.IdentifierFromProvider,
	"netbox_primary_ip":                 config.IdentifierFromProvider,
	"netbox_rack_reservation":           config.IdentifierFromProvider,
	"netbox_rack_role":                  config.NameAsIdentifier,
	"netbox_rack":                       config.NameAsIdentifier,
	"netbox_region":                     config.NameAsIdentifier,
	"netbox_rir":                        config.NameAsIdentifier,
	"netbox_route_target":               config.NameAsIdentifier,
	"netbox_service":                    config.NameAsIdentifier,
	"netbox_site_group":                 config.NameAsIdentifier,
	"netbox_site":                       config.NameAsIdentifier,
	"netbox_tag":                        config.NameAsIdentifier,
	"netbox_tenant_group":               config.NameAsIdentifier,
	"netbox_tenant":                     config.NameAsIdentifier,
	"netbox_token":                      config.IdentifierFromProvider,
	"netbox_user":                       config.IdentifierFromProvider,
	"netbox_virtual_chassis":            config.NameAsIdentifier,
	"netbox_virtual_disk":               config.NameAsIdentifier,
	"netbox_virtual_machine":            config.NameAsIdentifier,
	"netbox_vlan_group":                 config.NameAsIdentifier,
	"netbox_vlan":                       config.NameAsIdentifier,
	"netbox_vpn_tunnel":                 config.NameAsIdentifier,
	"netbox_vpn_tunnel_group":           config.NameAsIdentifier,
	"netbox_vpn_tunnel_termination":     config.IdentifierFromProvider,
	"netbox_vrf":                        config.NameAsIdentifier,
	"netbox_webhook":                    config.NameAsIdentifier,*/
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
