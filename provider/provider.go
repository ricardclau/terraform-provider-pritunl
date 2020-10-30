package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/resources"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,
		Schema: map[string]*schema.Schema{
			"pritunl_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pritunl_token": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pritunl_secret": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"pritunl_organization":             resources.Organization(),
			"pritunl_user":                     resources.User(),
			"pritunl_link":                     resources.Link(),
			"pritunl_link_location":            resources.LinkLocation(),
			"pritunl_link_host":                resources.LinkHost(),
			"pritunl_server":                   resources.Server(),
			"pritunl_route":                    resources.Route(),
			"pritunl_link_server_organization": resources.LinkServerOrganization(),
			"pritunl_settings":                 resources.Settings(),
			"pritunl_server_start":             resources.ServerStart(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return schemas.LoadProvider(d), nil
}
