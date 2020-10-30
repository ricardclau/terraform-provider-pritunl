package resources

import (
	"fmt"

	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/request"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
)

func Server() *schema.Resource {
	return &schema.Resource{
		Create: serverCreate,
		Read:   serverRead,
		Update: serverUpdate,
		Delete: serverDelete,

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

type serverData struct {
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

func serverGet(prvdr *schemas.Provider, sch *schemas.Server) (
	data *serverData, err error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/server/%s", sch.Id),
	}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 400 || resp.StatusCode == 400 {
		data = nil
	}

	return
}

func serverPut(prvdr *schemas.Provider, sch *schemas.Server) (
	data *serverData, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/server/%s", sch.Id),
		Json: &serverData{
			PortWg:           sch.PortWg,
			DNSServers:       sch.DNSServers,
			Protocol:         sch.Protocol,
			MaxDevices:       sch.MaxDevices,
			MaxClients:       sch.MaxClients,
			LinkPingTimeout:  sch.LinkPingTimeout,
			PingTimeout:      sch.PingTimeout,
			Ipv6:             sch.Ipv6,
			Vxlan:            sch.Vxlan,
			NetworkMode:      sch.NetworkMode,
			BindAddress:      sch.BindAddress,
			BlockOutsideDNS:  sch.BlockOutsideDNS,
			NetworkStart:     sch.NetworkStart,
			Name:             sch.Name,
			PingInterval:     sch.PingInterval,
			AllowedDevices:   sch.AllowedDevices,
			UsersOnline:      sch.UsersOnline,
			Ipv6Firewall:     sch.Ipv6Firewall,
			SessionTimeout:   sch.SessionTimeout,
			OtpAuth:          sch.OtpAuth,
			MultiDevice:      sch.MultiDevice,
			SearchDomain:     sch.SearchDomain,
			LzoCompression:   sch.LzoCompression,
			PreConnectMsg:    sch.PreConnectMsg,
			InactiveTimeout:  sch.InactiveTimeout,
			LinkPingInterval: sch.LinkPingInterval,
			Id:               sch.Id,
			PingTimeoutWg:    sch.PingTimeoutWg,
			Uptime:           sch.Uptime,
			NetworkEnd:       sch.NetworkEnd,
			Network:          sch.Network,
			DhParamBits:      sch.DhParamBits,
			Wg:               sch.Wg,
			Port:             sch.Port,
			DevicesOnline:    sch.DevicesOnline,
			NetworkWg:        sch.NetworkWg,
			Status:           sch.Status,
			DNSMapping:       sch.DNSMapping,
			Hash:             sch.Hash,
			Debug:            sch.Debug,
			RestrictRoutes:   sch.RestrictRoutes,
			UserCount:        sch.UserCount,
			Groups:           sch.Groups,
			InterClient:      sch.InterClient,
			ReplicaCount:     sch.ReplicaCount,
			Cipher:           sch.Cipher,
			MssFix:           sch.MssFix,
			JumboFrames:      sch.JumboFrames,
		},
	}

	data = &serverData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func serverPost(prvdr *schemas.Provider, sch *schemas.Server) (
	data *serverData, err error) {

	req := request.Request{
		Method: "POST",
		Path:   "/server",
		Json: &serverData{
			PortWg:           sch.PortWg,
			DNSServers:       sch.DNSServers,
			Protocol:         sch.Protocol,
			MaxDevices:       sch.MaxDevices,
			MaxClients:       sch.MaxClients,
			LinkPingTimeout:  sch.LinkPingTimeout,
			PingTimeout:      sch.PingTimeout,
			Ipv6:             sch.Ipv6,
			Vxlan:            sch.Vxlan,
			NetworkMode:      sch.NetworkMode,
			BindAddress:      sch.BindAddress,
			BlockOutsideDNS:  sch.BlockOutsideDNS,
			NetworkStart:     sch.NetworkStart,
			Name:             sch.Name,
			PingInterval:     sch.PingInterval,
			AllowedDevices:   sch.AllowedDevices,
			UsersOnline:      sch.UsersOnline,
			Ipv6Firewall:     sch.Ipv6Firewall,
			SessionTimeout:   sch.SessionTimeout,
			OtpAuth:          sch.OtpAuth,
			MultiDevice:      sch.MultiDevice,
			SearchDomain:     sch.SearchDomain,
			LzoCompression:   sch.LzoCompression,
			PreConnectMsg:    sch.PreConnectMsg,
			InactiveTimeout:  sch.InactiveTimeout,
			LinkPingInterval: sch.LinkPingInterval,
			Id:               sch.Id,
			PingTimeoutWg:    sch.PingTimeoutWg,
			Uptime:           sch.Uptime,
			NetworkEnd:       sch.NetworkEnd,
			Network:          sch.Network,
			DhParamBits:      sch.DhParamBits,
			Wg:               sch.Wg,
			Port:             sch.Port,
			DevicesOnline:    sch.DevicesOnline,
			NetworkWg:        sch.NetworkWg,
			Status:           sch.Status,
			DNSMapping:       sch.DNSMapping,
			Hash:             sch.Hash,
			Debug:            sch.Debug,
			RestrictRoutes:   sch.RestrictRoutes,
			UserCount:        sch.UserCount,
			Groups:           sch.Groups,
			InterClient:      sch.InterClient,
			ReplicaCount:     sch.ReplicaCount,
			Cipher:           sch.Cipher,
			MssFix:           sch.MssFix,
			JumboFrames:      sch.JumboFrames,
		},
	}

	data = &serverData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		err = errors.New("server: Not found on post")

		return
	}

	return
}

func serverDel(prvdr *schemas.Provider, sch *schemas.Server) (
	err error) {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/server/%s", sch.Id),
	}

	_, err = req.Do(prvdr, nil)

	if err != nil {
		return
	}

	return
}

func serverCreate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServer(d)

	data, err := serverGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.Id = data.Id

		data, err = serverPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	if data == nil {
		data, err = serverPost(prvdr, sch)
		if err != nil {
			return
		}
	}

	d.SetId(data.Id)

	return
}

func serverUpdate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServer(d)

	data, err := serverPut(prvdr, sch)
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

func serverRead(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServer(d)

	data, err := serverGet(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		return
	}

	d.Set("name", data.Name)
	d.SetId(data.Id)
	// d.Set("port", data.Port)

	return
}

func serverDelete(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServer(d)

	err = serverDel(prvdr, sch)
	if err != nil {
		return
	}

	d.SetId("")

	return
}
