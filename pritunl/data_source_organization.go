package pritunl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/client"
)

func DataSourceOrganization() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceOrganizationRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func DataSourceOrganizationRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.PritunlClient)
	organizationName := d.Get("name").(string)

	org, err := c.OrganizationGetByName(organizationName)
	if err != nil {
		return err
	}

	d.SetId(org.Id)

	return nil
}
