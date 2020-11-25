package pritunl

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
)

func ResourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"auth_type": {
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
			"pin": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"network_links": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"bypass_secondary": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"client_to_client": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dns_servers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dns_suffix": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	organizationId := d.Get("organization_id").(string)

	u := client.UserPostData{
		Name:            d.Get("name").(string),
		Email:           d.Get("email").(string),
		AuthType:        d.Get("auth_type").(string),
		Groups:          expandStringListFromSetSchema(d.Get("groups").(*schema.Set)),
		Pin:             d.Get("pin").(string),
		Disabled:        d.Get("disabled").(bool),
		NetworkLinks:    expandStringListFromSetSchema(d.Get("network_links").(*schema.Set)),
		BypassSecondary: d.Get("bypass_secondary").(bool),
		ClientToClient:  d.Get("client_to_client").(bool),
		DnsServers:      expandStringListFromSetSchema(d.Get("dns_servers").(*schema.Set)),
		DnsSuffix:       d.Get("dns_suffix").(string),
	}

	userData, err := c.UserCreate(organizationId, u)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s:%s", organizationId, userData.Id))

	return resourceUserRead(d, m)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	organizationId, userId, err := resourceUserParseId(d.Id())
	if err != nil {
		return err
	}

	data, err := c.UserGet(organizationId, userId)
	if err != nil {
		return err
	}

	d.Set("organization_id", organizationId)
	d.Set("name", data.Name)
	d.Set("auth_type", data.AuthType)
	d.Set("email", data.Email)
	d.Set("disabled", data.Disabled)
	d.Set("bypass_secondary", data.BypassSecondary)
	d.Set("client_to_client", data.ClientToClient)
	d.Set("groups", data.Groups)
	d.Set("network_links", data.NetworkLinks)
	d.Set("dns_servers", data.DnsServers)
	d.Set("dns_suffix", data.DnsSuffix)

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	organizationId, userId, err := resourceUserParseId(d.Id())
	if err != nil {
		return err
	}

	u := client.UserPostData{
		Name:            d.Get("name").(string),
		Email:           d.Get("email").(string),
		AuthType:        d.Get("auth_type").(string),
		Groups:          expandStringListFromSetSchema(d.Get("groups").(*schema.Set)),
		Pin:             d.Get("pin").(string),
		Disabled:        d.Get("disabled").(bool),
		NetworkLinks:    expandStringListFromSetSchema(d.Get("network_links").(*schema.Set)),
		BypassSecondary: d.Get("bypass_secondary").(bool),
		ClientToClient:  d.Get("client_to_client").(bool),
		DnsServers:      expandStringListFromSetSchema(d.Get("dns_servers").(*schema.Set)),
		DnsSuffix:       d.Get("dns_suffix").(string),
	}

	_, err = c.UserUpdate(organizationId, userId, u)
	if err != nil {
		return err
	}

	return resourceUserRead(d, m)
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	organizationId, userId, err := resourceUserParseId(d.Id())
	if err != nil {
		return err
	}

	err = c.UserDelete(organizationId, userId)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func resourceUserParseId(id string) (organizationId, userId string, err error) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 {
		err = fmt.Errorf("user id must be of the form <org_id>:<user_id>")
		return
	}

	organizationId = parts[0]
	userId = parts[1]
	return
}
