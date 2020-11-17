package pritunl

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
)

func ResourceOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourceOrganizationCreate,
		Read:   resourceOrganizationRead,
		Update: resourceOrganizationUpdate,
		Delete: resourceOrganizationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceOrganizationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	o := client.OrganizationPostData{
		Name: d.Get("name").(string),
	}

	data, err := c.OrganizationCreate(o)
	if err != nil {
		return err
	}

	d.SetId(data.Id)

	return resourceOrganizationRead(d, m)
}

func resourceOrganizationUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	o := client.OrganizationPostData{
		Name: d.Get("name").(string),
	}

	_, err := c.OrganizationUpdate(d.Id(), o)
	if err != nil {
		return err
	}

	return resourceOrganizationRead(d, m)
}

func resourceOrganizationRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	data, err := c.OrganizationGet(d.Id())
	if err != nil {
		return err
	}

	if data == nil {
		return errors.New(fmt.Sprintf("Cannot Read Organization %v", d.Id()))
	}

	d.Set("name", data.Name)

	return nil
}

func resourceOrganizationDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	err := c.OrganizationDelete(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
