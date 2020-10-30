package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Route struct {
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

func LoadRoute(d *schema.ResourceData) (sch *Route) {
	sch = &Route{
		Id:      d.Id(),
		Server:  d.Get("server").(string),
		Network: d.Get("network").(string),
	}

	return
}
