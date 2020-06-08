package resources

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/kihahu/terraform-provider-pritunl/request"
	"github.com/kihahu/terraform-provider-pritunl/schemas"
)

func ServerStart() *schema.Resource {
	return &schema.Resource{
		Create: serverStartCreate,
		Read:   serverStartRead,
		Update: serverStartUpdate,
		Delete: serverStartDelete,

		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type serverStartData struct {
	PortWg           interface{} `json:"port_wg"`
	DNSServers       []string    `json:"dns_servers"`
	Protocol         string      `json:"protocol"`
	MaxDevices       int         `json:"max_devices"`
	MaxClients       int         `json:"max_clients"`
	LinkPingTimeout  int         `json:"link_ping_timeout"`
	PingTimeout      int         `json:"ping_timeout"`
	Ipv6             bool        `json:"ipv6"`
	Vxlan            bool        `json:"vxlan"`
	NetworkMode      string      `json:"network_mode"`
	BindAddress      string      `json:"bind_address"`
	BlockOutsideDNS  bool        `json:"block_outside_dns"`
	NetworkStart     string      `json:"network_start"`
	Name             string      `json:"name"`
	PingInterval     int         `json:"ping_interval"`
	AllowedDevices   interface{} `json:"allowed_devices"`
	UsersOnline      int         `json:"users_online"`
	Ipv6Firewall     bool        `json:"ipv6_firewall"`
	SessionTimeout   interface{} `json:"session_timeout"`
	OtpAuth          bool        `json:"otp_auth"`
	MultiDevice      bool        `json:"multi_device"`
	SearchDomain     interface{} `json:"search_domain"`
	LzoCompression   bool        `json:"lzo_compression"`
	PreConnectMsg    interface{} `json:"pre_connect_msg"`
	InactiveTimeout  interface{} `json:"inactive_timeout"`
	LinkPingInterval int         `json:"link_ping_interval"`
	Id               string      `json:"id"`
	PingTimeoutWg    int         `json:"ping_timeout_wg"`
	Uptime           interface{} `json:"uptime"`
	NetworkEnd       string      `json:"network_end"`
	Network          string      `json:"network"`
	DhParamBits      int         `json:"dh_param_bits"`
	Wg               bool        `json:"wg"`
	Port             int         `json:"port"`
	DevicesOnline    int         `json:"devices_online"`
	NetworkWg        interface{} `json:"network_wg"`
	Status           string      `json:"status"`
	DNSMapping       bool        `json:"dns_mapping"`
	Hash             string      `json:"hash"`
	Debug            bool        `json:"debug"`
	RestrictRoutes   bool        `json:"restrict_routes"`
	UserCount        int         `json:"user_count"`
	Groups           []string    `json:"groups"`
	InterClient      bool        `json:"inter_client"`
	ReplicaCount     int         `json:"replica_count"`
	Cipher           string      `json:"cipher"`
	MssFix           interface{} `json:"mss_fix"`
	JumboFrames      bool        `json:"jumbo_frames"`
}

func serverStartGet(prvdr *schemas.Provider, sch *schemas.Server) (
	data *serverStartData, err error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/server/%s", sch.Id),
	}

	xdata := &serverStartData{}
	resp, err := req.Do(prvdr, xdata)
	if err != nil {
		return
	}

	if resp.StatusCode == 400 || resp.StatusCode == 400 {
		data = nil
	}

	return xdata, err
}

func serverStartPut(prvdr *schemas.Provider, sch *schemas.Server) (
	data *serverStartData, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/server/%s/operation/start", sch.Id),
		Json: &serverData{
			Id: sch.Id,
		},
	}

	data = &serverStartData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func serverStopPut(prvdr *schemas.Provider, sch *schemas.Server) (
	data *serverStartData, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/server/%s/operation/stop", sch.Id),
		Json: &serverData{
			Id: sch.Id,
		},
	}

	data = &serverStartData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func serverStartCreate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServerStart(d)

	data, err := serverStartGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.Id = data.Id

		data, err = serverStartPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	d.SetId(data.Id)

	return
}

func serverStartUpdate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServerStart(d)

	data, err := serverStartPut(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		return
	}

	d.SetId(data.Id)

	return
}

func serverStartRead(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServerStart(d)

	data, err := serverStartGet(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		return
	}

	d.SetId(data.Id)

	return
}

func serverStartDelete(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServerStart(d)

	data, err := serverStartGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.Id = data.Id

		data, err = serverStopPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	return
}
