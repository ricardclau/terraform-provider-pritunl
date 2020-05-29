package resources

import (
	"fmt"

	"github.com/dropbox/godropbox/errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/kihahu/terraform-provider-pritunl/errortypes"
	"github.com/kihahu/terraform-provider-pritunl/request"
	"github.com/kihahu/terraform-provider-pritunl/schemas"
)

// LinkServerOrganization Link server to orgaminzation
func LinkServerOrganization() *schema.Resource {
	return &schema.Resource{
		Create: linkServerOrganizationCreate,
		Read:   linkServerOrganizationRead,
		Update: linkServerOrganizationUpdate,
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

func linkServerOrganizationGet(prvdr *schemas.Provider, sch *schemas.LinkServerOrganization) (
	data *linkServerOrganizationData, err error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/server/%s/organization/%s", sch.Server, sch.OrganizationId),
	}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode < 405 {
		return
	}

	return
}

func linkServerOrganizationPut(prvdr *schemas.Provider, sch *schemas.LinkServerOrganization) (
	data *linkServerOrganizationData, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/server/%s/organization/%s", sch.Server, sch.OrganizationId),
		Json: &linkServerOrganizationData{
			Server: sch.Server,
			Id:     sch.OrganizationId,
		},
	}

	data = &linkServerOrganizationData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func linkServerOrganizationPost(prvdr *schemas.Provider, sch *schemas.LinkServerOrganization) (
	data *linkServerOrganizationData, err error) {

	req := request.Request{
		Method: "POST",
		Path:   fmt.Sprintf("/server/%s/organization/%s", sch.Server, sch.OrganizationId),
		Json: &linkServerOrganizationData{
			Server: sch.Server,
			Id:     sch.OrganizationId,
		},
	}

	data = &linkServerOrganizationData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		err = &errortypes.RequestError{
			errors.New("server: Not found on post"),
		}
		return
	}

	return
}

func linkServerOrganizationDel(prvdr *schemas.Provider, sch *schemas.LinkServerOrganization) (
	err error) {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/server/%s/organization/%s", sch.Server, sch.OrganizationId),
	}

	_, err = req.Do(prvdr, nil)

	if err != nil {
		return
	}

	return
}

func linkServerOrganizationCreate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkServerOrganization(d)

	data, err := linkServerOrganizationGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.OrganizationId = data.Id

		data, err = linkServerOrganizationPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	if data == nil {
		data, err = linkServerOrganizationPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	d.SetId(data.Id)

	return
}

func linkServerOrganizationUpdate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkServerOrganization(d)

	data, err := linkServerOrganizationPut(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		// d.SetId("")
		return
	}

	d.SetId(data.Id)

	return
}

func linkServerOrganizationRead(d *schema.ResourceData, m interface{}) (err error) {
	// prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkServerOrganization(d)

	// data, err := linkServerOrganizationGet(prvdr, sch)
	// if err != nil {
	// 	return
	// }

	// if data == nil {
	// 	return
	// }

	d.Set("server", sch.Server)
	d.Set("organization_id", sch.OrganizationId)
	d.SetId(sch.Id)

	return
}

func linkServerOrganizationDelete(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkServerOrganization(d)

	err = linkServerOrganizationDel(prvdr, sch)
	if err != nil {
		return
	}

	d.SetId("")

	return
}
