package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/kihahu/terraform-provider-pritunl/resources"
	"github.com/kihahu/terraform-provider-pritunl/schemas"
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
			"pritunl_organization":  resources.Organization(),
			"pritunl_user":          resources.User(),
			"pritunl_link":          resources.Link(),
			"pritunl_link_location": resources.LinkLocation(),
			"pritunl_link_host":     resources.LinkHost(),
			"pritunl_server":        resources.Server(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return schemas.LoadProvider(d), nil
}
