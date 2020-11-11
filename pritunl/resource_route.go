package pritunl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
	"sync"
)

// The Pritunl routes API cannot create multiple routes concurrently and this is why this
// mutex semaphore is created on create, update and delete operations
var mutex = &sync.Mutex{}

func ResourceRoute() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRouteCreate,
		Read:   ResourceRouteRead,
		Update: ResourceRouteUpdate,
		Delete: ResourceRouteDelete,

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
				Optional: true,
			},
			"nat": {
				Type:     schema.TypeBool,
				Optional: true,
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

	mutex.Lock()
	routeData, err := c.RouteCreate(serverId, r)
	mutex.Unlock()

	if err != nil {
		return err
	}

	d.SetId(routeData.Id)

	return ResourceRouteRead(d, m)
}

func ResourceRouteUpdate(d *schema.ResourceData, m interface{}) error {
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

	mutex.Lock()
	_, err := c.RouteUpdate(serverId, d.Id(), r)
	mutex.Unlock()

	if err != nil {
		return err
	}

	return ResourceRouteRead(d, m)
}

func ResourceRouteRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	serverId := d.Get("server").(string)
	data, err := c.RouteGet(serverId, d.Id())
	if err != nil {
		return err
	}

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

	serverId := d.Get("server").(string)

	mutex.Lock()
	err := c.RouteDelete(serverId, d.Id())
	mutex.Unlock()

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
