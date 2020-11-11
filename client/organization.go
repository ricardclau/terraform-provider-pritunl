package client

import (
	"fmt"
)

type OrganizationPostData struct {
	Name string `json:"name"`
}

type OrganizationData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c *PritunlClient) OrganizationGetByName(name string) (*OrganizationData, error) {
	req := Request{
		Method: "GET",
		Path:   fmt.Sprintf("/organization"),
	}

	var data []OrganizationData

	err := c.Do(req, &data)
	if err != nil {
		return nil, err
	}

	for _, org := range data {
		if org.Name == name {
			return &org, nil
		}
	}

	return nil, fmt.Errorf("Cannot find organization with name: %v", name)
}

func (c *PritunlClient) OrganizationGet(id string) (*OrganizationData, error) {

	req := Request{
		Method: "GET",
		Path:   fmt.Sprintf("/organization/%s", id),
	}

	data := &OrganizationData{}
	err := c.Do(req, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *PritunlClient) OrganizationUpdate(id string, o OrganizationPostData) (*OrganizationData, error) {

	req := Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/organization/%s", id),
		Json:   &o,
	}

	data := &OrganizationData{}

	err := c.Do(req, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *PritunlClient) OrganizationCreate(o OrganizationPostData) (*OrganizationData, error) {

	req := Request{
		Method: "POST",
		Path:   "/organization",
		Json:   &o,
	}

	data := &OrganizationData{}

	err := c.Do(req, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *PritunlClient) OrganizationDelete(id string) error {

	req := Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/organization/%s", id),
	}

	return c.Do(req, nil)
}
