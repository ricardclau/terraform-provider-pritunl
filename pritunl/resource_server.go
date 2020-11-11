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
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dns_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceServerCreate(d *schema.ResourceData, m interface{}) error {
	//c := m.(*client.PritunlClient)

	return ResourceServerRead(d, m)
}

func ResourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	//c := m.(*client.PritunlClient)

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
	d.Set("groups", data.Groups)
	d.Set("dns_servers", data.DNSServers)

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
