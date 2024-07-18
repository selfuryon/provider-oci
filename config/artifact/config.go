package artifact

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_artifacts_generic_artifact", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
}
