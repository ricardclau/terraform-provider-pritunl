package resources

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/request"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
)

// LinkServerOrganization Link server to orgaminzation
func LinkServerOrganization() *schema.Resource {
	return &schema.Resource{
		Create: linkServerOrganizationCreateOrUpdate,
		Read:   linkServerOrganizationRead,
		Update: linkServerOrganizationCreateOrUpdate,
		Delete: linkServerOrganizationDelete,

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

type linkServerOrganizationData struct {
	Id             string      `json:"id"`
	OrganizationId string      `json:"organization_id"`
	Server         string      `json:"server"`
	Name           interface{} `json:"name"`
}

func linkServerOrganizationGet(prvdr *schemas.Provider, sch *schemas.LinkServerOrganization) (*linkServerOrganizationData, error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/server/%s/organization/%s", sch.Server, sch.OrganizationId),
	}

	data := &linkServerOrganizationData{}
	_, err := req.Do(prvdr, data)

	return data, err
}

func linkServerOrganizationPut(prvdr *schemas.Provider, sch *schemas.LinkServerOrganization) (*linkServerOrganizationData, error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/server/%s/organization/%s", sch.Server, sch.OrganizationId),
		Json: &linkServerOrganizationData{
			Server: sch.Server,
			Id:     sch.OrganizationId,
		},
	}

	data := &linkServerOrganizationData{}
	_, err := req.Do(prvdr, data)

	return data, err
}

func linkServerOrganizationPost(prvdr *schemas.Provider, sch *schemas.LinkServerOrganization) (*linkServerOrganizationData, error) {

	req := request.Request{
		Method: "POST",
		Path:   fmt.Sprintf("/server/%s/organization/%s", sch.Server, sch.OrganizationId),
		Json: &linkServerOrganizationData{
			Server: sch.Server,
			Id:     sch.OrganizationId,
		},
	}

	data := &linkServerOrganizationData{}
	_, err := req.Do(prvdr, data)

	return data, err
}

func linkServerOrganizationDel(prvdr *schemas.Provider, sch *schemas.LinkServerOrganization) error {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/server/%s/organization/%s", sch.Server, sch.OrganizationId),
	}

	_, err := req.Do(prvdr, nil)

	return err
}

func linkServerOrganizationCreateOrUpdate(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkServerOrganization(d)

	data, err := linkServerOrganizationPut(prvdr, sch)
	if err != nil {
		return err
	}

	d.SetId(data.Id)

	return nil
}

func linkServerOrganizationRead(d *schema.ResourceData, m interface{}) error {
	sch := schemas.LoadLinkServerOrganization(d)
	d.Set("server", sch.Server)
	d.Set("organization_id", sch.OrganizationId)
	d.SetId(sch.Id)

	return nil
}

func linkServerOrganizationDelete(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkServerOrganization(d)

	err := linkServerOrganizationDel(prvdr, sch)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
