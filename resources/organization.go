package resources

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/request"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
)

func Organization() *schema.Resource {
	return &schema.Resource{
		Create: organizationCreate,
		Read:   organizationRead,
		Update: organizationUpdate,
		Delete: organizationDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type organizationPostData struct {
	Name string `json:"name"`
}

type organizationPutData struct {
	Name string `json:"name"`
}

type organizationData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func organizationGet(prvdr *schemas.Provider, sch *schemas.Organization) (*organizationData, error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/organization/%s", sch.Id),
	}

	data := &organizationData{}
	_, err := req.Do(prvdr, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func organizationPut(prvdr *schemas.Provider, sch *schemas.Organization) (*organizationData, error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/organization/%s", sch.Id),
		Json: &organizationPutData{
			Name: sch.Name,
		},
	}

	data := &organizationData{}

	_, err := req.Do(prvdr, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func organizationPost(prvdr *schemas.Provider, sch *schemas.Organization) (*organizationData, error) {

	req := request.Request{
		Method: "POST",
		Path:   "/organization",
		Json: &organizationPostData{
			Name: sch.Name,
		},
	}

	data := &organizationData{}

	_, err := req.Do(prvdr, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func organizationDel(prvdr *schemas.Provider, sch *schemas.Organization) error {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/organization/%s", sch.Id),
	}

	_, err := req.Do(prvdr, nil)

	if err != nil {
		return err
	}

	return nil
}

func organizationCreate(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadOrganization(d)

	data, err := organizationPost(prvdr, sch)
	if err != nil {
		return err
	}

	d.SetId(data.Id)

	return nil
}

func organizationUpdate(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadOrganization(d)

	data, err := organizationPut(prvdr, sch)
	if err != nil {
		return err
	}

	if data == nil {
		d.SetId("")
		return nil
	}

	d.SetId(data.Id)

	return nil
}

func organizationRead(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadOrganization(d)

	data, err := organizationGet(prvdr, sch)
	if err != nil {
		return err
	}

	if data == nil {
		return errors.New(fmt.Sprintf("Cannot Read Organization %v", d.Id()))
	}

	d.Set("name", data.Name)
	d.SetId(data.Id)

	return nil
}

func organizationDelete(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadOrganization(d)

	err := organizationDel(prvdr, sch)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
