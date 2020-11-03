package resources

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/request"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
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

func routeGet(prvdr *schemas.Provider, sch *schemas.Route) (*routeData, error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/server/%s/route/%s", sch.Server, sch.Id),
	}

	data := &routeData{}
	_, err := req.Do(prvdr, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func routePut(prvdr *schemas.Provider, sch *schemas.Route) (*routeData, error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/server/%s/route/%s", sch.Server, sch.Id),
		Json: &routeData{
			Server:  sch.Server,
			Network: sch.Network,
			Comment: sch.Comment,
			Id:      sch.Id,
		},
	}

	data := &routeData{}

	_, err := req.Do(prvdr, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func routePost(prvdr *schemas.Provider, sch *schemas.Route) (*routeData, error) {

	req := request.Request{
		Method: "POST",
		Path:   fmt.Sprintf("/server/%s/route", sch.Server),
		Json: &routeData{
			Server:  sch.Server,
			Network: sch.Network,
			Comment: sch.Comment,
			Id:      sch.Id,
		},
	}

	data := &routeData{}

	_, err := req.Do(prvdr, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func routeDel(prvdr *schemas.Provider, sch *schemas.Route) error {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/server/%s/route/%s", sch.Server, sch.Id),
	}

	_, err := req.Do(prvdr, nil)

	if err != nil {
		return err
	}

	return nil
}

func routeCreate(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadRoute(d)

	data, err := routeGet(prvdr, sch)
	if err != nil {
		return err
	}

	// if data != nil {
	// 	sch.Id = data.Id

	// 	data, err = routePut(prvdr, sch)
	// 	if err != nil {
	// 		return
	// 	}
	// }

	// if data == nil {
	data, err = routePost(prvdr, sch)
	if err != nil {
		return err
	}
	// }

	// d.SetId(data.Id)
	d.Set("server", data.Server)
	d.SetId(data.Network)

	return nil
}

func routeUpdate(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadRoute(d)

	data, err := routePost(prvdr, sch)
	if err != nil {
		return err
	}

	if data == nil {
		return fmt.Errorf("Cannot update Route with id: %v", sch.Network)
	}

	// d.SetId(data.Id)
	d.Set("server", data.Server)
	d.SetId(data.Network)

	return nil
}

func routeRead(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadRoute(d)

	data, err := routeGet(prvdr, sch)
	if err != nil {
		return err
	}

	if data == nil {
		return fmt.Errorf("Cannot read Route with id: %v", sch.Network)
	}

	d.Set("server", data.Server)
	d.SetId(data.Network)
	// d.SetId(data.Id)

	return nil
}

func routeDelete(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadRoute(d)

	err := routeDel(prvdr, sch)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
