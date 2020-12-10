package pritunl

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
	"strings"
)

func ResourceRoute() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRouteCreate,
		Read:   ResourceRouteRead,
		Update: ResourceRouteUpdate,
		Delete: ResourceRouteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"server": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"nat_interface": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat_netmap": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"advertise": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"metric": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"net_gateway": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func ResourceRouteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)
	serverId := d.Get("server").(string)
	r := client.RoutePostData{
		Network:      d.Get("network").(string),
		Comment:      d.Get("comment").(string),
		Metric:       d.Get("metric").(string),
		Nat:          d.Get("nat").(bool),
		NatInterface: d.Get("nat_interface").(string),
		NatNetmap:    d.Get("nat_netmap").(string),
		NetGateway:   d.Get("net_gateway").(bool),
		Advertise:    d.Get("advertise").(bool),
	}

	routeData, err := c.RouteCreate(serverId, r)

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s:%s", serverId, routeData.Id))

	return ResourceRouteRead(d, m)
}

func ResourceRouteUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	serverId, routeId, err := resourceRouteParseId(d.Id())
	if err != nil {
		return err
	}

	r := client.RoutePostData{
		Network:      d.Get("network").(string),
		Comment:      d.Get("comment").(string),
		Metric:       d.Get("metric").(string),
		Nat:          d.Get("nat").(bool),
		NatInterface: d.Get("nat_interface").(string),
		NatNetmap:    d.Get("nat_netmap").(string),
		NetGateway:   d.Get("net_gateway").(bool),
		Advertise:    d.Get("advertise").(bool),
	}

	_, err = c.RouteUpdate(serverId, routeId, r)
	if err != nil {
		return err
	}

	return ResourceRouteRead(d, m)
}

func ResourceRouteRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	serverId, routeId, err := resourceRouteParseId(d.Id())
	if err != nil {
		return err
	}

	data, err := c.RouteGet(serverId, routeId)
	if err != nil {
		return err
	}

	d.Set("server", serverId)
	d.Set("network", data.Network)
	d.Set("comment", data.Comment)
	d.Set("nat", data.Nat)
	d.Set("nat_interface", data.NatInterface)
	d.Set("nat_netmap", data.NatNetmap)
	d.Set("advertise", data.Advertise)
	d.Set("net_gateway", data.NetGateway)
	d.Set("metric", data.Metric)

	return nil
}

func ResourceRouteDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	serverId, routeId, err := resourceRouteParseId(d.Id())
	if err != nil {
		return err
	}

	err = c.RouteDelete(serverId, routeId)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func resourceRouteParseId(id string) (serverId, routeId string, err error) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 {
		err = fmt.Errorf("user id must be of the form <org_id>:<user_id>")
		return
	}

	serverId = parts[0]
	routeId = parts[1]
	return
}
