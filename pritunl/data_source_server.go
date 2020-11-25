package pritunl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
)

func DataSourceServer() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceServerRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"groups": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dns_servers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ipv6": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"otp_auth": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"wg": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func DataSourceServerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)
	serverName := d.Get("name").(string)

	data, err := c.ServerGetByName(serverName)
	if err != nil {
		return err
	}

	d.SetId(data.Id)
	d.Set("name", data.Name)
	d.Set("network", data.Network)
	d.Set("port", data.Port)
	d.Set("protocol", data.Protocol)
	d.Set("groups", data.Groups)
	d.Set("dns_servers", data.DNSServers)
	d.Set("ipv6", data.Ipv6)
	d.Set("otp_auth", data.OtpAuth)
	d.Set("wg", data.Wg)

	return nil
}
