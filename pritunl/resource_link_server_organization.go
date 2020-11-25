package pritunl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
)

func ResourceLinkServerOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourceLinkServerOrganizationCreate,
		Read:   resourceLinkServerOrganizationRead,
		Update: resourceLinkServerOrganizationUpdate,
		Delete: resourceLinkServerOrganizationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceLinkServerOrganizationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	organizationId := d.Get("organizationId").(string)
	server := d.Get("server").(string)

	data, err := c.LinkServerOrganizationCreate(server, organizationId)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s:%s:%s", data.Id, server, organizationId))

	return resourceLinkServerOrganizationRead(d, m)
}

func resourceLinkServerOrganizationUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	_, server, organizationId, err := resourceLinkServerOrganizationParseId(d.Id())
	o := client.LinkServerOrganizationPostData{
		OrganizationId: d.Get("organizationId").(string),
		Server:         d.Get("server").(string),
	}

	_, err = c.LinkServerOrganizationUpdate(server, organizationId, o)
	if err != nil {
		return err
	}

	return resourceLinkServerOrganizationRead(d, m)
}

func resourceLinkServerOrganizationRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	_, server, organizationId, err := resourceLinkServerOrganizationParseId(d.Id())
	if err != nil {
		return err
	}

	data, err := c.LinkServerOrganizationGet(server, organizationId)
	if err != nil {
		return err
	}

	if data == nil {
		return errors.New(fmt.Sprintf("Cannot Read LinkServerOrganizationId %v", d.Id()))
	}

	d.Set("server", data.Server)
	d.Set("organization_id", data.OrganizationId)

	return nil
}

func resourceLinkServerOrganizationDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	_, server, organizationId, err := resourceLinkServerOrganizationParseId(d.Id())
	if err != nil {
		return err
	}

	err = c.LinkServerOrganizationDelete(server, organizationId)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func resourceLinkServerOrganizationParseId(id string) (resourceId, server, organizationId string, err error) {
	parts := strings.SplitN(id, ":", 3)
	if len(parts) != 3 {
		err = fmt.Errorf("linkServerOrganization id must be of the form <id>:<server>:<org_id>")
		return
	}

	resourceId = parts[0]
	server = parts[1]
	organizationId = parts[2]
	return
}
