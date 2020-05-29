package schemas

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type LinkServerOrganization struct {
	Id             string      `json:"id"`
	OrganizationId string      `json:"organization_id"`
	Server         string      `json:"server"`
	Name           interface{} `json:"name"`
}

func LoadLinkServerOrganization(d *schema.ResourceData) (sch *LinkServerOrganization) {
	sch = &LinkServerOrganization{
		Id:             d.Id(),
		OrganizationId: d.Get("organization_id").(string),
		Server:         d.Get("server").(string),
	}

	return
}
