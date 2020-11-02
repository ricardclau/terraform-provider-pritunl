package resources

import (
	"fmt"
	"log"

	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/request"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
)

func User() *schema.Resource {
	return &schema.Resource{
		Create: userCreate,
		Read:   userRead,
		Update: userUpdate,
		Delete: userDelete,
		Schema: map[string]*schema.Schema{
			"organization_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"network_links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"bypass_secondary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"client_to_client": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dns_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dns_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

type userPostData struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	AuthType        string   `json:"auth_type"`
	Groups          []string `json:"groups,omitempty"`
	Pin             string   `json:"pin"`
	Disabled        bool     `json:"disabled"`
	NetworkLinks    []string `json:"network_links,omitempty"`
	BypassSecondary bool     `json:"bypass_secondary"`
	ClientToClient  bool     `json:"client_to_client"`
	DnsServers      []string `json:"dns_servers,omitempty"`
	DnsSuffix       string   `json:"dns_suffix,omitempty"`
}

type userPutData struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	AuthType        string   `json:"auth_type"`
	Groups          []string `json:"groups,omitempty"`
	Pin             string   `json:"pin"`
	Disabled        bool     `json:"disabled"`
	NetworkLinks    []string `json:"network_links,omitempty"`
	BypassSecondary bool     `json:"bypass_secondary"`
	ClientToClient  bool     `json:"client_to_client"`
	DnsServers      []string `json:"dns_servers,omitempty"`
	DnsSuffix       string   `json:"dns_suffix,omitempty"`
}

type userData struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	AuthType        string   `json:"auth_type"`
	Groups          []string `json:"groups,omitempty"`
	Disabled        bool     `json:"disabled"`
	NetworkLinks    []string `json:"network_links,omitempty"`
	BypassSecondary bool     `json:"bypass_secondary"`
	ClientToClient  bool     `json:"client_to_client"`
	DnsServers      []string `json:"dns_servers,omitempty"`
	DnsSuffix       string   `json:"dns_suffix,omitempty"`
}

func userGet(prvdr *schemas.Provider, sch *schemas.User) (*userData, error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/user/%s/%s", sch.OrganizationId, sch.Id),
	}

	data := &userData{}
	_, err := req.Do(prvdr, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func userPut(prvdr *schemas.Provider, sch *schemas.User) (*userData, error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/user/%s/%s", sch.OrganizationId, sch.Id),
		Json: &userPutData{
			Name:            sch.Name,
			Email:           sch.Email,
			AuthType:        sch.AuthType,
			Groups:          sch.Groups,
			Pin:             sch.Pin,
			Disabled:        sch.Disabled,
			NetworkLinks:    sch.NetworkLinks,
			BypassSecondary: sch.BypassSecondary,
			ClientToClient:  sch.ClientToClient,
			DnsServers:      sch.DnsServers,
			DnsSuffix:       sch.DnsSuffix,
		},
	}

	data := &userData{}

	_, err := req.Do(prvdr, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func userPost(prvdr *schemas.Provider, sch *schemas.User) (*userData, error) {

	req := request.Request{
		Method: "POST",
		Path:   fmt.Sprintf("/user/%s", sch.OrganizationId),
		Json: &userPostData{
			Name:            sch.Name,
			Email:           sch.Email,
			AuthType:        sch.AuthType,
			Groups:          sch.Groups,
			Pin:             sch.Pin,
			Disabled:        sch.Disabled,
			NetworkLinks:    sch.NetworkLinks,
			BypassSecondary: sch.BypassSecondary,
			ClientToClient:  sch.ClientToClient,
			DnsServers:      sch.DnsServers,
			DnsSuffix:       sch.DnsSuffix,
		},
	}

	data := []userData{}

	_, err := req.Do(prvdr, &data)
	if err != nil {
		return nil, err
	}

	ret := data[0]
	return &ret, nil
}

func userDel(prvdr *schemas.Provider, sch *schemas.User) error {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/user/%s/%s", sch.OrganizationId, sch.Id),
	}

	_, err := req.Do(prvdr, nil)
	if err != nil {
		return err
	}

	return nil
}

func userCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Inside userCreate: %s", d)

	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadUser(d)

	log.Printf("[DEBUG] Before userPost: %s", d)
	data, err := userPost(prvdr, sch)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Before setting data: %s", data)
	d.SetId(data.Id)

	return nil
}

func userUpdate(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadUser(d)

	data, err := userPut(prvdr, sch)
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

func userRead(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadUser(d)

	data, err := userGet(prvdr, sch)
	if err != nil {
		return err
	}

	if data == nil {
		return errors.New(fmt.Sprintf("Cannot Read User %v", d.Id()))
	}

	d.Set("name", data.Name)
	d.Set("email", data.Email)
	d.Set("auth_type", data.AuthType)
	d.Set("groups", data.Groups)
	d.Set("disabled", data.Disabled)
	d.Set("network_links", data.NetworkLinks)
	d.Set("bypass_secondary", data.BypassSecondary)
	d.Set("client_to_client", data.ClientToClient)
	d.Set("dns_servers", data.DnsServers)
	d.Set("dns_suffix", data.DnsSuffix)
	d.SetId(data.Id)

	return nil
}

func userDelete(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadUser(d)

	err := userDel(prvdr, sch)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
