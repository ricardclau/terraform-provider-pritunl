package resources

import (
	"fmt"

	"github.com/dropbox/godropbox/errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/kihahu/terraform-provider-pritunl/errortypes"
	"github.com/kihahu/terraform-provider-pritunl/request"
	"github.com/kihahu/terraform-provider-pritunl/schemas"
)

func Route() *schema.Resource {
	return &schema.Resource{
		Create: routeCreate,
		Read:   routeRead,
		Update: routeUpdate,
		Delete: routeDelete,

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
		},
	}
}

type routeData struct {
	Id             string      `json:"id"`
	Server         string      `json:"server"`
	Network        string      `json:"network"`
	Comment        string      `json:"comment"`
	Metric         interface{} `json:"metric"`
	VirtualNetwork interface{} `json:"virtual_network"`
	WgNetwork      interface{} `json:"wg_network"`
	NetworkLink    interface{} `json:"network_link"`
	ServerLink     interface{} `json:"server_link"`
	Nat            bool        `json:"nat"`
	NatInterface   string      `json:"nat_interface"`
	NatNetmap      string      `json:"nat_netmap"`
	NetGateway     bool        `json:"net_gateway"`
	Advertise      bool        `json:"advertise"`
	VpcRegion      interface{} `json:"vpc_region"`
	VpcID          interface{} `json:"vpc_id"`
}

func routeGet(prvdr *schemas.Provider, sch *schemas.Route) (
	data *routeData, err error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/server/%s/route/%s", sch.Server, sch.Id),
	}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode < 405 {
		data = nil
	}

	return
}

func routePut(prvdr *schemas.Provider, sch *schemas.Route) (
	data *routeData, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/server/%s/route/%s", sch.Server, sch.Id),
		Json: &routeData{
			Server:  sch.Server,
			Network: sch.Network,
			Comment: sch.Comment,
		},
	}

	data = &routeData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func routePost(prvdr *schemas.Provider, sch *schemas.Route) (
	data *routeData, err error) {

	req := request.Request{
		Method: "POST",
		Path:   fmt.Sprintf("/server/%s/route", sch.Server),
		Json: &routeData{
			Server:  sch.Server,
			Network: sch.Network,
			Comment: sch.Comment,
		},
	}

	data = &routeData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		err = &errortypes.RequestError{
			errors.New("server: Not found on post"),
		}
		return
	}

	return
}

func routeDel(prvdr *schemas.Provider, sch *schemas.Route) (
	err error) {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/server/%s/route/%s", sch.Server, sch.Id),
	}

	_, err = req.Do(prvdr, nil)

	if err != nil {
		return
	}

	return
}

func routeCreate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadRoute(d)

	data, err := routeGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.Id = data.Id

		data, err = routePut(prvdr, sch)
		if err != nil {
			return
		}
	}

	if data == nil {
		data, err = routePost(prvdr, sch)
		if err != nil {
			return
		}
	}

	d.SetId(data.Id)

	return
}

func routeUpdate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadRoute(d)

	data, err := routePut(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		// d.SetId("")
		return
	}

	d.SetId(data.Id)

	return
}

func routeRead(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadRoute(d)

	data, err := routeGet(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		return
	}

	d.Set("server", data.Server)
	d.Set("network", data.Network)
	d.SetId(data.Id)

	return
}

func routeDelete(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadRoute(d)

	err = routeDel(prvdr, sch)
	if err != nil {
		return
	}

	d.SetId("")

	return
}
