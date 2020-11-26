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
		Delete: resourceLinkServerOrganizationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"server": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLinkServerOrganizationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	server := d.Get("server").(string)
	organizationId := d.Get("organization_id").(string)

	_, err := c.LinkServerOrganizationAttach(server, organizationId)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s:%s", server, organizationId))

	return resourceLinkServerOrganizationRead(d, m)
}

func resourceLinkServerOrganizationRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)

	server, organizationId, err := resourceLinkServerOrganizationParseId(d.Id())
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

	server, organizationId, err := resourceLinkServerOrganizationParseId(d.Id())
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

func resourceLinkServerOrganizationParseId(id string) (server, organizationId string, err error) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 {
		err = fmt.Errorf("linkServerOrganization id must be of the form <server>:<org_id>")
		return
	}

	server = parts[0]
	organizationId = parts[1]
	return
}
