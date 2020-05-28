package resources

import (
	"fmt"

	"github.com/dropbox/godropbox/errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/kihahu/terraform-provider-pritunl/errortypes"
	"github.com/kihahu/terraform-provider-pritunl/request"
	"github.com/kihahu/terraform-provider-pritunl/schemas"
)

func Server() *schema.Resource {
	return &schema.Resource{
		Create: serverCreate,
		Read:   serverRead,
		Update: serverUpdate,
		Delete: serverDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type serverPostData struct {
	Name string `json:"name"`
}

type serverPutData struct {
	Name string `json:"name"`
}

type serverData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func serverGet(prvdr *schemas.Provider, sch *schemas.Server) (
	data *serverData, err error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/server/%s", sch.Id),
	}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 || resp.StatusCode == 401 {
		data = nil
	}

	return
}

func serverPut(prvdr *schemas.Provider, sch *schemas.Server) (
	data *serverData, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/server/%s", sch.Id),
		Json: &serverPutData{
			Name: sch.Name,
		},
	}

	data = &serverData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func serverPost(prvdr *schemas.Provider, sch *schemas.Server) (
	data *serverData, err error) {

	req := request.Request{
		Method: "POST",
		Path:   "/server",
		Json: &serverPostData{
			Name: sch.Name,
		},
	}

	data = &serverData{}

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

func serverDel(prvdr *schemas.Provider, sch *schemas.Server) (
	err error) {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/server/%s", sch.Id),
	}

	_, err = req.Do(prvdr, nil)

	if err != nil {
		return
	}

	return
}

func serverCreate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServer(d)

	data, err := serverGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.Id = data.Id

		data, err = serverPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	if data == nil {
		data, err = serverPost(prvdr, sch)
		if err != nil {
			return
		}
	}

	d.SetId(data.Id)

	return
}

func serverUpdate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServer(d)

	data, err := serverPut(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		d.SetId("")
		return
	}

	d.SetId(data.Id)

	return
}

func serverRead(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServer(d)

	data, err := serverGet(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		return
	}

	d.Set("name", data.Name)
	d.SetId(data.Id)

	return
}

func serverDelete(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadServer(d)

	err = serverDel(prvdr, sch)
	if err != nil {
		return
	}

	d.SetId("")

	return
}
