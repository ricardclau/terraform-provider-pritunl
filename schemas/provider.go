package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Provider struct {
	PritunlHost   string
	PritunlToken  string
	PritunlSecret string
}

func LoadProvider(d *schema.ResourceData) *Provider {
	return &Provider{
		PritunlHost:   d.Get("pritunl_host").(string),
		PritunlToken:  d.Get("pritunl_token").(string),
		PritunlSecret: d.Get("pritunl_secret").(string),
	}
}
