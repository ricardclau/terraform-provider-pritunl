package resources

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/request"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
)

func Settings() *schema.Resource {
	return &schema.Resource{
		Create: settingsCreate,
		Read:   settingsRead,
		Update: settingsUpdate,
		Delete: settingsDelete,

		Schema: map[string]*schema.Schema{
			"sso": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sso_google_email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sso_match": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_org": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type SettingsData struct {
	Theme          string   `json:"theme"`
	SsoOrg         string   `json:"sso_org"`
	Sso            string   `json:"sso"`
	SsoMatch       []string `json:"sso_match"`
	ID             string   `json:"id"`
	SsoGoogleEmail string   `json:"sso_google_email"`
}

func settingGet(prvdr *schemas.Provider, sch *schemas.Settings) (
	data *schemas.Settings, err error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/settings"),
	}

	xdata := &schemas.Settings{}

	resp, err := req.Do(prvdr, xdata)

	if err != nil {
		return
	}

	if resp.StatusCode < 405 {
		data = nil
	}

	return xdata, err
}

func settingsPut(prvdr *schemas.Provider, sch *schemas.Settings) (
	data *schemas.Settings, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/settings"),
		Json: &SettingsData{
			Sso:            sch.Sso,
			SsoMatch:       sch.SsoMatch,
			SsoOrg:         sch.SsoOrg,
			SsoGoogleEmail: sch.SsoGoogleEmail,
			Theme:          sch.Theme,
			ID:             sch.ID,
		},
	}

	xdata := &schemas.Settings{}

	resp, err := req.Do(prvdr, xdata)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return xdata, err
}

func settingPost(prvdr *schemas.Provider, sch *schemas.Settings) (data *schemas.Settings, err error) {

	req := request.Request{
		Method: "POST",
		Path:   "/settings",
		Json: &SettingsData{
			Sso:            sch.Sso,
			SsoMatch:       sch.SsoMatch,
			SsoOrg:         sch.SsoOrg,
			SsoGoogleEmail: sch.SsoGoogleEmail,
			Theme:          sch.Theme,
			ID:             sch.ID,
		},
	}

	data = &schemas.Settings{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		err = errors.New("server: Not found on post")

		return
	}

	return
}

func settingDel(prvdr *schemas.Provider, sch *schemas.Settings) (err error) {
	return
}

func settingsCreate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadSettings(d)

	data, err := settingGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.ID = data.ID

		data, err = settingsPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	if data == nil {
		return
	}

	d.SetId(data.ID)

	return
}

func settingsUpdate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadSettings(d)

	data, err := settingsPut(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		return
	}

	d.SetId(data.ID)

	return
}

func settingsRead(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadSettings(d)

	data, err := settingGet(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		return
	}

	d.SetId(data.ID)

	return
}

func settingsDelete(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadSettings(d)

	err = settingDel(prvdr, sch)
	if err != nil {
		return
	}

	d.SetId("")

	return
}
