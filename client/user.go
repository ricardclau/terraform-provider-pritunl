package client

import (
	"fmt"
)

type UserPostData struct {
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

type UserData struct {
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

func (c *PritunlClient) UserGet(organizationId string, userId string) (*UserData, error) {
	req := Request{
		Method: "GET",
		Path:   fmt.Sprintf("/user/%s/%s", organizationId, userId),
	}

	data := &UserData{}
	err := c.Do(req, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *PritunlClient) UserCreate(organizationId string, u UserPostData) (*UserData, error) {
	req := Request{
		Method: "POST",
		Path:   fmt.Sprintf("/user/%s", organizationId),
		Json:   u,
	}

	var data []UserData

	err := c.Do(req, &data)
	if err != nil {
		return nil, err
	}

	ret := data[0]
	return &ret, nil
}

func (c *PritunlClient) UserUpdate(organizationId string, userId string, u UserPostData) (*UserData, error) {
	req := Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/user/%s/%s", organizationId, userId),
		Json:   u,
	}

	data := &UserData{}
	err := c.Do(req, data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (c *PritunlClient) UserDelete(organizationId string, userId string) error {

	req := Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/user/%s/%s", organizationId, userId),
	}

	return c.Do(req, nil)
}
