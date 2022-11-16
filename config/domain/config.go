package domain

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_domain", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		// We need to override the default group that upjet generated for
		// this resource, which would be "linode"
		// r.ShortGroup = "domain"
	})
}
