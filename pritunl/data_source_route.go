package pritunl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
)

func DataSourceRoute() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceRouteRead,
		Schema: map[string]*schema.Schema{
			"server": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nat": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"nat_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nat_netmap": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"advertise": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"metric": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"net_gateway": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func DataSourceRouteRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	network := d.Get("network").(string)
	serverId := d.Get("server").(string)

	data, err := c.RouteGetByNetwork(serverId, network)
	if err != nil {
		return err
	}

	d.SetId(data.Id)
	d.Set("comment", data.Comment)
	d.Set("nat", data.Nat)
	d.Set("nat_interface", data.NatInterface)
	d.Set("nat_netmap", data.NatNetmap)
	d.Set("advertise", data.Advertise)
	d.Set("net_gateway", data.NetGateway)
	d.Set("metric", data.Metric)

	return nil
}
