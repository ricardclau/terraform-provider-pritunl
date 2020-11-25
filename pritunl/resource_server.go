package pritunl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
)

func ResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServerCreate,
		Read:   ResourceServerRead,
		Update: ResourceServerUpdate,
		Delete: ResourceServerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
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
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
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

func ResourceServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	s := client.ServerPostData{
		Name: d.Get("name").(string),
		// This is autopopulated in the UI ????
		Network:    d.Get("network").(string),
		DNSServers: expandStringListFromSetSchema(d.Get("dns_servers").(*schema.Set)),
		Port:       d.Get("port").(int),
		Protocol:   d.Get("protocol").(string),
		Groups:     expandStringListFromSetSchema(d.Get("groups").(*schema.Set)),
		Ipv6:       d.Get("ipv6").(bool),
		OtpAuth:    d.Get("otp_auth").(bool),
		Wg:         d.Get("wg").(bool),
	}

	c.ServerCreate(s)

	return ResourceServerRead(d, m)
}

func ResourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	s := client.ServerPostData{
		Name: d.Get("name").(string),
		// This is autopopulated in the UI ????
		Network:    d.Get("network").(string),
		DNSServers: expandStringListFromSetSchema(d.Get("dns_servers").(*schema.Set)),
		Port:       d.Get("port").(int),
		Protocol:   d.Get("protocol").(string),
		Groups:     expandStringListFromSetSchema(d.Get("groups").(*schema.Set)),
		Ipv6:       d.Get("ipv6").(bool),
		OtpAuth:    d.Get("otp_auth").(bool),
		Wg:         d.Get("wg").(bool),
	}

	c.ServerUpdate(d.Id(), s)

	return ResourceServerRead(d, m)
}

func ResourceServerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	data, err := c.ServerGet(d.Id())
	if err != nil {
		return err
	}

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

func ResourceServerDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	err := c.ServerDelete(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
