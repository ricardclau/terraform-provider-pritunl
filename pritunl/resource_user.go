package pritunl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
)

func ResourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeList,
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
				Type:     schema.TypeList,
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
				Type:     schema.TypeList,
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
		Groups:          nil,
		Pin:             d.Get("pin").(string),
		Disabled:        false,
		NetworkLinks:    nil,
		BypassSecondary: false,
		ClientToClient:  false,
		DnsServers:      nil,
		DnsSuffix:       "",
	}

	userData, err := c.UserCreate(organizationId, u)
	if err != nil {
		return err
	}

	d.SetId(userData.Id)

	return resourceUserRead(d, m)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	organizationId := d.Get("organization_id").(string)
	data, err := c.UserGet(organizationId, d.Id())
	if err != nil {
		return err
	}

	d.Set("name", data.Name)
	d.Set("auth_type", data.AuthType)
	d.Set("email", data.Email)
	// d.Set("pin", data.Pin)

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	organizationId := d.Get("organization_id").(string)

	u := client.UserPostData{
		Name:            d.Get("name").(string),
		Email:           d.Get("email").(string),
		AuthType:        d.Get("auth_type").(string),
		Groups:          nil,
		Pin:             d.Get("pin").(string),
		Disabled:        false,
		NetworkLinks:    nil,
		BypassSecondary: false,
		ClientToClient:  false,
		DnsServers:      nil,
		DnsSuffix:       "",
	}

	_, err := c.UserUpdate(organizationId, d.Id(), u)
	if err != nil {
		return err
	}

	return resourceUserRead(d, m)
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	organizationId := d.Get("organization_id").(string)
	err := c.UserDelete(organizationId, d.Id())
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
