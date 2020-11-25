package pritunl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
)

func DataSourceServer() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceServerRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"groups": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dns_servers": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipv6": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"otp_auth": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"wg": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dh_param_bits": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cipher": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ping_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"link_ping_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"inactive_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"session_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_clients": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"network_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mss_fix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_devices": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"pre_connect_msg": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bind_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hash": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ping_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"link_ping_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"allowed_devices": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"search_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replica_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"multi_device": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"debug": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"restrict_routes": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"block_outside_dns": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dns_mapping": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"inter_client": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"vxlan": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func DataSourceServerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)
	serverName := d.Get("name").(string)

	data, err := c.ServerGetByName(serverName)
	if err != nil {
		return err
	}

	d.SetId(data.Id)
	d.Set("name", data.Name)
	d.Set("network", data.Network)
	d.Set("port", data.Port)
	d.Set("protocol", data.Protocol)
	d.Set("groups", data.Groups)
	d.Set("dns_servers", data.DNSServers)
	d.Set("ipv6", data.Ipv6)
	d.Set("otp_auth", data.OtpAuth)
	d.Set("wg", data.Wg)
	d.Set("dh_param_bits", data.DhParamBits)
	d.Set("cipher", data.Cipher)
	d.Set("ping_interval", data.PingInterval)
	d.Set("link_ping_interval", data.LinkPingInterval)
	d.Set("inactive_timeout", data.InactiveTimeout)
	d.Set("session_timeout", data.SessionTimeout)
	d.Set("max_clients", data.MaxClients)
	d.Set("network_mode", data.NetworkMode)
	d.Set("mss_fix", data.MssFix)
	d.Set("max_devices", data.MaxDevices)
	d.Set("pre_connect_msg", data.PreConnectMsg)
	d.Set("bind_address", data.BindAddress)
	d.Set("hash", data.Hash)
	d.Set("ping_timeout", data.PingTimeout)
	d.Set("link_ping_timeout", data.LinkPingTimeout)
	d.Set("allowed_devices", data.AllowedDevices)
	d.Set("search_domain", data.SearchDomain)
	d.Set("replica_count", data.ReplicaCount)
	d.Set("multi_device", data.MultiDevice)
	d.Set("debug", data.Debug)
	d.Set("restrict_routes", data.RestrictRoutes)
	d.Set("block_outside_dns", data.BlockOutsideDNS)
	d.Set("dns_mapping", data.DNSMapping)
	d.Set("inter_client", data.InterClient)
	d.Set("vxlan", data.Vxlan)

	return nil
}
