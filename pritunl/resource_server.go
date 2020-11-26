package pritunl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/pritunl/terraform-provider-pritunl/client"
	"strings"
)

func ResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServerCreate,
		Read:   ResourceServerRead,
		Update: ResourceServerUpdate,
		Delete: ResourceServerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
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
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dns_servers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ipv6": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"otp_auth": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"wg": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dh_param_bits": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      2048,
				ValidateFunc: validation.IntInSlice([]int{1024, 1536, 2048, 3072, 4096}),
			},
			"cipher": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "aes128",
				ValidateFunc: validation.StringInSlice([]string{"bf128", "bf256", "aes128", "aes192", "aes256"}, true),
			},
			"ping_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"link_ping_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"inactive_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"session_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_clients": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"network_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "tunnel",
				ValidateFunc: validation.StringInSlice([]string{"tunnel", "bridge"}, true),
			},
			"mss_fix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_devices": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pre_connect_msg": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bind_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hash": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "sha1",
				ValidateFunc: validation.StringInSlice([]string{"md5", "sha1", "sha256", "sha512"}, true),
			},
			"ping_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"link_ping_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"allowed_devices": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "any",
				ValidateFunc: validation.StringInSlice([]string{"any", "mobile", "desktop"}, true),
			},
			"search_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"replica_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"multi_device": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"debug": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"restrict_routes": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"block_outside_dns": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dns_mapping": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"inter_client": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"vxlan": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	s := client.ServerPostData{
		Name:             d.Get("name").(string),
		Network:          d.Get("network").(string),
		DNSServers:       expandStringListFromSetSchema(d.Get("dns_servers").(*schema.Set)),
		Port:             d.Get("port").(int),
		Protocol:         d.Get("protocol").(string),
		Groups:           expandStringListFromSetSchema(d.Get("groups").(*schema.Set)),
		Ipv6:             d.Get("ipv6").(bool),
		OtpAuth:          d.Get("otp_auth").(bool),
		Wg:               d.Get("wg").(bool),
		DhParamBits:      d.Get("dh_param_bits").(int),
		Cipher:           strings.ToLower(d.Get("cipher").(string)),
		PingInterval:     d.Get("ping_interval").(int),
		LinkPingInterval: d.Get("link_ping_interval").(int),
		InactiveTimeout:  d.Get("inactive_timeout").(int),
		SessionTimeout:   d.Get("session_timeout").(int),
		MaxClients:       d.Get("max_clients").(int),
		NetworkMode:      d.Get("network_mode").(string),
		MssFix:           d.Get("mss_fix").(string),
		MaxDevices:       d.Get("max_devices").(int),
		PreConnectMsg:    d.Get("pre_connect_msg").(string),
		BindAddress:      d.Get("bind_address").(string),
		Hash:             d.Get("hash").(string),
		PingTimeout:      d.Get("ping_timeout").(int),
		LinkPingTimeout:  d.Get("link_ping_timeout").(int),
		AllowedDevices:   d.Get("allowed_devices").(string),
		SearchDomain:     d.Get("search_domain").(string),
		ReplicaCount:     d.Get("replica_count").(int),
		MultiDevice:      d.Get("multi_device").(bool),
		Debug:            d.Get("debug").(bool),
		RestrictRoutes:   d.Get("restrict_routes").(bool),
		BlockOutsideDNS:  d.Get("block_outside_dns").(bool),
		DNSMapping:       d.Get("dns_mapping").(bool),
		InterClient:      d.Get("inter_client").(bool),
		Vxlan:            d.Get("vxlan").(bool),
	}

	data, err := c.ServerCreate(s)
	if err != nil {
		return err
	}

	d.SetId(data.Id)

	return ResourceServerRead(d, m)
}

func ResourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	s := client.ServerPostData{
		Name: d.Get("name").(string),
		// This is autopopulated in the UI ????
		Network:          d.Get("network").(string),
		DNSServers:       expandStringListFromSetSchema(d.Get("dns_servers").(*schema.Set)),
		Port:             d.Get("port").(int),
		Protocol:         d.Get("protocol").(string),
		Groups:           expandStringListFromSetSchema(d.Get("groups").(*schema.Set)),
		Ipv6:             d.Get("ipv6").(bool),
		OtpAuth:          d.Get("otp_auth").(bool),
		Wg:               d.Get("wg").(bool),
		DhParamBits:      d.Get("dh_param_bits").(int),
		Cipher:           strings.ToLower(d.Get("cipher").(string)),
		PingInterval:     d.Get("ping_interval").(int),
		LinkPingInterval: d.Get("link_ping_interval").(int),
		InactiveTimeout:  d.Get("inactive_timeout").(int),
		SessionTimeout:   d.Get("session_timeout").(int),
		MaxClients:       d.Get("max_clients").(int),
		NetworkMode:      d.Get("network_mode").(string),
		MssFix:           d.Get("mss_fix").(string),
		MaxDevices:       d.Get("max_devices").(int),
		PreConnectMsg:    d.Get("pre_connect_msg").(string),
		BindAddress:      d.Get("bind_address").(string),
		Hash:             d.Get("hash").(string),
		PingTimeout:      d.Get("ping_timeout").(int),
		LinkPingTimeout:  d.Get("link_ping_timeout").(int),
		AllowedDevices:   d.Get("allowed_devices").(string),
		SearchDomain:     d.Get("search_domain").(string),
		ReplicaCount:     d.Get("replica_count").(int),
		MultiDevice:      d.Get("multi_device").(bool),
		Debug:            d.Get("debug").(bool),
		RestrictRoutes:   d.Get("restrict_routes").(bool),
		BlockOutsideDNS:  d.Get("block_outside_dns").(bool),
		DNSMapping:       d.Get("dns_mapping").(bool),
		InterClient:      d.Get("inter_client").(bool),
		Vxlan:            d.Get("vxlan").(bool),
	}

	_, err := c.ServerUpdate(d.Id(), s)
	if err != nil {
		return err
	}

	return ResourceServerRead(d, m)
}

func ResourceServerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	data, err := c.ServerGet(d.Id())
	if err != nil {
		return err
	}

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

func ResourceServerDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	err := c.ServerDelete(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
