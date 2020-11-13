package pritunl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/pritunl/terraform-provider-pritunl/client"
	"net/http"
	"time"
)

var defaultClient = &http.Client{
	Timeout: 10 * time.Second,
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
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
		DataSourcesMap: map[string]*schema.Resource{
			"pritunl_organization": DataSourceOrganization(),
			"pritunl_user":         DataSourceUser(),
			"pritunl_route":        DataSourceRoute(),
			"pritunl_server":       DataSourceServer(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"pritunl_organization": ResourceOrganization(),
			"pritunl_user":         ResourceUser(),
			"pritunl_route":        ResourceRoute(),
			"pritunl_server":       ResourceServer(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return client.NewPritunlClient(
		d.Get("pritunl_host").(string),
		d.Get("pritunl_token").(string),
		d.Get("pritunl_secret").(string),
		defaultClient,
	), nil
}
