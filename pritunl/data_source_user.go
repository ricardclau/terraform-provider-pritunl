package pritunl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
)

func DataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"groups": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pin": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"network_links": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"bypass_secondary": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"client_to_client": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dns_servers": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dns_suffix": {
				Type:     schema.TypeString,
				Computed: true,
			}},
	}
}

func DataSourceUserRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	organizationId := d.Get("organization_id").(string)
	name := d.Get("name").(string)

	data, err := c.UserGetByName(organizationId, name)
	if err != nil {
		return err
	}

	d.Set("name", data.Name)
	d.Set("auth_type", data.AuthType)
	d.Set("email", data.Email)
	d.Set("disabled", data.Disabled)
	d.Set("bypass_secondary", data.BypassSecondary)
	d.Set("client_to_client", data.ClientToClient)
	d.Set("groups", data.Groups)
	d.Set("network_links", data.NetworkLinks)
	d.Set("dns_servers", data.DnsServers)
	d.Set("dns_suffix", data.DnsSuffix)

	return nil
}
