package schemas

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type Server struct {
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

func LoadServer(d *schema.ResourceData) (sch *Server) {
	sch = &Server{
		Id:               d.Id(),
		Protocol:         "udp",
		MaxDevices:       0,
		MaxClients:       2000,
		LinkPingTimeout:  5,
		PingTimeout:      60,
		Ipv6:             false,
		Vxlan:            false,
		NetworkMode:      "tunnel",
		BindAddress:      "",
		BlockOutsideDNS:  false,
		NetworkStart:     "",
		Name:             d.Get("name").(string),
		PingInterval:     10,
		UsersOnline:      0,
		Ipv6Firewall:     true,
		OtpAuth:          false,
		MultiDevice:      false,
		LzoCompression:   false,
		LinkPingInterval: 1,
		PingTimeoutWg:    360,
		NetworkEnd:       "",
		Network:          d.Get("network").(string),
		DhParamBits:      2048,
		Wg:               false,
		Port:             d.Get("port").(int),
		DevicesOnline:    0,
		Status:           "offline",
		DNSMapping:       false,
		Hash:             "sha1",
		Debug:            false,
		RestrictRoutes:   true,
		UserCount:        0,
		// Groups:           utils.ExpandStringList(d.Get("groups").(*schema.Set).List()),
		InterClient:  true,
		ReplicaCount: 1,
		Cipher:       "aes128",
		JumboFrames:  false,
	}

	dnsServers := d.Get("dns_servers").([]interface{})
	if dnsServers != nil {
		sch.DNSServers = []string{}
		for _, dnsServer := range dnsServers {
			sch.DNSServers = append(sch.DNSServers, dnsServer.(string))
		}
	}

	groups := d.Get("groups").([]interface{})
	if groups != nil {
		sch.Groups = []string{}
		for _, group := range groups {
			sch.Groups = append(sch.Groups, group.(string))
		}
	}

	return
}
